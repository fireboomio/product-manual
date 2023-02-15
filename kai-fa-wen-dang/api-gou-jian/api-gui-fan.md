# API规范

飞布API由GraphQL OPERATION编译而成，不同OPERATION会生成不同类型的API。接下来，我们具体介绍OPERATION的编译规范。

## OPERATION类型

### URL 结构

URL结构如下：

```
https://<hostname>/operations/<operation>
```

假定你的域名是\`example.com\`，对于 getUser OPERATION，它的URL为：

```
https://example.com/operations/getUser
```

### 状态码

* 200：Operation执行成功
* 404：Operation未找到
* 401：身份验证或身份鉴权失败
* 400：入参校验失败
* 500：Operation执行失败

### 查询Queries

查询对应GET请求。对于发送变量，有两种方式：

在URL查询字符串中：

```
GET https://<hostname>/operations/<operationName>?name=Jannik
```

将参数用URL编码为JSON对象，然后赋值给`wg_variables`查询参数：

```
GET https://<hostname>/operations/<operationName>?wg_variables={"name":"Jannik"}
```

如果你的变量是扁平的，并且想要使用像Postman或curl这样的客户端，推荐用第一种方式。如果是根据OPERATION自动生成的客户端，推荐使用第二种方式，因为它支持嵌套变量。

此外，客户端也可以通过`wg_api_hash`查询参数，主动触发缓存失效。

### 变更Mutations

查询对应POST请求。变量作为JSON在请求体中发送。

```
POST https://<hostname>/operations/<operationName>
Content-Type: application/json

{
  "name": "Jannik"
}
```

对于使用基于Cookie验证的客户端，需要设置`X-CSRF-Token`请求头。不然，服务器将拒绝该请求。

### 订阅Subscriptions

查询也对应GET请求。响应是一个JSON对象流，由两个换行符分割。

发送入参的方式和查询Queries一样。

```
GET https://<hostname>/operations/<operationName>?name=Jannik
```

当客户端连接后，服务器将持续发送数据，直到客户端断开连接。

客户端也可以为URL增加可选字段`wg_subscribe_once` 。如果值为`true`，服务端将只发送一次消息，然后断开连接。当服务端第一次渲染页面时，该方式很有用。

客户端也可以为URL增加可选字段`wg_see`。如果值为`true`，服务端将使用 [Server-Sent Events ](https://en.wikipedia.org/wiki/Server-sent\_events) 协议发送消息。每个消息都以`data：`开头。当你想增加更好的调试能力时，这很有用，因为浏览器能把消息解析为  Server-Sent Events。

### 实时查询

实时查询也对应GET请求。响应是一个JSON对象流，由两个换行符分割。

发送入参的方式和查询Queries一样。

客户端必须为URL增加wg\_live字段，表明他们想接收实时更新。

```
GET https://<hostname>/operations/<operationName>?name=Jannik&wg_live=true
```

## 响应格式

对于所有operations，响应格式都是包含两个字段的JSON对象：`data`和`errors`。

有可能两个字段同时存在。该情况，意味着操作部分失败了。

只有errors，代表Operation失败了。

只有data，代表Operation成功了。

data对象的类型是从操作(Operation)推断出来的。

errors对象的类型是一个包含以下字段的对象数组：

```graphql
type Error {
  message: String
  path: [String]
}
```

## 授权

飞布支持基于cookie和基于令牌的身份验证。

### 基于Cookie的验证

为了启动身份验证流程，客户端应该向URL发送一个GET请求：

```
GET https://<hostname>/auth/cookie/authorize/<authProviderID>?redirect_uri=[redirect_uri]
```

客户端必须向URL发送一个`redirect_uri`查询参数。这是客户端在身份验证流程完成后应该重定向到的URL。

{% hint style="info" %}
注意，必须把重定向URI列入白名单。
{% endhint %}

授权流成功后，会向`redirect_uri` 所在页面注入登录用户的cookie，然后可以通过下述URL获取当前用户信息：

```
GET https://<hostname>/auth/cookie/user
```

客户端可以向URL添加一个可选的`revalidate=true`查询参数。如果将此设置为true，服务器将触发用户的身份验证状态的重新验证，允许后端更新或撤销用户的身份验证状态。

### 基于Token的验证

对于非基于浏览器的客户端，也可以使用基于令牌的身份验证。

在这种情况下，客户端需要向请求中添加以下请求头：

```
Authorization: Bearer <token>
```

Bearer Token需要从身份提供者处获得，该过程飞布无法控制。

## CSRF保护

飞布自动保护变更Mutation免受CSRF攻击。客户端需要获取一个CSRF令牌，并将其设置在`X-CSRF-Token`报头中。

若想获得CSRF令牌，客户端需要调用csrf端点。

```
GET https://<hostname>/auth/cookie/csrf
```

响应包含文本格式的cookie。身份验证后，客户端需要再次调用CSRF端点来获取一个新的CSRF令牌。

## 文件上传&#x20;

飞布支持文件上传。文件以`multipart/form-data`编码的HTTP请求形式发送。

文件需要作为`files`字段添加到multipart对象中，并作为POST请求发送到以下URL：

```
POST https://<hostname>/s3/<storageID>/upload
```

响应包含一个带有字段`fileKeys`的json对象，该对象是上传文件生成的id列表。
