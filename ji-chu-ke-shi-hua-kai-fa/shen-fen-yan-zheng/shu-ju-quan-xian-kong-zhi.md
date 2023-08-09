---
description: '@fromClaim指令'
---

# 数据权限控制

最后，我们学习接口的数据权限控制，其核心是`@fromclaim`指令。

`@fromclaim`指令，又叫 [#deng-lu-xiao-yan-zhi-ling](../api-gou-jian/api-zhi-ling.md#deng-lu-xiao-yan-zhi-ling "mention") ，是注入指令的一种，适用于查询、变更和订阅OPERATION。

使用该指令后，有2个影响：

* 登录校验：登录后才能访问当前接口，这部分和单独在设置中开启授权功能一致。
* 参数注入：将登录用户的信息注入到入参中，保证该变量不被客户端修改。

<figure><img src="../../.gitbook/assets/image (54).png" alt=""><figcaption><p>@fromclaim指令</p></figcaption></figure>

该指令用于不同类型OPERATION，有不同用途：

* 用于查询或订阅：则保证只能查看当前用户拥有的数据，如获取登录用户的所有文章。
* 用于变更：则保证创建的记录中包含登录用户的信息，如post表的uid。

该指令支持注入如下信息：

* USERID：
* EMAIL
* EMAIL\_VERIFIED
* NAME
* NICKNAME
* LOCATION
* PROVIDER &#x20;
* ROLES：ROLES is string array, Please use in \[in, notIn]

roles比较特殊，我们后续讲到角色鉴权时再展开。
