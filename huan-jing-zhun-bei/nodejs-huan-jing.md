# NodeJs环境

在 Fireboom 中使用 NodeJs 钩子时，你需要提前准备 NodeJs 环境。

* 如果你使用的是 Windows 系统，请前往[https://nodejs.org/en/download/](https://nodejs.org/en/download/)下载安装最新 NodeJs
* 如果你是 MacOs 或 Linux 系系统，建议使用 `nvm` 进行安装

```console
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash
export NVM_DIR="$([ -z "${XDG_CONFIG_HOME-}" ] && printf %s "${HOME}/.nvm" || printf %s "${XDG_CONFIG_HOME}/nvm")"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" # This loads nvm
nvm install stable
```

安装完成后使用命令检查

```console
# 尽量大于16
node -v
```

在 Fireboom 的最佳实践中，用户登录、授权、校验、角色管理等都应该交由 OIDC 服务来处理，Fireboom 支持常见的一些 OIDC 服务商，在使用前你需要先准备好其中的一个或多个服务，下面是部分常见服务商的配置获取方法。
