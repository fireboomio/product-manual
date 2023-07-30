# 自定义逻辑

### 注册graphql服务

1. web界面（GET请求）

* 路径：/gqls/${apiNamespace}/graphql
* 入参：无
* 出参：返回html页面，将下面的文件读取动态修改其中${graphqlEndpoint}为web界面请求路径
* html文件&#x20;

{% file src="../.gitbook/assets/helix.html" %}

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
* 注意：请求的参数需要结合 [#jie-xi-fei-bu-sheng-cheng-de-json-pei-zhi-wen-jian](zi-ding-yi-luo-ji.md#jie-xi-fei-bu-sheng-cheng-de-json-pei-zhi-wen-jian "mention")中graphql配置的graphql.schema结合使用，组装对应语言的graphql-sdk需要的params
