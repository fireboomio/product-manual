# 手动部署

部署Fireboom一般意味着要同时部署：Fireboom服务和钩子服务。

部署时需要将Fireboom服务依赖的配置文件和钩子服务（二进制或代码）上传到服务器。

## 目录介绍

### Fireboom服务依赖的目录

Fireboom服务启动时需要依赖如下目录，详情见下文：

```
├─ .gitignore 
├─ .env # 环境变量，开发环境和部署环境一般不同
├─ custom-go # golang钩子目录，详情见下文
├─ custom-ts # Nodejs钩子目录，详情见下文
├─ log # 日志目录，无需上传
├─ exported # 部署时无需上传，可用./fireboom build 命令生成
├─ store # 必须上传
├─ template # 部署时若不上传，则fireboom启动时会自动下载
└─ upload # 必须上传
```

### 钩子服务依赖的目录

钩子服务依赖的目录，取决于钩子服务的语言类型。

* **脚本语言**：需要上传代码，例如：nodejs typescript
* **编译语言**：只需要上传二进制，例如：golang

{% tabs %}
{% tab title="golang" %}
```
├─ custom-go
│  ├─ main # golang 钩子的二进制
│  ├─ helix.html # 钩子中graphql服务的面板静态文件
│  ├─ generated # 部署时无需上传，可用./fireboom build 命令生成
```
{% endtab %}

{% tab title="nodejs" %}
```
├─ custom-ts
│  ├─ node_modules # 无需上传，依赖 npm install 下载
│  ├─ generated # 部署时无需上传，可用./fireboom build 命令生成
│  ├─ authentication
│  │  └─ mutatingPostAuthentication.ts
│  ├─ ecosystem.config.js
│  ├─ fireboom.server.ts
│  ├─ nodemon.json
│  ├─ operation
│  │  └─ User
│  │     └─ GetOneUser
│  │        └─ postResolve.ts
│  ├─ operations.tsconfig.json
│  ├─ package.json
│  ├─ storage
│  │  └─ tengxunyun
│  │     └─ avatar
│  │        ├─ postUpload.ts
│  │        └─ preUpload.ts
│  ├─ tsconfig.json
```
{% endtab %}
{% endtabs %}

{% hint style="info" %}
golang可跨平台编译，可在开发环境编译出对应版本直接上传！
{% endhint %}

## 同步目录

同步上述目录有两种方案：github（推荐）和工具同步。

### github（推荐）

推荐使用github管理项目，支持版本控制与多人协作。

部分文件不需要上传到开发环境，可以用.gitignore忽略，参考如下：

{% code title=".gitignore" %}
```ignore
# bin file
fireboom
# hook
node_modules
yarn.lock
package-lock.json
pnpm-lock.yaml

# log
log/*
```
{% endcode %}

在生产服务器上用 [`git`](https://git-scm.com/book/zh/v2/%E8%B5%B7%E6%AD%A5-%E5%AE%89%E8%A3%85-Git) 命令拉取项目即可。

### 工具同步

使用IDE或`rsync`将文件同步到服务器，例如：

```bash
rsync -avr  --exclude 'node_modules' --exclude 'fireboom' ./* user@server.ip:/path/to/publish
```

## 环境准备

### 钩子启动环境

不同语言的钩子服务有不同的环境，需要参考各编程语言进行准备：

* typescript钩子：安装nodejs环境，[前往学习](https://github.com/nvm-sh/nvm#installing-and-updating)

```bash
‌curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash
nvm install stable
node -v
```

* golang钩子：无需准备环境

### 环境变量

Fireboom的环境变量存储在根目录下的`.env`文件中，生产环境需要根据实际情况修改。

```properties
FB_API_PUBLIC_URL="http://localhost:9991" # 修改为公网IP
FB_API_LISTEN_HOST="localhost" # 用IP访问，请修改为 0.0.0.0
FB_API_LISTEN_PORT="9991" # 用IP访问，请放开防火墙
FB_API_INTERNAL_URL="http://localhost:9991" 
FB_SERVER_URL="http://localhost:9992"
FB_SERVER_LISTEN_HOST="localhost"
FB_SERVER_LISTEN_PORT="9992"
```

其他环境变量，如数据库URL等，请按照实际情况修改。

## 启动服务

### 构建依赖

启动钩子服务前，需要先执行Fireboom构建命令：

```
# cd‌ [project-name] 进入项目根目录
./fireboom build 
```

它会生成两类产物：

* Fireboom服务**生产模式**依赖的产物，如`exported`目录下的文件
* 钩子服务依赖的文件
  * 所有钩子都依赖的配置，包括custom-x/generated/目录下的  fireboom.config.json和fireboom.operations.json
  * 脚本钩子依赖的代码，例如 nodejs依赖的custom-ts/generated/\*.ts

钩子服务依赖配置，例如

### 启动钩子

根据钩子服务产物的类型执行不同命令，启动钩子服务。

对于编译语言，直接执行二进制即可，例如golang钩子：

```bash
cd custom-go
./main # 启动服务
```

对于脚本语言，一般需要先安装依赖，然后再启动，例如 typescript钩子：

```bash
‌cd custom-ts
# 安装依赖
npm install
# 构建产物
npm run build
# 安装pm2
npm i -g pm2
# 以上命令都只需要执行一次
pm2 start
# 查看启动日志
pm2 logs 0
# 后续重启使用
pm2 restart 0
# 更多pm2使用方法请参考 https://pm2.keymetrics.io/docs/usage/quick-start/
```

### 启动Fireboom

为保障安全，需要用生成模式启动Fireboom：

```bash
# cd‌ [project-name] 进入项目根目录
# 以生产模式启动服务
./fireboom start 
```

该方式将挂起命令行，可以用其他命令来作为守护进程。

#### systemctl守护

在`/usr/lib/systemd/system/` 目录中新建`fb.service` ，内容如下

```sh
‌[Unit]
Description=Fireboom server
After=syslog.target network.target

[Service]
Type=simple
# 根据实际路径来修改
WorkingDirectory=[project-name]
# 根据实际路径来修改
ExecStart=[project-name]/fireboom start
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

然后执行

```sh
# 重新加载
systemctl daemon-reload
# 开机自启
systemctl enable fb
# 启动
systemctl start fb
# 查看日志
systemctl status fb
```

#### PM2守护

```bash
cd [project-name]
pm2 start ./fireboom start
cd custom-ts
pm2 start
# cd custom-go
# pm2 start ./main
```

## Nginx配置

如果想用域名访问，可用nginx代理服务，参考如下配置：

<mark style="color:red;">待修改</mark>

```nginx
upstream backend {
  # server localhost:9991;
  server 127.0.0.1:9991; # Fireboom API 服务监听地址
}

server {
    listen 80;
    server_name your.domain;
    rewrite ^(.*) https://$server_name$1 permanent;
}

# https 配置，没有则将后续配置剪切到上面的80服务里
server {
  listen       443 ssl;
  server_name  your.domain;
  charset utf-8;

  gzip_static on;

  gzip_proxied        expired no-cache no-store private auth;
  gzip_disable        "MSIE [1-6]\.";
  gzip_vary           on;

  # 略，ssl证书配置

  # 配置前端的访问路径
  location / {
    root   /path/to/deploy/web; # 根据实际情况修改
    index  index.html;
    try_files   $uri $uri/ /index.html;
  }
  
  # 配置Fireboom OPERATION 的路由
  location /operations {
    proxy_pass       http://backend/operations;
    proxy_set_header X-Real_IP $remote_addr;
    proxy_set_header Host $host;
    proxy_set_header X_Forward_For $proxy_add_x_forwarded_for;
    client_max_body_size 0;
  }
  # 配置Fireboom 授权相关的路由
  location /auth {
    proxy_pass       http://backend/auth;
    proxy_set_header X-Real_IP $remote_addr;
    proxy_set_header Host $host;
    proxy_set_header X_Forward_For $proxy_add_x_forwarded_for;
    client_max_body_size 0;
  }
}
```

访问：

* Fireboom控制台：http://dashboard.fireboom.io
* Fireboom API：http://api.fireboom.io

