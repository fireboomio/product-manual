# 钩子规范

钩子服务本质上是一个实现了飞布钩子规范的WEB服务，可以用任意后端语言实现。

如果你希望实现其他语言的 hook SDK，需要遵从如下协议。

根据用途划分，钩子可分为4大类：局部钩子、全局钩子、授权钩子、文件钩子。

### 局部钩子（OPERATION钩子）

局部钩子目的是扩展OPEARTION的能力，分别在“OPEARTION执行”前后执行，主要用途是参数校验和副作用触发，如创建文章后发送邮件通知审核。

详情见如下流程图。

![](../../assets/hook-flow.png)

前置钩子在 "执行OPERATION"前执行，可修改校验或修改输入参数。

{% tabs %}
{% tab title="参数校验" %}
```go
func PreResolve(hook *base.HookRequest, body generated.Todo__CreateOneTodoBody) (res generated.Todo__CreateOneTodoBody, err error) {
    if body.Input.Title == "" {
	return nil, errors.New("标题不能为空")
    }
    return body, nil
}
```
{% endtab %}

{% tab title="修改入参" %}
```go
func MutatingPreResolve(hook *base.HookRequest, body generated.Todo__CreateOneTodoBody) (res generated.Todo__CreateOneTodoBody, err error) {
    if body.Input.Title == "" {
	body.Input.Title = "默认标题"
    }
    return body, nil
}
```
{% endtab %}
{% endtabs %}

其中mutating开头的钩子可以修改 request 入参。

后置钩子在 "执行OPERATION" 后执行，可触发副作用（发邮件）或修改响应参数。

{% tabs %}
{% tab title="触发副作用" %}
```go
func PostResolve(hook *base.HookRequest, body generated.Todo__CreateOneTodoBody) (res generated.Todo__CreateOneTodoBody, err error) {
	fmt.Println("我要发一封邮件xxx,标题是：", body.Input.Title, data)
	return body, nil
}
```
{% endtab %}

{% tab title="修改响应参数" %}

{% endtab %}
{% endtabs %}

其中mutating开头的钩子可以修改 response 结果。

除了上述局部钩子，还有两个特殊的局部钩子：自定义处理钩子和模拟钩子。

自定义处理钩子（customResolve）：

如果该钩子有返回值，那么将跳过“执行OPERATION”，直接返回当前钩子的返回值 ，否则继续执行后续流程。使用该钩子，可以修改默认 “执行OPERATION”的逻辑。

模拟钩子（mockResolve）：

用于返回模拟值，使用时会短路其余所有局部钩子。



| 路径                                             | 入参                                                                 | 成功出参                                                                                                                              | 失败出参                                                      | 说明                                                    |
| ---------------------------------------------- | ------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------- | ----------------------------------------------------- |
| /operation/{operationName}/mockResolve         | { ...ctx, input: req.body.input }                                  | { op: operationName, hook: 'mock', response: ret, setClientRequestHeaders: request.ctx.clientRequest.headers }                    | { op: operationName, hook: 'mock', error }                | 模拟钩子，直接返回模拟数据而不经过其它流程                                 |
| /operation/{operationName}/preResolve          | { ...ctx, input: req.body.input }                                  | { op: operationName, hook: 'preResolve', setClientRequestHeaders: request.ctx.clientRequest.headers }                             | { op: operationName, hook: 'preResolve', error }          | 前置钩子，operation 处理前执行                                  |
| /operation/{operationName}/postResolve         | { ...ctx, input: req.body.input, response: request.body.response } | { op: operationName, hook: 'postResolve', setClientRequestHeaders: request.ctx.clientRequest.headers }                            | { op: operationName, hook: 'postResolve', error }         | 后置钩子，operation 处理后执行                                  |
| /operation/{operationName}/mutatingPreResolve  | { ...ctx, input: req.body.input }                                  | { op: operationName, hook: 'mutatingPreResolve', input: ret, setClientRequestHeaders: request.ctx.clientRequest.headers }         | { op: operationName, hook: 'mutatingPreResolve', error }  | 前置可修改钩子，可以修改 request 入参                               |
| /operation/{operationName}/mutatingPostResolve | { ...ctx, input: req.body.input, response: request.body.response } | { op: operationName, hook: 'mutatingPostResolve', response: ret, setClientRequestHeaders: request.ctx.clientRequest.headers }     | { op: operationName, hook: 'mutatingPostResolve', error } | 后置可修改钩子，可以修改返回的response                               |
| /operation/{operationName}/customResolve       | { ...ctx, input: req.body.input }                                  | { op: operationName, hook: 'customResolve', response: ret \|\| null, setClientRequestHeaders: request.ctx.clientRequest.headers } | { op: operationName, hook: 'customResolve', error }       | 自定义处理钩子，如果该钩子有返回值，那么将跳过后续的流程，直接返回 response，否则继续执行后续流程 |

其中`{operationName}`为`api.operations`遍历时的`operation.name`

### 全局钩子（数据源钩子）



| 路径                                     | 入参 | 成功出参 | 失败出参 | 说明                                                                         |
| -------------------------------------- | -- | ---- | ---- | -------------------------------------------------------------------------- |
| /global/httpTransport/onOriginRequest  | -  | -    | -    | 全局钩子 - 前置拦截                                                                |
| /global/httpTransport/onOriginResponse | -  | -    | -    | 全局钩子 - 后置拦截                                                                |
| /global/wsTransport/onConnectionInit   | -  | -    | -    | subscription 钩子， 需根据 `config.global?.wsTransport?.onConnectionInit` 判断是否开启 |

### 授权钩子



| 路径                                         | 入参  | 成功出参                                                                                           | 失败出参                                        | 说明                                                  |
| ------------------------------------------ | --- | ---------------------------------------------------------------------------------------------- | ------------------------------------------- | --------------------------------------------------- |
| /authentication/postAuthentication         | ctx | { hook: 'postAuthentication' }                                                                 | { hook: 'postAuthentication', error }       | OIDC流程用户登录成功后，执行该钩子，不可修改user对象，成功200，失败500，下同       |
| /authentication/mutatingPostAuthentication | ctx | { hook: 'postAuthentication', response: 函数返回值, setClientRequestHeaders: 参考flattenHeaders }     | { hook: 'postAuthentication', error }       | OIDC流程用户登录成功后，执行该钩子。主要用于修改登录对象user的值，实现特定逻辑，如绑定用户角色 |
| /authentication/revalidateAuthentication   | ctx | { hook: 'revalidateAuthentication', response: ret, setClientRequestHeaders: 参考flattenHeaders } | { hook: 'revalidateAuthentication', error } | 重校验钩子                                               |

### 文件钩子





### 数据代理

飞布服务不仅可以按照约定调用钩子服务，钩子也可以调用其它钩子，此时飞布服务变身为数据代理。



### NodeJs 上下文参考

```ts
req.ctx = {
  log: pino({
    level: PinoLogLevel.Debug
  }),
  user: req.body.__wg.user!,
  // clientRequest represents the original client request that was sent initially to the WunderNode.
  clientRequest: {
    headers: new Headers(req.body.__wg.clientRequest?.headers),
    requestURI: req.body.__wg.clientRequest?.requestURI || '',
    method: req.body.__wg.clientRequest?.method || 'GET',
  },
  internalClient: clientFactory({ 'x-request-id': req.id }, req.body.__wg.clientRequest),
}
```

### NodeJs 钩子参考

参考[ NodeJs 钩子](node-gou-zi.md)实现其它语言的钩子服务。
