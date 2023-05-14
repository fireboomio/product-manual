# 身份验证

飞布支持OIDC进行身份验证，实现了OIDC中定义的两种授权流程：基于cookie登录-授权码模式（Authorization Code）和基于Token登录-隐式模式（Implicit）。

## 支持OIDC Provider

飞布能与任意实现OIDC规范的供应商集成。目前主流OIDC供应商如下：

* IDaaS服务商auth0、authing等
* 开源OIDC服务：okta、[casdoor](../../huan-jing-zhun-bei.md#zi-bu-shu-casdoor)等。

接下来，我们学习下如何配置，我们以AUTHING为例。

{% embed url="https://www.bilibili.com/video/BV1wk4y1i7rD/?share_source=copy_web&vd_source=f8709d15baaa835ea2d0bb3bcc6857da" %}
15功能介绍-飞布如何添加身份验证器？
{% endembed %}



## 快速操作

### 基本设置

1. 在身份验证面板中点击“+”，进入OIDC新建页
2. 首先，输入供应商名称，自动生成 “登录回调 URL”。
3. 然后，前往AUTHING应用配置页，获取APP ID 、App Secret和Issuer ，分别填入身份验证器表单。
4. 随后，输入APP ID。
5. 接着，输入Issuer。输入后，系统自动生成服务发现地址，并从中获取jwksURL和用户端点
6. 当开发WEB应用时，开启基于cookie的模式，同时填入App Secret。
7. 当开发移动应用时，开启基于Token的模式。

若JwksURL无法访问，可将JWKS切换到JSON模式，然后输入JSON字符串。

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

8. 最后，保存表单，完成配置。

值得注意的是，基于Cookie的登录，需要OIDC供应商（OIDC Provider）和飞布服务器后端（Relying-Party）同时配置回调地址。

首先是，复制 登录回调 URL， 前往AUTHING设置“登录回调 URL”，多个URL可用"英文逗号"分开。

接着，点击“配置登录回调”按钮，前往"设置->安全"，设置 "重定向URL"。

系统提供了两个默认值：

* localhost:9123/#/workbench/userInfo：用户详情页回调URL，用于测试OIDC
* localhost:9123/#/workbench/rapi/loginBack：API预览页回调URL，用于测试需要授权的API接口。

后续，可根据集成的前端项目，添加对应URL。

回到详情页，点击右上角“测试”按钮。跳转至authing提供的登录页，登录后，可查看当前用户信息。

{% hint style="info" %}
只有开启“基于Cookie”模式后，才能直接测试。
{% endhint %}

出于安全考虑，在回调至OIDC供应商URL时，系统会自动跳转到HTTPS链接。若想关闭该功能，可关闭 “强制 HTTPS 跳转”。

### API设置

1. 前往API管理面板，选择需要设置的API
   1. 登录访问：切换到设置面板，开启授权，限制API必须登录才能访问
   2. 数据权限：用@fromClaim修饰入参，限制API的数据权限
2. 点击顶部菜单栏的“<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAMAAAC7IEhfAAAAY1BMVEUAAADU1NRjZmxvcnePkZVgY2rPz9BfY2poaHTAwMOAhIxoa3FgYmpgYmlgY2nFxcbU1NRmZm+Ag427u73U1NRfYml/g4zt7e3k5eXW1te9vsCztLefoaWanKCOkJVydXtucXecDQKGAAAAFXRSTlMAzP336NDOiAvTz/rn2tjSph7Qs6d9epWLAAAAjElEQVQ4y+2T2Q6EIAxFK+A6mzMj4q7//5VaYngCG2N8cDkvNOlJSG9TuCq+XMQ3oiQ4p0jGsx+/fCIByDwrqRFzDYDn4BatYiw4Y1zEhBgIJjUsjJbED5eG19ctBtrr66rD9x05RYH9oVBKtViFTvGB7UZNlFg9N4n01/QwdDwrA0/mU0jtK/zDYRgBwgsrsPomQg4AAAAASUVORK5CYII=" alt="预览" data-size="line">”，前往API预览页，选择当前API
3. 输入参数，测试接口，你会发现，接口返回401（这是因为你没登录）
4. 在预览页顶部，选择OIDC供应商，点击前往登录，登录后可查看用户信息
5. 重复步骤3，可以看到接口执行成功，有三种情形
   1. 登录访问：未限制数据权限，正常执行，唯一区别是需要用户登录才能执行
   2. 查询请求：限制数据权限，只返回当前用户拥有的数据
   3. 变更请求：限制数据权限，插入数据时绑定当前用户的标识，如UID或EMAIL等

## 客户端如何使用

### 基于COOKIE登录

构建如下URL，在网页上点击跳转即可。

http://localhost:9991/api/auth/cookie/authorize/<mark style="color:purple;"><供应商ID></mark>?redirect\_uri=<mark style="color:purple;"><当前页URL></mark>

* 供应商ID：对应OIDC表单中的供应商ID
* 当前页URL：对应"设置->安全"中的"重定向URL"

### 基于TOKEN登录

客户端需要向请求中添加以下请求头：

```
Authorization: Bearer <token>
```

如何获取TOKEN，可参考[Authing的SDK](https://docs.authing.cn/v2/reference/sdk-for-node/authentication/AuthenticationClient.html#%E4%BD%BF%E7%94%A8%E7%94%A8%E6%88%B7%E5%90%8D%E7%99%BB%E5%BD%95)。

## 工作原理

### 专业术语

1. `EU（End-User）`：终端用户
2. `RP（Relying-Party）`：服务器后端，这里指飞布服务器后端
3. `OP（OIDC Provider）`： 提供身份验证的服务器，例如Authing 服务器

### 基于cookie登录-授权码模式

> 授权码模式是 OIDC 授权登录中最常用的模式，OP 服务器返回一个授权码 code 给开发者后端服务器，在后端完成 code 换取 access\_token，再用 access\_token 换取用户信息的操作，从而实现用户的身份认证。

#### 1. 发起登录请求

发起授权需要拼接一个用来授权的 URL，并让终端用户在浏览器中访问，具体参数如下：

```
<Issuer>/auth
 ?response_type=code
 # RP身份标识，对应App ID
 &client_id=29352915982374239857  
 # 授权服务器接收请求后返回给浏览器的跳转访问地址，对应 设置->安全->重定向URL
 &redirect_uri=https%3A%2F%2Fexample-client.com%2Fcallback  
 &scope=create+delete
 &state=xcoiv98y2kd22vusuye3kch
```

{% hint style="info" %}
```
Issuer：OP服务的连接，例如：https://<应用域名>.authing.cn/oidc/
```
{% endhint %}

#### 2.用户登录

发起 OIDC 登录之后，如果用户先前**未在 OP 登录过**，OP 会将用户重定向到登录页面，引导用户完成在 OP 的认证，此时用户需要选择一种方式进行登录：

![authing登录页](https://cdn.authing.cn/blog/20200927203336.png)

#### 3.获取code

OP将验证此用户是否合法，验证通过后会将浏览器重定向到**发起授权登录请求时指定**的 **redirect\_uri** 并通过 URL query 传递授权码 code 参数。

```
https://example-client.com/redirect
 ?code=g0ZGZmNjVmOWIjNTk2NTk4ZTYyZGI3  # 授权码，用于换取access_token，用一次就失效
 &state=xcoiv98y2kd22vusuye3
```

#### 4.使用code换取token

飞布默认换取 token 身份验证方式为 client\_secret\_post，需要向`token_endpoint`发送POST请求，具体如下：

```json
# POST <token_endpoint> // 
{
    code:"g0ZGZmNjVmOWIjNTk2NTk4ZTYyZGI3" # 授权码
    client_id: "29352915982374239857", # RP身份标识，对应App ID
    client_secret: "xxxx", # RP密钥，对应App Secret。
    grant_type: "authorization_code",# 指定RP正在使用的授权流程。
    # 与请求authorization code时使用的redirect_uri相同。
    redirect_uri:"https%3A%2F%2Fexample-client.com%2Fcallback ",
}
```

{% hint style="info" %}
`token_endpoint一般从服务发现地址中获取，格式：<`Issuer`>/.well-known/openid-configuration`
{% endhint %}

#### 5.签发访问令牌

OP将会验证第4步中的请求参数，当验证通过后（校验`authorization code`是否过期，`client id`和`client secret`是否匹配等），OP将向`RP`返回`access token`。

```json
{
    # 访问令牌
    "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjEifQ.eyJqdGkiOiJQZU41YXg1b3FabGRhcUJUMzQzeUkiLCJzdWIiOiI1Y2U1M2FlYTlmODUyNTdkZDEzMmQ3NDkiLCJpc3MiOiJodHRwczovL29hdXRoLmF1dGhpbmcuY24vb2F1dGgvb2lkYyIsImlhdCI6MTU4MTQyMDk1NywiZXhwIjoxNTgxNDI0NTU0LCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGF1dGhpbmdfdG9rZW4gZW1haWwgcGhvbmUgYWRkcmVzcyBvZmZsaW5lX2FjY2VzcyIsImF1ZCI6IjVkMDFlMzg5OTg1ZjgxYzZjMWRkMzFkZSJ9.rtpRSL3_U03zXShZUCILquSR_KEDuS-OldWpy8RLztWUNG_tMyrg2g9CG4hC7pJUwmgzZKtp7vsVrj6W0eyo_ehE4KGz9iKnyd46DFbx9W9pi-mieRW5HuVMGL2zvDH8zF467WXET2SVB3LUhFLNmEbxpvjPZ5Ksvbcd7nqHfnUN4-z3SqAvhGWWfcmt7QDFlLtWPw4LzyznEqmM9sDkNiNDnTkjmcjm7yHJR-yv5FvpzQB2kraQVOrrdAixbHf29ihOVO25CrjmgeKemg1vuLNGUcOrr_XWn7xaCSvyAfXrBuRalecW9RA4p_Cp6YslHc_572awekt3kUO2TebUQA",
    "expires_in": 3597,
    "id_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjEifQ.eyJzdWIiOiI1Y2U1M2FlYTlmODUyNTdkZDEzMmQ3NDkiLCJiaXJ0aGRhdGUiOiIiLCJmYW1pbHlfbmFtZSI6IiIsImdlbmRlciI6IiIsImdpdmVuX25hbWUiOiIiLCJsb2NhbGUiOiIiLCJtaWRkbGVfbmFtZSI6IiIsIm5hbWUiOiIiLCJuaWNrbmFtZSI6IiIsInBpY3R1cmUiOiJodHRwczovL3VzZXJjb250ZW50cy5hdXRoaW5nLmNuL2F1dGhpbmctYXZhdGFyLnBuZyIsInByZWZlcnJlZF91c2VybmFtZSI6IiIsInByb2ZpbGUiOiIiLCJ1cGRhdGVkX2F0IjoiIiwid2Vic2l0ZSI6IiIsInpvbmVpbmZvIjoiIiwiY29tcGFueSI6IiIsImJyb3dzZXIiOiIiLCJsb2dpbnNfY291bnQiOjEwMywicmVnaXN0ZXJfbWV0aG9kIjoiZGVmYXVsdDp1c2VybmFtZS1wYXNzd29yZCIsImJsb2NrZWQiOmZhbHNlLCJsYXN0X2lwIjoiMTIxLjIxLjU2LjE3MSIsInJlZ2lzdGVyX2luX3VzZXJwb29sIjoiNWM5NTkwNTU3OGZjZTUwMDAxNjZmODUzIiwibGFzdF9sb2dpbiI6IjIwMjAtMDItMTFUMTE6MzU6MTUuNjk2WiIsInNpZ25lZF91cCI6IjIwMTktMDUtMjJUMTI6MDQ6NTguMjk0WiIsInRva2VuIjoiZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SmtZWFJoSWpwN0ltVnRZV2xzSWpvaWRHVnpkRE5BTVRJekxtTnZiU0lzSW1sa0lqb2lOV05sTlROaFpXRTVaamcxTWpVM1pHUXhNekprTnpRNUlpd2lZMnhwWlc1MFNXUWlPaUkxWXprMU9UQTFOVGM0Wm1ObE5UQXdNREUyTm1ZNE5UTWlmU3dpYVdGMElqb3hOVGd4TkRJd09URTFMQ0psZUhBaU9qRTFPREkzTVRZNU1URjkuM0l0X0NJQTNFbUpoYWcyMW92WjNwd0RfY0owcTVTZkJjSURSZThRX3FoayIsImVtYWlsIjoidGVzdDNAMTIzLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicGhvbmVfbnVtYmVyIjoiMTMxMTIzNDEyMzQiLCJhZGRyZXNzIjoiIiwiYXRfaGFzaCI6IjV6QnNUOHF4RHc1RmNYdU55UFg4YUEiLCJzaWQiOiJkNmZiOTE5Ny00NmE3LTQ1ZGEtOGVkMC05ODhjZjg0ZjQwZWUiLCJhdWQiOiI1ZDAxZTM4OTk4NWY4MWM2YzFkZDMxZGUiLCJleHAiOjE1ODE0MjQ1NTQsImlhdCI6MTU4MTQyMDk1NywiaXNzIjoiaHR0cHM6Ly9vYXV0aC5hdXRoaW5nLmNuL29hdXRoL29pZGMifQ.VZzqULytIteyBfouww5TsHQ50gEhM06kUWMeDiO3FVFSCW9ys2bFPos5p6LFzliK4Ce09ypOwVQiRnE2gNYsukLvlUPlKDIP_Xk5W19frKi1Z8ImuIPvUqVMKbFutVNS0TfIPCPJVBl8C1j5OXeIs6z0V90QrvyJao6FqVEa3axOHxbhpo1fH2hP04-wkGOp_l10d7RFhGcnPyPnz9-C5X6A4UEsCSDCVw1mDQHxDSFP9OPaB_OlCG_Bi6G-CeLhPa3V5hyIefdBvxC9SIpK-6qY-_BfsNKkBHDVKMb0xodgN2hzn3UTUGBuuoiaB4JhCv72EZ7eiXKIXFz6zVcogA",
    "refresh_token": "DuSPlrUFPAvCZ1WQKarv5MbEsXN",
    "scope": "openid profile authing_token email phone address offline_access",
    "token_type": "Bearer"
}
```

### 基于Token登录-隐式模式

> OIDC 隐式模式不会返回授权码 code，而是直接将 `access_token` 和 `id_token` 通过 **URL hash** 发送到**回调地址前端**，**后端无法获取**到这里返回的值，因为 URL hash 不会被直接发送到后端。该模式常结合移动应用或 Web App 使用。

#### 1. 用户授权请求

发起隐式模式的授权登录**需要拼接一个 URL**，并让终端用户在浏览器中访问，**不能直接输入**认证地址域名。具体参数如下：

<pre><code>&#x3C;authorization_endpoint>
# RP身份标识，对应App ID
<strong>?client_id=0oabv6kx4qq6h1U5l0h7
</strong>&#x26;response_type=token # 为token 或 id_token
# 回调链接，用户在 OP 认证成功后，OP 会将 id_token、access_token 以 URL hash 的形式发送到这个地址。
&#x26;redirect_uri=http%3A%2F%2Flocalhost%3A8080
&#x26;state=state-296bc9a0-a2a2-4a57-be1a-d0e2fd9bb601
<strong>&#x26;nonce=foo
</strong></code></pre>

{% hint style="info" %}
authorization\_endpoint`一般从服务发现地址中获取，常见格式：<`Issuer`>/`authorize

`服务发现地址：<`Issuer`>/.well-known/openid-configuration`
{% endhint %}

#### 2.用户授权应用（略）

#### 3.用访问令牌重定向URI

假设用户授予访问权限，跳转后链接示例：

```
http://localhost:8080/
# 访问令牌, 以 URL hash 形式传递
#access_token=eyJhb[...]erw
# 当且仅当response_type设置为 token 时返回，值恒为 Bearer
&token_type=Bearer 
&expires_in=3600
&scope=openid
&state=state-296bc9a0-a2a2-4a57-be1a-d0e2fd9bb601
```

> 为什么信息在 URL hash 里而不是 query 里？因为 hash 内容不会直接发送到服务器，避免 id\_token、access\_token 被盗用。

{% hint style="info" %}
1-3步骤为标准流程，不同客户端获取`access_token 的流程不同，需要根据实际情况处理。从工程实践中看，常用的另一种方式是直接调用OIDC供应商的登录接口，从中获取access_token。`
{% endhint %}

#### 4.传递给应用程序的访问令牌

浏览器向RP发送`access token`。RP采用两种方式校验令牌：

* 公钥签名校验：优先使用使用**公钥**验证签名。公钥地址（jwks\_uri）一般为： `<Issuer>/.well-known/jwks.json` 。
* **在线接口校验：**若公钥验签失败，则调用供应商的<mark style="color:red;">token验证接口</mark>进行在线验证。userinfo\_endpoint?

{% hint style="info" %}
**参考链接**

1. JWKS [参考规范 (opens new window)](https://openid.net/specs/openid-connect-registration-1\_0.html#ClientMetadata)；
2. 可以在线检验 JWT 的签名的网站：[https://jwt.io (opens new window)](https://jwt.io/)；
3. RSA 公私钥 PEM 格式 与 JWK 格式互转：[https://8gwifi.org/jwkconvertfunctions.jsp (opens new window)](https://8gwifi.org/jwkconvertfunctions.jsp)；
4. 生成 JWK：[https://mkjwk.org (opens new window)](https://mkjwk.org/)；
{% endhint %}

### 获取用户信息 <a href="#06-shi-yong-accesstoken-huan-qu-yong-hu-xin-xi" id="06-shi-yong-accesstoken-huan-qu-yong-hu-xin-xi"></a>

直到`access token` 过期或失效之前，`RP`使用access\_token，通过OP的`userinfo_endpoint` API换取用户信息。如果发起授权登录时的 scope 参数不同，这里的返回信息也会不同，返回信息中的字段取决于 scope 参数。字段符合 [OIDC 规范 (opens new window)](https://openid.net/specs/openid-connect-core-1\_0.html#AuthorizationExamples)，用户信息字段与 scope 对应关系请参考 [scope 参数对应的用户信息](https://old-docs.authing.cn/authentication/oidc/oidc-params.html#scope-%E5%8F%82%E6%95%B0%E5%AF%B9%E5%BA%94%E7%9A%84%E7%94%A8%E6%88%B7%E4%BF%A1%E6%81%AF)。

具体请求如下：

```
GET <userinfo_endpoint>?access_token=<access_token>
```

{% hint style="info" %}
```
userinfo_endpoint一般从服务发现地址中获取，常见格式：<Issuer>/userinfo
```
{% endhint %}

返回值示例：

```json
{
  "sub": "5f7174df27e0eb9c6d21436d",
  "birthdate": null,
  "family_name": null,
  "gender": "U",
  "given_name": null,
  "locale": null,
  "middle_name": null,
  "name": null,
  "nickname": null,
  "picture": "https://usercontents.auth0.cn/avatar.png",
  "preferred_username": null,
  "profile": null,
  "updated_at": "2020-09-28T05:33:15.892Z",
  "website": null,
  "zoneinfo": null
}
```

{% hint style="info" %}
OIDC规范中不包含角色的描述，因此返回值不涉及`roles`字段
{% endhint %}
