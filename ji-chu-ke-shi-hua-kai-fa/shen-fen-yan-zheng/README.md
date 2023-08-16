# 身份验证

限制接口只能被登录用户访问，且只能获取当前用户所拥有的数据行或字段，是WEBAPI开发的常见需求。

飞布通过自定义GraphQL指令：`@fromClaim`，结合OIDC协议，实现了API数据权限控制。

飞布支持OIDC进行身份验证，实现了OIDC中定义的两种授权流程：

* 基于cookie登录-授权码模式（Authorization Code）
* 基于Token登录-隐式模式（Implicit）

两种模式适用场景不同。
