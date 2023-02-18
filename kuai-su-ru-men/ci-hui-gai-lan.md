---
description: 飞布中涉及的名词，持续更新
---

# 词汇概览

### GraphQL

GraphQL是一种用于API的查询语言和运行时。它允许客户端指定所需的数据，而不是服务器决定要返回的数据。这样，客户端可以有效地请求所需的数据，减少了浪费和冗余。GraphQL由Facebook开发，并于2015年开源。它目前是一种流行的API技术，被许多公司和开发人员所使用。它跟 SQL 的关系是共用 QL 后缀，就好像「汉语」和「英语」共用后缀一样，但他们本质上是不同的语言。

学习GraphQL：[GraphQL速学（译）](https://blog.biglion.top/2019/12/08/GraphQL%E9%80%9F%E5%AD%A6/)

### 数据源

数据的来源，指的是数据存储的地方。开发人员可以使用数据源中的数据创建应用程序，并向用户提供有关信息。数据源可以是本地或远程的，并且可以是静态或动态的。飞布中的数据源包含：数据库、REST API、GraphQL API。其中数据库包含：![](http://localhost:9123/assets/PostgreSQL.2a7e38b3.svg)PostgreSQL、![](http://localhost:9123/assets/MySQL.1461110d.svg)MySQL、![](http://localhost:9123/assets/MongoDB.491466e8.svg)MongoDB、![](http://localhost:9123/assets/SQLite.48a4dbe0.svg)Sqlite、![](http://localhost:9123/assets/CockroachDB.2c178614.svg)CockroachDB、![](http://localhost:9123/assets/SQLServer.bda97784.svg)SQL Server、![](http://localhost:9123/assets/Planetscale.7a27b09b.svg)Plantscale、![](http://localhost:9123/assets/MariaDB.6fe1963e.svg)MariaDB。

### 超图

本质上是由所有数据源内省出的GraphQL  schema组成的超级GraphQL  schema。目的是，用GraphQL协议统一不同类型数据源，便于实现跨源编排和查询。

### OIDC

[OpenID Connect](https://www.cnblogs.com/linianhui/p/openid-connect-core.html) (OIDC) 是一种用于身份验证和授权的开放式协议。它建立在OAuth 2.0基础上，并为第三方应用程序提供了一种方便的方法来验证用户身份并获取用户信息，例如名称、邮件地址等。OIDC还支持单点登录（SSO），以便用户只需在一个地方登录，就可以访问多个应用程序。

### OAuth 2.0 <a href="#toc_0" id="toc_0"></a>

[OAuth 2.0](https://www.cnblogs.com/Wddpct/p/8976480.html#52-implicit-flow)是一种开放的授权协议，允许用户将他们的数据（例如用户名和密码）委托给第三方应用程序，以便在不需要用户为该第三方应用程序提供敏感信息的情况下访问其他资源。OAuth 2.0提供了一种安全、透明的方法，用于授权第三方应用程序访问用户数据，而不需要用户提供其凭据。例如，当你想要登录某个论坛，但没有账号，而这个论坛接入了如 QQ、Facebook 等登录功能，在你使用 QQ 登录的过程中就使用的 OAuth 2.0 协议。

### RBAC

它是一种基于角色的访问控制方法，通过分配组织中个体用户的角色来规范对资源或操作的访问。这允许管理员向特定角色授予权限，而不是个体用户，从而更容易管理大型系统的访问控制。

### HOOKS（钩子）

在编程中，钩子（hooks）是一种机制，允许您在程序执行特定操作时插入代码，以更改或扩展程序的行为。钩子通常用于在不修改程序源代码的情况下对程序进行自定义，或对程序进行调试或监控。







