# 常见用例

## 组装查询

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

在一个接口中，同时变更多个数据源的数据。

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

### 事务

用`@transaction`指令修饰mutation OPERATION，保证原子性

```graphql
mutation MyQuery @transaction {
  rb_createOneT(data: { name: "22211122"}) {
    id
    name
  }
  rb_createOneRole(data: {code: "a111111", name: "1111"}) {
    code
    name
  }
} 
```

## 分页

分页查询，支持排序和查询条件（例如模糊搜索）。具体用法参考：[#query-operation](shi-yong-api.md#query-operation "mention")

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

同一数据源的关联查询有两种情况：

* 有外键：建立表之间的外键关联
* 无外键： [prisma-shu-ju-yuan.md](../shu-ju-yuan/prisma-shu-ju-yuan.md "mention")

其用法相同，如下：

```graphql
query MyQuery {
  rb_findFirstUser {
    name
    uid
    Role {
      code
      name
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

## 可为空字段

对于可为空字段，支持用null来查询或更新。

```prisma
model T {
  id   Int     @id @unique
  name String
  des  String? // 注意，这里的?
}
```

### null查询

将null作为查询条件，例如：

1，获取所有des=null的数据

```graphql
# 方法1
query MyQuery {
  rb_findManyT(where: {des: null}) {
    des
    id
  }
} 

# 方法2，不推荐使用
query MyQuery {
  rb_findManyT(where: {des: {equals: null}}) {
    des
    id
  }
}
```

2，获取所有des is not null的数据

```graphql
query MyQuery {
  rb_findManyT(where: {des: {not:null}}) {
    des
    id
  }
}
```

### null更新

设置某字段的值为null，例如：设置id为10的记录，des=null

```graphql
mutation MyQuery($des: String = null) {
  rb_updateOneT(data: {des: {set: $des}}, where: {id: 10}) {
    id
    des
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

## 变量默认值

### 默认值

设置`$skip`的默认值为0，请求接口时可忽略。

```graphql
query MyQuery($skip: Int = 0 ) {
  rb_findManyUser(skip: $skip, take: 10) {
    uid
    name
  }
}
```

### 必填项

设置`$skip`为必填项，请求接口时必填。

```graphql
query MyQuery($skip: Int! ) {
  rb_findManyUser(skip: $skip, take: 10) {
    uid
    name
  }
}
```
