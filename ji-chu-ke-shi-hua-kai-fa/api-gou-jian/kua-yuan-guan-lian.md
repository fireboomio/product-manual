# 关联查询

关联查询数据是API开发中的常见用例。有两种形式的关联查询：有外键关联和跨源关联。

## 有外键关联（嵌套关联）

在查询某个对象的同时，获取其关联对象，对应sql left join。使用关联查询，需要在数据库表间建立外键。

基于超图面板对 gql 的可视化封装，能充分发挥其嵌套查询的优势。甚至不需要掌握 SQL，就能构建多级的嵌套查询。

<figure><img src="../../.gitbook/assets/image (48).png" alt=""><figcaption></figcaption></figure>

如图所示有4张主表，其中 `User` 和 `Profile` 是1:1关联，`User` 和 `Post` 是 1对多关联，`Post` 和`Category` 是多对多关联。

其 Prisma Model 如下：

<details>

<summary>prisma model</summary>

```prisma
model Category {
  id   Int    @id @default(autoincrement())
  name String
  Post Post[]
}

model Post {
  id        Int        @id @default(autoincrement())
  createdAt DateTime   @default(now())
  title     String
  published Boolean    @default(false)
  authorId  Int
  User      User       @relation(fields: [authorId], references: [id])
  Category  Category[]

  @@index([authorId], map: "Post_authorId_fkey")
}

model Profile {
  id     Int    @id @default(autoincrement())
  bio    String
  userId Int    @unique
  User   User   @relation(fields: [userId], references: [id])
}
model User {
  id      Int      @id @default(autoincrement())
  email   String   @unique
  name    String?
  role    String   @default("admin")
  Post    Post[]
  Profile Profile?
}
```

</details>

{% hint style="info" %}
如何用prisma model建立关联关系，可参考[Prisma文档](https://www.prisma.io/docs/concepts/components/prisma-schema/relations/one-to-one-relations)
{% endhint %}

我们构建了如图所示多级嵌套查询 OPERATION，以 `User` 表为主体，查询到的 `Profile` 是对象，查询到的 `Post` 是数组，`Post` 中的 `Category` 也是数组。

```graphql
query MyQuery {
  mysql_findFirstUser {
    id
    name
    email
    # 1:n关联：post为对象数组
    Post(where: {title: {contains: "test"}}) {
      id
      title
      # n:n关联：category为对象数组
      Category {
        id
        name
      }
    }
    # 1:1关联：profile为对象
    Profile {
      id
      userId
    }
  }
} 
```

### N+1 查询

在架构图中，我们提到过飞布引擎做了很多性能优化，其中之一是N+1查询优化。

我们以示例讲解下，飞布如何实现查询优化。

<figure><img src="../../.gitbook/assets/image (49).png" alt=""><figcaption></figcaption></figure>

还以上述表为例，`User` 和 `Post` 为1对多关联，这时候我们构建了一个新的 `Operation` ，查询 `Post` 的 `User`，即查询文章的同时拿到作者信息。可以看到如下响应，其中 `User` 为对象。

这里有个问题，3条 `Post` 的作者是同一个，是否意味着：**飞布会查询3次用户？**类似上图，先查询所有文章，然后遍历文章，挨个查询作者，最后再合并。

实际上，在底层我们做了N+1查询优化，用 where in 的方式避免了无用的查询。将上述 sql 改成了：**先查询所有文章，然后用文章的作者id注入where in语句中**。

这样能提升查询性能，避免不必要的数据库请求。

## 跨源关联

默认情况下，OPRATION 都是**并行**执行，但很多场景需要先查询数据，然后将返回值作为条件，执行下一个查询/变更。

跨源关联有两种常见用例：

* 同一数据库表间未建立外键（虚拟外建更适合，详情见 [prisma 数据源](../shu-ju-yuan/prisma-shu-ju-yuan.md)）
* 跨数据源关联查询/变更

{% embed url="https://www.bilibili.com/video/BV1iM4y1U7mE/" fullWidth="false" %}
11功能介绍-飞布如何实现跨源关联？
{% endembed %}

> 举个物联网的例子：设备列表页，不仅要展示设备列表，还要展示设备在线状态。
>
> 传统模式下，我们需要先从数据库获取设备列表，然后遍历数据，逐个调用物联网平台接口获取设备在线状态，最后拼接数据返回给客户端。用编码的方式，大概需要几百行代码。
>
> 而利用飞布的跨源关联功能，只需要几行 GraphQL OPERATION 就能实现上述需求。

**跨源关联**本质上是一种流程编排，将通常情况下并行的请求，改造成串行。

<figure><img src="../../.gitbook/assets/operation-export.gif" alt=""><figcaption><p>跨源关联时序图</p></figcaption></figure>

使用跨源关联，至少需要配置两个数据源（或同一数据源两个不同函数）。

### 跨源查询

跨源查询，将上一个返回值作为参数传给下一个**查询。**

例如，db为数据库，iot为物联网 REST API。以下示例展示了：先用 `db_findManyDevice` 查询设备列表，然后用 `iot_deviceState` 查询设备在线状态。

```graphql
query MyQuery($device_id:Int! @internal) {  # 声明：定义变量 $device_id
  db_findManyDevice {
    id @export(as:"device_id")# 赋值： $device_id = id
    name
  # _join 字段返回类型Query!
  # 它存在于每个对象类型上，所以你可以在Query文档的任何地方看到它
    _join{
      iot_deviceState(device_id: $device_id) { # 使用：将 $device_id 设置给 device_id 变量 
        status
      }
    }
  } 
}
```

上述示例，主要分为三个环节：

* 声明：`@internal` 指令从公开API中移除 `$device_id` 变量。这意味，用户不能手工设置它。我们称它为关联键（join key）。
* 赋值：使用 `@export` 指令，我们可以将字段  `id` 的值导出给关联键(`$device_id`)
* 使用：一旦我们进入 `_join` 字段，我们可以使用 `$device_id` 变量去关联物联网 API

其中涉及到两个知识点：

* 指令： `@internal` 用于定义变量；`@export` 用于赋值变量。
* `_join`字段：`_join` 字段是一个特殊的字段，是 QUERY 类型。

{% hint style="info" %}
`_join` 这种定义方式，实现了循环嵌套，详情可查看文件：`exported/generated/fireboom.app.schema.graphql`
{% endhint %}

### 跨源更新

跨源更新，将上一个返回值作为参数传给下一个**变更**。

例如，blog为数据库1，todo为数据库2。以下示例展示了：先用 `blog_findUniqueUser` 查询用户（id=1），然后用 `todo_updateManyTodo` 更新该用户的所有待做事项为 未完成 （completed=false）。

```graphql
query MyQuery($authorId: Int! @internal) {
  blog_findUniqueUser(where: {id: 1}) {
    email
    id @export(as: "authorId")
    _join_mutation {
      todo_updateManyTodo(where:{authorId:{equals:$authorId}},data:{completed:{set:false}}) {
        count
      }
    }
  }
}
```

跨源更新与跨源查询的指令基本一致，除了 `_join` 变成了 `_join_mutation` 。`_join_mutation`是 MUTATION 类型。
