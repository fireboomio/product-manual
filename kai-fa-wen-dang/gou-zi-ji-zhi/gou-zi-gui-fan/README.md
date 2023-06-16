# 钩子规范

钩子服务本质上是一个实现了飞布钩子规范的WEB服务，可以用任意后端语言实现。

如果你希望实现其他语言的 hook SDK，需要遵从如下协议。

### 搭建一个可以提供Rest服务的项目

常用的编程语言实现动态注册HTTP对外访问的框架有很多，比如：

1. Java：Spring Boot、Spring Cloud、JAX-RS、Vert.x等；
2. Python：Flask、Django、Tornado、FastAPI等；
3. Go：Gin、Echo、Beego等；
4. Node.js：Express、Koa、Hapi等。
5. Ruby：Ruby on Rails、Sinatra、Roda等；
6. PHP：Laravel、Symfony、Slim等；
7. Rust：Rocket、Actix-web等；
8. Kotlin：Ktor、Spring Boot等；
9. Swift：Vapor、Kitura等。

这些框架都提供了丰富的功能和插件，可以方便地实现HTTP对外访问。通常来说，这些框架都支持动态注册路由和中间件，可以根据不同的URL和HTTP方法，将请求路由到相应的处理函数，从而实现动态注册HTTP对外访问。使用这些框架可以很好地提高开发效率和代码质量。

### 熟悉并使用handlerbars语法生成代码 [sdk-sheng-cheng](../../sdk-sheng-cheng/ "mention")

1. 生成钩子函数时建议使用对应语言的范型来约束出入参
2. 模版生成最终应该包含

* 钩子函数出入参的对象/结构体定义，用来约束类型
* 全局钩子，用来 [#zhu-ce-quan-ju-gou-zi](./#zhu-ce-quan-ju-gou-zi "mention")
* 认证钩子，用来 [#zhu-ce-ren-zheng-gou-zi](./#zhu-ce-ren-zheng-gou-zi "mention")
* 查询钩子，用来 [#zhu-ce-operation-gou-zi](./#zhu-ce-operation-gou-zi "mention")中查询类型的钩子
* 变更钩子，用来 [#zhu-ce-operation-gou-zi](./#zhu-ce-operation-gou-zi "mention")中变更类型的钩子
* 订阅钩子，用来 [#zhu-ce-operation-gou-zi](./#zhu-ce-operation-gou-zi "mention")中订阅类型的钩子（暂时未支持调用）
* 上传钩子，用来 [#zhu-ce-shang-chuan-gou-zi](./#zhu-ce-shang-chuan-gou-zi "mention")
* graphql配置，用来 [#zhu-ce-graphql-fu-wu](./#zhu-ce-graphql-fu-wu "mention")

3. 建议最终生成的对象按照key-value的格式

*   全局钩子，value为注册函数${function}，key为

    ```
    BeforeOriginRequest/OnOriginRequest/OnOriginResponse
    ```
*   认证钩子，value为注册函数${function}，key为

    ```
    PostAuthentication/MutatingPostAuthentication/Revalidate/PostLogout
    ```
*   operation钩子，key为${operationPath}, value为

    ```json
    {
        "mockResolve": ${function},
        "preResolve": ${function},
        "postResolve": ${function},
        "MutatingPreResolve": ${function},
        "MutatingPostResolve": ${function},
        "CustomResolve": ${function}
    }
    ```
* 上传钩子，key为${provider}，value为

```json
{
    "image": { // ${profile}
        "preUpload": ${function},
        "postUpload": ${funtion}
    }
}
```

* graphql配置，主要包含以下信息

```json
[
    {
        "apiNamespace": "todo_gql", // 自定义数据源名称
        "enableGraphQLEndpoint": true, // 是否开启graphql端点访问
        "schema": ${graphql.Schema} // 自定义的graphql解析，需要使用对应语言的sdk
    }
]
```

4. ${function}为各个创建的不同目录下的文件，其中包含可执行函数，例如

* 全局钩子文件global/beforeRequest.(go/java/py/ts等)
* 认证钩子文件authentication/postAuthentication.(go/java/py/ts等)
* operation钩子hooks/${operationPath}/postResolve.(go/java/py/ts等)
* 上传钩子uploads/${provider}/${profile}/preUpload.(go/java/py/ts等)

5. ${function}最终需要注册到各个的接口时使用，最终流程为：

```
钩子接口(1.根据请求路径在注册的key/function中寻找对应钩子函数; 2.处理请求并利用范型转换入参) => 
钩子函数(自定义逻辑) => 
钩子接口(处理钩子函数返回最终提供响应)
```

### 解析飞布生成的json配置文件

1. 文件路径${钩子项目目录}/generated/fireboom.config.json
2. 文件内容示例如下：

```json
{
    
    "api": {
        "operations": [
            {
                "name": "Todo__CreateOne", 
                "path": "Todo/CreateOne", // 文档中${operationPath}使用此值
                "operationType": 1 // 0 QUERY, 1 MUTATION, 2 SUBSCRIPTION
            }
        ],
        "s3UploadConfiguration": [
            {
                "name": "oss-todo", // provider name 文档中${provider}使用此值
                "uploadProfiles": {
                    "audio": { // profile name 文档中${profile}使用此值 
                        "preUpload": true,
                        "postUpload": false
                    }
                }
            }
        ],
        "serverOptions": {
            "listen": {
                "host": {
                    "staticVariableContent": "0.0.0.0" // 钩子服务启动地址
                },
                "port": {
                    "staticVariableContent": "" // 钩子服务启动端口
                }
            },
            "logger": {
                "level: {
                    "staticVariableContent": "INFO" // 钩子服务日志级别
                }
            }
        },
        "nodeOptions": {
            "nodeUrl": {
                "staticVariableContent": "http://localhost:9991" // 飞布服务内网访问地址
            },
            "publicNodeUrl": {
                "staticVariableContent": "http://ip:port" // 飞布服务对外访问地址
            }
        }
    }
}
```

3. api.operations\[\*].path用来过滤 [#shou-xi-bing-shi-yong-handlerbars-yu-fa-sheng-cheng-dai-ma](./#shou-xi-bing-shi-yong-handlerbars-yu-fa-sheng-cheng-dai-ma "mention")生成的operation钩子函数
4. api.s3UploadConfiguration.name和api.s3UploadConfiguration.uploadProfiles.\*用来过滤 [#shou-xi-bing-shi-yong-handlerbars-yu-fa-sheng-cheng-dai-ma](./#shou-xi-bing-shi-yong-handlerbars-yu-fa-sheng-cheng-dai-ma "mention")生成的上传钩子函数
5. api.serverOptions.listen.port用来指定钩子服务启动端口号
6. api.nodeOptions.nodeUrl用来 [#gou-jian-internalclient](./#gou-jian-internalclient "mention")指定baseNodeUrl

### 解析全局参数

1. 所有的钩子请求都是POST请求，并且Content-Type=application/json
2. 解析参数"\_\_wg"，json结构如下（请dump body数据，防止后续请求使用body因为流关闭导致异常）

```json
{
    "__wg": {
        // 原有请求的详细信息
        "clientRequest": {
            "method": "POST", // GET/POST
            "requestURI": "/operations/Todo", // request.urlPath
            "headers": {
                "Content-Type": "application/json",
                "Authorization": "Bearer st-123123"
                ...
            }, // request.headers
            "body": [11, 12, 13, 44] // request.body 转为json的body
            "originBody": [11, 12, 13, 44] // request.body 初始的body(在beforeRequest中使用)
        },
        "user": {
            ”provider“: "authing",
            "providerId": "",
            "userId": "adfa7dafy9adfha1f", // 用户ID
            ”name“: "fireboom",
            "firstName": "",
            "lastName": "",
            ”middleName“: "",
            "nickName": "fireboom.io", // 用户昵称
            "preferredUsername": "",
            ”profile“: "",
            "picture": "",
            "website": "",
            ”email“: "",
            "emailVerified": false,
            "birthDate": "",
            ”zoneInfo“: "",
            "locale": "",
            "location": "",
            ”roles“: ["admin", "user"], // 用户权限角色
            "customAttributes": [""],
            "customClaims": {}, // 自定义属性，扩展用户信息，可以通过认证钩子修改
            ”accessToken“: "", // accessToken字符串
            "rawAccessToken": [11, 12], // accessToken解析成json的字节数组
            "idToken": "", // idToken字符串
            "rawIdToken": [11, 12], // idToken解析成json的字节数组
        }
    }
}
```

3. 构建internalClient [#gou-jian-internalclient](./#gou-jian-internalclient "mention")
4. 将全局参数和internalClient传递到后续请求中，方便后续钩子函数直接访问
5. 全局错误统一处理，返回json格式如下

```json
{
    "error": "error message"
}
```

### 构建internalClient

1. 使用http框架封装一个构造请求和处理响应的函数
2. 设置请求Method=POST和Content-Type=application/json
3. 设置请求URL=${baseNodeUrl}/internal/operations/${operationPath}，其中baseNodeUrl为 [#jie-xi-fei-bu-sheng-cheng-de-json-pei-zhi-wen-jian](./#jie-xi-fei-bu-sheng-cheng-de-json-pei-zhi-wen-jian "mention")serverOptions.nodeUrl的值，operationPath为 [#jie-xi-fei-bu-sheng-cheng-de-json-pei-zhi-wen-jian](./#jie-xi-fei-bu-sheng-cheng-de-json-pei-zhi-wen-jian "mention")api.operations\[\*].path路径。
4. 请求参数，如下。其中input的入参为 [#shou-xi-bing-shi-yong-handlerbars-yu-fa-sheng-cheng-dai-ma](./#shou-xi-bing-shi-yong-handlerbars-yu-fa-sheng-cheng-dai-ma "mention")生成出operation入参的结构体/对象

```json
{
    "input": {"name": "fireboom"}, // operation请求入参 
    "__wg": {
        "clientRequest": {
            "method": "POST",
            "requestURI": ${operationPath}, // operation请求路径
            "headers"：${__wg.clientRequest.headers} // 与全局参数__wg.clientRequest.headers格式一致，可以自定义请求头
        },
        "user": ${__wg.clientRequest.user} // 可以直接使用全局参数__wg.clientRequest.user
    }
}
```

4. 响应结果，如下。其中data的返回结果为 [#shou-xi-bing-shi-yong-handlerbars-yu-fa-sheng-cheng-dai-ma](./#shou-xi-bing-shi-yong-handlerbars-yu-fa-sheng-cheng-dai-ma "mention")生成出operation返回的结构体/对象

```json
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

5. 用途：在钩子服务中使用飞布发布的接口，若存在钩子会调用 [#zhu-ce-operation-gou-zi](./#zhu-ce-operation-gou-zi "mention")的函数

### 注册全局钩子

1. 预执行钩子

* 路径：/global/httpTransport/beforeOriginRequest
* 入参：json格式如下，body数据请使用request.originBody

```json
{
    "request": ${__wg.clientRequest} // 与全局参数路径__wg.clientRequest格式一致
    “operationName”: "Todo", // operation名字（多级目录以"__"分割）
    “operationType”: "QUERY" // QUERY/MUTATION/SUBSCRIPTION
}
```

* 出参：

```json
{
    "response": {
        "request": ${__wg.clientRequest} // 与全局参数路径__wg.clientRequest格式一致
    }
}
```

* 用途：在最初请求接受到的时候，修改出参中${response.request}的body和headers实现请求的改写

2. 前置钩子

* 路径：/global/httpTransport/onOriginRequest
* 入参：json格式如下，body数据请使用request.body

```json
{
    "request": ${__wg.clientRequest} // 与全局参数路径__wg.clientRequest格式一致
    “operationName”: "Todo", // operation名字（多级目录以"__"分割）
    “operationType”: "QUERY" // QUERY/MUTATION/SUBSCRIPTION
}
```

* 出参：

```json
{
    "response": {
        "request": ${__wg.clientRequest} // 与全局参数路径__wg.clientRequest格式一致
    }
}
```

* 用途：在operation执行前（前置钩子执行后），修改出参中${response.request}的body和headers实现请求的改写

3. 后置钩子

* 路径：/global/httpTransport/onOriginResponse
* 入参：json格式如下，body数据请使用response.body

```json
{
    "response": {
        "status": "200",
        "statusCode": 200,
        ...${__wg.clientRequest}
    }, // 全局参数路径__wg.clientRequest一一复制到此
    “operationName”: "Todo", // operation名字（多级目录以"__"分割）
    “operationType”: "QUERY" // QUERY/MUTATION/SUBSCRIPTION
}
```

* 出参：

```json
{
    "response": {
        "response": {
            "status": "200",
            "statusCode": 200,
            ...${__wg.clientRequest}
        } // 全局参数路径__wg.clientRequest一一复制到此
    }
}
```

* 用途：在operation执行后（后置钩子执行后），修改出参中${response.response}的statusCode、body和headers实现响应的改写

### 注册认证钩子

1. 后置普通钩子

* 路径：/authentication/postAuthentication
* 入参：请使用全局参数${\_\_wg.user}
* 出参：

```json
{
    "hook": "postAuthentication"
}
```

* 用途：认证成功后同步用户信息或记录用户访问日志

2. 后置修改信息钩子

* 路径：/authentication/mutatingPostAuthentication
* 入参：请使用全局参数${\_\_wg.user}
* 出参：

```json
{
    "hook": "mutatingPostAuthentication",
    "response": {
        "user": ${__wg.user} // 与全局参数路径__wg.user格式一致
        "status": "ok", // 状态不为ok时，使用message作为错误抛出
        "message": "not ok message"
    }
}
```

* 用途：认证成功后修改用户信息或中断认证

3. 后置重新校验钩子

* 路径：/authentication/revalidateAuthentication
* 入参：请使用全局参数${\_\_wg.user}
* 出参：

```json
{
    "hook": "revalidateAuthentication",
    "response": {
        "user": ${__wg.user} // 与全局参数路径__wg.user格式一致
        "status": "ok", // 状态不为ok时，使用message作为错误抛出
        "message": "not ok message"
    }
}
```

* 用途：请求携带revalidate参数会每次重走认证，默认从缓存获取user，根据参数选择是否进行重新认证校验或改写

### 注册上传钩子

1. 前置钩子

* 路径：/upload/${provider}/${profile}/preUpload
* 入参：

```json
{
    "file": {
        "Name": "TEST.JPG", // 文件名
        "size": 256, // 文件大小
        "type": "jpg" // content-type
    }
}
```

* 出参：

```json
{
    "fileKey": "test_modify.jpg" // 修改后的文件名
}
```

* 用途：校验文件name、size、type等信息，并返回自定义文件名（用户oss上传后显示的名称，默认随机字符串）

2. 后置钩子

* 路径：/upload/${provider}/${profile}/postUpload
* 入参：

```json
{
    "file": {
        "Name": "TEST.JPG", // 文件名
        "size": 256, // 文件大小
        "type": "jpg" // content-type
    }
}
```

* 出参：使用全局错误统一处理，正常返回200
* 用途：上传文件成功后自定义处理，可以用来记录上传日志

### 注册operation钩子

前置和后置的四个钩子中返回均会改写clientRequestHeaders（后续全局参数中headers均会改变）

局部钩子目的是扩展OPEARTION的能力，分别在“OPEARTION执行”前后执行，主要用途是参数校验和副作用触发，如创建文章后发送邮件通知审核。

详情见如下流程图。

![](../../../assets/hook-flow.png)

前置钩子在 "执行OPERATION"前执行，可校验参数或修改输入参数。

1. 前置普通钩子

* 路径：/operation/${operationPath}/preResolve
* 入参：

```json
{
   "op": ${operationPath}, // operationPath
   "hook": "preResolve", // hookName
   "input": {"name": "fireboom"} // operation入参
}
```

* 出参：

```json
{
    "setClientRequestHeaders": ${__wg.clientRequest.headers} // 与全局参数__wg.clientRequest.headers格式保持一致
}
```

* 用途：作operation入参校验使用

2. 前置修改入参钩子

* 路径：/operation/${operationPath}/mutatingPreResolve
* 入参：

```json
{
   "op": ${operationPath}, // operationPath
   "hook": "mutatingPreResolve", // hookName
   "input": {"name": "fireboom"} // operation入参
}
```

* 出参：

```json
{
    "input": {"name": "fireboom"}, // operation入参
    "setClientRequestHeaders": ${__wg.clientRequest.headers} // 与全局参数__wg.clientRequest.headers格式保持一致
}
```

* 用途：作operation入参改写使用

3. 自定义解析钩子

* 路径：/operation/${operationPath}/customResolve
* 入参：

```json
{
   "op": ${operationPath}, // operationPath
   "hook": "customResolve", // hookName
   "input": {"name": "fireboom"} // operation入参
}
```

* 出参：

```json
{
    "response": {"name": "fireboom"}, // operation出参
    "setClientRequestHeaders": ${__wg.clientRequest.headers} // 与全局参数__wg.clientRequest.headers格式保持一致
}
```

* 用途：自定义返回值，若返回response不为空会中断后置钩子执行

4. 模拟数据钩子

* 路径：/operation/${operationPath}/mockResolve
* 入参：

```json
{
   "op": ${operationPath}, // operationPath
   "hook": "mockResolve", // hookName
   "input": {"name": "fireboom"} // operation入参
}
```

* 出参：

```json
{
   "response": {"name": "fireboom"}, // operation出参
   "setClientRequestHeaders": ${__wg.clientRequest.headers} // 与全局参数__wg.clientRequest.headers格式保持一致
}
```

* 用途：模拟数据，此钩子打开会中断后置钩子执行

后置钩子在 "执行OPERATION" 后执行，可触发自定义操作或修改响应结果。

5. 后置普通钩子

* 路径：/operation/${operationPath}/postResolve
* 入参：

```json
{
   "op": ${operationPath}, // operationPath
   "hook": "postResolve", // hookName
   "input": {"name": "fireboom"}, // operation入参
   "response": {"name": "fireboom"} // operation出参
}
```

* 出参：

```json
{
    "setClientRequestHeaders": ${__wg.clientRequest.headers} // 与全局参数__wg.clientRequest.headers格式保持一致
}
```

* 用途：作operation后置通知使用

6. 后置修改出参钩子

* 路径：/operation/${operationPath}/mutatingPostResolve
* 入参：

```json
{
   "op": ${operationPath}, // operationPath
   "hook": "mutatingPostResolve", // hookName
   “input": {"name": "fireboom"}, // operation入参
   "response": {"name": "fireboom"} // operation出参
}
```

* 出参：

```json
{
    "response": {"name": "fireboom"}, // operation出参
    "setClientRequestHeaders": ${__wg.clientRequest.headers} // 与全局参数__wg.clientRequest.headers格式保持一致
}
```

* 用途：作operation出参修改使用

### 注册graphql服务

1. web界面（GET请求）

* 路径：/gqls/${apiNamespace}/graphql
* 入参：无
* 出参：返回html页面，将下面的文件读取动态修改其中${graphqlEndpoint}为web界面请求路径
* html文件&#x20;

{% file src="../../../.gitbook/assets/helix.html" %}

2. 内省和访问（POST请求）

* 路径：/gqls/${apiNamespace}/graphql
* 入参：

```json
{
    "query": "", // 飞布发过来的query，内省或访问
    "variables": {}, // graphql请求参数
    "operationName": "IntrospectionQuery" // 操作名，内省/“”
}
```

* 出参：使用对应sdk的返回graphql.result
* 用途：自定义数据源，处理复杂业务，内省后可以在飞布中提供api以供使用
* 注意：请求的参数需要结合 [#jie-xi-fei-bu-sheng-cheng-de-json-pei-zhi-wen-jian](./#jie-xi-fei-bu-sheng-cheng-de-json-pei-zhi-wen-jian "mention")中graphql配置的graphql.schema结合使用，组装对应语言的graphql-sdk需要的params

### 注册健康检查

* 路径：/health
* 入参：无
* 出参：

```json
{
    "status": "ok"
}
```

* 用途：钩子服务健康检查使用



