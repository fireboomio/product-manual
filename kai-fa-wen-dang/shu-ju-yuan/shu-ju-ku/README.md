---
description: 绝大多数应用都离不开数据库，本章将介绍如何使用飞布连接数据库，并从中构建API。
---

# 数据库

飞布控制台提供了友好的界面，无需借助三方工具，就能快速完成数据建模。你也可以使用现有数据库连接，进行API开发。~~飞布CLI还可以生成迁移文件，你可以直接编辑并进行数据库版本控制。~~

飞布使用Prisma连接到数据库，如SQLite, MySQL, PostgreSQL, SQL Server或MongoDB。飞布底层封装了[prisma引擎](https://www.prisma.io/)，利用prisma的内省能力，将数据库中的表结构内省成GraphQL Schema，并利用prisma的查询能力解析GraphQL Opeartion读取/更新数据库数据。此外，飞布还利用prisma的迁移能力，实现了数据建模功能。

![多数据源支持](https://www.fireboom.io/images/gif/01-01%E5%A4%9A%E6%95%B0%E6%8D%AE%E6%BA%90%E6%94%AF%E6%8C%81.gif)

## 支持数据库

理论上飞布支持prisma兼容的所有数据库，所支持的数据源的完整列表可以在[这里](https://www.prisma.io/docs/concepts/database-connectors/mysql)查看，当前正在按照优先级持续兼容中。

* [x] [MySQL](https://www.prisma.io/docs/concepts/database-connectors/mysql)
* [x] [SQLite](https://www.prisma.io/docs/concepts/database-connectors/sqlite)
* [ ] [PostgreSQL](https://www.prisma.io/docs/concepts/database-connectors/postgresql)
* [ ] [MongoDB](https://www.prisma.io/docs/concepts/database-connectors/mongodb)
* [ ] [SQL Server](https://www.prisma.io/docs/concepts/database-connectors/sql-server)
* [ ] [CockroachDB](https://www.prisma.io/docs/concepts/database-connectors/cockroachdb)



如果你想了解支持进度，可[联系我们](https://github.com/fireboomio/product-manual/discussions/1)。

