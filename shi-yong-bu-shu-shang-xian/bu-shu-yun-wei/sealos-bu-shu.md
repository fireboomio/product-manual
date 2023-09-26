# sealos部署

### 一键部署

1，前往[Sealos模板市场](https://fastdeploy.cloud.sealos.io/)，找到Fireboom，点击“**Deploy**”

<figure><img src="https://bar9vnf09af.feishu.cn/space/api/box/stream/download/asynccode/?code=ODBhZTVkZDc3ZGQ1YmE0NDJjNzkzMTlhMWM5ZWI2ZWVfRjZsYWppYk1GVFR3WkZrbk0wTDg0MzU4bG03cXdvRVFfVG9rZW46VzVuVmJPS3RNb2ZWRlp4dWNuRmNIV2lCbnhoXzE2OTU3MjA1NDk6MTY5NTcyNDE0OV9WNA" alt=""><figcaption></figcaption></figure>

2，进入sealos控制台，查看应用

<figure><img src="https://bar9vnf09af.feishu.cn/space/api/box/stream/download/asynccode/?code=N2Q4NGE1YmJjOWNlYTMwNTI4MTU4ZTAwMjQyZmNmNjFfNmw4ZUNvcHdnZjYzQVo5REdMYWtZMG5XS2xHRVE0Y3BfVG9rZW46R1RHVWJ5ODFwb3pmYW14UUp4cmNwcXZybjFlXzE2OTU3MjA1NDk6MTY5NTcyNDE0OV9WNA" alt=""><figcaption></figcaption></figure>

### 访问Fireboom

<figure><img src="https://bar9vnf09af.feishu.cn/space/api/box/stream/download/asynccode/?code=YjdhZTM5YTNkZTJkYzVkZDg2NDU1M2UwYTcyZGZhNmZfaFJWN0swMGpvTUl1cTA2dkRzSE5vUkZ5bVBNdklYTHRfVG9rZW46WjdxU2IzYmNEbzB4U0d4QmFlQmN5RTJCbnRiXzE2OTU3MjE4NDM6MTY5NTcyNTQ0M19WNA" alt=""><figcaption></figcaption></figure>

1，访问控制台

<figure><img src="https://bar9vnf09af.feishu.cn/space/api/box/stream/download/asynccode/?code=OGQwMjM3YTY1NzY3YmE2MDk4NGZiM2Q1NDAxMjA3OGZfSGtweUhDWk5lZHVUSFlKdnNYOVVyUGx5Q1FLOTRJcUhfVG9rZW46TE92VmJld2VKb3ZpRTV4b0REWGNwRTdTbkxkXzE2OTU3MjA1NDk6MTY5NTcyNDE0OV9WNA" alt=""><figcaption></figcaption></figure>

2，访问API端点

```
status ok
```

### 配置Fireboom（可选）

1，查看API外网地址： 设置-> 系统

<figure><img src="https://bar9vnf09af.feishu.cn/space/api/box/stream/download/asynccode/?code=MWI1MzkzYWNhNjcwYTA1MGFlZGVkZGM5ZmI3ZGYwNWFfSUpyVzhsQVhCcGFUT1FFeE1KWWZ1U2xJd08ydlRmTjVfVG9rZW46UmR6emJLZ0J3b1lZR2h4V0Rra2NuM29BbjhjXzE2OTU3MjA1NDk6MTY5NTcyNDE0OV9WNA" alt=""><figcaption></figcaption></figure>

2，修改为：API端点地址，9991对应的公网地址

* 静态值：选择静态值，设置为 **API端点** 公网域名
* 环境变量：前往 环境变量 ，找到 FB\_API\_PUBLIC\_URL 设置为 **API端点** 公网域名

<figure><img src="https://bar9vnf09af.feishu.cn/space/api/box/stream/download/asynccode/?code=N2RmMmFhZjczYmZjYTFiYWQ1MTI1MmE3YmYxZTU4Y2VfRmVGWTVxNTVYVnJFY0F0VFYxalVjU0FMcUtKTDBZRzlfVG9rZW46RDYyaWJiYkRqb0N2Q1B4eFFrQWNyOHlKbjVlXzE2OTU3MjA1NDk6MTY5NTcyNDE0OV9WNA" alt=""><figcaption></figcaption></figure>

### 测试API

点击Fireboom控制台右上角的swagger文档图标，进入文档页。

<figure><img src="https://bar9vnf09af.feishu.cn/space/api/box/stream/download/asynccode/?code=MDBlZDRiNjNkOWFiMjkxNThjOWM4NTBiMDVmNzI5M2NfVEMyeXBJNzEwSnp2amlhaWJsb1R5Z014a21JcllSZmhfVG9rZW46QncyaGJJSmtEb0ppZGt4bm5YS2NPVTdwbkZoXzE2OTU3MjA1NDk6MTY5NTcyNDE0OV9WNA" alt=""><figcaption></figcaption></figure>

### 高级配置

#### 秘钥保护

fireboom在sealos上的模板默认用dev模式启动，且未开启秘钥保护，公网访问不安全。

可参考如下过程开启秘钥保护。

1，在sealos控制台，打开启动的fireboom服务

2，点击“Update”进入设置页

<figure><img src="https://bar9vnf09af.feishu.cn/space/api/box/stream/download/asynccode/?code=MTQ5NmM0ODhhOTNiNmEzMTU5MGYzOWM4MzRhODViY2ZfU0E1QTZ4QmNoQXlMaGJHUmhGZ2tlanFjc3lrYWJ6OVZfVG9rZW46T3FJeWI5b29Kb0g1VUp4VVp0YWMyOUhjblhiXzE2OTU3MjA1NDk6MTY5NTcyNDE0OV9WNA" alt=""><figcaption></figcaption></figure>

3，修改Parameters的参数为：

```sh
dev --enable-auth // 开发模式，带秘钥保护

// or

strat // 生产模式，带秘钥保护
```

<figure><img src="https://bar9vnf09af.feishu.cn/space/api/box/stream/download/asynccode/?code=YTA0ZDJmNzllMDU4MWI3NGZhYmY3YjE2Mjc5MGJmYWZfNWpMOGMwOFRUSExzT3ZMTjV4NHR4NmZQUzN3TlNqcktfVG9rZW46RmNNemJzS0lQb0dtS294eDVXSWN4MzZzbjliXzE2OTU3MjA1NDk6MTY5NTcyNDE0OV9WNA" alt=""><figcaption></figcaption></figure>

4，再次访问Fireboom控制台，看到如下界面

<figure><img src="https://bar9vnf09af.feishu.cn/space/api/box/stream/download/asynccode/?code=OGRlODY5OGMzOTkzZTYxYjQ3ZWI4MjEwN2E0NmZhZjJfZ0hVR2NNcDhXMEFvc0poWUZ5M0VLQUlBelhOUUJkSUtfVG9rZW46TVE1cmJKN0tJbzF6dTB4M3podGNjZ0tsbnVoXzE2OTU3MjA1NDk6MTY5NTcyNDE0OV9WNA" alt=""><figcaption></figcaption></figure>

#### 查找秘钥

1，进入terminal

<figure><img src="https://bar9vnf09af.feishu.cn/space/api/box/stream/download/asynccode/?code=NTE5YTk2NjI3MDlmMGM1ZDc1ODllZTliZGVlOTg1ZDhfZ3lIeVN0SUFUbHR2ZG9icGdWQVpDTGN6RVI3Qm9BUkpfVG9rZW46RWg5bmI5alh6b1lQMEd4bThaaGNXSnd2bk9iXzE2OTU3MjA1NDk6MTY5NTcyNDE0OV9WNA" alt=""><figcaption></figcaption></figure>

2，查看authentication.key文件

<figure><img src="https://bar9vnf09af.feishu.cn/space/api/box/stream/download/asynccode/?code=Y2Y3NzNmODI2MzU4MzVlNTcyNGY3MjNmNDY0ZDMxYTRfellsSFpCRjdZU3ZHSWwwQWRaQVBhUWZXWGJ6NTBTUWlfVG9rZW46QlppamI4UnRBbzBhVlN4cHlFZmN0NjJHbmdlXzE2OTU3MjA1NDk6MTY5NTcyNDE0OV9WNA" alt=""><figcaption></figcaption></figure>

3，在控制台输入秘钥，即可访问

