# 数据库

绝大多数应用都离不开数据库，本章将介绍如何使用飞布连接数据库，并从中构建API。

飞布使用Prisma ORM连接数据库，支持SQLite, MySQL, PostgreSQL, SQL Server或MongoDB。

飞布提供了友好的界面，无需借助三方工具，就能完成数据建模。你也可以连接现有数据库开发API。~~飞布CLI还可以生成迁移文件，你可以直接编辑并进行数据库版本控制。~~

## 原理浅析

飞布底层封装了[prisma引擎](https://github.com/prisma/prisma-engines)，利用prisma的内省能力\[内省引擎]，将数据库中的表结构内省成Prisma Model以及GraphQL Schema，并利用prisma的查询能力\[查询引擎]执行GraphQL Opeartion操作数据。此外，飞布还利用prisma的迁移能力\[迁移引擎]，实现了数据建模功能。

{% hint style="info" %}
Prisma ORM一般指的是基于Ts对Prisma Engine的封装，而飞布只使用了[Prisma  Engine](https://github.com/prisma/prisma-engines)，未使用[Ts Prisma](https://github.com/prisma/prisma)。
{% endhint %}

了解更多，请前往"[工作原理](../../../kuai-su-ru-men/gong-zuo-yuan-li.md)"。

## 查询引擎

Prisma查询引擎是一个常驻WEB SERVER，用于执行GraphQL Operation（通过HTTP协议或NODE API）。本质上是将GraphQL Operation转换成SQL语句，发送至数据库执行，并将查询结果拼装成JSON响应返回。下图是Prisma Client调用查询引擎的原理图。

<figure><img src="../../../.gitbook/assets/image (2) (4).png" alt=""><figcaption><p>查询引擎Ts调用原理图</p></figcaption></figure>

Prisma Client用TS函数封装了GraphQL Operation。下图的`findMany`，在飞布中对应如下GraphQL Operation。

```graphql
# todo_是命名空间，代表数据库名称
# Todo是数据库表名
query MyQuery {
  todo_findManyTodo {
    completed
    createdAt
    id
    title
  }
}
```

飞布GraphQL Operation和Prisma Client函数签名的对应规则为：

**GraphQL Operation=\[命名空间]+函数签名+\[表名]**

Prisma 查询引擎支持函数如下：

查询方法：

* findFirst：返回列表中符合条件的第一条记录。
* findMany：返回记录列表。
* findUnique：根据主键或唯一键查询一条记录。
* aggregate：聚合数据，包括avg、count、sum、min、max。
* [groupBy](https://www.prisma.io/docs/concepts/components/prisma-client/aggregation-grouping-summarizing#groupby-and-ordering)：结合聚合函数，根据一个或多个列对结果集进行分组。

变更方法：

* createOne：创建一条记录，对应sql中的insert。
* deleteMany：批量删除记录，对应sql中的delete。
* deleteOne：删除一条记录，对应sql中的delete。
* updateMany：更新多条记录，对应sql的update。
* updateOne：更新一条记录，对应sql的update。
* upsertOne：更新或插入一条记录。

了解更多， 前往 [Prisma文档](https://www.prisma.io/docs/reference/api-reference/prisma-client-reference#findunique) 查看。

## 内省引擎

Prisma内省引擎是一个二进制命令行，用于内省数据库，获得数据库schema，并映射成prisma schema。

对应`prisma db pull`命令。

<figure><img src="../../../.gitbook/assets/image (6) (4) (1).png" alt=""><figcaption><p>内省引擎Ts调用原理图</p></figcaption></figure>

了解更多，前往 [Prisma 文档](https://www.prisma.io/docs/concepts/components/introspection) 查看。

## 迁移引擎

Prisma迁移引擎是一个二进制命令行，用于设计或迁移数据库，本质上是将prisma schema转换成建表或更新表语句。

迁移分为两种情况：

#### 原型设计

对应`prisma db push`命令，适用于开发阶段，可能造成数据丢失。

{% hint style="info" %}
在飞布数据建模功能模块中的”迁移“操作，底层调用的是该命令，因此不建议生产环境中使用。
{% endhint %}

#### 数据库迁移

对应 [`prisma migrate dev`](https://www.prisma.io/docs/concepts/components/prisma-migrate/mental-model#track-your-migration-history-with-prisma-migrate-dev)，适用于生产阶段。其原理如下图：

<figure><img src="../../../.gitbook/assets/image (40).png" alt=""><figcaption></figcaption></figure>

对该命令的集成，暂未实现，可先用[Prisma官方的命令](https://www.prisma.io/docs/concepts/components/prisma-migrate/migrate-development-production)进行支持。

## 支持数据库

理论上飞布支持prisma兼容的所有数据库，所支持的数据源的完整列表可以在[这里](https://www.prisma.io/docs/concepts/database-connectors/mysql)查看，当前正在按照优先级持续兼容中。

* [x] [MySQL](https://www.prisma.io/docs/concepts/database-connectors/mysql)
* [x] [SQLite](https://www.prisma.io/docs/concepts/database-connectors/sqlite)
* [x] [PostgreSQL](https://www.prisma.io/docs/concepts/database-connectors/postgresql)
* [ ] [MongoDB](https://www.prisma.io/docs/concepts/database-connectors/mongodb)
* [ ] [SQL Server](https://www.prisma.io/docs/concepts/database-connectors/sql-server)
* [ ] [CockroachDB](https://www.prisma.io/docs/concepts/database-connectors/cockroachdb)



如果你想了解支持进度，可[联系我们](https://github.com/fireboomio/product-manual/discussions/1)。

