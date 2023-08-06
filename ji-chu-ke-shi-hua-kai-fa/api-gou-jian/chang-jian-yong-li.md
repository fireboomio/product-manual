# 常见用例

## 组装查询&#x20;

### QUERY

在一个接口中，同时拿到多个数据源的数据。

```graphql
query MyQuery{
  # mysql 数据源
  mysql_findFirstPost {
    createdAt
    id
  }
  # pgsql 数据源
  pgsql_findFirstTodo {
    completed
    content
    createdAt
  }
  # rest 数据源
  system_getAllRoles {
    code
    remark
  }
}
```

### MUTATION

在一个接口中，同时变更多个数据源的数据。同步执行，暂不支持事务。

```graphql
mutation MyQuery {
  # mysql创建数据
  mysql_createOneTodo(data: {title: "test"}) {
    id
  }
  # mysql删除数据
  mysql_deleteOneTodo(where: {id: 10}) {
    id
  }
  # pgsql删除数据
  pgsql_deleteOneTodo(where: {id: 10}) {
    id
  }
}
```

## 分页

分页查询，支持排序和查询条件（例如模糊搜索）。具体用法参考：[#query-operation](ke-shi-hua-gou-jian/shi-yong-api.md#query-operation "mention")

```graphql
query GetTodoList(
  $take: Int = 10, $skip: Int = 0, # 分页 skip=(page-1)*take
  $orderBy: [todo_TodoOrderByWithRelationInput], # 排序
  $query: todo_TodoWhereInput) { # 查询条件，支持模糊搜索
  # 获取列表
  data: todo_findManyTodo(skip: $skip take: $take orderBy: $orderBy where: {AND: $query}) {
    id
    title
    completed
    createdAt
  }
  # 获取记录数
  total: todo_aggregateTodo(where: {AND: $query}) @transform(get: "_count.id") {
    _count {
      id
    }
  }
}
```



## 模糊搜索

模糊搜索，对应sql like 匹配，例如：根据$title模糊搜索待做事项列表

```graphql
query MyQuery($title: String, $skip: Int!) {
  todo_findFirstTodo(skip: $skip, take: 10, where: {title: {contains: $title}}) {
    createdAt
    completed
    id
    title
  }
}gr
```

## 关联查询

在查询某个对象的同时，获取其关联对象，对应sql left join。使用关联查询，需要在数据库表间建立外键。若无外键，请使用 [kua-yuan-guan-lian.md](kua-yuan-guan-lian.md "mention")

<figure><img src="../../.gitbook/assets/image.png" alt=""><figcaption></figcaption></figure>

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

## 关联更新

在更新某个对象数据的同时，关联更新其关联对象的数据，仅适用于1:1关联。

```graphql
# 更新用户$name的同时，更新其profile的$bio字段。
mutation MyQuery( $id: Int!, $name: String, $bio: String) {
  mysql_updateOneUser(
    data: {name: {set: $name}, Profile: {update: {bio: {set: $bio}}}}
    where: {id: $id}
  ) {
    id
  }
}
```

## 跨源关联

通过 @export 和 \_join 可以进行跨数据源的关联查询。详情见 [跨源关联](chang-jian-yong-li.md#kua-yuan-guan-lian)。

## **原生sql**

数据库内省的“函数”基于prisma构建，有些场景无法支持，如数据统计类需求。因此，支持raw sql，以实现任意复杂度的需求。

### queryRaw

执行查询SQL，返回值是对象数组

```graphql
# 用法1：界面上语法会报错，但实际上支持
mutation MyQuery($id:Int=1) {
  todo_queryRaw(query: "SELECT *,rowid \"NAVICAT_ROWID\" FROM \"main\".\"Todo\"  WHERE id=$1", parameters: [$id])
} 
# 用法2
mutation MyQuery($parameters: todo_Json = [1]) {
  todo_queryRaw(
    query: "SELECT *,rowid \"NAVICAT_ROWID\" FROM \"main\".\"Todo\"  WHERE id=$1"
    parameters: $parameters
  )
}
```

### executeRaw

执行变更SQL，返回值是包含count字段的对象

```graphql
# 用法1：界面上语法会报错，但实际上支持
mutation MyQuery($title: String = "beijing",$id:Int=1) {
  todo_executeRaw(
    query: "UPDATE \"main\".\"Todo\" SET \"title\" = $1 WHERE id=$2"
    parameters: [$title,$id]
  )
}
# 用法2
mutation MyQuery($parameters: todo_Json = ["beijing", 1]) {
  todo_executeRaw(
    query: "UPDATE \"main\".\"Todo\" SET \"title\" = $1 WHERE id=$2"
    parameters: $parameters
  )
}
```

## 默认值

必填项与默认值？
