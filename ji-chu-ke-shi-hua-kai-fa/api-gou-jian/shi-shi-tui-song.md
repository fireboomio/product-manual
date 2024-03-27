# 实时推送

除了实时查询外，飞布还支持实时推送功能，常见用例：接入消息队列。相对于实时查询的<mark style="color:purple;">准实时</mark>更新，实时推送是实时更新。

实时推送基于graphql的订阅 operation实现。订阅类似查询，但不是在一次读取中返回数据，而是持续获取服务器推送的数据。这适用于实时应用场景，如IM或物联网应用。

## GraphQL订阅

一般情况下，GraphQL 订阅是发送到 WebSocket 端点的订阅查询字符串。 每当后端出现数据变化，新数据都会通过 WebSocket 从服务器向客户端推送。

例如，下面是一个基于[NodeJs构建的GraphQL 服务](../../he-xin-gai-nian/graphql.md#graphql-server)，正在执行订阅请求，数据通过WebSocket不断推送到客户端。

<figure><img src="../../.gitbook/assets/image (46).png" alt=""><figcaption></figcaption></figure>

## 实时推送

与通常graphql 服务不同，飞布基于SSE实现了实时推送，将上游的 WebSocket 订阅事件，转换为HTTP长连接。

当客户端访问时，飞布先与具有订阅接口的gql数据源建立连接，然后等待 WebSocket 推送数据，并将数据在通过SSE推送给客户端。

其流程图如下所示：

<figure><img src="../../.gitbook/assets/image (4) (1) (2) (1).png" alt=""><figcaption><p>订阅流程图</p></figcaption></figure>

首先，当客户端订阅服务时①，飞布服务同步订阅事件源②。

然后，等待事件源推送消息给飞布服务③，飞布将其转发给客户端④。

最后，当客户端取消订阅时⑤，飞布也取消订阅⑥。

其中，③和④在取消订阅前，循环执行。

构建如下SUBSCRIPTION OPERATION：

{% code title="Sub.graphql" %}
```graphql
subscription MyQuery {
  gql_messageCreated {
    content
    id
  }
}
```
{% endcode %}

点击顶部工具栏的 ”复制“ICON，获取访问地址：`http://localhost:9991/operations/Sub?`<mark style="color:red;">`wg_see=true`</mark> 。

<figure><img src="../../.gitbook/assets/image (47).png" alt=""><figcaption></figcaption></figure>

前往访问，可以看到数据实时推送到界面。响应是一个JSON对象流，由<mark style="color:purple;">两个换行符分割</mark>。

此外，若只想获取一次数据，则需要用 `wg_subscribe_once` 代替 `wg_see`。

* 订阅1次：wg\_subscribe\_once=true
* 持续订阅：wg\_see=true

同理，订阅OPERATION编译的REST API也可复用HTTP的身份鉴权。在该OPERTION中设置的登录校验或者权限控制，无需任何额外操作，都可以应用到实时推送接口中。

{% hint style="info" %}
SUBSCRIPTION OPERATION不能被声明为内部。
{% endhint %}

## 客户端使用

* SSE TypeScript示例：[前往](https://github.com/fireboomio/fb-admin/blob/46c919afd4fe80ab2ee89560ba394cc5ae3f9da7/front/src/layout/components/notice/index.vue#L29C16-L29C33)
