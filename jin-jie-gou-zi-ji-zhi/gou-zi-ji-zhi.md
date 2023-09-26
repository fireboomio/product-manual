# 钩子概览

本章将介绍，如何使用飞布的钩子机制扩展API，实现自定义逻辑。

> 在编程中，钩子（hooks）是一种机制，允许您在程序执行特定操作时插入代码，以更改或扩展程序的行为。钩子通常用于在不修改程序源代码的情况下对程序进行自定义，或对程序进行调试或监控。

飞布提供了各种类型的钩子，包括API请求生命周期的钩子、授权生命周期的钩子、文件上传生命周期的钩子，用于扩展逻辑。

钩子本质上是http请求生命周期的“切面”​，在请求经过“切面”时，通过自定义函数，修改请求的输入参数和响应结果。

飞布服务与钩子服务相互独立，服务间通过HTTP协议通讯，可分别部署。钩子服务本质上是一个实现了飞布钩子规范的WEB服务。因此，可以用任意后端开发语言实现钩子，真正做到多语言兼容。

<figure><img src="../.gitbook/assets/image (2) (1) (1) (1) (1) (1) (1) (1).png" alt=""><figcaption><p>飞布服务与钩子服务调用关系</p></figcaption></figure>

不仅，飞布可以调用钩子，钩子也可以调用飞布。此时，飞布相对于钩子是一个数据代理服务，同时飞布的所有OPERATION都可以供钩子使用。在另一个角度上，钩子也可以认为是serverless架构。

## 安装钩子

接着，我们学习如何安装钩子。

在状态栏中有一项叫做“<mark style="color:red;">钩子模板</mark>”，默认为空。点击后，可打开模板页，默认为空。点击右上角“浏览模板市场”，打开模板下载页，里面展示了两种类型的模板：

* 钩子模板，用于生成钩子服务的脚手架
* 客户端模板，用于生成客户端的SDK，详情见 [sdk-sheng-cheng](../shi-yong-bu-shu-shang-xian/sdk-sheng-cheng/ "mention")

今天我们只介绍钩子模板。我们以golang钩子模板为例。

<figure><img src="../.gitbook/assets/image (55).png" alt=""><figcaption></figcaption></figure>

1，在模板下载页，选择 `golang-server` ，点击下载

1.1 Fireboom将从下述仓库中下载模板：

* V2.0版本：[https://github.com/fireboomio/sdk-template\_go-server/tree/V2.0](https://github.com/fireboomio/sdk-template\_go-server/tree/V2.0)

下载后，有两处变更。

1.2 在`store/sdk`目录中新增`golang-server.json`文件，其格式如下：

```json
{
  "name": "golang-server", # 钩子SDK的名称
  "enabled": true,
  "type": "server", # SDK类型，server表示服务端钩子
  "language": "go", # 钩子的语言
  "extension": ".go", # 钩子的后缀
  "gitUrl": "https://code.100ai.com.cn/fireboomio/sdk-template_go-server.git",
  "gitBranch": "V2.0",
  "gitCommitHash": "1e85822a9f2db63a4d97e434765738b5b35e5515", # 记录HASH值，保证稳定拉取
  "outputPath": "./custom-go", # 生成目录
  "codePackage": "",
  "createTime": "2023-08-27T17:36:45.207601+08:00",
  "updateTime": "2023-09-06T17:14:48.241365+08:00",
  "deleteTime": "",
  "generateTime": "2023-09-06T17:14:47.930345005+08:00",
  "icon": "xxx",
  "title": "Golang server",
  "author": "fireboom",
  "version": "latest",
  "description": "Golang hook server SDK template for fireboom"
}
```

1.3 在 `template` 目录下增加 `golang_server` 目录，包含有3部分：

* README.md：自述文件
* files：模板文件
  * .hbs：[handlerbars](https://www.handlebarsjs.cn/)格式，每次编译后替换目标文件
  * 其他文件，例如`.go`、`.html`，目标文件不存在时，生成（否则不生成）
  * <mark style="color:red;">.fbignore</mark>：用于指定强制忽略的文件或目录，格式与.gitignore一致
* partials（可选）：每个文件都是一个模板片段，主要用于复用逻辑，可在files的hbs中引用

{% tabs %}
{% tab title="golang" %}
```
template
├─ golang-server
│  ├─ README.md
│  ├─ files
│  │  ├─ generated
│  │  │  └─ models.go.hbs
│  │  ├─ .fbignore # 用于指定要忽略的文件或目录
│  │  ├─ go.mod
│  │  ├─ helix.html
│  │  ├─ main.go
│  │  ├─ pkg
│  │  │  ├─ base
│  │  │  │  ├─ client.go
│  │  │  │  ├─ operation.go
│  │  │  │  ├─ request.go
│  │  │  │  ├─ upload.go
│  │  │  │  └─ user.go
│  │  │  ├─ consts
│  │  │  │  └─ env.go
│  │  │  ├─ plugins
│  │  │  │  ├─ graphqls.go
│  │  │  │  └─ internal_request.go
│  │  │  ├─ types
│  │  │  │  ├─ configure.go
│  │  │  │  └─ server.go
│  │  │  ├─ utils
│  │  │  │  ├─ config.go
│  │  │  │  ├─ file.go
│  │  │  │  ├─ http.go
│  │  │  │  ├─ random.go
│  │  │  │  └─ strings.go
│  │  │  └─ wgpb
│  │  │     └─ wundernode_config.pb.go
│  │  ├─ scripts
│  │  │  ├─ install.sh
│  │  │  ├─ run-build.sh
│  │  │  ├─ run-dev.sh
│  │  │  └─ run-prod.sh
│  │  └─ server
│  │     ├─ fireboom_server.go.hbs # 入口文件模板
│  │     └─ start.go
```
{% endtab %}

{% tab title="nodejs" %}
```
template
├─ node-server
│  ├─ README.md
│  ├─ files
│  │  ├─ ecosystem.config.js
│  │  ├─ fireboom.server.ts.hbs
│  │  ├─ generated
│  │  │  ├─ claims.ts.hbs
│  │  │  ├─ client.legacy.ts.hbs
│  │  │  ├─ client.ts.hbs
│  │  │  ├─ fireboom.factory.ts
│  │  │  ├─ fireboom.internal.client.ts.hbs
│  │  │  ├─ fireboom.internal.operations.client.ts.hbs
│  │  │  ├─ fireboom.operations.ts.hbs
│  │  │  ├─ fireboom.server.ts.hbs
│  │  │  ├─ linkbuilder.ts.hbs
│  │  │  ├─ models.ts.hbs
│  │  │  ├─ prisma.ts.hbs
│  │  │  └─ testing.ts.hbs
│  │  ├─ nodemon.json
│  │  ├─ operations.tsconfig.json
│  │  ├─ package.json
│  │  ├─ scripts
│  │  │  ├─ buildOperations.ts
│  │  │  ├─ install.sh
│  │  │  ├─ run-build.sh
│  │  │  ├─ run-dev.sh
│  │  │  └─ run-prod.sh
│  │  └─ tsconfig.json
│  ├─ manifest.json
│  └─ partials
│     ├─ operation_partial.hbs
│     └─ schema_partial.hbs
```
{% endtab %}
{% endtabs %}

2，在模板页修改“生成路径”，并开启钩子开关（钩子模板同时只能开启1个）

3，后续每次触发“编译”，都会重新生成钩子文件（<mark style="color:orange;">非</mark><mark style="color:orange;">`.hbs`</mark><mark style="color:orange;">文件或.fbignore指定的文件，只生成1次</mark>）

{% tabs %}
{% tab title="golang" %}
```
├─ custom-go
│  ├─ .fbignore # 在里面指定覆盖时，需要排除的文件
│  ├─ go.mod
│  ├─ go.sum
│  ├─ helix.html
│  ├─ main.go
│  ├─ pkg
│  │  ├─ base
│  │  │  ├─ client.go
│  │  │  ├─ operation.go
│  │  │  ├─ request.go
│  │  │  ├─ upload.go
│  │  │  └─ user.go
│  │  ├─ consts
│  │  │  └─ env.go
│  │  ├─ plugins
│  │  │  ├─ auth_hooks.go
│  │  │  ├─ global_hooks.go
│  │  │  ├─ graphqls.go
│  │  │  ├─ internal_request.go
│  │  │  ├─ operation_hooks.go
│  │  │  ├─ proxy_hooks.go
│  │  │  └─ upload_hooks.go
│  │  ├─ types
│  │  │  ├─ configure.go
│  │  │  └─ server.go
│  │  ├─ utils
│  │  │  ├─ config.go
│  │  │  ├─ file.go
│  │  │  ├─ http.go
│  │  │  ├─ random.go
│  │  │  └─ strings.go
│  │  └─ wgpb
│  │     └─ wundernode_config.pb.go
│  ├─ scripts
│  │  ├─ install.sh
│  │  ├─ run-build.sh
│  │  ├─ run-dev.sh
│  │  └─ run-prod.sh
│  ├─ server
│  │  └─ start.go
│  │  └─ fireboom_server.go # 生成的入口文件（重要！！！）
```
{% endtab %}

{% tab title="nodejs" %}
```
todo
```
{% endtab %}
{% endtabs %}

4，按照各语言的方法安装依赖，并启动钩子服务即可

钩子服务本质上是web服务，该服务会按照钩子规范注册对应路由，我们只需要在路由的控制器中完善业务逻辑即可，详情见 [qi-dong-gou-zi](qi-dong-gou-zi/ "mention")

## 升级钩子

接着，基于上述原理 ，学习如何升级模板。

### 1，检查更新

若钩子模板有更新，则会在界面上展示 `new` 标签，点击可对比模板变更

<figure><img src="../.gitbook/assets/image (1) (1) (1).png" alt=""><figcaption></figcaption></figure>

例如：[https://github.com/fireboomio/sdk-template\_go-server/compare/ab1427e..e7fa762](https://github.com/fireboomio/sdk-template\_go-server/compare/ab1427e..e7fa762)

### 2，升级模板

在模板页选择 `golang-server` ，点击“...”，选择“升级”，将下载最新模板，并覆盖旧模板。

注意 `.fbignore` 中声明的文件不会被更新！

### 3，重装依赖

根据钩子语言，重装依赖，例如：

```bash
go mod tidy
```
