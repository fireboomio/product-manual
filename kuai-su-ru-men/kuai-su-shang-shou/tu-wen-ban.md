# 图文版

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
