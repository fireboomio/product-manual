# 钩子规范

钩子服务本质上是一个实现了飞布钩子规范的WEB服务，可以用任意后端语言实现。

如果你希望实现其他语言的 hook SDK，需要遵从如下协议。

根据用途划分，钩子可分为4大类：局部钩子、全局钩子、授权钩子、文件钩子。

### 局部钩子（OPERATION钩子）

局部钩子目的是扩展OPEARTION的能力，分别在“OPEARTION执行”前后执行，主要用途是参数校验和副作用触发，如创建文章后发送邮件通知审核。

详情见如下流程图。

![](../../assets/hook-flow.png)

前置钩子在 "执行OPERATION"前执行，可校验参数或修改输入参数。

{% tabs %}
{% tab title="前置普通钩子" %}
```go
// 自定义参数校验
func PreResolve(hook *base.HookRequest, body generated.Todo__CreateOneTodoBody) (res generated.Todo__CreateOneTodoBody, err error) {
    if body.Input.Title == "" {
	return nil, errors.New("标题不能为空")
    }
    return body, nil
}
```
{% endtab %}

{% tab title="前置修改入参钩子" %}
```go
// 修改operation的input入参
func MutatingPreResolve(hook *base.HookRequest, body generated.Todo__CreateOneTodoBody) (res generated.Todo__CreateOneTodoBody, err error) {
    if body.Input.Title == "" {
	body.Input.Title = "默认标题"
    }
    return body, nil
}
```
{% endtab %}
{% endtabs %}

后置钩子在 "执行OPERATION" 后执行，可触发自定义操作或修改响应结果。

{% tabs %}
{% tab title="后置普通钩子" %}
```go
// 执行自定义操作
func PostResolve(hook *base.HookRequest, body generated.Todo__CreateOneTodoBody) (res generated.Todo__CreateOneTodoBody, err error) {
    fmt.Println("我要发一封邮件xxx,标题是：", body.Input.Title, data)
    return body, nil
}
```
{% endtab %}

{% tab title="后置修改响应钩子" %}
```go
// 修改operation响应结果
func MutatingPostResolve(hook *base.HookRequest, body generated.Todo__CreateOneTodoBody) (res generated.Todo__CreateOneTodoBody, err error) {
    body.Response.Data.Data.UpdateAt = time.Now()
    return body, nil
}
```
{% endtab %}
{% endtabs %}

除了上述局部钩子，还有两个特殊的局部钩子：自定义处理钩子和模拟钩子。

{% tabs %}
{% tab title="自定义处理钩子" %}
```go
// 若该钩子有返回值，那么将跳过“执行OPERATION”，直接返回当前钩子的返回值 
func CustomResolve(hook *base.HookRequest, body generated.$HOOK_NAME$Body) (res generated.$HOOK_NAME$Body, err error) {
    hook.Logger().Info("CustomResolve")
    return body, nil
}
```
{% endtab %}

{% tab title="模拟数据钩子" %}
```go
// 用于返回模拟值，使用时会短路其余所有局部钩子。
func MockResolve(hook *base.HookRequest, body generated.$HOOK_NAME$Body) (res generated.$HOOK_NAME$Body, err error) {
    hook.Logger().Info("MockResolve")
    return body, nil
}
```
{% endtab %}
{% endtabs %}

### 全局钩子

全局钩子目的是改写原有的request和response，分别在请求最初、opertaion触发前后，可以用来实现全局的参数/响应改写或发送全局的通知，包括预执行、前置、后置钩子。

{% tabs %}
{% tab title="预执行钩子" %}
```go

// 在最初请求，可以修改body和header（请使用OriginBody）
func BeforeOriginRequest(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientRequest, error) {
    return modifyForPayNotify(body)
}

func modifyForPayNotify(body *plugins.HttpTransportBody) (*base.ClientRequest, error) {
    u, err := url.Parse(body.Request.RequestURI)
    if err != nil {
	return nil, err
    }

    if u.Path != "/operations/Payment/PayNotify" {
	return body.Request, nil
    }

    // 1. 从body中取值
    modifyBody, err := sjson.Set("{}", "data", string(body.Request.OriginBody))
    if err != nil {
	return nil, err
    }

    // 2. 从url中取值
    for key, valArr := range u.Query() {
	modifyBody, _ = sjson.Set(modifyBody, key, valArr[0])
    }
    body.Request.Body = []byte(modifyBody)
    return body.Request, nil
}
```
{% endtab %}

{% tab title="前置钩子" %}
```go
// 在operation执行前，可以修改请求的body和header
func OnOriginRequest(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientRequest, error) {
    modifyBody := string(body.Request.Body)
    modifyBody, err := sjson.Set(modifyBody, "name", "admin")
    body.Request.Body = []byte(modifyBody)
    return body.Request, nil
}
```
{% endtab %}

{% tab title="后置钩子" %}
```go
// 在operation执行后，可以修改响应的body和header
func OnOriginResponse(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientResponse, error) {
   if hook.User == nil {
       body.Response.StatusCode = 401
   }
   
   return body.Response, nil
}
```
{% endtab %}
{% endtabs %}

### 授权钩子

认证钩子目的是在认证成功后执行自定义操作，可以用来实现用户信息同步，用户信息改写，用户重新校验（需要请求携带参数revalidate）

{% tabs %}
{% tab title="后置普通钩子" %}
<pre class="language-go"><code class="lang-go">// 在认证后做自定义处理，比如同步用户信息等
func PostAuthentication(hook *base.AuthenticationHookRequest) error {
    hook.Context.Logger().Infof("用户%s已同步", hook.User.NickName)
    return nil
<strong>}
</strong></code></pre>
{% endtab %}

{% tab title="后置修改认证信息钩子" %}
```go
// 在认证后修改用户信息
func MutatingPostAuthentication(hook *base.AuthenticationHookRequest) (*plugins.AuthenticationResponse, error) {
    hook.User.Name = "admin"
    return &plugins.AuthenticationResponse{User: hook.User, Status: "ok"}, nil
}
```
{% endtab %}

{% tab title="重新校验钩子" %}
```go
// 重新认证，默认从缓存中获取，当请求中携带revalidate参数时，所有认证钩子会依次重新执行一次
func Revalidate(hook *base.AuthenticationHookRequest) (*plugins.AuthenticationResponse, error) {
    return &plugins.AuthenticationResponse{User: hook.User, Status: "ok"}, nil
}
```
{% endtab %}
{% endtabs %}

### 文件钩子

文件钩子目的是在文件上传前后进行自定义操作，可以用来实现重写文件名和自定义处理

{% tabs %}
{% tab title="前置钩子" %}
```go
// 改写上传后的文件名，默认随机字符串
func PreUpload(request *base.PreUploadHookRequest, body *plugins.UploadBody[any]) (*base.UploadHookResponse, error) {
    return &base.UploadHookResponse{FileKey: body.File.Name}, nil
}
```
{% endtab %}

{% tab title="后置钩子" %}
```go
// 自定义后置操作（记录上传日志）
func PostUpload(request *base.PreUploadHookRequest, body *plugins.UploadBody[any]) (*base.UploadHookResponse, error) {
    request.Context.Logger().Infof("【%s】上传了文件【%s】", hook.User.NickName, body.File.Name)
    return nil, nil
}
```
{% endtab %}
{% endtabs %}

### 数据代理

飞布服务不仅可以按照约定调用钩子服务，钩子也可以调用其它钩子，此时飞布服务变身为数据代理。
