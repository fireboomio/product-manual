# 跨源关联

{% embed url="https://www.bilibili.com/video/BV1iM4y1U7mE/" %}
11功能介绍-飞布如何实现跨源关联？
{% endembed %}

某些场景下，我们需要跨数据源实现关联查询，例如获取物联网设备列表，并查看设备在线状态。

传统模式下，我们需要先从数据库获取设备列表，然后遍历数据，逐个调用物联网平台接口获取设备在线状态，最后拼接数据返回给客户端。用编码的方式，大概需要几百行代码。

而利用飞布的跨源关联功能，只需要几行graphql描述就能实现上述需求。

跨源关联本质上是一种流程编排，将通常情况下并行的请求，改造成串行。

<figure><img src="../../.gitbook/assets/operation-export.gif" alt=""><figcaption><p>跨源关联时序图</p></figcaption></figure>

使用跨源关联，至少需要配置两个数据源。例如，db为数据库，iot为物联网REST API。

```graphql
query MyQuery($device_id:Int! @internal) {  # 声明
  db_findManyDevice {
    id @export(as:"device_id")# 赋值
    name
  # _join 字段返回类型Query!
  # 它存在于每个对象类型上，所以你可以在Query文档的任何地方看到它
    _join{
      iot_deviceState(device_id: $device_id) { # 使用
        status
      }
    }
  } 
}

```

上述示例，主要分为三个环节：

* 声明：@internal 指令从公开API中移除 $capital 变量。这意味，用户不能手工设置它。我们称它为关联键（JOIN key）。
* 赋值：使用 @export 指令，我们可以将字段 \`id\`的值导出给关联键($device\_id)
* 使用：一旦我们进入 \_join 字段，我们可以使用 $device\_id 变量去关联物联网 API

![](https://cdn.nlark.com/yuque/0/2023/png/8370227/1679561691909-4d31fb4a-ba16-480b-aefa-47f62e6648b2.png)
