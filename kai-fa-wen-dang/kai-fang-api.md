# 开放API（开发中）

在 Fireboom 的 API 设计理念里，接口权限控制是由 Operation 上的 `@rbac` 指令来指定的，比如配置了 `@rbac(requireMatchAll: [admin])` 则表示这个接口要求必须是具有`admin`角色的用户才可以访问。

Fireboom 本身提供了角色的增删改查，但我们没有提供内置的用户，因为用户在 Fireboom 的设计里是由 OIDC 服务商来提供的。这时要实现 Fireboom 角色和 OIDC 用户绑定就很困难。



鉴于此，我们提供了开放的 API 来帮助我们完成用户和角色、接口的绑定。

## 对外接口

1. 根据角色获取 API 列表

GET http://localhost:9991/roles/apis?code={role\_code}

入参`role_code`为角色名称

成功 200，返回数据结构为

```json
{
  "roles": [
    {
      "name": "admin",
      "apis": ["/operations/Hello"]
    },
    {
      "name": "admin",
      "apis": ["/operations/World"]
    },
  ]
}
```

2. 角色绑定 API 列表

POST http://localhost:9991/roles/bindApis

body 参数结构为

```json
{
  "code": "admin",
  "apis": ["/operations/Hello", "/operations/World"]
}
```

成功 200，无返回值

## 接口安全

上述开放的 2 个接口权限较高，需要设置角色控制。默认都是 `requireMatchAll: [admin]`，你可以在 Fireboom 控制台页面点击“设置”，选择“开放API”，分别为2个接口配置角色限制，具体操作和[新建 API 时相似](yan-zheng-he-shou-quan/shou-quan-yu-fang-wen-kong-zhi.md)。

## 接口使用

要实现前面说的角色绑定接口功能，我们需要先在查询角色列表时调用接口 1，在给具体角色绑定 API 接口列表时调用接口 2。

具体使用方法可参考 [refine 管理后台案例](../zui-jia-shi-jian/guan-li-hou-tai-refine.md)
