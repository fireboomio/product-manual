# 数据库连接

你可以快速且便捷的连接飞布到新的、现存的或者示例数据库上。你也可以同时连接多个数据库，实现跨数据源的数据编排。

## 数据准备

首先你要准备好数据库的连接配置。

## 新建数据库

例如：新建数据源->MySQL，填写数据库配置信息->测试连接，连接成功后，选择保存，该MySQL数据源即新建完成，并保存到了数据库列表中。

数据库配置信息有两种方式：

### 连接URL

#### MySQL

```
mysql://USER:PASSWORD@HOST:PORT/DATABASE
```

![mysql connector](https://www.prisma.io/docs/static/a3179ecce1bf20faddeb7f8c02fb2251/4c573/mysql-connection-string.png)

#### PostgreSQL

```
postgresql://USER:PASSWORD@HOST:PORT/DATABASE
```

![postgresql connector](https://www.prisma.io/docs/static/13ad9000b9d57ac66c16fabcad9e08b7/4c573/postgresql-connection-string.png)

#### SQLite

```
file:./dev.db
```

#### MongoDB

![](https://www.prisma.io/docs/static/b5ef4062c4686c772571b3079ba1331c/4c573/mongodb.png)

```
mongodb://USERNAME:PASSWORD@HOST/DATABASE
```

### 连接参数

| 名称   | 占位符        | 描述                         |
| ---- | ---------- | -------------------------- |
| 主机   | `HOST`     | 数据服务的IP地址或域名，例如`localhost` |
| 端口   | `PORT`     | 数据库服务运行的端口，例如`3600`        |
| 用户   | `USER`     | 数据库用户的名称，例如`janedoe`       |
| 密码   | `PASSWORD` | 数据库用户的密码                   |
| 数据库名 | `DATABASE` | 你想使用的数据库名称，例如 `mydb`       |



### 高级设置

在数据库中存储JSON列是很常见的用例。如果你使用数据库，例如PostgreSQL，你可以使用 `json` 或`jsonb` 类型存储JSON列。在GraphQL schema中，该列将被表示为`JSON` 标量类型。在内部，把值存储到数据库之前，飞布将json值编码为字符串，并在从数据库读取时解码它。这样，`JSON`标量类型非常容易使用。

然而，该方法有个缺点。如果你打算存储复杂JSON对象，你无法利用GraphQL的类型系统优势。不容易从JSON对象中选择特定的字段。你必须要手动解析它，并且选择你想要的字段。

高级设置就是解决该问题的机制。飞布允许您扩展GraphQL Schema并用自定义类型替换特定的JSON字段。通过这种方式，你可以利用GraphQL的类型系统，同时能够将数据作为JSON对象存储在数据库中。



```graphql
mutation (
  $email: String! @fromClaim(name: EMAIL)
  $name: String! @fromClaim(name: NAME)
  $message: String!
  $payload: JSON! # 这是一个JSON变量
) @rbac(requireMatchAll: [user]) {
  createOnemessages: db_createOnemessages(
    data: {
      message: $message
      payload: $payload
      users: {
        connectOrCreate: {
          create: { name: $name, email: $email }
          where: { email: $email }
        }
      }
    }
  ) {
    id
    message
    payload
  }
}
```



```
{
  findManymessages: db_findManymessages(take: 20, orderBy: [{ id: desc }]) {
    id
    message
    payload
    users {
      id
      name
    }
  }
}
```



```graphql
 type MessagePayload {
            extra: String!
        }
        input MessagePayloadInput {
            extra: String!
        }
```



```json
[
    {
      entityName: 'messages',
      fieldName: 'payload',
      responseTypeReplacement: 'MessagePayload',
      inputTypeReplacement: 'MessagePayloadInput',
    },
  ]
```





```
mutation (
  $email: String! @fromClaim(name: EMAIL)
  $name: String! @fromClaim(name: NAME)
  $message: String!
  $payload: db_MessagePayloadInput!
) @rbac(requireMatchAll: [user]) {
  createOnemessages: db_createOnemessages(
    data: {
      message: $message
      payload: $payload
      users: {
        connectOrCreate: {
          create: { name: $name, email: $email }
          where: { email: $email }
        }
      }
    }
  ) {
    id
    message
    payload {
      extra
    }
  }
}
```



```
{
  findManymessages: db_findManymessages(take: 20, orderBy: [{ id: desc }]) {
    id
    message
    payload {
      extra
    }
    users {
      id
      name
    }
  }
}
```



在数据库详情页，点击右上角“高级设置”，可



