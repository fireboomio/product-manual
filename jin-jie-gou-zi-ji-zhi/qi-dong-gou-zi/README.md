# 启动钩子

任何语言实现的Fireboom钩子，本质上都是一个WEB服务，其遵循Fireboom规范注册对应路由。

任意语言的钩子服务启动时，都遵循如下流程。

<figure><img src="../../.gitbook/assets/image (1) (1) (1) (1) (1) (1) (1) (1).png" alt=""><figcaption></figcaption></figure>

{% hint style="info" %}
Fireboom 同时只兼容一种语言的钩子！！！
{% endhint %}

## 读取配置文件

钩子服务需要依赖Fireboom服务的某些配置，因此需要读取Fireboom的配置文件：`exported/generated/fireboom.config.json` 。

为了便于读取，且减少冗余。Fireboom为`fireboom.config.json`创建了一个软连接，并生成到开启钩子的指定路径，例如：`custom-go/generated/fireboom.config.json` 。

其中，包含钩子启动所依赖的大部分信息，如

* 钩子监听端口：`serverOptions.listen.port`
* S3配置信息：`s3UploadConfiguration`
* ...

读取该文件的代码如下：

{% tabs %}
{% tab title="golang" %}
{% code title="pkg/types/configure.go" %}
```go
var configJsonPath = filepath.Join("generated", "fireboom.config.json")

func init() {
	_ = utils.ReadStructAndCacheFile(configJsonPath, &WdgGraphConfig)
}
```
{% endcode %}
{% endtab %}

{% tab title="nodejs" %}

{% endtab %}
{% endtabs %}

## 读取环境变量

钩子服务还依赖Fireboom服务的环境变量，使用相对路径读取： `../.env`。

{% tabs %}
{% tab title="golang" %}
{% code title="server/fireboom_server.go" %}
```go
const nodeEnvFilepath = "../.env"

func init() {
    _ = godotenv.Overload(nodeEnvFilepath)
```
{% endcode %}
{% endtab %}

{% tab title="nodejs" %}

{% endtab %}
{% endtabs %}

## 注册中间件

1，解析Fireboom调用钩子时携带的全局参数 `_wg`

```json
"__wg": { # 全局参数
    "clientRequest": { # 原始客户端请求，即请求9991端口的request对象
      "method": "GET",
      "requestURI": "/operations/Weather?code=beijing",
      "headers": {
        "Accept": "application/json",
        "Content-Type": "application/json"
      }
    },
    "user": { # （可选）授权用户的信息
      "userID": "1",
      "roles": ["user"]
    }
  }
```

2，为上下文ctx注入`User`对象，用于获取登录用户的信息

3，为上下文ctx注入`InternalClient`对象（用于[内部调用](../nei-bu-tiao-yong.md)）

{% tabs %}
{% tab title="golang" %}
{% code title="server/start.go" fullWidth="true" %}
```go
e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == http.MethodGet {
			return next(c)
		}
		// 1，解析Fireboom调用时携带的全局参数 _wg
		var body base.BaseRequestBody
		err := utils.CopyAndBindRequestBody(c.Request(), &body)
		if err != nil {
			return err
		}
	
		if body.Wg.ClientRequest == nil {
			body.Wg.ClientRequest = &base.ClientRequest{
				Method:     c.Request().Method,
				RequestURI: c.Request().RequestURI,
				Headers:    map[string]string{},
			}
		} else {
			for name, value := range body.Wg.ClientRequest.Headers {
				c.Request().Header.Set(name, value)
			}
		}
		reqId := c.Request().Header.Get("x-request-id")
		internalClient := base.InternalClientFactoryCall(map[string]string{"x-request-id": reqId}, body.Wg.ClientRequest, body.Wg.User)
		internalClient.Queries = internalQueries
		internalClient.Mutations = internalMutations
		brc := &base.BaseRequestContext{
			Context:        c,
		//2，为上下文ctx注入User对象，用于获取登录用户的信息
			User:           body.Wg.User,
		// 3,为上下文ctx注入InternalClient对象
			InternalClient: internalClient,
		}
		return next(brc)
	}
})
```
{% endcode %}
{% endtab %}

{% tab title="nodejs" %}

{% endtab %}
{% endtabs %}

## 注册钩子

Fireboom中有各种类型的钩子，主要包括：

[operation-gou-zi.md](../operation-gou-zi.md "mention")

[shen-fen-yan-zheng-gou-zi.md](../shen-fen-yan-zheng-gou-zi.md "mention")

[graphql-gou-zi.md](../graphql-gou-zi.md "mention")

[wen-jian-shang-chuan-gou-zi.md](../wen-jian-shang-chuan-gou-zi.md "mention")

每个钩子开启后，都会生成对应模板，并按照规范注册路由。

### 开启钩子

开启钩子很简单，有两种方法：

* 主功能区-> <mark style="color:orange;">概览面板</mark> ，比较直观的展示了钩子所处的注入点，见上图
* 主功能区-> <mark style="color:orange;">钩子面板</mark> ，简单罗列了所有的钩子

**所有钩子都有约定的目录结构，**各类型钩子对应目录：

```
├─ custom-*
│  ├─ auth  # 授权钩子目录
│  │  ├─ mutatingPostAuthentication.*
│  │  ├─ postAuthentication.*
│  │  ├─ postLogout.*
│  │  └─ revalidate.*
│  ├─ customize # graphql钩子目录
│  │  ├─ gql1.*
│  ├─ global  # OPERATION全局钩子目录
│  │  ├─ beforeRequest.*
│  │  ├─ onRequest.*
│  │  └─ onResponse.*
│  ├─ hooks # OPERATION 局部钩子目录
│  │  └─ Simple
│  │     ├─ customResolve.*
│  │     ├─ mockResolve.*
│  │     ├─ mutatingPostResolve.*
│  │     ├─ mutatingPreResolve.*
│  │     ├─ postResolve.*
│  │     └─ preResolve.*
│  ├─ functions # 函数钩子目录
│  │  └─ fun1.*
│  ├─ proxys # 代理钩子目录
│  │  └─ p1.*
│  └─ uploads # 上传文件钩子目录
│     └─ tengxunyun
│        └─ avatar
│           ├─ postUpload.*
│           └─ preUpload.*
```

以golang钩子为例， 上图中的Simple OPERATION的 `postResolve` 钩子，对应 `custom-go/hooks/Simple/postResolve.go` 文件。

开启钩子时，若对应钩子的文件不存在，则会从github仓库对应文件中获取钩子模板，并在对应目录创建默认钩子文件。

所有类型钩子的默认模板，都存储在如下仓库：

* golang：[https://github.com/fireboomio/files/blob/main/hook.templates.go.json](https://github.com/fireboomio/files/blob/main/hook.templates.go.json)
* nodejs：[https://github.com/fireboomio/files/blob/main/hook.templates.ts.json](https://github.com/fireboomio/files/blob/main/hook.templates.ts.json)

其中golang钩子模板仓库如下：

```json
{
  "global": {
    "beforeRequest": "package global\n\nimport (\n\t\"custom-go/pkg/base\"\n\t\"custom-go/pkg/plugins\"\n)\n\nfunc BeforeOriginRequest(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientRequest, error) {\n\treturn body.Request, nil\n}\n",
    "onRequest": "package global\n\nimport (\n\t\"custom-go/pkg/base\"\n\t\"custom-go/pkg/plugins\"\n)\n\nfunc OnOriginRequest(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientRequest, error) {\n\treturn body.Request, nil\n}\n",
    "onResponse": "package global\n\nimport (\n\t\"custom-go/pkg/base\"\n\t\"custom-go/pkg/plugins\"\n)\n\nfunc OnOriginResponse(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientResponse, error) {\n\treturn body.Response, nil\n}\n",
    "onConnectionInit": "import type { WsTransportHookRequest } from 'generated/fireboom.hooks'\nimport type { WsTransportOnConnectionInitResponse } from 'fireboom-wundersdk/server'\n\nexport default async function onConnectionInit(hook: WsTransportHookRequest): Promise<WsTransportOnConnectionInitResponse> {\n  return {\n    payload: {\n      // your code here\n    }\n  }\n}"
  },
  "auth": {
    "postAuthentication": "package auth\n\nimport \"custom-go/pkg/base\"\n\nfunc PostAuthentication(hook *base.AuthenticationHookRequest) error {\n\treturn nil\n}",
    "mutatingPostAuthentication": "package auth\n\nimport (\n\t\"custom-go/pkg/base\"\n\t\"custom-go/pkg/plugins\"\n)\n\nfunc MutatingPostAuthentication(hook *base.AuthenticationHookRequest) (*plugins.AuthenticationResponse, error) {\n\treturn &plugins.AuthenticationResponse{User: hook.User, Status: \"ok\"}, nil\n}\n",
    "revalidateAuthentication": "package auth\n\nimport (\n\t\"custom-go/pkg/base\"\n\t\"custom-go/pkg/plugins\"\n)\n\nfunc Revalidate(hook *base.AuthenticationHookRequest) (*plugins.AuthenticationResponse, error) {\n\treturn nil, nil\n}",
    "postLogout": "package auth\n\nimport \"custom-go/pkg/base\"\n\nfunc PostLogout(hook *base.AuthenticationHookRequest) error {\n\treturn nil\n}"
  },
  "hook": {  // operation局部钩子
  // 带入参的OPERATION
    "WithInput": {
      "mockResolve": "package $HOOK_PACKAGE$\n\nimport (\n\t\"custom-go/generated\"\n\t\"custom-go/pkg/base\"\n)\n\nfunc MockResolve(hook *base.HookRequest, body generated.$HOOK_NAME$Body) (res generated.$HOOK_NAME$Body, err error) {\n\thook.Logger().Info(\"MockResolve\")\n\treturn body, nil\n}\n",
      "preResolve": "package $HOOK_PACKAGE$\n\nimport (\n\t\"custom-go/generated\"\n\t\"custom-go/pkg/base\"\n)\n\nfunc PreResolve(hook *base.HookRequest, body generated.$HOOK_NAME$Body) (res generated.$HOOK_NAME$Body, err error) {\n\thook.Logger().Info(\"PreResolve\")\n\treturn body, nil\n}\n",
      "mutatingPreResolve": "package $HOOK_PACKAGE$\n\nimport (\n\t\"custom-go/generated\"\n\t\"custom-go/pkg/base\"\n)\n\nfunc MutatingPreResolve(hook *base.HookRequest, body generated.$HOOK_NAME$Body) (res generated.$HOOK_NAME$Body, err error) {\n\thook.Logger().Info(\"MutatingPreResolve\")\n\treturn body, nil\n}\n",
      "postResolve": "package $HOOK_PACKAGE$\n\nimport (\n\t\"custom-go/generated\"\n\t\"custom-go/pkg/base\"\n)\n\nfunc PostResolve(hook *base.HookRequest, body generated.$HOOK_NAME$Body) (res generated.$HOOK_NAME$Body, err error) {\n\thook.Logger().Info(\"PostResolve\")\n\treturn body, nil\n}\n",
      "customResolve": "package $HOOK_PACKAGE$\n\nimport (\n\t\"custom-go/generated\"\n\t\"custom-go/pkg/base\"\n)\n\nfunc CustomResolve(hook *base.HookRequest, body generated.$HOOK_NAME$Body) (res generated.$HOOK_NAME$Body, err error) {\n\thook.Logger().Info(\"CustomResolve\")\n\treturn body, nil\n}\n",
      "mutatingPostResolve": "package $HOOK_PACKAGE$\n\nimport (\n\t\"custom-go/generated\"\n\t\"custom-go/pkg/base\"\n)\n\nfunc MutatingPostResolve(hook *base.HookRequest, body generated.$HOOK_NAME$Body) (res generated.$HOOK_NAME$Body, err error) {\n\thook.Logger().Info(\"MutatingPostResolve\")\n\treturn body, nil\n}\n"
    },
  // 不带入参的OPERATION
    "WithoutInput": {
      "mockResolve": "package $HOOK_PACKAGE$\n\nimport (\n\t\"custom-go/generated\"\n\t\"custom-go/pkg/base\"\n)\n\nfunc MockResolve(hook *base.HookRequest, body generated.$HOOK_NAME$Body) (res generated.$HOOK_NAME$Body, err error) {\n\thook.Logger().Info(\"MockResolve\")\n\treturn body, nil\n}\n",
      "preResolve": "package $HOOK_PACKAGE$\n\nimport (\n\t\"custom-go/generated\"\n\t\"custom-go/pkg/base\"\n)\n\nfunc PreResolve(hook *base.HookRequest, body generated.$HOOK_NAME$Body) (res generated.$HOOK_NAME$Body, err error) {\n\thook.Logger().Info(\"PreResolve\")\n\treturn body, nil\n}\n",
      "mutatingPreResolve": "package $HOOK_PACKAGE$\n\nimport (\n\t\"custom-go/generated\"\n\t\"custom-go/pkg/base\"\n)\n\nfunc MutatingPreResolve(hook *base.HookRequest, body generated.$HOOK_NAME$Body) (res generated.$HOOK_NAME$Body, err error) {\n\thook.Logger().Info(\"MutatingPreResolve\")\n\treturn body, nil\n}\n",
      "postResolve": "package $HOOK_PACKAGE$\n\nimport (\n\t\"custom-go/generated\"\n\t\"custom-go/pkg/base\"\n)\n\nfunc PostResolve(hook *base.HookRequest, body generated.$HOOK_NAME$Body) (res generated.$HOOK_NAME$Body, err error) {\n\thook.Logger().Info(\"PostResolve\")\n\treturn body, nil\n}\n",
      "customResolve": "package $HOOK_PACKAGE$\n\nimport (\n\t\"custom-go/generated\"\n\t\"custom-go/pkg/base\"\n)\n\nfunc CustomResolve(hook *base.HookRequest, body generated.$HOOK_NAME$Body) (res generated.$HOOK_NAME$Body, err error) {\n\thook.Logger().Info(\"CustomResolve\")\n\treturn body, nil\n}\n",
      "mutatingPostResolve": "package $HOOK_PACKAGE$\n\nimport (\n\t\"custom-go/generated\"\n\t\"custom-go/pkg/base\"\n)\n\nfunc MutatingPostResolve(hook *base.HookRequest, body generated.$HOOK_NAME$Body) (res generated.$HOOK_NAME$Body, err error) {\n\thook.Logger().Info(\"MutatingPostResolve\")\n\treturn body, nil\n}\n"
    }
  },
  // graphql 钩子
  "custom": "package customize\n\nimport (\n\t\"custom-go/pkg/plugins\"\n\t\"github.com/graphql-go/graphql\"\n)\n\ntype Person struct {\n\tId        int    `json:\"id\"`\n\tFirstName string `json:\"firstName\"`\n\tLastName  string `json:\"lastName\"`\n}\n\nvar (\n\tpersonType = graphql.NewObject(graphql.ObjectConfig{\n\t\tName:        \"Person\",\n\t\tDescription: \"A person in the system\",\n\t\tFields: graphql.Fields{\n\t\t\t\"id\": &graphql.Field{\n\t\t\t\tType: graphql.Int,\n\t\t\t},\n\t\t\t\"firstName\": &graphql.Field{\n\t\t\t\tType: graphql.String,\n\t\t\t},\n\t\t\t\"lastName\": &graphql.Field{\n\t\t\t\tType: graphql.String,\n\t\t\t},\n\t\t},\n\t})\n\n\tfields = graphql.Fields{\n\t\t\"person\": &graphql.Field{\n\t\t\tType:        personType,\n\t\t\tDescription: \"Get Person By ID\",\n\t\t\tArgs: graphql.FieldConfigArgument{\n\t\t\t\t\"id\": &graphql.ArgumentConfig{\n\t\t\t\t\tType: graphql.Int,\n\t\t\t\t},\n\t\t\t},\n\t\t\tResolve: func(params graphql.ResolveParams) (interface{}, error) {\n\t\t\t\t_ = plugins.GetGraphqlContext(params)\n\t\t\t\tid, ok := params.Args[\"id\"].(int)\n\t\t\t\tif ok {\n\t\t\t\t\ttestPeopleData := []Person{\n\t\t\t\t\t\t{Id: 1, FirstName: \"John\", LastName: \"Doe\"},\n\t\t\t\t\t\t{Id: 2, FirstName: \"Jane\", LastName: \"Doe\"},\n\t\t\t\t\t}\n\t\t\t\t\tfor _, p := range testPeopleData {\n\t\t\t\t\t\tif p.Id == id {\n\t\t\t\t\t\t\treturn p, nil\n\t\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t\treturn nil, nil\n\t\t\t},\n\t\t},\n\t}\n\n\trootQuery = graphql.ObjectConfig{Name: \"RootQuery\", Fields: fields}\n\n\t^$CUSTOMIZE_NAME$_schema, _ = graphql.NewSchema(graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)})\n)\n",
  "example": {
    "custom": [
      {
        "name": "示例数据源",
        "code": "package customize\n\nimport (\n\t\"custom-go/pkg/plugins\"\n\t\"github.com/graphql-go/graphql\"\n)\n\ntype Person struct {\n\tId        int    `json:\"id\"`\n\tFirstName string `json:\"firstName\"`\n\tLastName  string `json:\"lastName\"`\n}\n\nvar (\n\tpersonType = graphql.NewObject(graphql.ObjectConfig{\n\t\tName:        \"Person\",\n\t\tDescription: \"A person in the system\",\n\t\tFields: graphql.Fields{\n\t\t\t\"id\": &graphql.Field{\n\t\t\t\tType: graphql.Int,\n\t\t\t},\n\t\t\t\"firstName\": &graphql.Field{\n\t\t\t\tType: graphql.String,\n\t\t\t},\n\t\t\t\"lastName\": &graphql.Field{\n\t\t\t\tType: graphql.String,\n\t\t\t},\n\t\t},\n\t})\n\n\tfields = graphql.Fields{\n\t\t\"person\": &graphql.Field{\n\t\t\tType:        personType,\n\t\t\tDescription: \"Get Person By ID\",\n\t\t\tArgs: graphql.FieldConfigArgument{\n\t\t\t\t\"id\": &graphql.ArgumentConfig{\n\t\t\t\t\tType: graphql.Int,\n\t\t\t\t},\n\t\t\t},\n\t\t\tResolve: func(params graphql.ResolveParams) (interface{}, error) {\n\t\t\t\t_ = plugins.GetGraphqlContext(params)\n\t\t\t\tid, ok := params.Args[\"id\"].(int)\n\t\t\t\tif ok {\n\t\t\t\t\ttestPeopleData := []Person{\n\t\t\t\t\t\t{Id: 1, FirstName: \"John\", LastName: \"Doe\"},\n\t\t\t\t\t\t{Id: 2, FirstName: \"Jane\", LastName: \"Doe\"},\n\t\t\t\t\t}\n\t\t\t\t\tfor _, p := range testPeopleData {\n\t\t\t\t\t\tif p.Id == id {\n\t\t\t\t\t\t\treturn p, nil\n\t\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t\treturn nil, nil\n\t\t\t},\n\t\t},\n\t}\n\n\trootQuery = graphql.ObjectConfig{Name: \"RootQuery\", Fields: fields}\n\n\tChatGPT_schema, _ = graphql.NewSchema(graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)})\n)\n"
      }
    ]
  },
  // 上传钩子
  "upload": {
    "preUpload": "package $PROFILE_NAME$\n\nimport (\n\t\"custom-go/generated\"\n\t\"custom-go/pkg/base\"\n\t\"custom-go/pkg/plugins\"\n)\n\nfunc PreUpload(request *base.UploadHookRequest, body *plugins.UploadBody[generated.^$STORAGE_NAME$_$PROFILE_NAME$ProfileMeta]) (*base.UploadHookResponse, error) {\n\treturn &base.UploadHookResponse{FileKey: body.File.Name}, nil\n}\n",
    "postUpload": "package $PROFILE_NAME$\n\nimport (\n\t\"custom-go/generated\"\n\t\"custom-go/pkg/base\"\n\t\"custom-go/pkg/plugins\"\n)\n\nfunc PostUpload(request *base.UploadHookRequest, body *plugins.UploadBody[generated.^$STORAGE_NAME$_$PROFILE_NAME$ProfileMeta]) (*base.UploadHookResponse, error) {\n\treturn nil, nil\n}\n"
  }
}
```

### 引入钩子

不同语言引入钩子文件的方式不同，但都会动态生成一个<mark style="color:purple;">入口文件</mark>，引入所有的钩子文件。

{% tabs %}
{% tab title="golang" %}
{% code title="custom-go/server/fireboom_server.go" %}
```go
package server
import (
	"custom-go/global"
	"github.com/joho/godotenv"
	"custom-go/auth"
	"custom-go/generated"
	// 引入 OPERATION钩子
	hooks_Single "custom-go/hooks/Single" 
	hooks_Weather "custom-go/hooks/Weather"
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"custom-go/pkg/types"
	// 引入 文件上传钩子
	uploads_tengxunyun_avatar "custom-go/uploads/tengxunyun/avatar"
	uploads_fireboom_avatar "custom-go/uploads/fireboom/avatar"
	// 引入 graphql钩子
	"custom-go/customize"
	_ "custom-go/proxys"
)
const nodeEnvFilepath = "../.env"
func init() {
	_ = godotenv.Overload(nodeEnvFilepath)
	types.WdgHooksAndServerConfig = types.WunderGraphHooksAndServerConfig{
		Hooks: types.HooksConfiguration{
			Global: plugins.GlobalConfiguration{
				HttpTransport: plugins.HttpTransportHooks{
					BeforeOriginRequest: global.BeforeOriginRequest,
					OnOriginRequest:     global.OnOriginRequest,
					OnOriginResponse:    global.OnOriginResponse,
				},
				WsTransport: plugins.WsTransportHooks{},
			},

			Authentication: plugins.AuthenticationConfiguration{
				MutatingPostAuthentication: auth.MutatingPostAuthentication,
			},

			Queries: base.OperationHooks{
				"Single": {
					PreResolve: plugins.ConvertBodyFunc[generated.SingleInternalInput, generated.SingleResponseData](hooks_Single.PreResolve),
				},
				"Weather": {
					CustomResolve: plugins.ConvertBodyFunc[generated.WeatherInternalInput, generated.WeatherResponseData](hooks_Weather.CustomResolve),
					PostResolve:   plugins.ConvertBodyFunc[generated.WeatherInternalInput, generated.WeatherResponseData](hooks_Weather.PostResolve),
				},
			},

			Mutations: base.OperationHooks{},

			Subscriptions: base.OperationHooks{},

			Uploads: map[string]plugins.UploadHooks{
				"tengxunyun": {
					"avatar": {
						PreUpload:  plugins.ConvertUploadFunc[generated.Tengxunyun_avatarProfileMeta](uploads_tengxunyun_avatar.PreUpload),
						PostUpload: plugins.ConvertUploadFunc[generated.Tengxunyun_avatarProfileMeta](uploads_tengxunyun_avatar.PostUpload),
					},
				},
				"ali": {},
				"fireboom": {
					"avatar": {
						PreUpload: plugins.ConvertUploadFunc[generated.Fireboom_avatarProfileMeta](uploads_fireboom_avatar.PreUpload),
					},
				},
			},
		},
		GraphqlServers: []plugins.GraphQLServerConfig{
			{
				ApiNamespace:          "Custom",
				ServerName:            "Custom",
				EnableGraphQLEndpoint: true,
				Schema:                customize.Custom_schema,
			},
		},
	}
}

```
{% endcode %}
{% endtab %}

{% tab title="nodejs" %}

{% endtab %}
{% endtabs %}

### 注册路由

引用<mark style="color:purple;">入口文件</mark>的变量，注册各种钩子的路由。

{% tabs %}
{% tab title="golang" %}
{% code title="server/start.go" %}
```go
// 注册proxy钩子
plugins.RegisterProxyHooks(e)
// 注册全局钩子
plugins.RegisterGlobalHooks(e, types.WdgHooksAndServerConfig.Hooks.Global)
// 注册授权钩子
plugins.RegisterAuthHooks(e, types.WdgHooksAndServerConfig.Hooks.Authentication)
// 注册上传钩子
plugins.RegisterUploadsHooks(e, types.WdgHooksAndServerConfig.Hooks.Uploads)

var internalQueries, internalMutations base.OperationDefinitions
nodeUrl := utils.GetConfigurationVal(types.WdgGraphConfig.Api.NodeOptions.NodeUrl)
queryOperations := filterOperationsHooks(types.WdgGraphConfig.Api.Operations, wgpb.OperationType_QUERY)
// 注册局部钩子
if queryLen := len(queryOperations); queryLen > 0 {
	internalQueries = plugins.BuildInternalRequest(e.Logger, nodeUrl, queryOperations)
	plugins.RegisterOperationsHooks(e, queryOperations, types.WdgHooksAndServerConfig.Hooks.Queries)
	e.Logger.Debugf(`Registered (%d) query operations`, queryLen)
}
mutationOperations := filterOperationsHooks(types.WdgGraphConfig.Api.Operations, wgpb.OperationType_MUTATION)
if mutationLen := len(mutationOperations); mutationLen > 0 {
	internalMutations = plugins.BuildInternalRequest(e.Logger, nodeUrl, mutationOperations)
	plugins.RegisterOperationsHooks(e, mutationOperations, types.WdgHooksAndServerConfig.Hooks.Mutations)
	e.Logger.Debugf(`Registered (%d) mutation operations`, mutationLen)
}
subscriptionOperations := filterOperationsHooks(types.WdgGraphConfig.Api.Operations, wgpb.OperationType_SUBSCRIPTION)
if subscriptionLen := len(subscriptionOperations); subscriptionLen > 0 {
	plugins.RegisterOperationsHooks(e, subscriptionOperations, types.WdgHooksAndServerConfig.Hooks.Subscriptions)
	e.Logger.Debugf(`Registered (%d) subscription operations`, subscriptionLen)
}
// 注册内部调用钩子
plugins.BuildDefaultInternalClient(internalQueries, internalMutations)
for _, registeredHook := range base.GetRegisteredHookArr() {
	go registeredHook(e.Logger)
}
// 注册graphql钩子
for _, gqlServer := range types.WdgHooksAndServerConfig.GraphqlServers {
	plugins.RegisterGraphql(e, gqlServer)
}
```
{% endcode %}
{% endtab %}

{% tab title="nodejs" %}

{% endtab %}
{% endtabs %}

## 启动钩子

### 启动服务

不同语言启动服务的方式不同，但一般都需要先安装依赖，再启动服务。

{% tabs %}
{% tab title="golang" %}
```bash
# 1.安装依赖
go mod tidy
# 2.启动服务
go run main.go
```
{% endtab %}

{% tab title="nodejs" %}
```bash
# 1.安装依赖
npm install
# 2.运行服务
npm run watch
```
{% endtab %}
{% endtabs %}

启动服务后，可通过健康检查接口，检测服务是否成功启动！

### 健康检查

健康检查接口，用于检查钩子服务健康状态，在界面上展示钩子是否已启动，见底部 <mark style="color:purple;">状态栏</mark>-> <mark style="color:purple;">钩子</mark> 状态。

{% swagger method="get" path="health" baseUrl="http://127.0.0.1:9992/" summary="健康检查接口" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-response status="200: OK" description="" %}
```json
{
    "status": "ok"
}
```
{% endswagger-response %}
{% endswagger %}

{% tabs %}
{% tab title="golang" %}
```go
// 健康检查
e.GET("/health", func(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
})
```
{% endtab %}

{% tab title="nodejs" %}

{% endtab %}
{% endtabs %}

