# function

function钩子注册到Fireboom中为一个API，且有出入参定义（json类型），此外还支持实时查询和权限控制。

见上图，API管理->function->login

{% code title="custom-go/function/login.go" %}
```go
package function

import (
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
)

func init() {
	plugins.RegisterFunction[loginReq, loginRes](login)
}

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Res      loginRes
}

type loginRes struct {
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func login(hook *base.HookRequest, body *base.OperationBody[loginReq, loginRes]) (*base.OperationBody[loginReq, loginRes], error) {
	body.Response = &base.OperationBodyResponse[loginRes]{Data: loginRes{Msg: "123"}}
	return body, nil
}
```
{% endcode %}

若想使该钩子生效，还需要在`main.go`文件中开启第6行注释，匿名导入钩子。

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
