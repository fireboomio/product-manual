# graphql钩子

为了进一步提升开发体验，Fireboom在HOOKS框架中集成了graphql端点。它本质上是一个内置的grphql数据源。可以用来实现复杂的业务逻辑。

{% hint style="info" %}
该方式有循环依赖，要遵循[下述方法](graphql-gou-zi.md#gou-jian-api)绕开依赖！你也可以自行在HOOKS SDK中实现GraphQL服务。
{% endhint %}

## 新建GraphQL钩子

1，在 Fireboom 控制台点击`数据源`面板的`+`号，进入数据源新建页。

2，在数据源新建页面，选择 自定义-> 脚本，设置起名称为：`Custom`。

3，系统自动初始化如下脚本。默认填充的是示例代码，你可以根据业务需求修改代码。

<figure><img src="../.gitbook/assets/image (1) (1) (1) (1) (1) (1).png" alt=""><figcaption><p>graphq钩子</p></figcaption></figure>

{% tabs %}
{% tab title="golang" %}
{% code title="customize/Custom.go" %}
```go
package customize

import (
	"custom-go/pkg/plugins"
	"github.com/graphql-go/graphql" # 自行学习该库的用法
)

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var (
	personType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Person",
		Description: "A person in the system",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"firstName": &graphql.Field{
				Type: graphql.String,
			},
			"lastName": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	fields = graphql.Fields{
		"person": &graphql.Field{
			Type:        personType,
			Description: "Get Person By ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				_ = plugins.GetGraphqlContext(params)
				id, ok := params.Args["id"].(int)
				if ok {
					testPeopleData := []Person{
						{Id: 1, FirstName: "John", LastName: "Doe"},
						{Id: 2, FirstName: "Jane", LastName: "Doe"},
					}
					for _, p := range testPeopleData {
						if p.Id == id {
							return p, nil
						}
					}
				}
				return nil, nil
			},
		},
	}
	rootQuery = graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	Custom_schema, _ = graphql.NewSchema(graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)})
)
```
{% endcode %}
{% endtab %}

{% tab title="nodejs" %}

{% endtab %}
{% endtabs %}

4，在数据源列表，右击 开启当前钩子，然后重新启动钩子服务

5，钩子服务启动后，将注册对应的graphql服务端点，其路由格式如下：

<pre class="language-http"><code class="lang-http"><strong>http://{serverAddress}/gqls/${gql-name}/graphql
</strong>
# ${gql-name}=graphql 数据源名称
# Example:: http://localhost:9992/gqls/Custom/graphql
</code></pre>

## graphql控制面板

Get请求上述端点时，默认打开 graphql 控制面板。

其原理的是：读取helix.html文件，修改其中`${graphqlEndpoint}`为web界面请求路径。

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

## 内省graphql

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

<figure><img src="../.gitbook/assets/image (1) (1) (1) (1) (1).png" alt=""><figcaption><p>内省</p></figcaption></figure>



## 构建API

由于Fireboom服务依赖钩子服务，因此，需要按照如下步骤操作。

* 开启钩子：开启graphql钩子后，Fireboom将重新生成HOOKS 项目，将该钩子合成到对应文件
* 重启钩子：手动重新启动钩子服务，保证钩子服务注册了graphql路由
* 重新编译：在Fireboom控制台，点击右上角的“编译”按钮，内省该数据源

通过上述方式拿到的GraphQL SCHEMA被作为子图合并到Fireboom 超图 中，详情 查看 [chao-tu.md](../he-xin-gai-nian/chao-tu.md "mention")。

{% tabs %}
{% tab title="golang" %}
{% code title="server/fireboom_server.go" %}
```go
GraphqlServers: []plugins.GraphQLServerConfig{
    {
        ApiNamespace:          "Custom",
        ServerName:            "Custom",
        EnableGraphQLEndpoint: true,
        Schema:                customize.Custom_schema,
    },
},
```
{% endcode %}
{% endtab %}

{% tab title="nodejs" %}

{% endtab %}
{% endtabs %}

接着，我们就可以在Fireboom的超图面板中构建OPERATION，生成REST API供客户端使用。

<figure><img src="../.gitbook/assets/image (2) (1) (1) (1).png" alt=""><figcaption><p>将GraphQL钩子转变成REST API</p></figcaption></figure>

通过该方式可以实现任意复杂的业务需求！且可以复用Fireboom的权限体系。

此外，通过该方式可以将消息队列接入到Fireboom中，将事件订阅转变成GraphQL SUBSCRIPTION，然后转变成SSE推送，供客户端消费！

## 示例

* Nodejs示例：[case-element-admin](https://github.com/fireboomio/case-element-admin/blob/main/server/custom-ts/customize/statistics.ts)
* golang示例：[fb-admin](https://github.com/fireboomio/fb-admin/blob/main/backend/custom-go/customize/statistics.go)
