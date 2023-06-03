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
2. 解析参数"\_\_wg"，json结构如下

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
            "customClaims": {},
            ”accessToken“: "", // accessToken字符串
            "rawAccessToken": [11, 12], // accessToken解析成json的字节数组
            "idToken": "", // idToken字符串
            "rawIdToken": [11, 12], // idToken解析成json的字节数组
        }
    }
}
```

3. 构建internalClient
4. 将全局参数和internalClient传递到后续请求中，方便后续钩子函数直接访问
5. 全局错误统一处理，返回json格式如下

```json
{
    "error": "error message"
}
```

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
    "hook": "beforeOriginRequest",
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
    "hook": "beforeOriginRequest",
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

### 注册上传钩子

### 注册operation钩子

### 注册graphql服务

### 注册健康检查



