---
description: 待完善
---

# 工作原理

本文主要介绍飞布的底层工作原理，帮你建立关于飞布底层工作机制的宏观印象。

对于B/S架构，WebAPI本质上是服务端和客户端进行数据交互的接口。服务端连接不同类型数据源，从中获取并处理数据，按照客户端需求拼接并返回字段。通常情况下，实现该过程需要后端开发者编写大量代码。但绝大多数WebAPI都是针对数据库表及其关联表的增删改查，还有一小部分只需要在查询数据之前或之后，增加一些自定义逻辑。

## 聚合数据源——超图

因此，我们可以提取共性，用声明式语法描述数据表操作，并集成钩子机制扩展业务逻辑。其实，SQL也是一种声明式语法，适用于关系型数据库查询，不适合作为API声明式语法。GraphQL也是一种声明式语法，是一种基于图形的查询语言，用于从API中检索数据。总的来说，SQL适用于关系数据库，GraphQL适用于API查询。

所以，飞布选择GraphQL作为声明式语言，用来替代数据库的SQL语法。此外，GraphQL还能表达REST API的OAS规范。因此，选择GraphQL作为声明式语言，可以统一数据库和REST API，当然也能统一本身就是GraphQL的GraphQL API。

为了简化操作，飞布支持自动内省数据库、REST API以及GraphQL API，获得对应数据源的GraphQL Schema，我们称之为“子图”。

总的来说，飞布以 API 为中心，将所有数据抽象为 API，包括 REST API，GraphQL API ，数据库甚至消息队列等，通过 GraphQL 协议把他们聚合在一起，形成具有数据全集的“超图”。

<figure><img src="../.gitbook/assets/image (2) (1) (1).png" alt=""><figcaption><p>飞布架构图</p></figcaption></figure>

## 服务端OPERATION

理论上，基于“超图”公开对外发布GraphQL API，能够一劳永逸。因为，“超图”中包含了数据源中的任意数据。前端开发者只需要从GraphQL API中选择所需数据，构建子集OPERATION，就能够满足绝大多数业务需求，这也是hasura的实现方式。但该方式除了给前端开发者增加<mark style="color:red;">额外学习成本</mark>外，还有如下弊端：<mark style="color:red;">无法复用HTTP基础设施</mark>、<mark style="color:red;">安全性差</mark>（攻击者在客户端构造深度OPERATION发起DDOS攻击）。

{% hint style="info" %}
OPERATION：GraphQL有三种类型的operation，分别为query（查询）, mutation（变更）以及 subscription（订阅）
{% endhint %}

飞布采用了一种新的实现形式：**服务端OPERATION**。相比于hasura对外暴露GraphQL API，然后让前端开发者编写OPERATION调用GraphQL端点的方式。飞布在生产环境下**不**对外暴露GraphQL API，而是让后端开发者在服务端编写OPERATION。然后，飞布引擎将OPERATION编译为 REST-API，暴露给前端开发者。该方式，不仅能避免客户端OPERATION的所有缺陷，而且能充分发挥GraphQL的优势。

首先，前端开发者无需感知GraphQL的存在，因此<mark style="color:orange;">无需任何学习成本</mark>。

其次，飞布对外暴露REST API，可以<mark style="color:orange;">复用HTTP基础设施</mark>，如CDN等。

再者，OPERATION保存在服务端，攻击者无法触达OPERATION，保证了<mark style="color:orange;">安全</mark>。

最后，无论客户端还是服务端OPERATION，都能利用GraphQL<mark style="color:orange;">按需取用、类型系统</mark>的优势。

<figure><img src="../.gitbook/assets/image (3) (1).png" alt=""><figcaption><p>数据流转图</p></figcaption></figure>

## 指令系统扩展能力

飞布采用服务端OPERATION架构还带来了额外的好处。利用GraphQL强大的指令系统，飞布可以通过指令注解的方式实现复杂业务逻辑，其中大部分逻辑都有安全性要求，因此不能用客户端OPERATION实现。

{% hint style="info" %}
Directives 可视为GraphQL 的一种语法蜜糖(sugar syntax)，通常用于调整query 及schema 的行为，不同场景下可以有以下功能：

1. 影响query原有行为，如@include, @skip为query增加条件判断
2. 为Schema加上描述性标签，如@deprecated可以用于废除schema的某field又避免breaking change
3. 为Schema 添加新功能，例如参数检查、简单计算、权限检查、错误处理等等。
{% endhint %}

### API接口权限

控制接口只能被拥有特定权限的用户访问，是WEBAPI开发过程中，最基本的需求。业内最通用接口权限控制方式是：RBAC模型。飞布通过自定义GraphQL指令：`@rbac`，实现了API接口的RBAC控制。

```graphql
query GetOnetodo($uid: Int!) @rbac(requireMatchAll: [admin]) # 拥有admin角色用户才能访问 {
  data: todo_findFirsttodo(where: {user_id: {equals: $uid}}) {
    id
    title
    user_id
  }
}
```

### API数据权限

限制接口只能被当前登录用户访问，且只能获取当前用户所拥有的数据行或字段，也是WEBAPI开发的常见需求。飞布通过自定义GraphQL指令：`@fromClaim`，实现了API数据权限控制。

```graphql
query GetOnetodo($uid: Int! @fromClaim(name: USERID) # 注入当前登录用户的ID) {
  data: todo_findFirsttodo(where: {user_id: {equals: $uid}}) {
    id
    title
    user_id
  }
}
```

### API入参校验

接口入参校验也是WEBAPI开发过程中较繁琐的部分。飞布通过自定义GraphQL指令：`@jsonSchema`，实现了API入参校验。

```graphql
query GetOnetodo($uid: Int! @jsonSchema(pattern: "^ [0-9]*$")# 正则表达式校验入参 ) {
  data: todo_findFirsttodo(where: {user_id: {equals: $uid}}) {
    id
    title
    user_id
  }
}
```

### API参数注入

很多场景下，接口的入参需要由服务端动态设置特定参数。飞布内置了如下指令，分别适用不同场景的需求。

* @injectGeneratedUUID：生成uuid注入到参数中
* @injectCurrentDateTime：获取当前时间，注入到参数中
* @injectEnvironmentVariable：获取环境变量，注入到参数中

```graphql
query GetOnetodo($uid: Int! @injectGeneratedUUID # 生成UUID) {
  data: todo_findFirsttodo(where: {user_id: {equals: $uid}}) {
    id
    title
    user_id
  }
}
```

### API响应转换

某些场景下，API所需的结构与数据库对应字段的层级不一致，因此要进行映射。飞布通过自定义GraphQL指令：`@transform`，实现了API入参校验。

```graphql
query GettodoList {
  total: todo_aggregatetodo @transform(get: "_count.id") # 将_count.id值赋值给total字段 {
    _count {
      id
    }
  }
}
```

### 跨数据源关联





## OPERATION配置



### 数据缓存



### N+1查询

### 服务端推送





## 钩子机制扩展逻辑







## 身份验证





##



飞布还充分利用了GraphQL的指令系统，通过指令注解实现了API权限和数据权限的控制，入参校验，以及跨数据源关联！



与GraphQL的指令系统结合，能够极大扩展API的能力，





飞布还充分利用了GraphQL的指令系统，通过指令注解实现了API权限和数据权限的控制，入参校验，以及跨数据源关联！但用户无需刻意学习，因为飞布提供了友好的交互，封装了这些技术细节。

飞布提供了开箱即用的API缓存、实时推送和实时查询功能。通过服务端轮询，可以实现任意数据源的实时查询！

此外，飞布基于 HTTP 协议+[WebAssembly](https://developer.mozilla.org/zh-CN/docs/WebAssembly)技术实现了 HOOKS 机制，方便开发者采用任何喜欢的语言实现自定义逻辑。同时，飞布内置了WebContainer，TypeScript开发者无需准备任何环境，即可进行nodejs钩子的开发。

飞布集成了众多行业规范，包括OIDC、S3存储、RBAC等！用户无需额外学习即可接入，快速完成业务需求。

最后，飞布还基于prisma设计了数据建模功能，实现开发流程的闭环。用户无需切换工具，即可完成数据建模和数据预览，且能跨数据库类型迁移表结构。
