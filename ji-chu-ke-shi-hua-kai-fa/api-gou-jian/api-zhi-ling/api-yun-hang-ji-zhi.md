# API运行机制（废弃）

飞布底层基于GraphQL构建，GraphQL OPERATION 分为三类，分别为：QUERY变更、MUTATION变更以及SUBSCRIPTION订阅。每种类型的OPERATION在编译为API时，有不同的表现。此外，使用特定指令修饰后，也会有不同表现，如 `@internalOperation` 指令。

接下来，我们分别进行介绍。

## 订阅流程图

飞布将 SUBSCRIPTION OPERATION 转换为[ 服务器推流 SSE](https://juejin.cn/post/6854573215516196878) 。客户端与服务端建立链接后，可以实时收到来自服务器的消息。

{% hint style="info" %}
SUBSCRIPTION OPERATION不能被声明为内部。
{% endhint %}

其流程图如下所示：

<figure><img src="../../../.gitbook/assets/image (4) (1) (2).png" alt=""><figcaption><p>订阅流程图</p></figcaption></figure>

### 流程讲解

首先，当客户端订阅服务时①，飞布服务同步订阅事件源②。

然后，等待事件源推送消息给飞布服务③，飞布将其转发给客户端④。

最后，当客户端取消订阅时⑤，飞布也取消订阅⑥。

其中，③和④在取消订阅前，循环执行。

### 钩子讲解

同时，该流程图也展示了钩子的执行时机，及钩子启停状态。

该流程图涉及两类钩子：全局钩子和局部钩子。

全局钩子：所有API共用，主要用于在请求数据源前后的逻辑处理，包括：onConnection。（<mark style="color:orange;">在代码逻辑中，根据数据源的不同执行不同逻辑！</mark>）

局部钩子：每个API独立编写和启用，包括前置：preResolve、mutatingPreResolve，后置：postResolve、mutatingPostResolve。

详情请前往[钩子章节](../../../jin-jie-gou-zi-ji-zhi/gou-zi-ji-zhi.md)。
