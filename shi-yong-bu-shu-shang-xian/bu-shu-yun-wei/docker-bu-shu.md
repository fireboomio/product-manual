# Docker部署

当前Docker镜像支持开发模式`dev`和生产模式`start`。

* 开发模式内置了golang和nodejs环境；
* 生产模式提供了飞布运行所必须得环境；

## 构建镜像

您可以按照下述方法自行构建镜像，也可以使用Fireboom官方构建的镜像，前往 [dockerhub查看](https://hub.docker.com/repository/docker/fireboomapi/fireboom\_server/general) 。

### Dockerfile

下面是构建Fireboom镜像依赖的所有文件，详情前往仓库查看：[https://github.com/lcRuri/fireboom-docker](https://github.com/lcRuri/fireboom-docker)

<details>

<summary>fireboom-docker目录</summary>

{% code title="Dockerfile" %}
```docker
FROM golang:1.20-alpine

MAINTAINER lcRuri

RUN apk update && \
    apk add --no-cache git bash curl

# 安装 Node.js
RUN apk add --no-cache nodejs npm

WORKDIR /fbserver

# 将代码复制到容器中
COPY . .

#RUN chmod +x update-fb.sh
RUN chmod +x host.sh
RUN chmod +x update-fb.sh

# 指定挂载目录
VOLUME /fbserver/log
VOLUME /fbserver/store
VOLUME /fbserver/template
VOLUME /fbserver/upload
VOLUME /fbserver/custom-go
VOLUME /fbserver/custom-ts

EXPOSE 9123
EXPOSE 9991

ENTRYPOINT ["/fbserver/host.sh"]
```
{% endcode %}

{% code title="host.sh" %}
```bash
#!/bin/bash
# 使用install.sh脚本安装Fireboom，并使用 init-todo 模板初始化
sh update-fb.sh

./fireboom build

start_command="/fbserver/fireboom $1"

eval "$start_command"
```
{% endcode %}

{% code title="update-fb.sh" %}
```bash
#!/usr/bin/env bash
curl -fsSL https://www.fireboom.io/update | bash
```
{% endcode %}

</details>

### 构建

在上述fireboom-docker目录下执行如下命令：

```bash
docker build -t fireboom_server:latest
```

## 使用方法

### 1. 拉取镜像

<pre class="language-bash"><code class="lang-bash"><strong># 拉取镜像
</strong>docker pull fireboomapi/fireboom_server:latest
</code></pre>

### 2. 运行容器

* 以开发模式启动，不可用于正式环境！

```bash
# 1，前往工作目录
cd workspace
# 2，以挂载目录的方式运行容器
docker run -it -v $(pwd)/store:/fbserver/store \
		-v $(pwd)/upload:/fbserver/upload \
		-v $(pwd)/template:/fbserver/template \
		-v $(pwd)/exported:/fbserver/exported \
		-v $(pwd)/custom-go:/fbserver/custom-go \
		-v $(pwd)/custom-ts:/fbserver/custom-ts \
		-p 9123:9123 -p 9991:9991 \
		 fireboomapi/fireboom_server:latest dev 
```

{% hint style="info" %}
若是windows系统，请将 **$(pwd)** 替换为\*\*`${pwd}`\*\*
{% endhint %}

* 以生产模式启动

```bash
docker run -it -p 9123:9123 -p 9991:9991 fireboom_server:latest start 
```

**挂载目录**

容器的工作目录为：`./fbserver`，根据需求挂载下述子目录。

* 存储目录： store、upload、exported
* 钩子目录：custom-go 或 custom-ts （不用钩子，无需暴露）
* 日志目录（可选）：log

**端口说明**

* 9123：飞布控制台的端口
* 9991：飞布处理所有api请求的端口

### 3. 使用Fireboom

访问地址：[http://localhost:9123](http://localhost:9123/)



更多内容，如在容器中使用钩子，请参考：[https://github.com/lcRuri/fireboom-doc/blob/main/docker/readme.md](https://github.com/lcRuri/fireboom-doc/blob/main/docker/readme.md)
