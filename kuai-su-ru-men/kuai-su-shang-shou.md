# 快速上手

本文主要介绍从初识飞布到快速了解飞布功能从而搭建第一个应用并有效访问的完整流程。

## 在线体验

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/fireboomio/fb-init-simple)

[gitpod 介绍](https://juejin.cn/post/6844903773878386701)

注意：启动成功后，在 gitpod 底部切换到`PORTS`面板，选择 `9123` 端口打开即可

## 本地安装

```bash
curl -fsSL https://www.fireboom.io/install.sh|bash -s fireboom-example-project

wget -qO- https://www.fireboom.io/install.sh|bash -s fireboom-example-project
```

### 运行

```shell
./fireboom.sh
# or run "./fireboom.sh init" to re-init
```

启动成功日志：

```
⇨ http server started on [::]:9123
```

打开控制面板

[http://localhost:9123](http://localhost:9123)

### 调试钩子

1. 前往配置修改钩子的启动模式为默认不启动（TODO:待实现该配置）
2. 打开./wundergraph/package.json 文件
3. 鼠标悬浮在 scripts.hook 上，点击`调试脚本`
4. 前往 wundergraph/.wundergraph/generated/bundle/server.js 中打断点

### 更新

```shell
# 同时更新命令行和前端资源
./fireboom.sh update
```

```shell
# 仅更新前端资源
./fireboom.sh updatefront
```

### 展示版本

```shell
./fireboom.sh version
```

### 快速使用

#### 1. 设置数据源

* 数据源
  * GraphQL: https://countries.trevorblades.com

![添加数据源](https://fireboom.oss-cn-hangzhou.aliyuncs.com/img/01-datasource.png)

#### 2. 新建 API

![新建API](https://fireboom.oss-cn-hangzhou.aliyuncs.com/img/02-api\_create.png) API 名称：GetCountry

```
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

#### 3. 扩展 API

![编写API钩子](https://fireboom.oss-cn-hangzhou.aliyuncs.com/img/02-api\_hooks.png) mutatingPostResolve.ts

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
    "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=69aa957f-7c05-49b3-9e9d-8859a53ea692",
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

#### 4. 身份验证（待完善提供前端示例）

![开启身份验证](https://fireboom.oss-cn-hangzhou.aliyuncs.com/img/02-api\_auth.png)

#### 5. 角色鉴权（待完善提供前端示例）

![开启角色鉴权](https://fireboom.oss-cn-hangzhou.aliyuncs.com/img/02-api\_rbac.png)

#### 6.实时 API

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

#### 7.其他特性

![API其他特性](https://fireboom.oss-cn-hangzhou.aliyuncs.com/img/02-api\_feature.png)

