# 自定义模板

### 生成对象定义

1. json结构说明

```json
// objectField
{
    "name": "TodoInput", // 对象/字段名
    "typeName": "object", // 字段类型名(integer/number/boolean/string/object/enum)
    "typeRefObject": ${objectField}, // 字段类型引用对象
    "typeRefEnum": ${enumField}, // 字段类型引用枚举
    "required": false, // 字段是否必须
    "isArray": false, // 字段类型是否数组
    "isDefinition": false, // 对象是否全局定义(orderBy, query这些数据库内省的)
    "documentPath": ["TodoInput"], // 文档路径(建议拼接后用来做对象名/字段类型名)
    "fields": [${objectField}], // 对象字段列表
    "root": "Input", // 顶层归属类型(Input/InternalInput/ResponseData/Definitions)
    "opeartionInfo": ${opeartionInfo} // 对应operation信息
}

// enumField
{
    "name": "DictValueType", // 枚举名称
    "values": ["site"] // 枚举值列表
}

// operationInfo
{
    "name": "Todo", // oeprationName
    "path": "Todo", // opeartionPath
    "IsInternal": false, // 是否内部
    "isLiveQuery": false, // 是否实时查询
    "isQuery": true, // 是否查询
    "isMutation": false, // 是否变更
    "isSubscription": false // 是否订阅
}
```

2. 定义对象类型

```handlebars
<!-- golang生成struct使用示例 -->
{{#each objectFieldArray}}
<!-- 使用documentPath拼接'_'对象名称 -->
type {{upperFirst (joinString '_' documentPath)}} struct {
    <!-- 遍历字段列表 -->
    {{#each fields}}
    <!-- 字段名首字母大写，并判断类型是否为数组 -->
    {{upperFirst (trimPrefix name '_')}} {{#if isArray}}[]{{~/if~}}
    {{~#if typeRefObject~}}
        <!-- 使用关联对象的documentPath拼接作为类型名称 -->
        {{#if typeRefObject.isDefinition}}*{{/if}}{{~upperFirst (joinString '_' typeRefObject.documentPath)~}}
    {{~else~}}
        {{~#if typeRefEnum~}}
            <!-- 使用关联枚举的name并大写首字母作为类型名称 -->
            {{~upperFirst typeRefEnum.name~}}
        {{~else~}}
            <!-- 普通类型做类型转换 -->
            {{~#equal typeName 'string'}}string{{/equal~}}
            {{~#equal typeName 'integer'}}int64{{/equal~}}
            {{~#equal typeName 'number'}}float64{{/equal~}}
            {{~#equal typeName 'boolean'}}bool{{/equal~}}
            {{~#equal typeName 'json'}}any{{/equal~}}
        {{~/if}}
    <!-- 到处使用字段名导出json，并判断字段是否必须 -->
    {{~/if}} `json:"{{name}}{{#unless required}},omitempty{{/unless}}"`
    {{/each}}
}
{{/each}}

<!-- typescript生成对象使用示例 -->
{{#each objectFieldArray}}
export interface {{joinString '_' documentPath}} {
    {{#each fields}}
    {{name}}{{#unless required}}?{{/unless}}:
    {{~#if typeRefObject~}}
        {{~joinString '_' typeRefObject.documentPath~}}
    {{~else~}}
        {{~#if typeRefEnum~}}
            {{~typeRefEnum.name~}}
        {{~else~}}
            {{~#equal typeName 'string'}}string{{/equal~}}
            {{~#equal typeName 'integer'}}number{{/equal~}}
            {{~#equal typeName 'number'}}number{{/equal~}}
            {{~#equal typeName 'boolean'}}boolean{{/equal~}}
            {{~#equal typeName 'json'}}any{{/equal~}}
        {{~/if}}
    {{~/if}}{{#if isArray}}[]{{~/if~}},
    {{/each}}
}
{{/each}}
```

3. 定义枚举类型

<pre class="language-handlebars"><code class="lang-handlebars"><strong>&#x3C;!-- 示例为生成go的枚举【go没有枚举类型，使用别名实现】-->
</strong>{{#each enumFieldArray}}
&#x3C;!-- 枚举名首字母大写 -->
type {{upperFirst name}} string
const (
    {{#each values}}
    &#x3C;!-- 遍历枚举值列表，使用枚举名作为前缀 -->
    {{upperFirst name}}_{{this}} {{upperFirst name}} = "{{this}}"
    {{/each}}
)
{{/each}}

&#x3C;!-- 示例为生成typescript的枚举-->
{{#each enumFieldArray}}
enum {{name}} {
    {{#each values}}
    '{{this}}',
    {{/each}}
}
{{/each}}
</code></pre>

4. 单对象模版生成多个文件（以java为例）

<pre class="language-handlebars"><code class="lang-handlebars">&#x3C;!-- ${objectFieldArray}.java.hbs (对象文件名必须以${objectFieldArray}开头) -->
<strong>package com.fireboom.entity.{{root}};
</strong>
import lombok.Data;

&#x3C;!-- &#x3C;#fileName#>标签用来标识文件名【支持路径】用来替换${objectFieldArray} -->
// &#x3C;#fileName#>{{root}}/{{upperFirst (joinString '_' documentPath)}}&#x3C;#fileName#>
@Data
public class {{upperFirst (joinString '_' documentPath)}} {
    public {{upperFirst (joinString '_' documentPath)}}(
    {{~#each fields}}
        {{~> field_type_java this}} {{name}}{{#unless @last}}, {{/unless~}}
    {{/each~}}
    ) {
        {{#each fields}}
        this.{{name}} = {{name}};
        {{/each}}
    }
    {{#each fields}}
    private {{> field_type_java this}} {{name}};
    {{/each}}
}

&#x3C;!-- field_type_java.hbs片段函数用来定义字段类型【多处使用，提取片段】 -->
{{#if isArray}}java.util.List&#x3C;{{~/if~}}
{{~#if typeRefObject~}}
    com.fireboom.entity.{{root}}.{{~upperFirst (joinString '_' typeRefObject.documentPath)~}}
    {{~else~}}
    {{~#if typeRefEnum~}}
        com.fireboom.entity.enums.{{~upperFirst typeRefEnum.name~}}
    {{~else~}}
        {{~#equal typeName 'string'}}String{{/equal~}}
        {{~#equal typeName 'integer'}}Integer{{/equal~}}
        {{~#equal typeName 'number'}}Double{{/equal~}}
        {{~#equal typeName 'boolean'}}Boolean{{/equal~}}
        {{~#equal typeName 'json'}}com.alibaba.fastjson.JSON{{/equal~}}
    {{~/if~}}
{{~/if~}}
{{#if isArray}}>{{~/if~}}
</code></pre>

5. 单枚举模板生成多个文件（以java为例）

```handlebars
<!-- ${enumFieldArray}.java.hbs (枚举文件名必须以${enumFieldArray}开头) -->
package com.fireboom.entity.enums;

import lombok.Getter;

<!-- <#fileName#>标签用来标识文件名【支持路径】用来替换${enumFieldArray} -->
// <#fileName#>enums/{{upperFirst name}}<#fileName#>
public enum {{upperFirst name}} {
    {{#each values}}
    {{this}}("{{this}}"){{#if @last}};{{else}},{{/if}}
    {{/each}}

    @Getter
    private final String value;

    {{upperFirst name}}(String value) {
        this.value = value;
    }
}
```
