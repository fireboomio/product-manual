# GraphQL

本文是其GraphQL的高度浓缩，若难以理解，推荐教程 [前往](https://graphql.cn/learn/) 。

GraphQL是一个用于 API 的查询语言。他是强类型的，可以校验入参，并确定响应类型。

GraphQL有两个核心概念：GraphQL schema和GraphQL OPERATION。

## GraphQL schema

<figure><img src="../.gitbook/assets/image (3) (1).png" alt=""><figcaption></figcaption></figure>

schema指的是graphql结构定义的集合，类似数据库schema或数据库建表语句，operation指的是从schema中取的子集，类似数据库的查询sql。

我们知道gql是一个强类型语言，它通过4个关键字：type、enum、scalar、input，定义了所有的数据结构。

其数据结构分为两种，一种叫做Scalar Type(标量类型)，另一种叫做Object Type(对象类型)。

GraphQL中的内建的标量包含：String、Int、Float、Boolean、Enum，这些概念想必大家都接触过。

其中，enum表示枚举，通过enum关键字，可以定义枚举标量，例如sortorder，花括号里面的asc和desc是它的 枚举值。

GraphQL也支持通过Scalar声明新的标量，比如，这里声明了一个DateTime标量，用来表示日期格式。

总之，我们只需要记住，标量是GraphQL类型系统中最小的颗粒。\


对象类型基于标量构建，其关键字为type。每一个对象有若干字段组成，字段都有类型。例如：TODO 对象类型。它有6个Field，分别是INT类型的id，String类型的Title和Boolean类型的clompeted，xxx。

gql支持对象嵌套，因此 字段不仅可以是标量，也可以是对象，甚至可以是自循环对象。

关于类型，还有一个较重要的概念，即类型修饰符，当前的类型修饰符有两种，分别是List和Required，它们的语法分别为\[Type]和Type!。前者代表数组，后者代表必填。

除了type定义对象外，input也用于定义对象，称为 参数类型。可以类比为函数入参的类型。很多编程语言没有区分对象类型和参数类型。

在这里，我们要注意下，type和input定义的对象，都支持嵌套定义，即todowhereinput类型的字段也可以是todowhereinput，可以是别的对象类型，如Intfilter。

因此，type和input，通过这种嵌套，声明了各个模型之间的内在关联（一对多、一对一或多对多）。

此外，还有3类特殊对象，query、mutation和subscription，对于传统的CRUD项目，我们只需要前两种类型就足够了，第三种是针对当前日趋流行的real-time应用提出的，这块后续会讲到。

我们一般叫这3中对象为：Query(大写)。

接下来，我们分别以REST和GraphQL的角度，以todo为数据模型，编写一系列CRUD的接口。例如上图这几个接口。

* GET /api/v1/todos/
* GET /api/v1/todo/:id/
* POST /api/v1/todo/
* DELETE /api/v1/todo/:id/

获取待做事项列表对应findmanytodo，其中findmanytodo为根字段，可以类比为一个函数名称。

我们还可以发现，GraphQL中用query和mutation两种类型代表了rest接口的众多请求类型，例如QUERY对应get请求，mutation对应POST\DELETE、patch请求。

findmanytodo既然是函数，自然有入参和出参（返回值）。其出参是数组对象。

入参也有类型，但不能用type定义的类型，而是必须要用input定义类型。例如，where参数的类型是todowhereinput类型。

其实，不只是根字段可以有入参，任意对象类型都可以有入参。这个后续我们会遇到。

至此，我们就基本掌握了如何定义gql schema。

## GraphQL OPERATION

接下来，我们学习operation。

<figure><img src="../.gitbook/assets/image (1) (1) (1) (1) (1).png" alt=""><figcaption></figcaption></figure>

operation指的是从schema中的Query里取的子集，如果把schema queryr中的跟字段比作函数定义，那operation就是函数调用。

operation和SCHEMA query类型一一对应，也分为3中类型，分别是query，mutation，subscription。

如该operation  调用了 findmanytodo函数，字段入参分别是take和skip。其中take为固定值10，skip为变量值，由变量$skip传入。

变量是一个新概念，是Operaiton上定义的参数，注意不是函数上的参数。变量主要用途是动态设置函数的参数。

变量支持默认值，如$skip默认值为10。此外，变量也支持用修饰符！修饰，表明该变量为必传字段。

接着我们看下operaiton的响应，例如右侧图的灰色部分，就是该operaion的响应。在gql中被称为作用于该mutation opration的选择集。

值得注意的是，一个opratinon中可以有多个函数，因此operation选择集可能会同时包含多个跟字段。

选择集大概率是对应字段的子集，例如mutation里面不仅包含del和cretat，还包括executeraw，而实际上并没有用到后者。

不仅有作用域opration上的选择集，也有作用域字段上的选择集，例如，count就是作用域deletemanytodo字段的选择集。

## GraphQL SERVER

接下来，我们学习：如何使用基于GraphQL协议构建的服务。

<figure><img src="../.gitbook/assets/image (2) (1) (1).png" alt=""><figcaption></figcaption></figure>

Gql服务启动后，会对外暴露gql端点，其路由一般由graphql结尾。

以GET请求访问端点时，会返回一个gql操作界面。

基于界面可以方便的构建OPERATION，并执行该OPE拿到响应。分别，对应步骤2xxx和3xxx。

其中3的底层是向该Gql端点，发送POST请求。

我们以该opration为例，它有2个跟字段，message和detail，其中messsage字段的类型为对象，且无入参，detail的类型为标量，有一个入参title，并通过Operation的变量$titlevar为其赋值。

当我们点击执行按钮时，其请求如下：

```bash
curl 'https://fireboom-gql.ansoncode.repl.co/graphql' \
  -H 'content-type: application/json' \
  --data-raw $'{"operationName":"MyQuery","variables":{"titleVar":"ttttt"},"query":"query MyQuery($titleVar: String\u0021) {\\n  messages {\\n    content\\n    id\\n  }\\n  deatail(title: $titleVar)\\n}\\n"}' \
  --compressed
```

端点为：https://fireboom-gql.ansoncode.repl.co/graphql，请求为post类型。

请求头`content-type`为json类型。

请求体有3字段：

* operationName，操作名称，可以省略
* query：operation的字符串
* variables：operation的入参对象

该请求的响应体和operation的选择集一致，只不过会包裹在data对象里面。

{% code title="响应结构" %}
```json
{
  "data": {
    "messages": [
      {
        "id": "0",
        "content": "Hello!"
      },
      {
        "id": "1",
        "content": "Bye!"
      }
    ],
    "deatail": "echo:ssss"
  }
}
```
{% endcode %}

## GraphQL 内省

但有个问题需要回答： 该前端界面如何了解后端SCHEMA的呢？

这就要提到gql的“内省”能力，它是一种特殊的query operation 。能够根据步骤3一样的方式获取schema的结构。

之所以能实现该功能，主要得益于该端点的强类型特性。

向GraphQL端点发送下述请求，可以拿到GraphQL SCHEMA对应的JSON结构体，通过特定库可以将其转换为GraphQL SCHEMA。

{% code title="内省QUERY OPERATION" %}
```graphql
query IntrospectionQuery {
  __schema {
    queryType {
      name
    }
    mutationType {
      name
    }
    subscriptionType {
      name
    }
    types {
      ...FullType
    }
    directives {
      name
      description
      locations
      args {
        ...InputValue
      }
    }
  }
}
fragment FullType on __Type {
  kind
  name
  description
  fields(includeDeprecated: true) {
    name
    description
    args {
      ...InputValue
    }
    type {
      ...TypeRef
    }
    isDeprecated
    deprecationReason
  }
  inputFields {
    ...InputValue
  }
  interfaces {
    ...TypeRef
  }
  enumValues(includeDeprecated: true) {
    name
    description
    isDeprecated
    deprecationReason
  }
  possibleTypes {
    ...TypeRef
  }
}
fragment InputValue on __InputValue {
  name
  description
  type {
    ...TypeRef
  }
  defaultValue
}
fragment TypeRef on __Type {
  kind
  name
  ofType {
    kind
    name
    ofType {
      kind
      name
      ofType {
        kind
        name
        ofType {
          kind
          name
          ofType {
            kind
            name
            ofType {
              kind
              name
              ofType {
                kind
                name
              }
            }
          }
        }
      }
    }
  }
}
```
{% endcode %}



```javascript
// graphql server code
import express from 'express';
import { createServer } from 'http';
import { PubSub } from 'apollo-server';
import { ApolloServer, gql } from 'apollo-server-express';

const app = express();

const pubsub = new PubSub();
const MESSAGE_CREATED = 'MESSAGE_CREATED';

const typeDefs = gql`
  type Query {
    messages: [Message!]!
    deatail(title:String!):String
  }

  type Subscription {
    messageCreated: Message
  }

  type Message {
    id: String
    content: String
  }
`;

const resolvers = {
  Query: {
    messages: (ctx) => {
      console.log("test", ctx)
      return [
        { id: 0, content: 'Hello!' },
        { id: 1, content: 'Bye!' },
      ]
    },
    deatail: (ctx,{title}) => {
      console.log("test", ctx,title)
      return "echo:"+title
    },
    
  },
  Subscription: {
    messageCreated: {
      subscribe: () => pubsub.asyncIterator(MESSAGE_CREATED),
    },
  },
};
const myPlugin = {
  // Fires whenever a GraphQL request is received from a client.
  async requestDidStart(requestContext) {
    console.log('Request started!')
    console.log(requestContext.request.http.headers)
    return {
      async parsingDidStart(requestContext) {
        console.log('Parsing started!');
      },
      async validationDidStart(requestContext) {
        console.log('Validation started!');
      },
    }
  },
};
const server = new ApolloServer({
  typeDefs,
  resolvers,
  plugins: [
    myPlugin
  ]
});

server.applyMiddleware({ app, path: '/graphql' });

const httpServer = createServer(app);
server.installSubscriptionHandlers(httpServer);

httpServer.listen({ port: 8000 }, () => {
  console.log('Apollo Server on http://localhost:8000/graphql');
});

let id = 2;

setInterval(() => {
  pubsub.publish(MESSAGE_CREATED, {
    messageCreated: { id, content: new Date().toString() },
  });

  id++;
}, 1000);
```

## 参考

* [30分钟理解GraphQL核心概念](https://juejin.cn/post/6844903586548154376)

