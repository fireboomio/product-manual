# 模板规范

以下结构体定义为sdk模版内的变量

```go
type SDKTpl struct {
  BaseURL         string // 现在是写死的localhost:9991
  SdkVersion      string // 版本，现在也是写死的
  ApplicationHash string // 唯一哈希
  ReactNative     bool
  Webhooks        []string `json:"webhooks"`
  Roles           []string `json:"roles"` // 角色列表
  Fields          []Fields // linkerBuilder
  Types           []Types  // linkerBuilder

  Operations   []*JsonSchemaOperations // operationsSchema转换(老版使用)
  OnceMap      map[string]any
  MaxLengthMap map[string]*MaxLength

  AuthProviders    []*wgpb.AuthProvider          // 认证配置
  DataSources      []types.DataSource            // 数据源详情
  S3Providers      []*wgpb.S3UploadConfiguration // S3上传配置
  HooksConfig      *types.HooksConfiguration     // 钩子配置
  ProxyDirectories []string                      // 代理钩子目录
  EnumFieldArray   []*enumField                  // 枚举类型定义(新版使用)
  ObjectFieldArray []*objectField                // 对象类型定义(新版使用)
}

// 对象类型(新)
type objectField struct {
  Name          string         // 对象/字段名
  TypeName      string         // 类型名(为字段时使用)
  TypeRef       string         // 忽略
  TypeRefObject *objectField   // 类型引用(为字段时使用)
  TypeRefEnum   *enumField     // 枚举引用(为字段时使用)
  Required      bool           // 是否必须(为字段时使用)
  IsArray       bool           // 是否数组(为字段时使用)
  IsDefinition  bool           // 是否全局定义
  DocumentPath  []string       // 文档路径(建议拼接后用来做对象名)
  Fields        []*objectField // 字段列表(为对象时使用)
  Root          string         // 归属顶层结构体
  OperationInfo *operationInfo // operation信息
}

// 枚举类型(新)
type enumField struct {
  Name   string   // 枚举名称
  Values []string // 枚举值列表
}

// operation信息(新)
type operationInfo struct {
  Name           string
  Path           string
  IsInternal     bool
  IsQuery        bool
  IsLiveQuery    bool
  IsMutation     bool
  IsSubscription bool
}

type JsonSchemaOperations struct {
  Name                   string                 // operation名称
  Path                   string                 // operation路径
  Copy                   map[string]interface{} // 用来作筛选
  InputSchema            types.DataSchema
  InputSchemaString      string
  InjectedSchema         types.DataSchema
  InjectedSchemaString   string
  InternalSchema         types.DataSchema
  InternalSchemaString   string
  ResponseSchema         types.DataSchema
  ResponseSchemaString   string
  Engine                 wgpb.OperationExecutionEngine `json:"engine"`
  IsInternal             bool                          `json:"isInternal"`             // 是否内部
  IsQuery                bool                          `json:"isQuery"`                // 是否查询
  IsLiveQuery            bool                          `json:"isLiveQuery"`            // 是否实时查询
  IsMutation             bool                          `json:"isMutation"`             // 是否变更
  IsSubscription         bool                          `json:"isSubscription"`         // 是否订阅
  HasInput               bool                          `json:"hasInput"`               // 是否有输入
  HasInjectedInput       bool                          `json:"hasInjectedInput"`       // 是否有injected输入
  HasInternalInput       bool                          `json:"hasInternalInput"`       // 是否有internal输入
  RequiresAuthentication bool                          `json:"requiresAuthentication"` //是否需要认证
}

type Fields struct {
  Name string   `json:"name"`
  Args []string `json:"args"`
}

type Types struct {
  Name   string   `json:"name"`
  Fields []string `json:"fields"`
}

type HooksConfiguration struct {
  Global         *HooksGlobalConf
  Authentication wgpb.ApiAuthenticationHooks //不使用指针，需要初始化默认值，不然 uuid 没法使用
  Queries   map[string]*wgpb.OperationHooksConfiguration
  Mutations map[string]*wgpb.OperationHooksConfiguration
}

type DataSource struct {
  ApiId                string                      `json:"api_id"`
  ApiNameSpace         string                      `json:"api_name_space"`
  Kind                 wgpb.DataSourceKind         `json:"kind"`
  DataBaseURL          *wgpb.ConfigurationVariable `json:"data_base_url"`
  URL                  *wgpb.ConfigurationVariable `json:"url"`
  SkipRenameRootFields []string                    `json:"skip_rename_root_fields"`
  Source               Source                      `json:"source"`        // openApi json 文件
  BaseURL              string                      `json:"base_url"`      // openApi baseUrl
  IsFederation         bool                        `json:"is_federation"` // graphql 配置，默认 false,暂时没有理解, 貌似是看数据源是不是独立的子图
  Header               HeadValues                  `json:"header"`
  IsCustomized         bool                        `json:"is_customized"`
}

type AuthProvider struct {
  Id           string                           `json:"id"`
  Kind         AuthProviderKind                 `json:"kind"`
  GithubConfig *GithubAuthProviderConfig        `json:"githubConfig"`
  OidcConfig   *OpenIDConnectAuthProviderConfig `json:"oidcConfig"`
}

type S3UploadConfiguration struct {
  Name            string                      `json:"name"`
  Endpoint        *ConfigurationVariable      `json:"endpoint"`
  AccessKeyID     *ConfigurationVariable      `json:"accessKeyID"`
  SecretAccessKey *ConfigurationVariable      `json:"secretAccessKey"`
  BucketName      *ConfigurationVariable      `json:"bucketName"`
  BucketLocation  *ConfigurationVariable      `json:"bucketLocation"`
  UseSSL          bool                        `json:"useSSL"`
  UploadProfiles  map[string]*S3UploadProfile `json:"uploadProfiles" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}
```

�

1. 全局Helpers注册

```
// 注册函数（判断字符串是否在切片中）
handlebars.RegisterHelper("stringInArray", func(str string, strArr []string) bool {
    return utils.StringInArray(str, strArr)
})
1. 判断字符串是否在切片中
 
// 使用姿势
{{#if (stringInArray 'age' slice)}}123{{/if}}


// 注册函数（首字母小写）
handlebars.RegisterHelper("lowerFirst", func(str string) string {
    return strings.ToLower(str[:1]) + str[1:]
})
// 使用姿势
{{lowerFirst "Age"}}


// 注册函数（字符串连接）
handlebars.RegisterHelper("joinString", func(sep string, strArr []string) string {
    return strings.Join(strArr, sep)
})
// 使用姿势
{{joinString ',' slice}}

// 注册函数（替换特殊字符为指定字符）
handlebars.RegisterHelper("replaceSpecial", func(str, sep string) string {
    if strings.HasPrefix(str, "/") {
        str = str[1:]
    }
    reg := regexp.MustCompile("[^A-Za-z0-9_]+")
    return reg.ReplaceAllString(str, sep)
})
// 使用姿势
{{replaceSpecial operationName "$"}}


// 注册函数（判断字符串是否想等，target为","分隔多个）
handlebars.RegisterHelper("equalAny", func(source string, target string) bool {
    return utils.StringInArray(source, strings.Split(target, ","))
})
// 使用姿势
{{#if (equalAny hookType "Queries,Mutations")}}


// 注册函数（判断任意对象是否为空）
handlebars.RegisterHelper("isNotEmpty", func(val any) bool {
    return utils.IsNotEmpty(val)
})
// 使用姿势
{{#if (isNotEmpty operations)}}


// 注册函数（获取schemaType的真实类型）(老版使用)
handlebars.RegisterHelper("realType", func(val any) string {
    schemaType := ""
    switch ret := val.(type) {
        case string:
        schemaType = ret
        case []string:
        schemaType = ret[0]
        case []interface{}:
        schemaType = fmt.Sprint(ret[0])
    }
    return schemaType
})


// 注册函数（根据key和value过滤operations，key为","分隔多个）(老版使用)
handlebars.RegisterHelper("filterOperations", func(operations []*JsonSchemaOperations, key string, val any) []*JsonSchemaOperations {
    var result []*JsonSchemaOperations
    keyArr := strings.Split(key, ",")
    for _, op := range operations {
        ok := true
        for _, k := range keyArr {
            reverse := false
            if k[:1] == "!" {
                reverse = true
                k = k[1:]
            }
            v, err := utils.GetNestedValue(op.Copy, k)
            if err != nil || !reverse && !reflect.DeepEqual(v, val) {
                ok = false
                break
            }
        }
        if ok {
        result = append(result, op)
    }
  }
    return result
})
  // 使用姿势
  {{#each (filterOperations operations 'isQuery,!isInternal' true)}}
```

2. 片段函数注册及调用（扫描目录partials下的文件，文件名作为helper函数，文件内容为函数执行结果片段）

```
// operation_partial.hbs内容
{{#each operations~}}
    {{~name}}Response
    {{~#if hasInput~}}
        ,{{name}}Input
    {{~/if~}}
    {{~#if hasInternalInput~}}
        {{~#if includeInternal~}}
            ,Internal{{name}}Input
        {{~/if~}}
    {{~/if~}}
    {{~#if hasInjectedInput~}}
        {{~#if includeInject~}}
            ,Injected{{name}}Input
        {{~/if~}}
    {{~/if~}}
    {{~#if includeResponseData~}}
        ,{{name}}ResponseData
    {{~/if~}}
    {{#unless @last}},{{/unless}}
{{~/each}}

// operation_partial调用，指定参数operations，includeInternal，includeInject，includeResponseData
import { {{> operation_partial operations=operations includeInternal=true includeInject=true includeResponseData=false}} } from "./models"
```

3. 使用特殊姿势

* 子表达式

```
// isNotEmpty作为子表达式中执行的函数，以inputSchema为参数，执行结果会返回给内置的if函数处理
{{#if (isNotEmpty inputSchema)~}}{{/if}}
```

* 遍历map

```
// 遍历definitions，制定key为name，value为schema
{{~#each inputSchema.definitions as |schema name|}}
	"{{name}}": {{schema}}
{{/each}}
```

* 遍历slice

```
// 遍历operations，遍历operations，name和path为每个子元素的属性
// @first @last分别判断是否为第一个和最后一个元素
{{#each operations~}}
  {{#if @first}}{{{/if}}
	"{{name}}": {{path}}
  {{#if @last}}}{{/if}}
{{/each}}
```

* with临时变量

```
// 子表达式筛选出isQuery为true的operations
// 使用with blocker helper进行操作，this为筛选后的对象
// 适用场景为需要在operations为空的情况下设置默认值
{{#with (filterOperations operations 'isQuery' true)}}
	{{#each this}}{{name}}{{/each}}
{{/with}}
```

* 变长参数

```
// fork了handlebars的代码并支持了变长参数的传递
// 现在可以注册使用变长参数的helper
handlebars.RegisterHelper("equalAny", func(source string, target ...string) bool {
    return utils.StringInArray(source, target)
})
```
