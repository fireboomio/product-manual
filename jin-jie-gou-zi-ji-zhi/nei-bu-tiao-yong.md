# 内部调用

某些场景下，我们不仅要在钩子中编写业务逻辑，还希望在钩子中调用存储数据的逻辑。

有两种方式可以实现该功能：

* 在钩子服务中建立连接池，像传统开发方式那样操作数据库
* 复用Fireboom本身的数据操纵能力——数据代理

本节，我们重点介绍Fireboom的数据操纵能力——内部调用。

Fireboom将所有的`query`和`mutation`操作，都挂载到了一个特殊路由上：`/internal`，供钩子调用。

此时，飞布升级为数据代理层。Fireboom数据代理和Fireboom API监听端口相同，因此API内网地址一般为：`http://localhost:9991`。

`InternalClient`是使用数据代理的对象，其上挂载了所有的`query`和`mutation`。

<img src="../.gitbook/assets/image (1) (1) (1) (1) (1) (1) (1).png" alt="" data-size="original">

## 内部调用协议

在钩子中调用Fireboom OPERATION的协议如下：

```http
http://{nodeAddress}/internal/operations/{operationPath}

Example:: http://localhost:9991/internal/operations/Internal

Content-Type: application/json
X-Request-Id: "83821325-9638-e1af-f27d-234624aa1824"

# JSON request
{
    "input": {"name": "fireboom"}, // operation请求入参 
     "__wg": { // 全局参数
        "clientRequest": { // 原始客户端请求，即请求9991端口的request对象
          "method": "GET",
          "requestURI": "/operations/Weather?code=beijing",
          "headers": {
            "Accept": "application/json",
            "Content-Type": "application/json"
          }
        },
        "user": { // （可选）授权用户的信息
          "userID": "1",
          "roles": ["user"]
        }
      }
}

# JSON response
{
    "data": {} // operation返回结果
    "errors": [
        {
            "message": "error message",
            "path": "" // 可选，设置报错定位
        }
    ]
}
```

{% hint style="info" %}
{nodeAddress} 默认为：localhost:9991，而不是钩子的9992端口
{% endhint %}

### InternalClient实现及使用

{% tabs %}
{% tab title="golang" %}
{% code title="server/start.go" %}
```go
// 用中间件的方式挂载InternalClient
e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == http.MethodGet {
			return next(c)
		}

		var body base.BaseRequestBody
		err := utils.CopyAndBindRequestBody(c.Request(), &body)
		if err != nil {
			return err
		}

		if body.Wg.ClientRequest == nil {
			body.Wg.ClientRequest = &base.ClientRequest{
				Method:     c.Request().Method,
				RequestURI: c.Request().RequestURI,
				Headers:    map[string]string{},
			}
		} else {
			for name, value := range body.Wg.ClientRequest.Headers {
				c.Request().Header.Set(name, value)
			}
		}
		// 用来追踪调用过程
		reqId := c.Request().Header.Get("x-request-id")
		// 构建InternalClient
		internalClient := base.InternalClientFactoryCall(map[string]string{"x-request-id": reqId}, body.Wg.ClientRequest, body.Wg.User)
		internalClient.Queries = internalQueries
		internalClient.Mutations = internalMutations
		brc := &base.BaseRequestContext{
			Context:        c,
			User:           body.Wg.User,
			InternalClient: internalClient,
		}
		return next(brc)
	}
})
```
{% endcode %}

<pre class="language-go" data-title="hooks/Lesson0301/Simple/postResolve.go"><code class="lang-go"><strong>func PostResolve(hook *base.HookRequest, body generated.Lesson0301__SimpleBody) (res generated.Lesson0301__SimpleBody, err error) {
</strong>	hook.Logger().Info("PostResolve")
	// 通过ExecuteInternalRequestQueries方法调用QUERY OPERATION，其中
	// [OP_PATH]为OPERATION的路径，合成规则为DIC__OPName，例如：Lesson0301__Internal
	// [OP_PATH]Input 为入参对象
	// [OP_PATH]ResponseData为响应对象
	todoRes, _ := plugins.ExecuteInternalRequestQueries[generated.Lesson0301__InternalInput, generated.Lesson0301__InternalResponseData]
	(hook.InternalClient, generated.Lesson0301__Internal, 
	generated.Lesson0301__InternalInput{
		Id: 2,
	})
	fmt.Println(todoRes)
	return body, nil
}
</code></pre>
{% endtab %}
{% endtabs %}

### 内部调用安全

为了保证安全：`http://{nodeAddress}/internal` 只能被内网访问！！！

## 内部OPERATION

如果我们希望某个OPERATION只能被钩子访问，不对外暴露API，需要借助：`@internalOperation`指令。该指令仅能修饰QUERY 和 MUTATION OPERATION 。设置OPERATION为<mark style="color:orange;">内部</mark>，类似私有方法，只能被钩子调用，而不会编译为API。

在GraphQL编辑区上方的工具栏，点击“<mark style="color:blue;">@私有</mark>”，选择 <mark style="color:blue;">私有</mark>，为当前OPERATION添加 `@internalOperation`指令。

### 流程图

设置后，可在右侧概览面板看到对应流程图的变化。

```graphql
query GetOneTodo($id: Int!) @internalOperation {
  data: todo_findFirstTodo(where: {id: {equals: $id}}) {
    id
    title
  }
}
```

<div align="center">

<img src="../.gitbook/assets/image (2) (1) (4) (1) (1).png" alt="内部OPERATION流程图示意图" width="295">

</div>

在钩子服务中，可通过InternalClient对象访问飞布数据代理（data proxy）中的内部OPERATION。如图中②表示请求流程，③表示响应流程。

### 常见用例

```graphql
# @internalOperation 指令声明当前OPERATION为内部
query MyQuery($id: Int!) @internalOperation {
  todo_findUniqueTodo(where: {id: $id}) {
    id
    createdAt
    completed
    title
  }
}
# 将下列原生查询声明为内部，使用时，可以传递任意SQL语句...
mutation MyQuery($sql: String!) @internalOperation {
  sqlite_queryRaw(query: $sql)
}
mutation MyQuery($query: String!) @internalOperation {
  sqlite_executeRaw(query: $query)
}
```

内部OPERATION，也有钩子。

{% hint style="info" %}
不要在钩子调用了自己的OEPRATION。不然会循环调用。
{% endhint %}
