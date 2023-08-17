# 授权码模式

授权码模式又被称为基于cookie的模式，适用于具有后端服务器的场景，要求应用必须能够安全存储密钥，用于后续使用授权码code换 Access Token。

> 授权码模式是 OIDC 授权登录中最常用的模式，OP 服务器返回一个授权码 code 给开发者后端服务器，在后端完成 code 换取 access\_token，再用 access\_token 换取用户信息的操作，从而实现用户的身份认证。

## OIDC配置

首先，我们学习如何在飞布中配置OIDC。

OIDC是行业内的通用规范，飞布能与任意实现OIDC规范的供应商集成，例如：

* IDaaS服务商：auth0、authing等
* 开源OIDC服务：okta、casdoor等

{% embed url="https://www.bilibili.com/video/BV1wk4y1i7rD/?share_source=copy_web&vd_source=f8709d15baaa835ea2d0bb3bcc6857da" %}
15功能介绍-飞布如何添加身份验证器？
{% endembed %}



如图，是飞布中OIDC配置页和auth0应用详情页的参数对应图。

<figure><img src="../../../.gitbook/assets/image (52).png" alt=""><figcaption></figcaption></figure>

我们逐项学习。

1，供应商名称：供应商的值，会影响下述[步骤1中的跳转地址](./#1.-ke-hu-duan-dao-fireboom)和[步骤4的回调地址](./#4.oidc-tiao-zhuan-fireboom-hui-tiao-di-zhi)，例如auth0

在本图中对应点击登录按钮后跳转的地址和这里设置的登录回调地址。

2，App ID和App Secret：对应[步骤5](./#5.fireboom-huo-qu-oidc-ling-pai)中的appid 和secret。

分别从auth0的client id和client secret中获取

3，issuer：对应auth0的domain地址。

一旦issuer确定，服务发现地址也就确定了，即在该domain后增加.well-known/openid-configuration。

```http
https://dev-5kzk7gzc.us.auth0.com/.well-known/openid-configuration
```

服务发现地址中涵盖了很多信息，包括[步骤2](./#2.fireboom-dao-oidc-shou-quan-duan-dian)中的授权端点、[步骤5](./#5.fireboom-huo-qu-oidc-ling-pai)中的token端点、[步骤6](./#6.fireboom-huo-qu-yong-hu-xin-xi)中的用户端点等都是从这里取到的。

```json
{
    "issuer": "https://dev-5kzk7gzc.us.auth0.com/",
    // 步骤2中的授权端点
    "authorization_endpoint": "https://dev-5kzk7gzc.us.auth0.com/authorize",
    // 步骤5中的token端点
    "token_endpoint": "https://dev-5kzk7gzc.us.auth0.com/oauth/token",
    "device_authorization_endpoint": "https://dev-5kzk7gzc.us.auth0.com/oauth/device/code",
    // 步骤6中的用户端点
    "userinfo_endpoint": "https://dev-5kzk7gzc.us.auth0.com/userinfo",
    "mfa_challenge_endpoint": "https://dev-5kzk7gzc.us.auth0.com/mfa/challenge",
    # 这个值隐式模式需要用！！！
    "jwks_uri": "https://dev-5kzk7gzc.us.auth0.com/.well-known/jwks.json",
    "registration_endpoint": "https://dev-5kzk7gzc.us.auth0.com/oidc/register",
    "revocation_endpoint": "https://dev-5kzk7gzc.us.auth0.com/oauth/revoke",
    "scopes_supported": [
        "openid",
        "profile",
        "offline_access",
        "name",
        "given_name",
        "family_name",
        "nickname",
        "email",
        "email_verified",
        "picture",
        "created_at",
        "identities",
        "phone",
        "address"
    ],
    "response_types_supported": [
        "code",
        "token",
        "id_token",
        "code token",
        "code id_token",
        "token id_token",
        "code token id_token"
    ],
    "code_challenge_methods_supported": [
        "S256",
        "plain"
    ],
    "response_modes_supported": [
        "query",
        "fragment",
        "form_post"
    ],
    "subject_types_supported": [
        "public"
    ],
    "id_token_signing_alg_values_supported": [
        "HS256",
        "RS256"
    ],
    "token_endpoint_auth_methods_supported": [
        "client_secret_basic",
        "client_secret_post",
        "private_key_jwt"
    ],
    "claims_supported": [
        "aud",
        "auth_time",
        "created_at",
        "email",
        "email_verified",
        "exp",
        "family_name",
        "given_name",
        "iat",
        "identities",
        "iss",
        "name",
        "nickname",
        "phone_number",
        "picture",
        "sub"
    ],
    "request_uri_parameter_supported": false,
    "request_parameter_supported": false,
    "token_endpoint_auth_signing_alg_values_supported": [
        "RS256",
        "RS384",
        "PS256"
    ]
}
```

4，接着，开启基于cookie的登录开关，开启后才能填写secret，将在[步骤5](./#5.fireboom-huo-qu-oidc-ling-pai)的token端点中使用。

5，登录回调地址，需要填写到auth0的允许回调地址列表中，[步骤4](./#4.oidc-tiao-zhuan-fireboom-hui-tiao-di-zhi)需要用到

<mark style="color:red;">若不填写，将会有如下报错。</mark>

<figure><img src="../../../.gitbook/assets/image (1) (1) (1) (1).png" alt=""><figcaption></figcaption></figure>

6，最后，还有一个地方要记得设置：设置->安全 中的 重定向URL 白名单。

需要填入客户端的网址，如：

```
http://localhost:9173
```

设置后，[步骤8](./#8.fireboom-zhu-ru-cookie-dao-ke-hu-duan) 才能生效，不然Fireboom无法为其注入Cookie。即允许哪些页面中可以发起下述1-8的cookie授权流程。

系统提供了两个默认值：

* localhost:9123/#/workbench/userInfo：用户详情页回调URL，用于测试OIDC
* localhost:9123/#/workbench/rapi/loginBack：API预览页回调URL，用于测试需要授权的API接口。

本项设置，主要是出于安全考虑。

## OIDC使用

设置完成后，我们学习如何在网页中使用，例如，当前网页是：

&#x20;`http://localhost:9173`。

### 登录

使用登录的话，只需要构建链接：

```html
<a href="http://localhost:9991/auth/cookie/authorize/auth0?redirect_uri=http://localhost:5173/">
   login
</a>
```

* Fireboom访问域名：http://localhost:9991
* 授权端点：/auth/cookie/authorize/auth0 ，其中auth0和供应商名称一致
* 重定向URL：http://localhost:5173/，其值为当前网页的地址

登录完成后，当前网页中会注入cookie。

### 获取登录用户

此时带着cookie访问user端点，可以拿到当前用户。

```http
http://localhost:9991/auth/cookie/user
```

### 退出登录

如果要退出登录，需要带上cookie，访问退出登录端点，就会清空客户端与Fireboom之间的cookie，退出登录。

```http
http://localhost:9991/auth/cookie/user/logout
```

但此时，不会退出客户端与auth0之间的登录态，如若退出，需要带上参数：`logout_openid_connect_provider`。

```http
http://localhost:9991/auth/cookie/user/logout?logout_openid_connect_provider=true
```

此刻，不仅退出客户端与飞布之间的登录，而且会调用退出端点，退出客户端与auth0的登录态。

### 快速测试

上述介绍的功能都已经在Fireboom集成。

1，保存后，回到详情页，点击右上角“测试”按钮，跳转至auth0提供的登录页。

<figure><img src="../../../.gitbook/assets/image (3) (1) (1).png" alt=""><figcaption></figcaption></figure>

2，登录后，可查看当前用户信息。

<figure><img src="../../../.gitbook/assets/image (2) (1) (1).png" alt=""><figcaption></figcaption></figure>

3，点击”登出“，可清空cookie。

{% hint style="info" %}
只有开启“基于Cookie”模式后，才能直接测试。
{% endhint %}



图中也涉及到：基于token 登录，我们后续再介绍。

## 工作原理

### 登录流程

接下来，我们学习下OIDC协议授权码模式的工作原理。

<figure><img src="../../../.gitbook/assets/image (6).png" alt=""><figcaption><p>授权码模式登录流程时序图</p></figcaption></figure>

该时序图有3个主体：

* 客户端：一般是浏览器
* Fireboom服务器：又被称为RP (Relying-Party)
* Auth0：OIDC供应商 OP，实现了OIDC的IDAAS服务。这里只是示例，也可以是其他供应商，如authing等。

#### 1.客户端到Fireboom

在客户端构造下述URL，点击后跳转到Fireboom服务器。

```http
http://localhost:9991/auth/cookie/authorize/[auth0]?redirect_uri=[https://example.com]
```

该URL的组成为 ：

* Fireboom服务器域名：http://localhost:9991
* cookie授权路径：/auth/cookie/authorize/\[auth0]，其中\[auth0]为OIDC提供商名称，在Fireboom控制台定义
* 重定向域名URL：\[https://example.com]，当前客户端的前端页面

#### 2.Fireboom到OIDC授权端点

Fireboom 服务跳转到 OIDC 供应商 auth0 的授权端点，同时后面会携带一些参数，例如：

```http
# auth0的授权端点
https://xxx.auth0.com/authorize? 
# OIDC 提供的客户端ID，用于区分是哪个应用，根据auth0获取
client_id=xxx& 
# 重定向URL，跳转回Fireboom时的地址，根据auth0拼接
redirect_uri=[http://localhost:9991/auth/cookie/callback/auth0]&
# 授权范围，openID和profile（用户个人信息）
scope=openid+profile& 
# 响应类型：code，表示授权码模式
response_type=code& 
state=xxx&
nonce=xxx
```

之所以能跳转到正确页面`redirect_uri`，以及该用哪个`client_id`，是因为第一步中的auth0告诉了Fireboom，该用哪一个供应商的信息。

在配置供应商时，提供了`client_id`，<mark style="color:orange;">详情见</mark> [#oidc-pei-zhi](./#oidc-pei-zhi "mention")

#### 3.跳转到OIDC登录页

若用户未登录，则跳转到auth0的登录页，支持账户密码、手机验证码，甚至是社交登录，如github、谷歌等。

![](<../../../.gitbook/assets/image (7).png>)

#### 4.OIDC跳转Fireboom回调地址

登录成功后，auth0再跳转到Fireboom的回调地址上，同时携带授权码code。

```http
# 对应步骤3中的redirect_uri
http://localhost:9991/auth/cookie/callback/[auth0]?
# OIDC服务auth0提供的授权码，用1次后失效
code=xxxx&
state=xxxx
```

#### 5.Fireboom获取OIDC令牌

Fireboom服务使用授权码code以及auth0的客户端ID `client_id`&客户端秘钥 `secret`，请求auth0 的access\_token端点，获得access token值。

```bash
# access_token 端点
curl --location --request POST 'https://xxx.auth0.com/oauth/token' \
--header 'Accept: */*' \
# 请求类型 x-www-form-urlencoded
--header 'Content-Type: application/x-www-form-urlencoded' \
# 客户端ID，根据auth0获取
--data-urlencode 'client_id=644512df9e36033f7a40e3e4' \
# 客户端Secret，根据auth0获取
--data-urlencode 'client_secret=7fc6623axsd76d3xc7e059f12bf616de' \
# 授权类型，authorization_code表示授权码
--data-urlencode 'grant_type=authorization_code' \
# 授权码，上一步拿到的
--data-urlencode 'code=0lqtCqYyog35H0f-1z8tpNg6XsSOZf3mWKrEKUJJrnp' \
# 重定向URL，步骤1中的的redirect_uri
--data-urlencode 'redirect_uri=https://example.com'
```

```json
{
    "scope": "openid profile offline_access",
    "token_type": "Bearer",
    "access_token": "part1.part2.part3",
    "expires_in": 1209600,
    "id_token": "part1.part2.part3",
    "refresh_token": "xxxx"
}
```

#### 6.Fireboom获取用户信息

Fireboom使用`access_token`请求auth0 的userinfo端点，获得登录用户的信息。

```bash
# userinfo端点
curl --location --request GET 'https://xxx.auth0.com/userinfo' \
--header 'Authorization: Bearer part1.part2.part3' \
--header 'Accept: */*' \
--header 'Connection: keep-alive' \
--header 'Cookie: interaction-oidc-idp=93c3da20-ac2e-4cd5-9c6e-bc8fd0ab52ec'
```

```json
{
    "name": "anson",
    "given_name": null,
    "middle_name": null,
    "family_name": null,
    "nickname": "Anson",
    "preferred_username": "xxxx",
    "profile": null,
    "picture": null,
    "website": null,
    "birthdate": "2023-06-29T12:47:00.173Z",
    "gender": "male",
    "zoneinfo": null,
    "locale": null,
    "updated_at": "2023-08-08T13:10:16.835Z",
    "sub": "645b472d9df8868fe52074d6"
}
```



总结一下，上面6个步骤的主要过程就是：用户在auth0登录后，给了Fireboom一个授权码code，然后，Fireboom 用code换取token，然后用token换登录用户的信息。

#### 7.Fireboom调用授权钩子

Fireboom 调用授权钩子，在钩子中修改用户信息，返回后，由Fireboom 缓存用户信息，详情见 [shen-fen-yan-zheng-gou-zi.md](../../../jin-jie-gou-zi-ji-zhi/shen-fen-yan-zheng-gou-zi.md "mention")

#### 8.Fireboom注入Cookie到客户端

还记得流程的发起方吗？是客户端，例如\[https://example.com]，这时候Fireboom 会把cookie注入到发起方中。

#### 9.客户端获取登录用户

客户端携带Cookie访问Fireboom提供的user端点：

```
[http://localhost:9991]/auth/cookie/user
```

从缓存中拿到当前登录的用户。

若想更新当前用户的信息，可以增加`revalidate=true`后缀，将重新执行步骤6，更新缓存中的用户信息。

```
[http://localhost:9991]/auth/cookie/user?revalidate=true
```

总结下，后半部分的主要目的是，把Fireboom从auth0那里拿到的授权凭证，通过cookie授权给客户端。这样就变相实现了auth0授权给客户端。

之所以折腾这一下，还是出于安全的考虑，因为步骤5需要用到auth0的敏感信息。

### 退出登录流程

接着，我们学习下OIDC授权码模式，退出登录的流程。

<figure><img src="../../../.gitbook/assets/image (53).png" alt=""><figcaption><p>OIDC退出登录流程</p></figcaption></figure>

1，首先，客户端携带cookie访问飞布登出端点，如：

```
[http://localhost:9991]/auth/cookie/user/logout?logout_openid_connect_provider=true
```

2，接着，飞布服务清空当前的用户cookie

3，请求登出钩子，详情见 [shen-fen-yan-zheng-gou-zi.md](../../../jin-jie-gou-zi-ji-zhi/shen-fen-yan-zheng-gou-zi.md "mention")

如果`logout_openid_connect_provider=true`，则执行后续逻辑。

4，接着，请求OIDC的登出端点，可以在服务发现地址中获取，如：

```
https://xxx.authing.cn/oidc/session/end
```

5，返回结果到客户端，一般是 logout\_uri&#x20;

6，如果需要客户端协助，则自动跳转到authing的登出页

## 参考：

* 选择 OIDC 授权模式：[https://docs.authing.co/v2/concepts/oidc/choose-flow.html](https://docs.authing.co/v2/concepts/oidc/choose-flow.html)
