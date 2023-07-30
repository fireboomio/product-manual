# 内部调用

### 构建internalClient

1. 使用http框架封装一个构造请求和处理响应的函数
2. 设置请求Method=POST和Content-Type=application/json
3. 设置请求URL=${baseNodeUrl}/internal/operations/${operationPath}，其中baseNodeUrl为 [#jie-xi-fei-bu-sheng-cheng-de-json-pei-zhi-wen-jian](nei-bu-tiao-yong.md#jie-xi-fei-bu-sheng-cheng-de-json-pei-zhi-wen-jian "mention")serverOptions.nodeUrl的值，operationPath为 [#jie-xi-fei-bu-sheng-cheng-de-json-pei-zhi-wen-jian](nei-bu-tiao-yong.md#jie-xi-fei-bu-sheng-cheng-de-json-pei-zhi-wen-jian "mention")api.operations\[\*].path路径。
4. 请求参数，如下。其中input的入参为 [#shou-xi-bing-shi-yong-handlerbars-yu-fa-sheng-cheng-dai-ma](nei-bu-tiao-yong.md#shou-xi-bing-shi-yong-handlerbars-yu-fa-sheng-cheng-dai-ma "mention")生成出operation入参的结构体/对象

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

4. 响应结果，如下。其中data的返回结果为 [#shou-xi-bing-shi-yong-handlerbars-yu-fa-sheng-cheng-dai-ma](nei-bu-tiao-yong.md#shou-xi-bing-shi-yong-handlerbars-yu-fa-sheng-cheng-dai-ma "mention")生成出operation返回的结构体/对象

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

5. 用途：在钩子服务中使用飞布发布的接口，若存在钩子会调用 [#zhu-ce-operation-gou-zi](nei-bu-tiao-yong.md#zhu-ce-operation-gou-zi "mention")的函数
