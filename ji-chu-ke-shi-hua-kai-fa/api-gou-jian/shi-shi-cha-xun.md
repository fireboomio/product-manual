# 实时查询

很多场景下，客户端需要实时更新数据。当前，主流方式是客户端轮询，或者短轮询，即客户端每隔几秒请求一次接口，获取数据。当客户端数量较多时，会给服务端造成较大并发压力。

例如，有如下OPERATION。

{% code title="SSE.graphql" %}
```graphql
query MyQuery {
  todo_findFirstTodo {
    completed
    createdAt
    id
    title
  }
}
```
{% endcode %}

其将被编译为REST API，访问地址为： `http://localhost:9991/operations/SSE`

客户端轮询示例代码：

```javascript
// 每隔1秒请求1次接口
setInterval(()=>{
    fetch("http://localhost:9991/operations/SSE", {
  "method": "GET",
  "mode": "cors",
});
},1000)
```

## 服务端轮询

飞布采用了一种新的机制：**服务端轮询**。它能以较小的代价，解决客户端轮询造成的资源消耗问题，实现数据的准实时更新。

服务端轮询把轮询逻辑从客户端移动到服务端，由服务端定时请求数据，并比对前后两次数据是否一致，若数据变化，则推送数据到客户端。同时，只有当客户端打开链接时，服务端才会定时轮询数据，保证系统性能。

<figure><img src="../../.gitbook/assets/image (8).png" alt=""><figcaption></figcaption></figure>

该时序图包括客户端、服务端和数据库。

首先，客户端与服务端建立HTTP2长链接；

接着，服务端查询数据库并将第一次结果返回给客户端；

最后，服务端定时轮询数据库并检查数据是否发生变化，若变化，则将数据推送至客户端。

### SSE

服务端轮询使用了 HTTP2 的[Server sent event](https://www.ruanyifeng.com/blog/2017/05/server-sent\_events.html)机制，即SSE服务器推送。它是基于 HTTP 协议中的持久连接，仅从服务器向客户端发送消息。&#x20;

它和WEBSOCKET有两点核心不同：

* 数据推送方向：SSE是服务器向客户端的单向通信，服务器可以主动推送数据给客户端。而WebSocket是双向通信，允许服务器和客户端之间进行实时的双向数据交换。
* 连接建立：SSE使用基于HTTP的长连接，通过普通的HTTP请求和响应来建立连接，从而实现数据的实时推送。 WebSocket使用自定义的协议，通过建立WebSocket连接来实现双向通信。

## 实时查询

任意QUERY OPERATION，都具备实时查询能力。使用实时查询非常容易，只需要在配置中一键即可开启。

选中 API，在右侧面板中 点击 设置 ，切换到配置页面。

* 勾选 使用独立配置
* 开启 实时配置&#x20;
* 开启 实时查询
* 设置 轮训间隔，单位是秒，最小值为1秒

<figure><img src="../../.gitbook/assets/image (1) (1).png" alt=""><figcaption></figcaption></figure>

设置完毕后，可以看到流程图中多了一条循环线，上面标记了轮训间隔。

点击顶部工具栏的 ”复制“ICON，获取访问地址：`http://localhost:9991/operations/SSE?`<mark style="color:red;">`wg_live=true`</mark> ，前往访问。

尝试修改数据，例如用数据建模预览面板修改。若，数据发生变化，则可以看到数据实时推送到界面。响应是一个JSON对象流，由<mark style="color:purple;">两个换行符分割</mark>。

每次轮训时，钩子都会生效！！！

基于SSE的实时查询，相对于websocket协议，还具有另一个优势，即复用HTTP的身份鉴权。在该OPERTION中设置的登录校验或者权限控制，无需任何额外操作，都可以应用到实时查询接口中。

而webssocket若要实现权限控制，则需要客户端的额外编写代码支持。

## 客户端使用

* SSE TypeScript示例：[前往](https://github.com/fireboomio/fb-admin/blob/46c919afd4fe80ab2ee89560ba394cc5ae3f9da7/front/src/layout/components/notice/index.vue#L29C16-L29C33)
