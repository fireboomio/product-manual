# 快速上手

本文主要介绍从初识飞布到快速了解飞布功能从而搭建第一个API并有效访问的完整流程。

## 前置知识

### 公共知识

* GraphQL：了解什么是GraphQL，掌握基本概念就行，推荐教程 [前往](https://graphql.cn/learn/)
* Prisma ORM：了解Prisma的基本函数签名，推荐教程 [前往](https://www.prisma.io/docs/reference/api-reference/prisma-client-reference#findunique)

### 高级特性

如果你要使用钩子等高级特性，则需要掌握一种后端开发语言。

如果你是前端开发者，推荐：

* TypeScript：了解node.js并熟悉TypeScript语法，推荐教程 [前往](https://typescript.bootcss.com/tutorials/typescript-in-5-minutes.html)

如果你是后端开发者，推荐：

* Golang：了解Golang基本语法即可，推荐教程 [前往](https://www.runoob.com/go/go-tutorial.html)

## 环境准备

{% embed url="https://www.bilibili.com/video/BV1mo4y1B78g" %}
00功能介绍-如何安装或升级飞布？
{% endembed %}

### 在线体验

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/fireboomio/fb-init-simple)

> [gitpod 介绍](https://juejin.cn/post/6844903773878386701)：Gitpod是一个在线IDE，可以从任何GitHub页面启动。在几秒钟之内，Gitpod就可以为您提供一个完整的开发环境，包括一个VS Code驱动的IDE和一个可以由项目定制化配置的云Linux容器。

{% hint style="info" %}
启动成功后，在 gitpod 底部切换到`PORTS`面板，选择 `9123` 端口打开即可
{% endhint %}

### Docker运行

<pre class="language-bash"><code class="lang-bash"><strong># 拉取镜像
</strong><strong>docker pull fireboomapi/fireboom_server:latest
</strong><strong># 运行镜像
</strong>docker run  -p 9123:9123 -p 9991:9991 -p 9992:9992 fireboomapi/fireboom_server:latest test
</code></pre>

打开控制面板，使用如下地址进行访问：

[http://localhost:9123](http://localhost:9123)

### 本地安装

#### 脚本安装

{% hint style="info" %}
如果你使用的是Windows系统，建议使用 Git bash 执行脚本，或者在`MSYS2`等环境下执行脚本，不支持在`CMD`或者`PowerShell`终端中执行
{% endhint %}

```bash
curl -fsSL https://www.fireboom.io/install.sh | bash -s project-name -t fb-init-todo
```

`project-name`为项目名称，可根据需求更改。

`-t fb-init-todo`为初始化模板，省略后默认创建空项目。

{% hint style="info" %}
飞布采用golang语言编写，上述版本基于golang的跨平台编译构建。如果你的操作系统不在上述列表，请[联系我们](https://github.com/fireboomio/product-manual/discussions)兼容。
{% endhint %}

#### 升级飞布

```bash
# 升级飞布命令行
# cd project-name
curl -fsSL https://www.fireboom.io/update.sh | bash
```

### 运行飞布

```shell
# 开发环境
# cd project-name
./fireboom dev
```

启动成功日志：

```
⇨ http server started on [::]:9123
```

打开控制面板

[http://localhost:9123](http://localhost:9123)

## 快速使用

简单使用，只需要观看本视频即可，后面内容可忽略。

{% embed url="https://www.bilibili.com/video/BV1rM411u7e8" %}
01入门教程-如何快速上手飞布？
{% endembed %}

### 1. 设置数据源

* 数据源
  * GraphQL: https://countries.trevorblades.com

![添加数据源](https://fireboom.oss-cn-hangzhou.aliyuncs.com/img/01-datasource.png)

### 2. 新建 API

<figure><img src="https://fireboom.oss-cn-hangzhou.aliyuncs.com/img/02-api_create.png" alt=""><figcaption><p>API新建</p></figcaption></figure>

{% code title="API 名称：GetCountry" %}
```graphql
query MyQuery($code: ID!) {
  country: countries_country(code: $code) {
    capital
    code
    currency
    emoji
    emojiU
    native
    phone
    name
  }
}
```
{% endcode %}

### 3. 扩展 API

<figure><img src="https://fireboom.oss-cn-hangzhou.aliyuncs.com/img/02-api_hooks.png" alt=""><figcaption><p>钩子编写</p></figcaption></figure>

{% code title="mutatingPostResolve.ts" %}
```typescript
import type { Context } from "@wundergraph/sdk";
import type { User } from "generated/wundergraph.server";
import type { InternalClient } from "generated/wundergraph.internal.client";
import { InjectedGetCountryInput, GetCountryResponse } from "generated/models";

// 在左侧引入当前包
import axios from "axios";

export default async function mutatingPostResolve(
  ctx: Context<User, InternalClient>,
  input: InjectedGetCountryInput,
  response: GetCountryResponse
): Promise<GetCountryResponse> {
  var country = response.data?.country;
  if (country) {
    country.phone = "fireboom/test"; //这里可以修改返回值
  }
  ctx.log.info("test");

  //触发一个post请求，给企业机器人发送一个消息
  var res = await axios.post(
    "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=[YOUR KEY]",
    {
      msgtype: "markdown",
      markdown: {
        content: `<font color="warning">${
          ctx.clientRequest.method
        }</font>/n输入：${JSON.stringify(input)}/n响应：${JSON.stringify(
          response
        )}`,
      },
    }
  );
  ctx.log.info("mutatingPostResolve SUCCESS:", res.data);
  return response;
}
```
{% endcode %}

### 4. 身份验证（待完善提供前端示例）

![开启身份验证](https://fireboom.oss-cn-hangzhou.aliyuncs.com/img/02-api\_auth.png)

### 5. 角色鉴权（待完善提供前端示例）

![开启角色鉴权](https://fireboom.oss-cn-hangzhou.aliyuncs.com/img/02-api\_rbac.png)

### 6.实时 API

![实时API](https://fireboom.oss-cn-hangzhou.aliyuncs.com/img/02-api\_live.png)

实时推送数据源：https://hasura.io/learn/graphql/graphiql

```
> graphql 端点:https://hasura.io/learn/graphql
> 请求头：
> Authorization:xxxxxxxxx(前往查看↑)
> content-type:application/json
```

```graphql
# 在测试数据源中插入一条todo，可以看到当前端点会实时拿到新数据
subscription MySubscription {
  todo_todos(order_by: { id: desc }, limit: 10) {
    id
    is_completed
    is_public
  }
}

# 新开页签插入数据
# mutation {
#   insert_todos_one(object: {is_public: false, title: "sssss"}) {
#     id
#   }
# }
```

### 7.其他特性

![API其他特性](https://fireboom.oss-cn-hangzhou.aliyuncs.com/img/02-api\_feature.png)

## 下一步

### 学习教程

体验fireboom更多特性，可前往B站查看完整教学视频，[前往](https://space.bilibili.com/3493080529373820)。

### 数据库操作

* [数据库建模](../kai-fa-wen-dang/shu-ju-yuan/shu-ju-ku/shu-ju-jian-mo.md)：学习如何使用飞布建模数据库，参考[prisma文档](https://prisma.yoga/concepts/components/prisma-schema/data-model)
* [数据库CRUD](../kai-fa-wen-dang/api-gou-jian/ke-shi-hua-kai-fa.md#chao-tu-schema-mian-ban)：了解数据库表结构和graphql的映射关系，参考[prisma文档](https://prisma.yoga/concepts/components/prisma-client/crud) 。

### 业务逻辑

实现自定义业务逻辑有几种不同的选项，具体取决于你的用例。

* [API钩子](../kai-fa-wen-dang/gou-zi-ji-zhi/)：在请求API的生命周期中，插入代码，以更改或扩展API行为，例如用户新建文章后，通过后置钩子发送邮件通知管理员审核。
* [API数据源](../kai-fa-wen-dang/shu-ju-yuan/san-fang-api/)：除数据库外，飞布支持集成REST API和GraphQL API，开发者可以自行用喜欢的方式实现自定义逻辑的API，但无需考虑权限问题。飞布此时变身API网关，作为BFF层对外提供接口。
* [自定义数据源](../kai-fa-wen-dang/shu-ju-yuan/san-fang-api/zi-ding-yi-api.md)：飞布还内置了自定义数据源，开发者可以直接编写脚本扩展逻辑。它本质上也是一个GraphQL API。
* [组合式API](../kai-fa-wen-dang/api-gou-jian/zu-he-shi-api.md)：适用于复杂业务逻辑的构建，当前只支持TS hooks。
