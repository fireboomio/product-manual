# Prisma 数据源

Prisma数据源是一种特殊的数据库类型数据源，能实现所有数据库数据源拥有的功能！

其主要用途如下：

* 虚拟外键：数据库无需建立外键，只在Pirsma model中建立，支持同数据库的关联查询
* 支持视图：在prisma model建立视图表，实现视图的查询（<mark style="color:red;">更新待测试</mark>）
* 精简数据表：简化prisma model表或字段，只声明业务需要的表或字段，缩减超图大小，提高性能

## 快速使用

1，数据源新建页->prisma，设置名称后进入编辑页

2，在编辑区中输入datasource参数，例如：

```prisma
datasource db {
  provider = "mysql"
  url      = "mysql://root:xxx@localhost:3306/rbac"
}
# or
datasource db {
  provider = "mysql"
  url      = env("DB_URL")
}
```

3，点击“内省”获取全量 prisma model

```prisma
datasource db {
  provider = "mysql"
  url      = "mysql://root:xxx@localhost:3306/rbac"
}

model Role {
  code       String       @id
  name       String
}

model User {
  uid  String @id
  name String
}

model Permission {
  id            Int    @id @default(autoincrement())
  name          String @unique
  operationPath String @unique
  operationId   BigInt @unique
}

model post {
  id      Int    @id @default(autoincrement())
  content String
  cate    String
}
```

4，将上述prisma model修改如下：

```prisma
datasource db {
  provider = "mysql"
  url      = "mysql://root:xxx@localhost:3306/rbac"
}

model Role {
  code       String       @id
  name       String
  Permission Permission[] #增加外键关联
  User       User[]       #增加外键关联
}

model User {
  uid  String @id
  name String
  Role Role[]             #增加外键关联
}

model Permission {
  id            Int    @id @default(autoincrement())
  name          String @unique
  operationPath String @unique
  operationId   BigInt @unique
  Role          Role[]   #增加外键关联
}

# 删除post表
```

本质上是删减表、字段或增加外键关联！

5，保存并上线该数据库，即可在超图中使用

Fireboom根据手动编写的Prisma model生成子图，合并到超图中。
