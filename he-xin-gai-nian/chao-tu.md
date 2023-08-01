# 超图

如果我们重新审视 API 的作用，我们会发现，作为客户端和服务端数据的桥梁，API 解析客户端的请求，从服务端某个数据源（可能是数据库，也可能是其他服务的数据等），获取相应的数据，然后按照 API 的约定返回合适的结果。

既然 API 的目的是提供数据，而数据往往有其严苛的 schema，同时 API 的 schema 大多数时候就是数据 schema 的子集，那么，我们是不是可以从数据 schema 出发，反向生成 API 呢？

但不同数据源协议不同，如 对数据库而言，不同数据库有不同SQL方言，对REST api和gql api协议也不同。

而且聚合不同类型数据源的数据，也是API的常见用例。因此，需要先将不同协议的schema，统一为一种协议，然后再行编排，生成API，才能符合绝大多数应用场景。

## 统一协议

<figure><img src="../.gitbook/assets/image (3).png" alt=""><figcaption></figcaption></figure>

gql作为一种强类型API语言，非常适合作为该统一协议。

1. 借助Prisma引擎，将不同类型数据库的schma转成了gql schema，统一了不同数据库的sql方言。
2. 同时，利用openapi specification规范，将rest api也转成了gql schema
3. 而graphql api，本身就是gql协议，具备gql schema

我们将上述各数据源转换而来的GQL SCHEMA，称为子图。

各个子图合并后的产物称为超图，它本质上也是GQL SCHEMA，只不过聚合了不同的数据源。图中将3种类型数据源 数据库、REST API以及gql api产生的超图，都合并到了这个超图中。

但为了避免不同数据源相同表名或接口合并到超图后，产生冲突，因此需要给不同子图加上命名空间，一示区分。

上述所有的步骤都是自动执行的，无需手工编写gql schema，极大提升开发体验。

而实现上述自动化的核心需要依赖gql的内省能力。

## Prisma

接着我们学习下prisma mode到gql scheama的转换规则。

prisma的核心原理见 [shu-ju-ku](../ji-chu-ke-shi-hua-kai-fa/shu-ju-yuan/shu-ju-ku/ "mention")

<figure><img src="../.gitbook/assets/image (5).png" alt=""><figcaption><p>Prisma2GraphQL转换规则</p></figcaption></figure>

首先，prisma model与数据库建表语句等价。且prisma model能转换成多种数据库的建表语句。

例如，这里的prisma model转换为了pgsql的建表语句。

```prisma
model Todo {
  id        Int      @id @default(autoincrement())
  title     String
  completed Boolean  @default(false)
  createdAt DateTime @default(now())
  content   String?  @default("")
  test      Int      @default(1)
}
```

接着，我们看下prisma model到gql schema的转换规则。

这里有一个todo模型，右侧是其转换后的schema可视化展示。

规则为：命名空间+方法名+表名。命名空间是为了区分不同数据库实例的相同表名。

其中查询有如下方法：

* findFirst：返回列表中符合条件的第一条记录
* findMany：返回记录列表
* findUnique：根据主键或唯一键查询一条记录
* aggregate：聚合数据，包括 avg、count、sum、min、max
* groupBy：聚合函数，根据一或多列对结果集分组

变更有如下方法：

* createOne：创建一条记录，对应 sql 中的 insert
* createMany：创建多条记录
* deleteMany：批量删除记录，对应 sql 中的 delete
* deleteOne：删除一条记录，对应 sql 中的 delete
* updateMany：更新多条记录，对应 sql 的 update
* updateOne：更新一条记录，对应 sql 的 update
* upsertOne：更新或插入一条记录

里面还包含两个原生方法，用于编写复杂SQL，适用于上述方法无法覆盖的场景。

* queryRaw：执行查询类sql
* executeRaw：执行创建、变更类sql

对于有关联关系的模型，其函数名称一致，但入参会更加复杂，这块需要[前往](https://www.prisma.io/docs/reference/api-reference/prisma-client-reference#findunique)prisma官方文档查看。

{% hint style="info" %}
Prisma对不同数据库的函数支持是不同的，例如sqlite数据库缺少createMany方法。
{% endhint %}

## RSET API

飞布除了支持数据库数据源外，还支持REST 数据源。

<figure><img src="../.gitbook/assets/image (6).png" alt=""><figcaption><p>OAS2GraphQL转换规则</p></figcaption></figure>

如图所示，是OAS3.0规范的示例，可以看到它也定义了入参和出参，并且支持类型，可以转换为gql scheme。

所以只需要提供OAS文件，飞布就能将rest api转变为子图，然后合并到飞布的超图中。

注意看，子图没有命名空间前缀，超图中增加了前缀PET。

接着，我们看看OAS生成GQL SCHEMA子图的详细规则：

1. paths中请求类型对应gql schema中的query：get对应query，POST/put/patch对应mutation。例如，pet/{petid} get 转换为query类型。
2. paths的operationid，对应gql schemla 中的根字段名称，或者说函数名称，如果它为空，则用path去除特殊字符后函数名。例如getpetbyid。
3. parameters入参对应gql schemla 中函数的入参。例如，petid为integer类型，对应到gql schema为Petid int类型，required true，对应为！.
4. response出参对应gql schemla 中函数的返回值。例如，pet  schema转换为pet 类型。id字段为标量interge类型，转换为int。category为对象，在这里也是对应为cateory对象。

{% hint style="info" %}
推荐使用apifox等工具导出oas文件，优先导出为oas3.0。
{% endhint %}

## 架构图

最后我们查看下架构图，重点关注超图部分。

<figure><img src="../.gitbook/assets/image (7).png" alt=""><figcaption><p>架构图-超图</p></figcaption></figure>

在这里，飞布主要做了2件事：

1. 用graphql统一协议，屏蔽数据库、REST API等不同类型数据源的协议。其中又用prisma屏蔽了各数据库sql方言的差异。
2. 自动内省自行转换协议：将rest api的oas文件转换为graphql，自动内省数据库和graphql 数据源，避免手动编写graphql。

至此，你已经掌握了飞布最核心的概念——超图。
