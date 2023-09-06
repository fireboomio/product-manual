# 消息队列

<mark style="color:orange;">当前该特性暂未上线，正在规划中\~</mark>

Fireboom可接入GraphQL数据源，且支持GraphQL订阅，详情见[shi-shi-tui-song.md](../api-gou-jian/shi-shi-tui-song.md "mention")。

借助该能力可接入消息队列，特别适合处理物联网事件。

实现该功能需要2个步骤：

1. 将消息队列转换成GraphQL数据源
2. 将GraphQL数据源接入Fireboom

利用上述原理，飞布映射实时消息到graphql订阅，并通过REST API的方式呈现给客户端。

同时，开发者还可以使用自定义脚本处理订阅事件，实现数据落库等功能。



在该功能未实现前，可<mark style="color:red;">手动</mark>实现上述功能：

1，利用现成的库，将消息队列转换成GraphQL服务，例如：

* nodejs示例：[https://github.com/apollographql/graphql-subscriptions](https://github.com/apollographql/graphql-subscriptions)
* golang示例：[https://github.com/lcRuri/fireboom-doc/blob/main/%E6%B6%88%E6%81%AF%E9%98%9F%E5%88%97/kafka%20graphql.md](https://github.com/lcRuri/fireboom-doc/blob/main/%E6%B6%88%E6%81%AF%E9%98%9F%E5%88%97/kafka%20graphql.md)

2，在Fireboom数据源中添加GraphQL数据源，详情见 [graphql-api.md](graphql-api.md "mention")
