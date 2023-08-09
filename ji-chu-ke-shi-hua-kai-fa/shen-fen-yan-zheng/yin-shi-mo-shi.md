# 隐式模式

OIDC除了授权码模式，还有隐式模式，即基于token的登录。该模式常结合移动应用或 Web App 使用。

> OIDC 隐式模式不会返回授权码 code，而是直接将 `access_token` 和 `id_token` 通过 **URL hash** 发送到**回调地址前端**，**后端无法获取**到这里返回的值，因为 URL hash 不会被直接发送到后端。

## OIDC配置

我们仍以授权码模式中提到的这张图为例：

<figure><img src="../../.gitbook/assets/image (52).png" alt=""><figcaption></figcaption></figure>

隐式模式不需要配置`App Secret`，但开启该模式后，需要保证`JWKS`可用，有两种指定方式：

* JWKS端点：[https://dev-5kzk7gzc.us.auth0.com/.well-known/jwks.json](https://dev-5kzk7gzc.us.auth0.com/.well-known/jwks.json)
* JWKS json文件：

```json
{
    "kty": "RSA",
    "e": "AQAB",
    "n": "1GBBv-QOtbNIvgJZqvW2nvIrNx6-YNKJAD3L3WspAcx1y-RYctI2RBb4k4GN0du8AH2UUf8wBywONHplYAw1djkWAztHgj4cc_WxqKvD1t5bNNjRW7I5EPA9ZEkFblIAxZVwhOPK5H8KLgiVaD7y9fPEks6sVhu2VUQKC0Qr85-0WJVzmXP3QH_1yLn1qRpkJtjCW1I4DPsB0TrQC6WBMy99Io8zECraueLrJFApuRx1H_MwgDwnt4VlYuaoqU17TyBUQWO077mUB-FFI-s0jALuPAUuNWHFFogTq2cbydaSfPcWQPjylYcLcIt-bBBdedLqsTk_0nTXPqREMFwexw",
    "alg": "RS256",
    "use": "sig",
    "kid": "AUQR2TFiVexgvm0j0PrbZ3ofEz9R2eG7qvEJP9Ua2f0"
}
```

其中，JWKS URL可以从服务发现地址中获取。

```http
https://dev-5kzk7gzc.us.auth0.com/.well-known/openid-configuration
```

## OIDC使用

### 获取Token

获取Token有两种方式：id\_token Flow和SDK。

{% hint style="info" %}
<mark style="color:purple;">区别是id\_token Flow需要依赖OIDC的登录页，而SDK获取可以自定义登录页！</mark>
{% endhint %}

#### id\_token Flow

以Authing为例，构建如下 id\_token flow URL。

{% code overflow="wrap" %}
```http
# 授权地址，与授权码模式类似
https://xxxx.authing.cn/oidc/auth?
# 客户端ID，即应用ID
client_id=644512df9e360e3f7a40e1e4&
# 重定向URL,授权完成后，跳转的页面
redirect_uri=https://example.com&
scope=openid%20profile&
# 授权类型，注意与授权码模式不同，非常重要
response_type=id_token%20token&
state=6223573295&nonce=1831289
```
{% endcode %}

访问后，进入<mark style="color:red;">登录页</mark>，输入账号密码登录。

登录成功后，将跳转到如下链接：

```http
https://example.com/#
# 授权令牌
access_token=part1.part2.part3&
# 令牌类型，Bearer
token_type=Bearer&
# 令牌过期时间
expires_in=1209600&
scope=openid%20profile&
# id令牌
id_token=part1.part2.part3&
state=6223573295
```

将 `access_token` 保存到客户端使用，例如：`local storage`。

#### 登录接口

下面是Authing用账户密码登录的HTTP请求，详情可参考[Authing的SDK](https://docs.authing.cn/v2/reference/sdk-for-node/authentication/AuthenticationClient.html#%E4%BD%BF%E7%94%A8%E7%94%A8%E6%88%B7%E5%90%8D%E7%99%BB%E5%BD%95)。

```bash
curl --location --request POST 'https://xxxx.authing.cn/api/v3/signin' \
# 用户池ID
--header 'x-authing-userpool-id: 64452248da5930463241789f' \ 
# 应用ID
--header 'x-authing-app-id: 644513df9e260e3f7a10e3e4' \
--header 'Content-Type: application/json' \
--header 'Accept: */*' \
--header 'Connection: keep-alive' \
--data-raw '{
    "connection":"PASSWORD", # 登录模式
    "passwordPayload":{
        "password":"123456", # 密码
        "phone":"18189156130" # 账户
    }
}'
```

```json
{
    "statusCode": 200,
    "message": "",
    "data": {
        "scope": "profile openid tenant_id",
        "token_type": "Bearer",
        "access_token": "part1.part2.part3", // 保存使用
        "expires_in": 1209600,
        "id_token": "part1.part2.part3"
    }
}
```

{% hint style="info" %}
不同的OIDC供应商，实现不同！
{% endhint %}

Fireboom官方也开源一个简单的OIDC服务，当前仅支持隐式模式，详情请查看：[fb-oidc](https://github.com/fireboomio/fb-oidc)

### 使用Token

客户端需要向请求中添加以下请求头：

```http
Authorization: Bearer <access_token>
```

## 工作原理

接下来，我们学习下OIDC协议隐式模式的工作原理。

### 正常流程

<figure><img src="../../.gitbook/assets/image (4).png" alt=""><figcaption><p>隐式模式时序图</p></figcaption></figure>

该时序图有3个主体：客户端、Fireboom服务器、OIDC供应商 authing。

#### 1.客户端到OIDC(Authing)

在客户端构造下述URL，跳转到Authing授权端点。

{% code overflow="wrap" %}
```http
# Authing 授权端点
https://xxx.authing.cn/oidc/auth?
# 客户端ID,用于区分是哪个应用
client_id=xxxx&
# 重定向URL，当前客户端的前端页面(不是Fireboom的地址)
redirect_uri=[https://example.com]&
# 授权范围
scope=openid profile&
# 授权类型，这是与授权码模式的区别之处
response_type=id_token token&
state=xxx
```
{% endcode %}

#### 2.跳转到OIDC登录页

若用户未登录，则跳转到authing的登录页，支持账户密码、手机验证码，甚至是社交登录，如微信、QQ等。

![](../../.gitbook/assets/image.png)

#### 3.OIDC跳转到客户端

登录成功后，authing再跳转到客户端回调地址上，同时携带授权码`access_token`和`id_token`。

和授权码模式不同这里拿到的不是code，且此时未经过飞布服务。

```http
https://example.com/#
# 授权令牌
access_token=part1.part2.part3&
# 令牌类型，Bearer
token_type=Bearer&
# 令牌过期时间
expires_in=1209600&
scope=openid%20profile&
# id令牌
id_token=part1.part2.part3&
state=6223573295
```

#### 4.客户端保存令牌

客户端保存`access_token`及`expires_in` ，供下次使用，一般存储到`local storage`中。

#### 5.客户端获取OIDC用户

客户端使用`access_token`，访问OIDC的用户信息端点，获得当前用户。

```bash
curl --location --request GET 'https://xxx.authing.cn/oidc/me' \
--header 'Authorization: Bearer <access_token>' \
--header 'Accept: */*' \
```

```json
{
    "name": "anson",
    "given_name": null,
    "middle_name": null,
    "family_name": null,
    "nickname": "18189156130",
    "preferred_username": "18189156130",
    "profile": null,
    "picture": null,
    "website": null,
    "birthdate": "2023-06-29T12:47:00.173Z",
    "gender": "male",
    "zoneinfo": null,
    "locale": null,
    "updated_at": "2023-08-09T09:07:37.988Z",
    "tenant_id": "",
    "sub": "645b472d9df8868fe52074d6"
}
```

总结一下，1-5步骤，都不需要飞布服务参与，完全可以由客户端和OIDC服务器完成。

和授权码模式最大的不同是，隐式模式不需要client secret，也不需要使用 code 换 token，更无需请求 token 端点，access\_token 和 id\_token 会直接从授权端点返回。

#### 6.客户端请求Fireboom API

假设，有如下OPERATION，使用 `@fromClaim` 指令修饰，意思是根据当前登录用户的UID，查询待做事项列表：

{% code title="FromClaim.graphql" %}
```graphql
query MyQuery($uid: String! @fromClaim(name: USERID)) {
  todo_findManyTodo(where: {uid: {equals: $uid}}) {
    id
    title
  }
}
```
{% endcode %}

其将被编译为如下REST API，请求如下：

```bash
curl -X GET "http://localhost:9991/operations/FromClaim" \
--header 'Authorization: Bearer <access_token>' \
 -H "accept: application/json" \
```

```json
{
  "data": {
    "todo_findUniqueClaim": {
      "email": "test@example.com",
      "emailVerified": false,
      "location": "",
      "name": "test@example.com",
      "nickname": "test",
      "provider": "auth0",
      "roles": null,
      "userId": "auth0|637b4ab0c0cb508c49de7cf3"
    }
  }
}
```

#### 7.Fireboom校验令牌

Fireboom 根据 [配置的JWK公钥](yin-shi-mo-shi.md#oidc-pei-zhi) 验签`access_token`。该过程不需要OIDC服务参与，因此也意味着OIDC中的Token黑名单不会生效。需要使用Fireboom的 [#wang-guan-gou-zi](../../jin-jie-gou-zi-ji-zhi/operation-gou-zi.md#wang-guan-gou-zi "mention") 自行实现黑名单。

#### 8.客户端获取Fireboom用户

若Fireboom服务未缓存用户信息，则会用`access_token`请求OIDC用户端点，获得用户，同 [#5.-ke-hu-duan-huo-qu-oidc-yong-hu](yin-shi-mo-shi.md#5.-ke-hu-duan-huo-qu-oidc-yong-hu "mention")。

#### 9.Fireboom调用钩子

Fireboom 调用授权钩子，在钩子中修改用户信息，返回后，由Fireboom 缓存用户信息，详情见 [shen-fen-yan-zheng-gou-zi.md](../../jin-jie-gou-zi-ji-zhi/shen-fen-yan-zheng-gou-zi.md "mention")，并返回用户信息到客户端。



### 简化流程

刚刚介绍的，无论是授权码模式还是隐式模式都需要使用OIDC原生的登录页。那如果想自定义登录页，又该如何做呢？

其实只要基于隐式模式稍作调整，就能支持该需求。我们知道，隐式模式的本质是获得`access_token`，然后用其请求接口。

那如果OIDC服务提供了根据账户密码获得`accest_token`的接口，是不是就可以供我们在自己编写的界面中直接调用了呢？

<figure><img src="../../.gitbook/assets/image (5).png" alt=""><figcaption><p>隐式模式时序图-简化版</p></figcaption></figure>

我们仍以authing为例，上述流程图就展示了该过程。

#### 1.调用登录接口

通过post请求，输入账户、密码或手机号、验证码，获得`acces_token`。不同OIDC供应商有对应实现，详情可参考其文档。例如前文中提到的 [#deng-lu-jie-kou](yin-shi-mo-shi.md#deng-lu-jie-kou "mention")

#### 2.客户端存储令牌

客户端存储 `acces_token`，供后续使用。

其他流程没有变化，同 [#zheng-chang-liu-cheng](yin-shi-mo-shi.md#zheng-chang-liu-cheng "mention") 的步骤6-9。

### 总结

总结一下，隐式模式的流程更加简洁。

由于其不依赖cookie，所以适用范围更广，不仅支持浏览器，也能支持类似微信小程序或者原生APP。

但缺点是安全性不足，因为其直接将`access_token`泄露给了客户端。
