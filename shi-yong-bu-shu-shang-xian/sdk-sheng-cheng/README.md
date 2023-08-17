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

## SDK能力

### 查询和变更

todo

### 实时



### 文件上传



### 身份验证



### CSRF保护

