# graphql钩子

为了进一步提升开发体验，Fireboom在HOOKS框架中集成了graphql端点。它本质上是一个内置的grphql数据源。可以用来实现复杂的业务逻辑。

## 新建GraphQL钩子

1，在 Fireboom 控制台点击`数据源`面板的`+`号，进入数据源新建页。

2，在数据源新建页面，选择 脚本-> GraphQL，设置名称为：`custom`。

3，系统自动初始化如下脚本。

<figure><img src="../.gitbook/assets/image (13).png" alt=""><figcaption><p>graphql 钩子</p></figcaption></figure>

{% tabs %}
{% tab title="golang" %}
{% code title="customize/custom.go" %}
```go
package customize

import (
	"custom-go/pkg/plugins"
	"github.com/graphql-go/graphql" # 自行学习该库的用法
	"log"
	"time"
)
// ... 这里有省略
// 定义 graphql schema
var Custom_schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	// 查询
	Query: graphql.NewObject(graphql.ObjectConfig{
		Name:   "query",
		Fields: queryFields,
	}),
	// 变更
	Mutation: graphql.NewObject(graphql.ObjectConfig{
		Name:   "mutation",
		Fields: mutationFields,
	}),
	// 订阅
	Subscription: graphql.NewObject(graphql.ObjectConfig{
		Name:   "subscription",
		Fields: subscriptionFields,
	}),
})

func init() {
	// 注册 graphql 服务
	plugins.RegisterGraphql(&Custom_schema)
}
```
{% endcode %}
{% endtab %}
{% endtabs %}

默认填充的是示例代码，你可以根据业务需求修改代码。

4，在`main.go`中匿名引入该包，然后重新启动钩子服务

{% tabs %}
{% tab title="golang" %}
{% code title="custom-go/main.go" %}
```go
package main

import (
	// 匿名引入该库
	_ "custom-go/customize"
	// _ "custom-go/function"
	// _ "custom-go/proxy"
	"custom-go/server"
)

func main() {
	server.Execute()
}
```
{% endcode %}
{% endtab %}
{% endtabs %}

5，钩子服务启动后，将注册对应的graphql服务端点，其路由格式如下：

<pre class="language-http"><code class="lang-http"><strong>http://{serverAddress}/gqls/${gql-name}/graphql
</strong>
# ${gql-name}=graphql 数据源名称
# Example:: http://localhost:9992/gqls/Custom/graphql
</code></pre>

6，同时，钩子服务[自动内省](graphql-gou-zi.md#nei-sheng-graphql)该端点，并将内省产graphql schema物存储到文件：`custom-x/customize/custom.json`

```
customize          
├─ custom.go   # graphql 钩子代码   
└─ custom.json # 内省生成的结构化graphql schema
```

7，Fireboom服务每秒触发1次健康检查，将获取如下结果：

```json
{
    "report": {
        // graphql 钩子
        "customizes": [
            "custom" 
        ],
        "time": "2023-09-06T17:18:21.957519+08:00"
    },
    // 钩子服务的状态
    "status": "ok"
}
```

可以看到 `report.customizes` 中包含`custom`服务。接着读取步骤6中的子图graphql schema，合并到Fireboom [超图](../he-xin-gai-nian/chao-tu.md) 中。

{% hint style="info" %}
如果`report`没有变化，则不会触发编译！
{% endhint %}

8，最后，Fireboom发送通知，触发控制台更新，在数据源面板展示 `custom` 数据源。

### graphql控制面板

Get请求上述端点时，默认打开 graphql 控制面板。

其原理的是：读取`helix.html`文件，修改其中`${graphqlEndpoint}`为web界面请求路径。

```html
<script >
    GraphQLHelixGraphiQL.init({
      defaultQuery: undefined,
      defaultVariableEditorOpen: undefined,
      endpoint: '${graphqlEndpoint}', // 修改为 http://localhost:9992/gqls/Custom/graphql
      headers: undefined,
      headerEditorEnabled: undefined,
      subscriptionsEndpoint: undefined,
      useWebSocketLegacyProtocol: undefined,
      hybridSubscriptionTransportConfig: undefined,
      shouldPersistHeaders: undefined,
    });
</script>
```

{% file src="../.gitbook/assets/helix.html" %}

### 内省graphql

post 请求上述端点，按照如下格式，可以获得其graphql schema。

```http
http://{serverAddress}/gqls/${gql-name}/graphql

# Example:: http://localhost:9992/gqls/Custom/graphql

Content-Type: application/json

# JSON request
{
    "query": "IntrospectionQuery content", // 内省QUERY OPERATION，见下文
    "variables": {}, 
    "operationName": "IntrospectionQuery" 
}

# JSON response
{
    // ... graphql schema struct
}
```

<details>

<summary>IntrospectionQuery content</summary>

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

</details>

<figure><img src="../.gitbook/assets/image (1) (1) (1) (1) (1) (1) (1) (1).png" alt=""><figcaption><p>内省</p></figcaption></figure>

## 构建API

在Fireboom的超图面板中构建OPERATION，生成REST API供客户端使用。

<figure><img src="../.gitbook/assets/image (2) (1) (1) (1) (1).png" alt=""><figcaption><p>将GraphQL钩子转变成REST API</p></figcaption></figure>

通过该方式可以实现任意复杂的业务需求！且可以复用Fireboom的权限体系。

此外，通过该方式可以将消息队列接入到Fireboom中，将事件订阅转变成GraphQL SUBSCRIPTION，然后转变成SSE推送，供客户端消费！

## 示例

* Nodejs示例：[case-element-admin](https://github.com/fireboomio/case-element-admin/blob/main/server/custom-ts/customize/statistics.ts)
* golang示例：[fb-admin](https://github.com/fireboomio/fb-admin/blob/main/backend/custom-go/customize/statistics.go)
