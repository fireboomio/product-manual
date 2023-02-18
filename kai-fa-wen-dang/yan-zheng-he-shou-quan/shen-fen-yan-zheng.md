# 身份验证

飞布支持OIDC进行身份验证，实现了OIDC中定义的两种Grant Type：Authorization Code和Implicit。

* Authorization Code：结合普通服务器端应用使用。

> &#x20;[`Authorization Code`](https://www.cnblogs.com/Wddpct/p/8976480.html#5-%E6%8E%88%E6%9D%83%E8%AE%B8%E5%8F%AFauthorization-grant)是最常使用的一种授权许可类型，它适用于第三方应用类型为`server-side`型应用的场景。它的授权流程基于重定向跳转，客户端必须能够与`User-agent`（即用户的 Web 浏览器）交互并接收通过`User-agent`路由发送的实际`authorization code`值。

* Implicit：结合移动应用或 Web App 使用。

> ``[`Implicit`](https://www.cnblogs.com/Wddpct/p/8976480.html#5-%E6%8E%88%E6%9D%83%E8%AE%B8%E5%8F%AFauthorization-grant)授权流程和`Authorization Code`基于重定向跳转的授权流程十分相似，但它适用于移动应用和 Web App，这些应用与普通服务器端应用相比有个特点，即`client secret`不能有效保存和信任。
>
> 相比`Authorization Code`授权流程，`Implicit`去除了请求和获得`authorization code`的过程，而用户点击授权后，授权服务器也会直接把`access token`放在`redirect_uri`中发送给`User-agent`（浏览器）。 同时构造请求用户授权 url 中的`response_type`参数值也由 _code_ 更改为 _token_ 或 _id\_token_ 。

在飞布系统中，我们分别称作：基于cookie登录和基于TOKEN登录。

## 支持OIDC Provider

飞布能与任意实现OIDC规范的供应商集成。目前主流OIDC供应商如下：

<mark style="color:red;">补充内容</mark>



接下来，我们学习下如何配置。

## 快速操作







## 客户端如何使用？

### 基于COOKIE登录





### 基于TOKEN登录









