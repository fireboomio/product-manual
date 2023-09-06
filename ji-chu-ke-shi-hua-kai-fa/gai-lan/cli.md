# CLI

飞布 CLI 是一个命令行工具，可以帮助你创建一个新项目，管理元数据，~~应用迁移~~等。

## 本地安装

使用飞布 CLI 前，需要先完成安装。安装脚本如下：

{% hint style="info" %}
如果你使用的是Windows系统，建议使用 Git bash 执行脚本，或者在`MSYS2`等环境下执行脚本，不支持在`CMD`或者`PowerShell`终端中执行
{% endhint %}

```bash
curl -fsSL fireboom.io/install | bash -s project-name -t init-todo --cn
```

* 项目名称：`project-name`，可根据需求修改
* 初始化模板：`-t init-todo`，省略后默认创建空项目
* 选择源：`--cn` ，指定从国内源下载，省略后从 github源下载

执行上述命令后，将在执行目录下新建名称为`project-name`的文件夹，其目录结构如下：

```
project-name
├─ .gitignore
├─ .gitpod.yml
├─ README.md
├─ custom-go # golang 钩子的目录 见下文
├─ custom-ts # typescript 钩子的目录 见下文
├─ exported # 生成的临时目录
│  └─ introspection
│     ├─ system
│     │  └─ schema.graphql
│     └─ todo
│        ├─ schema.graphql
│        └─ schema.prisma
├─ store # 元数据存储目录
│  ├─ config # 全局配置
│  │  ├─ global.operation.json
│  │  └─ global.setting.json
│  ├─ datasource #数据源配置，每个数据源1个json
│  │  ├─ system.json
│  │  └─ todo.json
│  ├─ operation # OPERATION存储，.json和.graphql成对存在
│  │  └─ Todo
│  │     ├─ CreateOneTodo.graphql
│  │     ├─ CreateOneTodo.json
│  │     ├─ GetManyTodo.graphql
│  │     ├─ GetManyTodo.json
│  └─ role # 角色存储，每个角色1个json
│     ├─ admin.json
│     └─ user.json
├─ template # SDK模板目录
│  ├─ golang-server # golang 钩子模板
│  └─ node-server # TS 钩子模板
├─ update-fb.sh
└─ upload # 上传文件目录
   ├─ oas
   │  └─ system.json
   └─ sqlite
      └─ todo.db
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

```bash
cd project-name
./fireboom dev
```

飞布将执行下述逻辑：

* 启动控制台：启动控制台，默认访问地址为：`http://localhost:`<mark style="color:red;">`9123`</mark>
* 实时编译API：检测配置变更，并将其**实时**编译为REST API。（编译流程如下）

飞布开发的API，默认访问地址为：`http://localhost:`<mark style="color:orange;">`9991`</mark>。不同于控制台的访问地址！

{% hint style="info" %}
飞布控制台(9123)和API(9991)用不同端口暴露。为保证安全，建议开启9123的授权保护，参考 [#kai-qi-mian-ban-bao-hu](cli.md#kai-qi-mian-ban-bao-hu "mention")
{% endhint %}

在开发模式下，每次通过界面修改配置，都会触发核心引擎的实时编译流程。

API编译流程如下：

1. 读取`store`目录下的配置，编译到目录：`exported/generated`
2. 内省数据源，获得各数据源的graphql schema描述，存储在目录：`exported/introspection`
3. 根据启用的模板库，生成对应SDK到指定目录，模板位于目录：`template`
4. 重启核心引擎，注册GraphQL Operation路由，对外暴露REST API 服务

### 帮助命令

```bash
# 查看支持的所有参数
./fireboom dev --help  
Start the fireboom application in development mode and watch for changes

Usage:
  fireboom dev [flags]

Examples:
./fireboom dev

Flags:
      --active string     Mode active to run in different environment
      --enable-auth       Whether enable auth key on dev mode
  -h, --help              help for dev
      --web-port string   Web port to listen on (default "9123")
```

### 指定面板端口

默认情况下，控制台监听端口：`9123`，通过参数`--web-port`可指定端口。

```bash
fireboom start --web-port 9123
```

### 指定环境变量

默认情况下，环境变量用文件：`.env` 。

使用 `--active` 参数，指定启用的环境变量。

```bash
./fireboom dev --active test
```

例如，上述命令指定激活的环境变量文件为：`.env.test` 。

### 开启面板保护

使用参数`--enable-auth`，开启面板保护，开启后输入秘钥才能访问控制台。

```bash
./fireboom dev --enable-auth
```

秘钥位于根目录的`authentication.key`文件中。

## 生产模式

```bash
cd project-name
./fireboom start
```

飞布将执行下述逻辑：

* 启动控制台：启动控制台，但需要输入秘钥才能访问
* 启动API：根据历史配置启动核心引擎，暴露 REST API 服务。（见上述步骤4）

生产模式下，通过界面修改配置，不会触发自动编译，以保证服务稳定。

### 帮助命令

<pre class="language-bash"><code class="lang-bash"># 查看支持的所有参数
<strong>./fireboom start --help     
</strong>Start the fireboom application in production mode without watching

Usage:
  fireboom start [flags]

Examples:
./fireboom start

Flags:
      # 同上
      --active string        Mode active to run in different environment (default "prod")
      --enable-swagger       Whether enable swagger in production
      --enable-web-console   Whether enable web console page in production (default true)
  -h, --help                 help for start
      --regenerate-key       Whether to renew authentication key in production
      # 同上
      --web-port string      Web port to listen on (default "9123")
</code></pre>

### 开启swagger文档

生产模式默认关闭swagger文档，使用参数`--enable-swagger`开启：

```bash
./fireboom start --enable-swagger
```

### 重新生成秘钥

使用参数`--regenerate-key`，更新秘钥：

```
./fireboom start --regenerate-key
```

## 构建命令

构建 [#sheng-chan-mo-shi](cli.md#sheng-chan-mo-shi "mention")所需要的产物，与`fireboom start`配套使用。

```bash
cd project-name
./fireboom build
```

执行逻辑同 `fireboom dev` 的1-3步骤。

## 升级飞布

升级Fireboom到最新版本

```bash
# 升级飞布命令行
# cd project-name
curl -fsSL www.fireboom.io/update | bash
```
