# proxy

proxy钩子注册到Fireboom中也为一个API，和funciton的区别是：

* 没有出入参定义，可以为任意类型，如非结构化数据或xml数据
* 不支持实时查询

**推荐优先使用function，funciton满足不了的，再用proxy钩子**。

<figure><img src="../../.gitbook/assets/image (18).png" alt=""><figcaption><p>proxy API界面</p></figcaption></figure>

具体操作步骤同 [zu-he-shi-api.md](zu-he-shi-api.md "mention")

示例代码：

{% code title="custom-go/proxy/ping.go" %}
```go
package proxy

import (
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"net/http"
)

func init() {
	plugins.RegisterProxyHook(ping)
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

路由规则：

```http
http://localhost:9991/operations/proxy/[proxy-name]

example:: http://localhost:9991/operations/proxy/ping
```

参考：

* fb-admin：[https://github.com/fireboomio/fb-admin/blob/main/backend/custom-go/proxy/bindmenu.go](https://github.com/fireboomio/fb-admin/blob/main/backend/custom-go/proxy/bindmenu.go)
