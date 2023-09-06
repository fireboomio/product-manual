# function

function钩子注册到Fireboom中为一个API，且有出入参定义（json类型），此外还支持实时查询和权限控制。

## 新建function

1，在 Fireboom 控制台点击`数据源`面板的`+`号，进入数据源新建页。

2，在数据源新建页面，选择 脚本-> Function，设置名称为：`fun1`。

3，系统自动初始化如下脚本。

{% code title="custom-go/function/fun1.go" %}
```go
package function

import (
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"custom-go/pkg/wgpb"
)

func init() {
	// 注册 function
	plugins.RegisterFunction[helloReq, helloRes](hello, wgpb.OperationType_MUTATION)
}

type helloReq struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Info     helloInfo `json:"info,omitempty"`
}

type helloInfo struct {
	Code    string `json:"code,omitempty"`
	Captcha string `json:"captcha,omitempty"`
}

type helloRes struct {
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func hello(hook *base.HookRequest, body *base.OperationBody[helloReq, helloRes]) (*base.OperationBody[helloReq, helloRes], error) {
	if body.Input.Username != "John" || body.Input.Password != "123456" {
		body.Response = &base.OperationBodyResponse[helloRes]{
			Errors: []base.GraphQLError{{Message: "username or password wrong"}},
		}
		return body, nil
	}

	body.Response = &base.OperationBodyResponse[helloRes]{Data: helloRes{Msg: "hello success"}}
	return body, nil
}

```
{% endcode %}

默认填充的是示例代码，你可以根据业务需求修改代码。

4，在`main.go`中匿名引入该包，然后重新启动钩子服务

{% tabs %}
{% tab title="golang" %}
{% code title="custom-go/main.go" %}
```go
package main

import (
	// _ "custom-go/customize"
	// 匿名引入该库
	 _ "custom-go/function"
	// _ "custom-go/proxy"
	"custom-go/server"
)

func main() {
	server.Execute()
}
```
{% endcode %}
{% endtab %}
{% endtabs %}

5，钩子服务启动后，生成对应的配置文件`custom-x/function/fun1.json`

```
function          
├─ fun1.go   # function 钩子代码   
└─ fun1.json # 该钩子对应的配置文件
```

{% code title="fun1.json" %}
```json
{
    "name": "fun1",
    // 请求类型，1代表POST请求
    "operationType": 1,
    // 请求入参的JSON SCHEMA
    "variablesSchema": "{\"definitions\":{\"helloInfo\":{\"properties\":{\"captcha\":{\"type\":\"string\"},\"code\":{\"type\":\"string\"}},\"type\":\"object\"}},\"properties\":{\"info\":{\"$ref\":\"#/definitions/helloInfo\"},\"password\":{\"type\":\"string\"},\"username\":{\"type\":\"string\"}},\"required\":[\"username\",\"password\"],\"type\":\"object\"}",
    // 响应结果的JSON SCHEMA
    "responseSchema": "{\"properties\":{\"data\":{\"type\":\"string\"},\"msg\":{\"type\":\"string\"}},\"required\":[\"msg\",\"data\"],\"type\":\"object\"}",
    "cacheConfig": null,
    // 授权配置
    "authenticationConfig": null,
    // 实时查询配置，仅限于GET请求
    "liveQueryConfig": null,
    // RBAC配置
    "authorizationConfig": null,
    "hooksConfiguration": null,
    "variablesConfiguration": null,
    "engine": 0,
    "path": "/function/fun1", // 路由
}
```
{% endcode %}

6，Fireboom服务每秒触发1次健康检查，将获取如下结果：

```json
{
    "report": {
        // function 钩子
        "functions": [
            "fun1"
        ],
        "time": "2023-09-06T17:18:21.957519+08:00"
    },
    // 钩子服务的状态
    "status": "ok"
}
```

可以看到 `report.`functions 中包含`fun1`接口。接着读取步骤5中的json配置，将该function注册到API列表中，类型为`function`。

请访问路由如下：

```bash
curl http://localhost:9991/operations/function/fun1
-X POST
-H 'Content-Type: application/json'
--data-raw '{"info":{"captcha":"string","code":"string"},"password":"string","username":"string"}'
--compressed
```

其路由规则为：

```
http://localhost:9991/operations/function/fun1
```

{% hint style="info" %}
如果`report`没有变化，则不会触发编译！
{% endhint %}

7，最后，Fireboom发送通知，触发控制台更新，在API管理面板展示API： `fun1` 。

<figure><img src="../../.gitbook/assets/image (15).png" alt=""><figcaption></figcaption></figure>

8，前往swagger文档可测试API：`fun1`
