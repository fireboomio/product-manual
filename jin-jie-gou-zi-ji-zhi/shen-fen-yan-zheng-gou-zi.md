# 身份验证钩子

### 注册认证钩子

1. 后置普通钩子

* 路径：/authentication/postAuthentication
* 入参：请使用全局参数${\_\_wg.user}
* 出参：

```json
{
    "hook": "postAuthentication"
}
```

* 用途：认证成功后同步用户信息或记录用户访问日志

2. 后置修改信息钩子

* 路径：/authentication/mutatingPostAuthentication
* 入参：请使用全局参数${\_\_wg.user}
* 出参：

```json
{
    "hook": "mutatingPostAuthentication",
    "response": {
        "user": ${__wg.user} // 与全局参数路径__wg.user格式一致
        "status": "ok", // 状态不为ok时，使用message作为错误抛出
        "message": "not ok message"
    }
}
```

* 用途：认证成功后修改用户信息或中断认证

3. 后置重新校验钩子

* 路径：/authentication/revalidateAuthentication
* 入参：请使用全局参数${\_\_wg.user}
* 出参：

```json
{
    "hook": "revalidateAuthentication",
    "response": {
        "user": ${__wg.user} // 与全局参数路径__wg.user格式一致
        "status": "ok", // 状态不为ok时，使用message作为错误抛出
        "message": "not ok message"
    }
}
```

* 用途：请求携带revalidate参数会每次重走认证，默认从缓存获取user，根据参数选择是否进行重新认证校验或改写
