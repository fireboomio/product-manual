# REST API

REST是构建API最常用的风格。REST API无处不在。您自己的应用程序很可能依赖它。

使用飞布，你可以将OpenAPI规范指定的任何REST API转换为GraphQL API。通过这种方式，您可以像对待其他GraphQL API一样对待它，甚至可以将它与其他API拼接在一起。

## 新建REST API

**新建数据源** -> **REST API**，设置参数，如名称、OAS文件和rest端点。一般OAS文件中共包含rest端点信息，系统自动解析OAS并填写rest端点默认值。

<figure><img src="../../../.gitbook/assets/image (8).png" alt=""><figcaption><p>REST数据源</p></figcaption></figure>

> OAS（全称OpenAPI Specification），旧称Swagger Specification，是一个开放式、跨语言、跨平台的API描述语言，用于定义RESTful API的接口，包括请求、响应、模型、错误代码等内容。它提供了一种标准化的方法来描述API，可以使API设计和实现更加明确和一致，提高API的可读性和可维护性。

简单来说，OAS是一个强类型的函数定义，里面包含了API接口的所有元数据描述。

飞布引擎能够将其解析为GraphQL Schema（子图），并以名称作为命名空间，合并到“超图”中。

## 认证方式

当配置基于http的数据源(如GraphQL或OpenAPI)时，你可以配置飞布服务器是否以及如何将头发送到源。

有两种模式来区分，静态和动态。使用静态模式，可以设置静态标头，如API密钥或令牌。动态标头允许您根据客户端请求设置标头。

### 静态头

上图也是一个如何配置静态头文件的例子。

![](<../../../.gitbook/assets/image (2).png>)

* 值：字面量
* 环境变量：读取环境变量

针对该REST数据源，飞布将始终发送`X-API-KEY:xxxx`请求头。由于该信息存储在服务器上，因此是一种与受保护数据源进行通讯的安全方式。

### 动态头

某些场景，REST数据源的授权头会过期失效，需要动态生成授权信息。针对该用例有两种解决方案：

* 动态头（转发自客户端）：客户端设置请求头，经过飞布服务端，透传到上游数据源
* 授权钩子：在钩子中编写脚本的方式，动态生成授权头，透传至上游数据源，见钩子章节

下面是一个动态头示例：

<figure><img src="../../../.gitbook/assets/image (10).png" alt=""><figcaption><p>动态头示例</p></figcaption></figure>

第一个参数是要添加的头的名称。第二个参数是要从客户端请求中获取值的头的名称。使用此配置，我们将发送带有客户端请求的Authorization头的值的报头X-Authorization。

{% hint style="info" %}
动态头待确认。
{% endhint %}

\
