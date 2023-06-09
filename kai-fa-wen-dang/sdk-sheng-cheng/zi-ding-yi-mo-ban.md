# 自定义模板（）

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
    "fields": [${objectField}] // 对象字段列表
}

// enumField
{
    "name": "DictValueType", // 枚举名称
    "values": ["site"] // 枚举值列表
}
```

2. 定义对象类型（示例为生成go的strcut）

```handlebars
{{#each objectFieldArray}}
<!-- 使用documentPath拼接'_'对象名称 -->
type {{joinString '_' documentPath}} struct {
    <!-- 遍历字段列表 -->
    {{#each fields}}
    <!-- 字段名首字母大写，并判断类型是否为数组 -->
    {{upperFirst name}} {{#if isArray}}[]{{~/if~}}
    {{~#if typeRefObject~}}
        <!-- 使用关联对象的documentPath拼接作为类型名称 -->
        {{#if typeRefObject.isDefinition}}*{{/if}}{{~joinString '_' typeRefObject.documentPath~}}
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
```

3. 定义枚举类型（示例为生成go的枚举【go没有枚举类型，使用别名实现】）

```handlebars
{{#each enumFieldArray}}
<!-- 枚举名首字母大写 -->
type {{upperFirst name}} string
const (
    {{#each values}}
    <!-- 遍历枚举值列表，使用枚举名作为前缀 -->
    {{upperFirst name}}_{{this}} {{upperFirst name}} = "{{this}}"
    {{/each}}
)
{{/each}}
```
