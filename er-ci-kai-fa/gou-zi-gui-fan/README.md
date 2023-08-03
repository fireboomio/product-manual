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

### 熟悉并使用handlerbars语法生成代码 [sdk-sheng-cheng](../../shi-yong-bu-shu-shang-xian/sdk-sheng-cheng/ "mention")

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

