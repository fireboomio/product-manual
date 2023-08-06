# 可视化构建

首先，我们学习如何可视化构建API。

## 构建OPERATION

飞布的主功能区基于超图的gql 服务构建，其核心目标是，如何快速构建和测试GraphQL operation。

它包含四部分：超图面板、OPERATION编辑区、指令面板、参数输入区（响应区）。

<figure><img src="../../../.gitbook/assets/image (44).png" alt=""><figcaption></figcaption></figure>

其中，超图面板和OPERATION编辑区功能类似，且互相联动，都用于构建operation。

超图面板以可视化的方式展示当前项目所有的函数（根字段），勾选函数后可在编辑区生成OPERATION，在operation 编辑区手动修改OPERATION后，也会实时反应在超图面板中。

超图面板的所有”函数“，通过发送 GraphQL内省请求拿到，具体如下：

```bash
# 超图服务的gql服务端点是：http://localhost:9123/app/main/graphql
curl 'http://localhost:9123/app/main/graphql' \
  -H 'Accept: application/json' \
  --data-raw '{"query":"\n    query IntrospectionQuery {\n      __schema {\n        \n        queryType { name }\n        mutationType { name }\n        subscriptionType { name }\n        types {\n          ...FullType\n        }\n        directives {\n          name\n          description\n          \n          locations\n          args {\n            ...InputValue\n          }\n        }\n      }\n    }\n\n    fragment FullType on __Type {\n      kind\n      name\n      description\n      \n      fields(includeDeprecated: true) {\n        name\n        description\n        args {\n          ...InputValue\n        }\n        type {\n          ...TypeRef\n        }\n        isDeprecated\n        deprecationReason\n      }\n      inputFields {\n        ...InputValue\n      }\n      interfaces {\n        ...TypeRef\n      }\n      enumValues(includeDeprecated: true) {\n        name\n        description\n        isDeprecated\n        deprecationReason\n      }\n      possibleTypes {\n        ...TypeRef\n      }\n    }\n\n    fragment InputValue on __InputValue {\n      name\n      description\n      type { ...TypeRef }\n      defaultValue\n      \n      \n    }\n\n    fragment TypeRef on __Type {\n      kind\n      name\n      ofType {\n        kind\n        name\n        ofType {\n          kind\n          name\n          ofType {\n            kind\n            name\n            ofType {\n              kind\n              name\n              ofType {\n                kind\n                name\n                ofType {\n                  kind\n                  name\n                  ofType {\n                    kind\n                    name\n                  }\n                }\n              }\n            }\n          }\n        }\n      }\n    }\n  "}' \
  --compressed
```

然后，经由[GraphiQL Explorer](https://github.com/OneGraph/graphiql-explorer)渲染为如图所示的函数列表。

此外，参数输入区和响应区用于测试OPERATION。

具体操作步骤如下：

1，在"API管理"面板，新建一个API

2，在"超图面板"找到对应"[函数](../../../he-xin-gai-nian/graphql.md#graphql-operation)"，勾选对应字段，在OPERATION编辑区生成对应的OPERAITON

<figure><img src="../../../.gitbook/assets/image (43).png" alt=""><figcaption></figcaption></figure>

1. 勾选 函数底部蓝色字段，生成为 OPERATION选择集
2. 勾选 函数顶部紫色字段，生成为 OPERATION参数
3. 点击参数后的$符，参数变成 OPERATION 变量
4. 若生成有误，也可以在OPERATION编辑区手动修改

3，在指令面板， 点击对应按钮 为OPERATION增加指令，详情见 [api-zhi-ling](../api-zhi-ling/ "mention")

4，OPERATION构建完毕后，在参数输入区录入OPERATION入参，支持两种模式：

1. 可视化录入：标量正常录入，对象同"源码录入"
2. 源码录入：以JSON的方式录入变量，输入"双引号"可以触发语法提醒

5，最后，点击指令面板的 "<mark style="color:blue;">测试</mark>"按钮，执行该OPERATION，可在“响应”TAB中查看测试结果。

点击“测试”调用的是GraphQL端点，其执行格式为：

```bash
curl 'http://localhost:9123/app/main/graphql'
-H 'Accept: application/json'
--data-raw '{"query":"query MyQuery {\n bg_findFirstPost {\n authorId\n createdAt\n id\n published\n title\n auhor:User {\n email\n id\n name\n role\n }\n }\n}","variables":{},"operationName":"MyQuery"}'
--compressed
```

测试端点仅用于测试 GraphQL OPEARTION 到数据源的执行情况，未兼容指令。除跨源关联指令外，其他指令均不生效，如角色、响应转换和入参指令等。

{% hint style="info" %}
这里的”测试“，不同于复制API链接！！！
{% endhint %}

总的来说，飞布的主功能区就是 gql 服务 控制台的升级版，提供了更加友好的交互。

了解更多，请查看 [#api-guan-li-mian-ban](ke-shi-hua-kai-fa.md#api-guan-li-mian-ban "mention")

## 路由规则

接下来，我们学习Fireboom的路由规则，掌握不同OPERATION 对应的路由。

如上图所示，在根目录下有一个叫做CreateTodo的OPERATION。

{% code title="Todo/CreateTodo.graphql" %}
```graphql
mutation MyQuery($title: String = "") {
  todo_createOneTodo(data: {title: $title}) {
    id
    completed
    title
  }
}
```
{% endcode %}

当我们保存并上线该OPERATION后，复制其链接可以拿到如下URL。

```bash
curl 'http://localhost:9991/operations/Todo/CreateTodo' \
  -X POST  \
  -H 'Content-Type: application/json' \
  --data-raw '{"title":"learn fireboom"}' \
  --compressed
```

该URL分为3部分：

* 请求类型：POST，规则：<mark style="color:blue;">MUTATION</mark>对应<mark style="color:blue;">POST</mark>；<mark style="color:purple;">QUERYR</mark>对应<mark style="color:purple;">GET</mark>；<mark style="color:orange;">SUBSCRIPTION</mark>对应<mark style="color:orange;">GET</mark>，且左上角会有一个闪电标识。
* 请求域名：http://localhost:9991，可在 设置->系统->外网地址 中修改
* 请求路径：operations/Todo/CreateTodo，规则为：operations/+目录+OPEARTION名称

下面，我们举一些示例：

### Query Operation

{% code title="Page.graphql" %}
```graphql
query GetTodoList(
  $take: Int = 10, $skip: Int = 0,
  $orderBy: [todo_TodoOrderByWithRelationInput], 
  $query: todo_TodoWhereInput) {
  data: todo_findManyTodo(skip: $skip take: $take orderBy: $orderBy where: {AND: $query}) {
    id
    title
    completed
    createdAt
  }
  total: todo_aggregateTodo(where: {AND: $query}) @transform(get: "_count.id") {
    _count {
      id
    }
  }
}
```
{% endcode %}

对应为GET请求，复制为普通URL，如下：

```http
http://localhost:9991/operations/Page?
# GET QUERY 标量入参
take=10&skip=0&
# GET QUERY 对象入参
orderBy=[{"createdAt":"asc"}]&
query={"title":{"contains":"fireboom"}}
```

若Query Operation开启了实时查询，则复制为如下URL：

```http
# 加上了?wg_live=true后缀
http://localhost:9991/operations/Page?take=10&skip=0&orderBy=[{"createdAt":"asc"}]&query={"title":{"contains":"fireboom"}}&
wg_live=true
```

### Mutation Operation

{% code title="Update.graphql" %}
```graphql
mutation MyQuery($id: Int!,$update: mysql_TodoUpdateInput!, $create: mysql_TodoCreateInput!) {
  mysql_upsertOneTodo(create: $create, update: $update, where: {id: $id}) {
    id
  }
} 
```
{% endcode %}

对应为POST请求，复制为curl，如下：

```bash
curl 'http://localhost:9991/operations/Lesson0202/Basic/Update' \
  -X POST  \
  -H 'Content-Type: application/json' \
  # POST 对象入参
  --data-raw '{"id":1,"update":{"completed":{"set":true}},"create":{"title":"测试","completed":false}}' \
  --compressed
```

Mutation 特殊标量入参

{% code title="RawExecute.graphql" %}
```graphql
mutation MyQuery($parameters: todo_Json = ["beijing", 1]) {
  todo_executeRaw(
    query: "UPDATE \"main\".\"Todo\" SET \"title\" = $1 WHERE id=$2"
    parameters: $parameters
  )
}
```
{% endcode %}

```bash
curl 'http://localhost:9991/operations/Lesson0202/Basic/RawExecute2' \
  -X POST  \
  -H 'Content-Type: application/json' \
  # POST 特殊标量JSON 入参
  --data-raw '{"parameters":["beijing",1]}' \
  --compressed
```

### Subscription Operation

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

对应为GET请求，复制为如下URL：

```http
# 加上了?wg_sse=true
http://localhost:9991/operations/Sub?wg_sse=true
```

{% hint style="info" %}
复制连接中对应的域名：http://localhost:9991，需要前往设置->系统 ->外网地址 中修改。
{% endhint %}

更多路由规则，详情查看 [api-gui-fan.md](api-gui-fan.md "mention")

## 状态码

最后，我们学习下状态码。

OPERATION上线后，将被编译为REST API。当用户访问接口时，其对应HTTP流程如上图右侧所示。

常见状态码有如下几种：

* 200：Operation执行成功
* 500：Operation执行失败，例如：数据源无法访问时
* 404：Operation未找到，访问未上线或不存在的OPERATION
* 401：身份验证或身份鉴权失败，例如OPEARTION开启了授权登录或使用了[ fromclaim指令](../api-zhi-ling/#fromclaim)

