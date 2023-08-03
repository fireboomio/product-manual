# 使用钩子

任何语言实现的Fireboom钩子，本质上都是一个WEB服务。但要遵循Fireboom规范注册对应路由。

任意语言的钩子服务启动时，都遵循如下流程。

<figure><img src="../../.gitbook/assets/image (1) (1) (1).png" alt=""><figcaption></figcaption></figure>

{% hint style="info" %}
Fireboom 同时只兼容一种语言的钩子！！！
{% endhint %}

## 读取配置文件

1. 配置文件：custom-go/generated/fireboom.config.json 是一个指向exported/generated/fireboom.config.json的软连接
2. 包含钩子启动所依赖的大部分信息，如钩子监听端口serverOptions.listen.port，S3配置信息s3UploadConfiguration等

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

{% hint style="info" %}
启动钩子前要检查custom-go/generated/fireboom.config.json是否存在，否则钩子无法启动。部署时，可借助 ./fireboom build 命令，生成上述文件。
{% endhint %}

## 读取环境变量

使用相对路径 `../.env`，和Fireboom服务共用

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

1，解析Fireboom调用时携带的全局参数 \_wg

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

2，为上下文ctx注入User对象，用于获取登录用户的信息

3，为上下文ctx注入InternalClient对象（用于[内部调用](../nei-bu-tiao-yong.md)）

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

## 注册路由

Fireboom根据开启的钩子，结合对应钩子模板生成对应的Hooks SDK。

### 钩子模板文件目录

{% tabs %}
{% tab title="nodejs" %}
```
template
├─ node-server
│  ├─ README.md
│  ├─ files
│  │  ├─ ecosystem.config.js
│  │  ├─ fireboom.server.ts.hbs
│  │  ├─ generated
│  │  │  ├─ claims.ts.hbs
│  │  │  ├─ client.legacy.ts.hbs
│  │  │  ├─ client.ts.hbs
│  │  │  ├─ fireboom.factory.ts
│  │  │  ├─ fireboom.internal.client.ts.hbs
│  │  │  ├─ fireboom.internal.operations.client.ts.hbs
│  │  │  ├─ fireboom.operations.ts.hbs
│  │  │  ├─ fireboom.server.ts.hbs
│  │  │  ├─ linkbuilder.ts.hbs
│  │  │  ├─ models.ts.hbs
│  │  │  ├─ prisma.ts.hbs
│  │  │  └─ testing.ts.hbs
│  │  ├─ nodemon.json
│  │  ├─ operations.tsconfig.json
│  │  ├─ package.json
│  │  ├─ scripts
│  │  │  ├─ buildOperations.ts
│  │  │  ├─ install.sh
│  │  │  ├─ run-build.sh
│  │  │  ├─ run-dev.sh
│  │  │  └─ run-prod.sh
│  │  └─ tsconfig.json
│  ├─ manifest.json
│  └─ partials
│     ├─ operation_partial.hbs
│     └─ schema_partial.hbs
```
{% endtab %}

{% tab title="golang" %}
```
template
├─ golang-server
│  ├─ README.md
│  ├─ files
│  │  ├─ generated
│  │  │  └─ models.go.hbs
│  │  ├─ go.mod
│  │  ├─ helix.html
│  │  ├─ main.go
│  │  ├─ pkg
│  │  │  ├─ base
│  │  │  │  ├─ client.go
│  │  │  │  ├─ operation.go
│  │  │  │  ├─ request.go
│  │  │  │  ├─ upload.go
│  │  │  │  └─ user.go
│  │  │  ├─ consts
│  │  │  │  └─ env.go
│  │  │  ├─ plugins
│  │  │  │  ├─ graphqls.go
│  │  │  │  └─ internal_request.go
│  │  │  ├─ types
│  │  │  │  ├─ configure.go
│  │  │  │  └─ server.go
│  │  │  ├─ utils
│  │  │  │  ├─ config.go
│  │  │  │  ├─ file.go
│  │  │  │  ├─ http.go
│  │  │  │  ├─ random.go
│  │  │  │  └─ strings.go
│  │  │  └─ wgpb
│  │  │     └─ wundernode_config.pb.go
│  │  ├─ scripts
│  │  │  ├─ install.sh
│  │  │  ├─ run-build.sh
│  │  │  ├─ run-dev.sh
│  │  │  └─ run-prod.sh
│  │  └─ server
│  │     ├─ fireboom_server.go.hbs # 生成的入口文件
│  │     └─ start.go
│  └─ manifest.json
```
{% endtab %}
{% endtabs %}

### 生成的钩子项目目录：

{% tabs %}
{% tab title="nodejs" %}
```
├─ custom-ts
│  ├─ customize
│  │  └─ A.ts
│  ├─ ecosystem.config.js
│  ├─ nodemon.json
│  ├─ operations.tsconfig.json
│  ├─ package.json
│  ├─ scripts
│  │  ├─ buildOperations.ts
│  │  ├─ install.sh
│  │  ├─ run-build.sh
│  │  ├─ run-dev.sh
│  │  └─ run-prod.sh
│  └─ tsconfig.json
```
{% endtab %}

{% tab title="golang" %}
```
├─ custom-go
│  ├─ auth
│  │  ├─ mutatingPostAuthentication.go
│  │  ├─ postAuthentication.go
│  │  ├─ postLogout.go
│  │  └─ revalidate.go
│  ├─ customize
│  │  ├─ S3.go
│  ├─ global
│  │  ├─ beforeRequest.go
│  │  ├─ onRequest.go
│  │  └─ onResponse.go
│  ├─ go.mod
│  ├─ go.sum
│  ├─ helix.html
│  ├─ hooks
│  │  └─ Weather
│  │     ├─ customResolve.go
│  │     ├─ mockResolve.go
│  │     ├─ mutatingPostResolve.go
│  │     ├─ mutatingPreResolve.go
│  │     ├─ postResolve.go
│  │     └─ preResolve.go
│  ├─ main.go
│  ├─ pkg
│  │  ├─ base
│  │  │  ├─ client.go
│  │  │  ├─ operation.go
│  │  │  ├─ request.go
│  │  │  ├─ upload.go
│  │  │  └─ user.go
│  │  ├─ consts
│  │  │  └─ env.go
│  │  ├─ plugins
│  │  │  ├─ auth_hooks.go
│  │  │  ├─ global_hooks.go
│  │  │  ├─ graphqls.go
│  │  │  ├─ internal_request.go
│  │  │  ├─ operation_hooks.go
│  │  │  ├─ proxy_hooks.go
│  │  │  └─ upload_hooks.go
│  │  ├─ types
│  │  │  ├─ configure.go
│  │  │  └─ server.go
│  │  ├─ utils
│  │  │  ├─ config.go
│  │  │  ├─ file.go
│  │  │  ├─ http.go
│  │  │  ├─ random.go
│  │  │  └─ strings.go
│  │  └─ wgpb
│  │     └─ wundernode_config.pb.go
│  ├─ proxys
│  │  └─ S3Presigned.go
│  ├─ scripts
│  │  ├─ install.sh
│  │  ├─ run-build.sh
│  │  ├─ run-dev.sh
│  │  └─ run-prod.sh
│  ├─ server
│  │  └─ start.go
│  └─ uploads
│     └─ tengxunyun
│        └─ avatar
│           ├─ postUpload.go
│           └─ preUpload.go
```
{% endtab %}
{% endtabs %}

### 注册路由

{% tabs %}
{% tab title="golang" %}
{% code title="server/start.go" %}
```go
# 注册proxy钩子
plugins.RegisterProxyHooks(e)
# 注册全局钩子
plugins.RegisterGlobalHooks(e, types.WdgHooksAndServerConfig.Hooks.Global)
# 注册授权钩子
plugins.RegisterAuthHooks(e, types.WdgHooksAndServerConfig.Hooks.Authentication)
# 注册上传钩子
plugins.RegisterUploadsHooks(e, types.WdgHooksAndServerConfig.Hooks.Uploads)

var internalQueries, internalMutations base.OperationDefinitions
nodeUrl := utils.GetConfigurationVal(types.WdgGraphConfig.Api.NodeOptions.NodeUrl)
queryOperations := filterOperationsHooks(types.WdgGraphConfig.Api.Operations, wgpb.OperationType_QUERY)
# 注册局部钩子
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
# 注册内部调用钩子
plugins.BuildDefaultInternalClient(internalQueries, internalMutations)
for _, registeredHook := range base.GetRegisteredHookArr() {
	go registeredHook(e.Logger)
}
# 注册graphql钩子
for _, gqlServer := range types.WdgHooksAndServerConfig.GraphqlServers {
	plugins.RegisterGraphql(e, gqlServer)
}
```
{% endcode %}
{% endtab %}

{% tab title="nodejs" %}

{% endtab %}
{% endtabs %}

### 健康检查

在开始前我们先学习钩子服务的第一个路由：健康检查接口

{% swagger method="get" path="health" baseUrl="http://127.0.0.1:9992/" summary="健康检查接口" %}
{% swagger-description %}
检查钩子服务健康状态，用于在界面上展示钩子是否已启动
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

### 剩余钩子



[operation-gou-zi.md](../operation-gou-zi.md "mention")

[shen-fen-yan-zheng-gou-zi.md](../shen-fen-yan-zheng-gou-zi.md "mention")

[graphql-gou-zi.md](../graphql-gou-zi.md "mention")

[wen-jian-shang-chuan-gou-zi.md](../wen-jian-shang-chuan-gou-zi.md "mention")

[nei-bu-tiao-yong.md](../nei-bu-tiao-yong.md "mention")

