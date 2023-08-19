# 客户端SDK

飞布以开发者体验优先，为提升开发者procode的效率，根据模板实时生成各语言的SDK，且支持用户自行扩展。

飞布的SDK包含两种类型：服务端钩子SDK和客户端SDK。

* 服务端钩子SDK：生成的代码运行在服务端上，主要用途是钩子开发，详情见 [gou-zi-ji-zhi.md](../../jin-jie-gou-zi-ji-zhi/gou-zi-ji-zhi.md "mention")
* 客户端SDK：生成调用REST API的客户端SDK，主要运行在浏览器中

## 安装客户端SDK

安装客户端SDK和安装钩子类似。

在状态栏中有一项叫做“<mark style="color:red;">客户端模板</mark>”，默认为0。点击后，可打开模板页，默认为空。点击右上角“浏览模板市场”，打开模板下载页，选择客户端模板。

我们以`ts-client`客户端模板为例。

1，在模板下载页，选择 `ts-client` ，点击下载

1.1 Fireboom将从下述仓库中下载模板：

* V1.0版本：[https://github.com/fireboomio/sdk-template\_ts-client](https://github.com/fireboomio/sdk-template\_ts-client)
* V2.0版本：[https://github.com/fireboomio/sdk-template\_ts-client/tree/test](https://github.com/fireboomio/sdk-template\_ts-client/tree/test)

不同Fireboom版本，下载模板的分支不同！

1.2 下载后，将在 `template` 目录下增加`ts-client` 目录

其生成规则和 钩子SDK 一致，详情参考 [#an-zhuang-gou-zi](../../jin-jie-gou-zi-ji-zhi/gou-zi-ji-zhi.md#an-zhuang-gou-zi "mention")

2，在模板页修改“生成路径”，并开启开关

3，后续每次触发“编译”，都会重新生成文件（<mark style="color:orange;">非</mark><mark style="color:orange;">`.hbs`</mark><mark style="color:orange;">文件，只生成1次</mark>）

```
sdk
└─ ts-client
   ├─ claims.ts
   ├─ client.ts
   ├─ index.ts  # 入口文件，暴露 client
   └─ models.ts # 结构体定义
```

4，在客户端引入该SDK即可使用

```typescript
import { client } from "./ts-client"
```

## 使用SDK

`ts-client`是Fireboom [HTTP协议](../../ji-chu-ke-shi-hua-kai-fa/api-gou-jian/ke-shi-hua-gou-jian/api-gui-fan.md)在TypeScript中的基本实现，可以在浏览器和服务器环境中使用。它被用作Web客户端实现的基本接口。

### 引用SDK

#### 配置基本请求头

在index.ts文件中增加baseURL配置，自定义访问域名。

{% code title="index.ts" lineNumbers="true" %}
```typescript
import { createClient } from './client';

export const client = createClient({
  // 如果需要修改访问域名，可以在这里配置
  baseURL: 'https://my-custom-base-url.com', 
});
```
{% endcode %}

#### Nodejs支持

使用customFetch配置选项，在没有内置fetch实现的服务器环境中使用SDK。&#x20;

安装node-fetch：

```bash
npm i node-fetch
```

并将其添加到SDK配置中。

{% code title="index.ts" %}
```typescript
import fetch from 'node-fetch';

const client = createClient({
  customFetch: fetch,
});
```
{% endcode %}

#### 自定义请求头

{% code title="index.ts" %}
```typescript
const client = createClient({
  extraHeaders: {
    customHeader: 'value',
  },
});

// or

client.setExtraHeaders({
  customHeader: 'value',
});
```
{% endcode %}

### 查询和变更

#### 查询

```typescript
const response = await client.query({
  operationName: 'Hello',
  input: {
    hello: 'World',
  },
});
```

#### 变更

```typescript
const response = await client.mutate({
  operationName: 'SetName',
  input: {
    name: 'WunderGraph',
  },
});
```

### 实时

#### 实时查询

```typescript
client.subscribe(
  {
    operationName: 'Hello',
    input: {
      name: 'World',
    },
    liveQuery: true, //  开启实时查询
  },
  (response) => {}
);
```

#### 订阅

```typescript
client.subscribe(
  {
    operationName: 'Countdown',
    input: {
      from: 100,
    },
  },
  (response) => {}
);
```

#### 订阅一次

使用`subscribeOnce`运行订阅，这将直接返回订阅响应，而不会推流。适用于SSR场景。

```typescript
const response = await client.subscribe(
  {
    operationName: 'Countdown',
    input: {
      from: 100,
    },
    subscribeOnce: true,
  },
  (response) => {}
);
```

### 文件上传

```typescript
async function handleFiles(e: Event) {
  const files = (e.target as HTMLInputElement)?.files;
  if (files == null) return
  // 调用上传函数
  const res = await client.uploadFiles({
    provider: "tengxunyun",
    files: files,
    profile: "avatar" // （可选）
  })
  console.log("https://test-1314985928.cos.ap-nanjing.myqcloud.com/"+res.fileKeys[0])
}
```

### 身份验证

身份认证包含两种模式：授权码模式（基于cookie）和隐式模式（基于token）。

#### 授权码模式

**登录**

```typescript
client.login('auth0');
```

**获取用户**

```typescript
const user = await client.fetchUser();
```

**退出登录**

```typescript
client.logout({
  logoutOpenidConnectProvider: true,
});
```

#### 隐式模式

**获取Token**

隐式模式登录的核心是获取`access_token`，有两种模式：使用[OIDC登录页](../../ji-chu-ke-shi-hua-kai-fa/shen-fen-yan-zheng/yin-shi-mo-shi.md#id\_token-flow)和[不使用登录页](../../ji-chu-ke-shi-hua-kai-fa/shen-fen-yan-zheng/yin-shi-mo-shi.md#deng-lu-jie-kou)。

一般不使用OIDC的登录页，而是根据各供应商的实现，调用其登录接口获取access\_token，例如：

* Authing： [#deng-lu-jie-kou](../../ji-chu-ke-shi-hua-kai-fa/shen-fen-yan-zheng/yin-shi-mo-shi.md#deng-lu-jie-kou "mention")
* fb-oidc：[前往](https://github.com/fireboomio/fb-oidc)

**使用Token**

{% code title="index.ts" %}
```typescript
export const client = createClient({
    extraHeaders: {
        // 这里注入请求头
        Authorization: 'Bearer <access_token>',
      },
})
```
{% endcode %}

### CSRF保护

// todo

### 中断请求

如果要中断请求，请使用AbortController实例。

```typescript
const controller = new AbortController();

const { fileKeys } = await client.uploadFiles({
  abortSignal: abortController.signal,
  provider: S3Provider.minio,
  files,
});

// cancel the request
controller.abort();
```

## 错误处理

OPERATION API错误分为3类：`InputValidationError`、`GraphQLResponseError`、`ResponseError`。

详情见，AI规范-> [#zhuang-tai-ma](../../ji-chu-ke-shi-hua-kai-fa/api-gou-jian/ke-shi-hua-gou-jian/api-gui-fan.md#zhuang-tai-ma "mention")

### InputValidationError

状态码为`400`时，表示入参校验错误，入参不符合规范，或不满足`@jsonSchema`指令校验规则。

```json
# statusCode=400
{
  "code": "InputValidationError",
  "message": "Bad Request: Invalid input",
  "input": {
    "id": 11
  },
  "errors": [
    {
      "propertyPath": "/id",
      "invalidValue": 11,
      "message": "must be less than or equal to 10"
    }
  ]
}
```

### GraphQLResponseError

状态码为`200`，但执行OPERATION失败，常见错误场景如下：

1，数据源中断

```json
# statusCode=200
{
    "errors": [
        {
            "message": "Error in connector: Error creating a database connection. (Error creating a database connection.)",
            "locations": null,
            "path": null
        }
    ],
    "data": {
        "rb_findUniqueT": null
    }
}
```

2，主键冲突

<pre class="language-json"><code class="lang-json"><strong># statusCode=200
</strong>{
  "errors": [
    {
      "message": "Error occurred during query execution:\nConnectorError(ConnectorError { user_facing_error: Some(KnownError { message: \"Unique constraint failed on the constraint: `PRIMARY`\", meta: Object {\"target\": String(\"PRIMARY\")}, error_code: \"P2002\" }), kind: UniqueConstraintViolation { constraint: Index(\"PRIMARY\") } })",
      "locations": null,
      "path": null
    }
  ],
  "data": {
    "rb_createOneT": null
  }
}
</code></pre>

### ResponseError

状态码非200时的其他错误场景，例如：

* 404：接口临时关闭时，请求不存在
* 503：服务不可用时

### 用例

```typescript
import {GraphQLResponseError,ResponseError,InputValidationError} from "fireboom-wundersdk/client";

const { data, error } = await client.query({
  operationName: 'Hello',
  input: {
    hello: 'World',
  },
});
console.log(data)
if (error != undefined) {
  if (error instanceof InputValidationError) {
  // 1，处理入参校验错误
    console.log("meet InputValidationError:", error.message)
  } else if (error instanceof GraphQLResponseError) {
  // 2，处理执行OPERATION错误
    console.log("meet GraphQLResponseError:", error.errors[0].location)
  } else if (error instanceof ResponseError) {
  // 3，处理其他错误
    console.log("meet ResponseError:", error.statusCode)
  } else {
    console.log("meet UnknownError:" , JSON.stringify(error))
  }
}
```
