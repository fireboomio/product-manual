# 接口安全

除了身份验证和身份授权外，Fireboom还内置更多的安全机制，来保证接口安全。

## 服务监听

### 监听地址

Fireboom启动时，默认监听 `localhost:9991`，只能被本机访问，以保证安全。若想通过公网访问，有两种方法：

* 将设置->系统-><mark style="color:purple;">API服务监听Host</mark> 修改为 `0.0.0.0`
* 通过nginx代理转发到 `localhost:9991`

### 内网访问

Fireboom不仅对外暴露API，而且还能作为数据代理层为钩子提供服务。

为保证数据代理接口的安全，`/internal` 路由只能通过内网访问。Fireboom数据代理和Fireboom API监听端口相同，因此API内网地址一般为：`http://localhost:9991`。

详情查看，钩子章节->[nei-bu-tiao-yong.md](../../jin-jie-gou-zi-ji-zhi/nei-bu-tiao-yong.md "mention")

### 钩子地址

Fireboom钩子服务默认监听 `localhost:9992` ，为保证钩子服务安全，请不要开启9992端口的防火墙。

## 允许HOST

不同的域名通过 A 记录或者 CNAME 方式可以连接都同一个 IP 下。

出于安全的考虑，有时需要限制API服务只能被特定域名或IP访问。使用设置->安全-><mark style="color:purple;">允许Host</mark>可以设置。

默认情况下，允许HOST为 `*` ，表示不限制。

<figure><img src="../../.gitbook/assets/image.png" alt=""><figcaption><p>允许HOST对应请求头中的Host字段</p></figcaption></figure>

## 环境变量

处于保证安全和环境切换的用途，Fireboom支持用环境变量作为参数。例如：数据库密码、OIDC秘钥、API外网地址等。

前往 设置-><mark style="color:purple;">环境变量</mark> 设置，其底层对应根目录下的`.env`文件。

