# 钩子规范

钩子服务本质上是一个实现了飞布钩子规范的WEB服务，可以用任意后端语言实现。

当前飞布已经支持了TS语法的nodejs钩子，下一步计划实现golang语言的钩子SDK。

如果你希望实现一个自己的 hook  SDK，需要遵从如下协议。

## OPERATION钩子

为扩展OPERATION的能力而设计的钩子，主要用于定制请求OPERATION API的行为。

按照切入位置不同，分别如下：

### 全局钩子

#### 前置拦截



#### 后置拦截



### 局部钩子

#### 前置钩子



#### 后置钩子



#### 模拟钩子



#### 自定义钩子



## 登录钩子

为扩展登录逻辑而设计的钩子，主要用于自定义登录行为。

### 登录后置钩子

#### mutatingPostAuthentication

OIDC流程用户登录成功后，执行该钩子。主要用于修改登录对象user的值，实现特定逻辑，如绑定用户角色。



#### postAuthentication

OIDC流程用户登录成功后，执行该钩子。与上述钩子的区别是，无返回值，无法修改user对象。



### 登出后置钩子



### 重校验钩子





## 文件钩子



## 内部钩子









1. 生成一个自己的入口文件，ts 版本为 `fireboom.server.ts`，参考 ts 的实现生成`authentication` `queries` `mutations` `global`等钩子的引用和其它必要配置
2. 提供一个 SDK 给上述入口文件使用，接受上述配置，启动API服务器
3. 【可选】SDK 提供日志输出，保存注册的路由等日志信息，方便后续调试
4. 根据环境变量`START_HOOKS_SERVER`决定是否启动钩子服务
5. 首页路径`/`指向`INDEX_PAGE` 环境变量指定的`index.html`文件内容，一般用`static`模块或者`sendFile`
6. 读取`{WG_ABS_DIR}/generated/wundergraph.config.json`中的配置文件
7. 构造请求上下文`ctx`，后续文中`ctx`皆指此。可通过中间件或者拦截器等方法实现，每个路由都需要，构造方法参考如下

```ts
req.ctx = {
  log: req.log,
  user: req.body.__wg.user,
  // clientRequest represents the original client request that was sent initially to the WunderNode.
  clientRequest: {
      headers: new headers_polyfill_1.Headers((_a = req.body.__wg.clientRequest) === null || _a === void 0 ? void 0 : _a.headers),
      requestURI: ((_b = req.body.__wg.clientRequest) === null || _b === void 0 ? void 0 : _b.requestURI) || '',
      method: ((_c = req.body.__wg.clientRequest) === null || _c === void 0 ? void 0 : _c.method) || 'GET',
  },
  internalClient: clientFactory({}, req.body.__wg.clientRequest),
}
```

大致是根据参数里的`__wg`数据展开并提供一个内部 client。同时外面生成的钩子调用处的函数签名要保持一致（函数参数一般就是上述的`ctx`）

下述的`ret`代表用户钩子执行返回值

~~如果`ctx.user`为空，返回400 reply.code(400).send({ error: "User context doesn't exist" });~~

8. 注册钩子路由



| 方法   | 路径                                             | 入参                                                                 | 成功出参                                                                                                                              | 失败出参                                                      | 说明                                                                       |
| ---- | ---------------------------------------------- | ------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------- | ------------------------------------------------------------------------ |
| post | /authentication/postAuthentication             | ctx                                                                | { hook: 'postAuthentication' }                                                                                                    | { hook: 'postAuthentication', error }                     | 有用户钩子函数则调用该函数，成功200，失败500，下同                                             |
| post | /authentication/mutatingPostAuthentication     | ctx                                                                | { hook: 'postAuthentication', response: 函数返回值, setClientRequestHeaders: 参考flattenHeaders }                                        | { hook: 'postAuthentication', error }                     |                                                                          |
| post | /authentication/revalidateAuthentication       | ctx                                                                | { hook: 'revalidateAuthentication', response: ret, setClientRequestHeaders: 参考flattenHeaders }                                    | { hook: 'revalidateAuthentication', error }               |                                                                          |
| post | /global/httpTransport/onOriginRequest          |                                                                    |                                                                                                                                   |                                                           | 不常用，参考 ts 源码实现                                                           |
| post | /global/httpTransport/onOriginResponse         |                                                                    |                                                                                                                                   |                                                           | 不常用，参考 ts 源码实现                                                           |
| post | /global/wsTransport/onConnectionInit           |                                                                    |                                                                                                                                   |                                                           | 需根据 `config.global?.wsTransport?.onConnectionInit` 判断是否开启，不常用，参考 ts 源码实现 |
| post | /operation/{operationName}/mockResolve         | { ...ctx, input: req.body.input }                                  | { op: operationName, hook: 'mock', response: ret, setClientRequestHeaders: request.ctx.clientRequest.headers }                    | { op: operationName, hook: 'mock', error }                |                                                                          |
| post | /operation/{operationName}/preResolve          | { ...ctx, input: req.body.input }                                  | { op: operationName, hook: 'preResolve', setClientRequestHeaders: request.ctx.clientRequest.headers }                             | { op: operationName, hook: 'preResolve', error }          |                                                                          |
| post | /operation/{operationName}/postResolve         | { ...ctx, input: req.body.input, response: request.body.response } | { op: operationName, hook: 'postResolve', setClientRequestHeaders: request.ctx.clientRequest.headers }                            | { op: operationName, hook: 'postResolve', error }         |                                                                          |
| post | /operation/{operationName}/mutatingPreResolve  | { ...ctx, input: req.body.input }                                  | { op: operationName, hook: 'mutatingPreResolve', input: ret, setClientRequestHeaders: request.ctx.clientRequest.headers }         | { op: operationName, hook: 'mutatingPreResolve', error }  |                                                                          |
| post | /operation/{operationName}/mutatingPostResolve | { ...ctx, input: req.body.input, response: request.body.response } | { op: operationName, hook: 'mutatingPostResolve', response: ret, setClientRequestHeaders: request.ctx.clientRequest.headers }     | { op: operationName, hook: 'mutatingPostResolve', error } |                                                                          |
| post | /operation/{operationName}/customResolve       | { ...ctx, input: req.body.input }                                  | { op: operationName, hook: 'customResolve', response: ret \|\| null, setClientRequestHeaders: request.ctx.clientRequest.headers } | { op: operationName, hook: 'customResolve', error }       |                                                                          |

其中`{operationName}`为`api.operations`遍历时的`operation.name`

9. 【可选】根据配置`serverConfig.graphqlServers`注册对应数量的路由`/gqls/{server.serverName}/graphql`
10. 生产环境（根据环境变量）优雅关闭
