# 快速上手

本文主要介绍从初识飞布到快速了解飞布功能从而搭建第一个应用并有效访问的完整流程。

## 环境准备

### 在线体验

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/fireboomio/fb-init-simple)

[gitpod 介绍](https://juejin.cn/post/6844903773878386701)

注意：启动成功后，在 gitpod 底部切换到`PORTS`面板，选择 `9123` 端口打开即可

### 本地安装

```bash
curl -fsSL https://www.fireboom.io/install.sh | bash -s fireboom-example-project
```

```bash
wget -qO- https://www.fireboom.io/install.sh | bash -s fireboom-example-project
```

### 运行飞布

```shell
./fireboom dev
```

启动成功日志：

```
⇨ http server started on [::]:9123
```

打开控制面板

[http://localhost:9123](http://localhost:9123)

## 快速使用

### 1. 设置数据源

* 数据源
  * GraphQL: https://countries.trevorblades.com

![添加数据源](https://fireboom.oss-cn-hangzhou.aliyuncs.com/img/01-datasource.png)

### 2. 新建 API

&#x20;

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

&#x20;

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
  // TODO: 在此处添加代码
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

要完整体验fireboom，请查看我们的30分钟[fireboom基础教程](https://www.bilibili.com/video/BV1w24y1U7fx/)。

### 数据库操作

数据库建模：学习如何使用飞布建模数据库，参考[prisma文档](https://prisma.yoga/concepts/components/prisma-schema/data-model)

数据CRUD：了解数据库表结构和graphql的映射关系，参考[prisma文档](https://prisma.yoga/concepts/components/prisma-client/crud) ，待补充教程。

### 业务逻辑

实现自定义业务逻辑有几种不同的选项，具体取决于你的用例。

* API钩子：在请求API的生命周期中，插入代码，以更改或扩展API行为，例如用户新建文章后，通过后置钩子发送邮件通知管理员审核
* API数据源：除数据库外，飞布支持集成REST API和GraphQL API，开发者可以自行用喜欢的方式实现自定义逻辑的API，但无需考虑权限问题。飞布此时变身API网关，作为BFF层对外提供接口。
* 自定义数据源：飞布还内置了自定义数据源，开发者可以直接编写脚本扩展逻辑。它本质上也是一个GraphQL API。

