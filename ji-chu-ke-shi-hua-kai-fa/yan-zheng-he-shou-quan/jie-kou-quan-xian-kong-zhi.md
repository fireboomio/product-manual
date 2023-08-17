---
description: '@rbac指令'
---

# 接口权限控制

接下来，我们先学习如何在Fireboom中为权限绑定角色，实现该过程需要用到`@rbac`指令。

该指令为全局指令，作用于整个OPERATIN上。

## 添加指令

1. 前往API管理面板，选择需要设置权限的API
2. 在GraphQL编辑区的工具栏中点击“@角色”，选择匹配模式并添加角色，如添加admin角色

<figure><img src="../../.gitbook/assets/image (3).png" alt=""><figcaption></figcaption></figure>

使用该指令后，有2个影响。

1. **登录校验**：用户登录后才能访问当前接口，这部分和单独在设置中开启授权功能一致。
2. **权限控制**：访问接口时，判断当前登录用户的角色是否与匹配规则一致，不一致时返回401错误。

## 指令介绍

该指令有4种类型的匹配规则，匹配规则本质上对应数学中集合的概念。

首先，需要知道三个集合域：

* 全部角色：角色管理列表中配置的角色

<figure><img src="../../.gitbook/assets/image (10) (2).png" alt=""><figcaption><p>全部角色</p></figcaption></figure>

* 用户角色：用户拥有的角色，在钩子中为用户设置的角色
* API角色：OPERATION上 `@RBAC`指令声明的角色

角色列表为全集，用户拥有的角色为其子集，API拥有的角色也是子集。

* requireMatchAll：全部匹配，用户角色包含API角色时，可访问
* <mark style="color:red;">requireMatchAny</mark>：任意匹配，用户角色与API角色有交集时，可访问
* denyMatchAny：互斥匹配，用户角色与API角色互斥时，可访问
* denyMatchAll：非全部匹配，当任意匹配或互斥匹配时，可访问

最后，以全部角色作为全集，结合四种关系，看用户角色和API角色的交并补情况，确定当前用户是否能访问当前API。



实际情况下，我们一般使用`requireMatchAny`匹配规则。

![](<../../.gitbook/assets/image (2).png>)

它表示角色拥有某些接口，当用户拥有该角色时，就能访问该接口，也即：用户通过角色拥有该角色的所有接口权限！<mark style="color:orange;">和RBAC0规范一致</mark>！

一方面，一个角色会有多个接口。对应着就是，多个OPERTION使用`requireMatchAny`绑定同一个角色。

例如：A、B OPERATION都绑定了admin角色。

{% code title="A.graphql" %}
```graphql
query A($id: Int = 1) @rbac(requireMatchAny: [admin]) {
  todo_findUniqueTodo(where: {id: $id}) {
    completed
    createdAt
    id
    title
  }
}
```
{% endcode %}

{% code title="B.graphql" %}
```graphql
query B($id: Int = 2) @rbac(requireMatchAny: [admin]) {
  todo_findUniqueTodo(where: {id: $id}) {
    completed
    createdAt
    id
    title
  }
} 
```
{% endcode %}

反过来，一个接口也会属于多个角色。对应就是，一个OPERTION使用`requireMatchAny`绑定了多个角色。

例如：C OPEARTION绑定了admin和user角色。

{% code title="C.graphql" %}
```graphql
query C($id: Int = 3) @rbac(requireMatchAny: [admin, user]) {
  todo_findUniqueTodo(where: {id: $id}) {
    completed
    createdAt
    id
    title
  }
}
```
{% endcode %}

* 如果，用户只拥有admin角色，它就能访问A、B、C三个接口
* 如果，用户只拥有user角色，就只能访问C接口。

