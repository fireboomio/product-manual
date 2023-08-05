# 数据库连接

你可以快速且便捷的连接飞布到新的、现存的或者示例数据库上。你也可以同时连接多个数据库，实现跨数据源的数据编排。

## 数据准备

首先你要准备好数据库的连接配置。

## 新建数据库

例如：新建数据源->选择MySQL，填写数据库配置信息->测试连接，连接成功后，选择保存，该MySQL数据源即新建完成，并保存到了数据库列表中。

{% embed url="https://www.bilibili.com/video/BV1fL411C72D/" %}
04功能介绍 飞布如何新建数据源？
{% endembed %}

数据库保存后，默认为关闭，若想使用数据库，需要在右上角“<mark style="color:purple;">开启</mark>”。开启后，将以数据库名称作为命名空间，将当前数据库合并到[超图](../../../../he-xin-gai-nian/chao-tu.md)中。

![连接数据库](https://www.fireboom.io/images/gif/01-02%E8%BF%9E%E6%8E%A5%E6%95%B0%E6%8D%AE%E5%BA%93.gif)

数据库配置信息有两种方式，一种是连接URL，另一种是连接参数。

{% tabs %}
{% tab title="连接URL" %}
该方式比较通用，所有数据库类型都支持该方式。

#### MySQL

```
mysql://USER:PASSWORD@HOST:PORT/DATABASE
```

![mysql connector](https://www.prisma.io/docs/static/a3179ecce1bf20faddeb7f8c02fb2251/4c573/mysql-connection-string.png)

#### PostgreSQL

```
postgresql://USER:PASSWORD@HOST:PORT/DATABASE
```

![postgresql connector](https://www.prisma.io/docs/static/13ad9000b9d57ac66c16fabcad9e08b7/4c573/postgresql-connection-string.png)

#### SQLite

```
file:./dev.db
```

#### MongoDB

![](https://www.prisma.io/docs/static/b5ef4062c4686c772571b3079ba1331c/4c573/mongodb.png)

```
mongodb://USERNAME:PASSWORD@HOST/DATABASE
```
{% endtab %}

{% tab title="连接参数" %}
| 名称   | 占位符        | 描述                         |
| ---- | ---------- | -------------------------- |
| 主机   | `HOST`     | 数据服务的IP地址或域名，例如`localhost` |
| 端口   | `PORT`     | 数据库服务运行的端口，例如`3600`        |
| 用户   | `USER`     | 数据库用户的名称，例如`janedoe`       |
| 密码   | `PASSWORD` | 数据库用户的密码                   |
| 数据库名 | `DATABASE` | 你想使用的数据库名称，例如 `mydb`       |


{% endtab %}
{% endtabs %}

{% hint style="info" %}
你可以使用环境变量或字面量连接数据库。为了安全起见，推荐使用环境变量连接。在连接URL或用户名+密码字段前选择“环境变量”，即可使用环境变量连接数据库。
{% endhint %}

{% hint style="info" %}
SSH隧道模式

Fireboom暂时不支持直接配置SSH Tunnel连接，你可以通过执行ssh脚本将远程的数据库端口映射到本地，然后连接本地映射后的端口即可。这里有个示例 ssh -L 3306:localhost:3306 database-machine.org 然后Fireboom中使用localhost:3306进行连接即可
{% endhint %}





