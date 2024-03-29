# CLI

飞布 CLI 是一个命令行工具，可以帮助你创建一个新项目，管理元数据，~~应用迁移~~等。

## 本地安装

使用飞布 CLI 前，需要先完成安装。安装脚本如下：

{% hint style="info" %}
如果你使用的是Windows系统，建议使用 Git bash 执行脚本，或者在`MSYS2`等环境下执行脚本，不支持在`CMD`或者`PowerShell`终端中执行
{% endhint %}

```bash
curl -fsSL https://www.fireboom.io/install.sh | bash -s project-name -t fb-init-todo
```

`project-name`为项目名称，可根据需求更改。

`-t fb-init-todo`为初始化模板，省略后默认创建空项目。

目前系统有如下模板：

* TODO示例：[https://github.com/fireboomio/fb-init-todo](https://github.com/fireboomio/fb-init-todo)
* 空项目：[https://github.com/fireboomio/fb-init-simple](https://github.com/fireboomio/fb-init-simple)

执行上述命令后，将在执行目录下新建名称为`project-name`的文件夹，其目录结构如下：

```
project-name
├─ .env
├─ .gitignore
├─ .gitpod.yml
├─ README.md
├─ fireboom # CLI 命令
├─ custom-go # golang 钩子的目录 见下文
├─ custom-ts # typescript 钩子的目录 见下文
├─ exported 
│  └─ generated # 飞布引擎启动依赖的所有元数据，每次编译都会重新生成
│  └─ operations
│     └─ Todo
│        ├─ CreateOneTodo.graphql
│        ├─ DeleteManyTodo.graphql.off
│        ├─ DeleteOneTodo.graphql
│        ├─ GetManyTodo.graphql
│        ├─ GetOneTodo.graphql.off
│        ├─ GetTodoList.graphql.off
│        ├─ UpdateOneTodo.graphql
│        └─ UpdateTodoCompleted.graphql
├─ hook.sh
├─ log
├─ store # 元数据存储目录
│  ├─ hooks # 钩子相关配置
│  │  ├─ auth # 身份验证钩子配置
│  │  │  └─ mutatingPostAuthentication.config.json
│  │  ├─ customize # 自定义数据源钩子配置
│  │  │  └─ chatGPT.config.json
│  │  ├─ global # 全局钩子配置
│  │  ├─ hooks # 每个OPERATION 钩子配置
│  │  │  └─ Todo
│  │  │     └─ CreateOneTodo
│  │  │        └─ postResolve.config.json
│  │  └─ uploads
│  ├─ list # 列表数据
│  │  ├─ FbAuthentication # 身份验证配置
│  │  ├─ FbDataSource # 数据源配置
│  │  ├─ FbOperation # API配置
│  │  ├─ FbRole # 用户角色配置
│  │  └─ FbStorageBucket # 文件存储配置
│  └─ object # 对象数据
│     ├─ global_config.json
│     ├─ global_operation_config.json
│     ├─ global_system_config.json
│     └─ operations # 每个OPERATION的独立配置
│        └─ Todo
│           ├─ CreateOneTodo.json
│           ├─ DeleteManyTodo.json
│           ├─ DeleteOneTodo.json
│           ├─ GetManyTodo.json
│           ├─ GetOneTodo.json
│           ├─ GetTodoList.json
│           ├─ UpdateOneTodo.json
│           └─ UpdateTodoCompleted.json
├─ template # SDK模板目录
│  ├─ golang-server # golang 钩子模板
│  └─ node-server # TS 钩子模板
└─ upload # 上传文件目录
   ├─ db # sqlite数据库目录
   │  └─ todo.db
   ├─ graphql # graphql 数据源 schema 文件目录
   ├─ oas # rest api数据源 oas 文件目录
   │  ├─ example_rest.json
   │  └─ openapi.json
   ├─ oss 
   └─ swagger

```

如果你用golang开发钩子服务，则需要关注如下目录：

```
├─ custom-go
│  ├─ auth # 授权钩子目录
│  │  └─ mutatingPostAuthentication.go
│  ├─ customize # 自定义数据源钩子目录
│  │  └─ chatGPT.go
│  ├─ go.mod
│  ├─ go.sum
│  ├─ helix.html
│  ├─ hooks # 每个OPERATION 钩子对应目录
│  │  └─ Todo
│  │     └─ CreateOneTodo
│  │        └─ postResolve.go
│  ├─ main.go
│  ├─ pkg
│  ├─ scripts
│  │  ├─ install.sh
│  │  ├─ run-build.sh
│  │  ├─ run-dev.sh
│  │  └─ run-prod.sh
│  └─ generated # 生成的文件，每次配置修改触发编译后，都会覆盖该目录
│  └─ server
│     └─ start.go # 入口文件，每次配置修改触发编译后，都会覆盖该文件
```

如果你用TypeScript开发钩子服务，则需要关注如下目录：（对应文件的功能同上）

```
├─ custom-ts
│  ├─ README.md
│  ├─ auth
│  ├─ customize
│  ├─ global
│  ├─ hooks
│  │  └─ Todo
│  │     └─ CreateOneTodo
│  │        └─ postResolve.ts
│  ├─ nodemon.json
│  ├─ operations
│  ├─ operations.tsconfig.json
│  ├─ package.json
│  ├─ scripts
│  │  └─ buildOperations.ts
│  ├─ tsconfig.json
│  └─ generated # 生成的文件，每次配置修改触发编译后，都会覆盖该目录
│  ├─ fireboom.server.ts # 入口文件，每次配置修改触发编译后，都会覆盖该文件
│  └─ uploads
```

## 开发模式

在`project-name`目录下执行`fireboom dev`命令，飞布将执行下述逻辑：

* 启动控制台：启动控制台，默认访问地址为：http://localhost:<mark style="color:red;">9123</mark>
* 实时编译API：检测配置变更，并将其**实时**编译为REST API。（编译流程如下）

飞布开发的API，默认访问地址为：http://localhost:<mark style="color:orange;">9991</mark>。不同于控制台的访问地址！

{% hint style="info" %}
出于安全考虑，飞布控制台(9123)和API(9991)用不同端口暴露。你可以通过防火墙设置安全策略，以保证安全，如对9991全部放行，对9123设置IP白名单放行。
{% endhint %}

在开发模式下，每次通过界面修改配置，都会触发核心引擎的实时编译流程。

API编译流程如下：

1. 读取store目录下的配置，生成元数据到exported/generated目录
2. 内省数据源，获得各数据源的graphql schema描述
3. 根据启用的模板库，生成对应SDK到指定目录
4. 重启核心引擎，暴露REST API 服务

## 生产模式

在`project-name`目录下执行`fireboom start`命令，飞布将执行下述逻辑：

* 启动控制台：启动控制台，但需要输入秘钥才能访问
* 启动API：根据历史配置启动核心引擎，暴露 REST API 服务。（见上述步骤4）

生产模式下，通过界面修改配置，不会触发自动编译，以保证服务稳定。

## 构建命令

在`project-name`目录下执行`fireboom build`命令，飞布将执行下述逻辑：

* 配置生成，打包生成生产模式依赖的配置文件（见上述步骤1-3）

通常与fireboom start配套使用。

## 升级飞布

```bash
# 升级飞布命令行
# cd project-name
curl -fsSL https://www.fireboom.io/update.sh | bash
```
