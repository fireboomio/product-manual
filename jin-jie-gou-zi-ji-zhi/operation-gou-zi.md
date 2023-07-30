# OPERATION钩子

### 注册operation钩子

前置和后置的四个钩子中返回均会改写clientRequestHeaders（后续全局参数中headers均会改变）

局部钩子目的是扩展OPEARTION的能力，分别在“OPEARTION执行”前后执行，主要用途是参数校验和副作用触发，如创建文章后发送邮件通知审核。

详情见如下流程图。

![](../assets/hook-flow.png)

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
