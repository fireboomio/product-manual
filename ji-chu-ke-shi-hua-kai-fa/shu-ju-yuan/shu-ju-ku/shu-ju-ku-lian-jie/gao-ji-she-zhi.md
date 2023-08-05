# 高级设置

在数据库中存储JSON列是很常见的用例。如果你使用数据库，例如PostgreSQL，你可以使用 `json` 或`jsonb` 类型存储JSON列。在GraphQL schema中，该列将被表示为`JSON` 标量类型。在内部，把值存储到数据库之前，飞布将json值编码为字符串，并在从数据库读取时解码它。这样，`JSON`标量类型非常容易使用。

然而，该方法有个缺点。如果你打算存储复杂JSON对象，你无法利用GraphQL的类型系统优势。不容易从JSON对象中选择特定的字段。你必须要手动解析它，并且选择你想要的字段。

高级设置就是解决该问题的机制。飞布允许您扩展GraphQL Schema并用自定义类型替换特定的JSON字段。通过这种方式，你可以利用GraphQL的类型系统，同时能够将数据作为JSON对象存储在数据库中。

接下来让我们用具体示例说明下：

{% code title="增加自定义类型前的OPERATION" %}
```graphql
mutation (
  $message: String!
  $payload: JSON! # 这是JSON字段
)  {
  createOnemessages: db_createOnemessages(
    data: {
      message: $message
      payload: $payload
    }
  ) {
    id
    message
    payload # 这里是JSON字段，是标量类型，客户端需要知道如何去解析它
  }
}
```
{% endcode %}

该操作声明了创建消息的接口。`$message`是字符串类型。`$payload`是JSON类型，没有方法可以校验入参。

现在我们用高级设置功能扩展GraphQL Schema。

首先，在自定义类型中，填写如下GraphQL Schema。&#x20;

```graphql
# 响应类型结构体
type MessagePayload {
    extra: String!
}
# 入参类型结构体
input MessagePayloadInput {
    extra: String!
}
```

它将为"GraphQL Schema"增加两个结构体，`type`开头为响应类型结构体，`input`开头为入参类型结构体。

接下来，定义想要替换的JSON字段。

在字段类型映射中，选择表`messages`，字段`payload`，响应类型`MessagePayload`，输入类型`MessagePayloadInput`。

系统底层构建如下结构体：

```json
[
    {
      entityName: 'messages',
      fieldName: 'payload',
      responseTypeReplacement: 'MessagePayload',
      inputTypeReplacement: 'MessagePayloadInput',
    },
  ]
```

<figure><img src="../../../../.gitbook/assets/image (3) (2).png" alt=""><figcaption><p>数据库高级设置</p></figcaption></figure>

最终，我们定义了输入字段替换和响应字段替换的类型。



现在，我们可以构建如下OPERATION，解决上述问题。

* payload入参：替换成了db\_MessagePayloadInput对象，基于该对象可以实现入参校验
* payload响应：替换成了db\_MessagePayload对象，可以“炸开”该对象，选择所需字段

{% code title="增加自定义类型后的OPERATION" %}
```graphql
mutation (
  $message: String!
  $payload: db_MessagePayloadInput! # 这里的JSON变成了对象
)  {
  createOnemessages: db_createOnemessages(
    data: {
      message: $message
      payload: $payload
    }
  ) {
    id
    message
    payload { # 可以看到这里可以"炸开"了
      extra
    }
  }
}
```
{% endcode %}

