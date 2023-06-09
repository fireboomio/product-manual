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
<!-- 使用documentPath拼接'_'对象 -->
type {{joinString '_' documentPath}} struct {
    {{#each fields}}
    {{upperFirst name}} {{#if isArray}}[]{{~/if~}}
    {{~#if typeRefObject~}}
        {{#if typeRefObject.isDefinition}}*{{/if}}{{~joinString '_' typeRefObject.documentPath~}}
    {{~else~}}
        {{~#if typeRefEnum~}}
            {{~upperFirst typeRefEnum.name~}}
        {{~else~}}
            {{~#equal typeName 'string'}}string{{/equal~}}
            {{~#equal typeName 'integer'}}int64{{/equal~}}
            {{~#equal typeName 'number'}}float64{{/equal~}}
            {{~#equal typeName 'boolean'}}bool{{/equal~}}
            {{~#equal typeName 'json'}}any{{/equal~}}
        {{~/if}}
    {{~/if}} `json:"{{name}}{{#unless required}},omitempty{{/unless}}"`
    {{/each}}
}
{{/each}}
```

