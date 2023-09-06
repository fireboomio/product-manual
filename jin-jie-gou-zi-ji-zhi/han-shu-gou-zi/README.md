# 函数钩子

函数钩子是一种特殊的钩子，功能和 [graphql-gou-zi.md](../graphql-gou-zi.md "mention")类似，都用于自行定义业务逻辑。

但区别在于，graphql钩子本质上是graphql数据源，将被合并到超图中，然后基于构建的OPERATION，暴露到客户端。

<figure><img src="../../.gitbook/assets/image (2).png" alt=""><figcaption></figcaption></figure>

而每个函数钩子都会按照规则，注册对应路由，直接暴露到客户端。
