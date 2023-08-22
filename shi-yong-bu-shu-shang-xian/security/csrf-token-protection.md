# CSRF token 保护

在有用户身份登录且安全敏感的应用场景中，为了保护用户信息被盗用，我们需要对表单提交类请求添加 CSRF 保护。

Fireboom 默认关闭了 CSRF 保护，如果需要，请前往 Fireboom 控制台，点击“设置”，选择“安全”，打开“CSRF 保护”。

## 获取 CSRF token

```console
GET https://<hostname>/auth/cookie/csrf
```

相应为文本格式的 CSRF token

## 使用 CSRF token

对于`Mutation`类型对 API，我们需要在每次请求前添加`X-CSRF-Token`请求头，值为上一步获取到的结果。

## 使用 SDK

我们生成的[客户端 SDK](../sdk-sheng-cheng/) 已自动实现了 CSRF token 保护，你可以放心的直接使用 SDK。
