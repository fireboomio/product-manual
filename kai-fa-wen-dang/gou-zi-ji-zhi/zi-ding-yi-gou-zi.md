# 自定义钩子

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

### 熟悉并使用handlerbars语法生成代码 [sdk-sheng-cheng](../sdk-sheng-cheng/ "mention")

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

3. 构建internalClient [#gou-jian-internalclient](zi-ding-yi-gou-zi.md#gou-jian-internalclient "mention")
4. 将全局参数和internalClient传递到后续请求中，方便后续钩子函数直接访问
5. 全局错误统一处理，返回json格式如下

```json
{
    "error": "error message"
}
```

### 构建internalClient



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
    "op": "Todo", // operation名字（多级目录以"__"分割）
    "hook": "beforeOriginRequest",
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
    "op": "Todo", // operation名字（多级目录以"__"分割）
    "hook": "onOriginRequest",
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
    "op": "Todo", // operation名字（多级目录以"__"分割）
    "hook": "onOriginResponse",
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

1. 前置普通钩子

* 路径：/operation/${operationPath}/preResolve
* 入参：

```json
{
   “input": {"name": "fireboom"} // operation入参
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
   “input": {"name": "fireboom"} // operation入参
}
```

* 出参：

```json
{
    “input": {"name": "fireboom"}, // operation入参
    "setClientRequestHeaders": ${__wg.clientRequest.headers} // 与全局参数__wg.clientRequest.headers格式保持一致
}
```

* 用途：作operation入参改写使用

3. 自定义解析钩子

* 路径：/operation/${operationPath}/customResolve
* 入参：

```json
{
   “input": {"name": "fireboom"} // operation入参
}
```

* 出参：

```json
{
    “response": {"name": "fireboom"}, // operation出参
    "setClientRequestHeaders": ${__wg.clientRequest.headers} // 与全局参数__wg.clientRequest.headers格式保持一致
}
```

* 用途：自定义返回值，若返回response不为空会中断后置钩子执行

4. 模拟数据钩子

* 路径：/operation/${operationPath}/mockResolve
* 入参：

```json
{
   “input": {"name": "fireboom"} // operation入参
}
```

* 出参：

```json
{
    “response": {"name": "fireboom"}, // operation出参
    "setClientRequestHeaders": ${__wg.clientRequest.headers} // 与全局参数__wg.clientRequest.headers格式保持一致
}
```

* 用途：模拟数据，此钩子打开会中断后置钩子执行

5. 后置普通钩子

* 路径：/operation/${operationPath}/postResolve
* 入参：

```json
{
   “input": {"name": "fireboom"}, // operation入参
   “response": {"name": "fireboom"} // operation出参
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
   “input": {"name": "fireboom"}, // operation入参
   “response": {"name": "fireboom"} // operation出参
}
```

* 出参：

```json
{
    “response": {"name": "fireboom"}, // operation出参
    "setClientRequestHeaders": ${__wg.clientRequest.headers} // 与全局参数__wg.clientRequest.headers格式保持一致
}
```

* 用途：作operation出参修改使用

### 注册graphql服务



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



