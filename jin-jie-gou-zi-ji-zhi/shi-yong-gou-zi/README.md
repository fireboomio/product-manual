# 使用钩子

任何语言实现的Fireboom钩子，本质上都是一个WEB服务。但要遵循Fireboom规范注册对应路由。

任意语言的钩子服务启动时，都遵循如下流程。

<figure><img src="../../.gitbook/assets/image (1) (1) (1).png" alt=""><figcaption></figcaption></figure>

## 读取配置文件：

1. 配置文件：custom-go/generated/fireboom.config.json 是一个指向exported/generated/fireboom.config.json的软连接
2. 包含钩子启动所依赖的大部分信息，如钩子监听端口serverOptions.listen.port，S3配置信息s3UploadConfiguration等

{% tabs %}
{% tab title="golang" %}
{% code title="pkg/types/configure.go" %}
```go
var configJsonPath = filepath.Join("generated", "fireboom.config.json")

func init() {
	_ = utils.ReadStructAndCacheFile(configJsonPath, &WdgGraphConfig)
}
```
{% endcode %}
{% endtab %}

{% tab title="nodejs" %}

{% endtab %}
{% endtabs %}

{% hint style="info" %}
启动钩子前要检查custom-go/generated/fireboom.config.json是否存在，否则钩子无法启动。部署时，可借助 ./fireboom build 命令，生成上述文件。
{% endhint %}

## 读取环境变量：

使用相对路径 `../.env`，和Fireboom服务共用

{% tabs %}
{% tab title="golang" %}
```go
const nodeEnvFilepath = "../.env"

func init() {
    _ = godotenv.Overload(nodeEnvFilepath)
```
{% endtab %}

{% tab title="nodejs" %}

{% endtab %}
{% endtabs %}

## 注册中间件：

1，解析Fireboom调用时携带的全局参数 \_wg



2，为上下文ctx注入User对象，用于获取登录用户的信息



3，为上下文ctx注入InternalClient对象（用于[内部调用](../nei-bu-tiao-yong.md)）



4，注册钩子路由，这块后面章节会详细介绍





在开始前我们先学习钩子服务的第一个协议：健康检查接口

{% swagger method="get" path="health" baseUrl="http://127.0.0.1:9992/" summary="健康检查接口" %}
{% swagger-description %}
检查钩子服务健康状态，用于在界面上展示钩子是否已启动
{% endswagger-description %}

{% swagger-response status="200: OK" description="" %}
```json
{
    "status": "ok"
}
```
{% endswagger-response %}
{% endswagger %}

剩余钩子如下：

[operation-gou-zi.md](../operation-gou-zi.md "mention")

[shen-fen-yan-zheng-gou-zi.md](../shen-fen-yan-zheng-gou-zi.md "mention")

[graphql-gou-zi.md](../graphql-gou-zi.md "mention")

[wen-jian-shang-chuan-gou-zi.md](../wen-jian-shang-chuan-gou-zi.md "mention")

[nei-bu-tiao-yong.md](../nei-bu-tiao-yong.md "mention")

