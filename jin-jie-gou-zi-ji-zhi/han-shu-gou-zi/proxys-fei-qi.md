# proxy

proxy钩子注册到Fireboom中也为一个API，和funciton的区别是，它没有出入参定义，可以为任意类型，如非结构化数据或xml数据，同时不支持实时查询。

见上图，API管理->proxy->test

{% code title="custom-go/proxy/test.go" %}
```go
package proxy

import (
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"custom-go/pkg/wgpb"
	"net/http"
)

func init() {
	plugins.RegisterProxyHook(ping)
}

var conf = &plugins.HookConfig{
	AuthRequired: true,
	AuthorizationConfig: &wgpb.OperationAuthorizationConfig{
		RoleConfig: &wgpb.OperationRoleConfig{
			RequireMatchAny: []string{"admin", "user"},
		},
	},
	EnableLiveQuery: false,
}

func ping(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientResponse, error) {
	// do something here ...
	body.Response = &base.ClientResponse{
		StatusCode: http.StatusOK,
	}
	body.Response.OriginBody = []byte("ok")
	return body.Response, nil
}

```
{% endcode %}

推荐优先使用function，funciton满足不了的，再用proxy钩子。

若想使上述该钩子生效，还需要在`main.go`文件中开启第7行注释，匿名导入proxy钩子。

{% code title="custom-go/main.go" lineNumbers="true" %}
```go
package main

import (
	// 根据需求，开启注释
	_ "custom-go/customize"
	_ "custom-go/function"
	_ "custom-go/proxy"
	"custom-go/server"
)

func main() {
	server.Execute()
}
```
{% endcode %}



参考：

* fb-admin：[https://github.com/fireboomio/fb-admin/blob/main/backend/custom-go/proxy/bindmenu.go](https://github.com/fireboomio/fb-admin/blob/main/backend/custom-go/proxy/bindmenu.go)
