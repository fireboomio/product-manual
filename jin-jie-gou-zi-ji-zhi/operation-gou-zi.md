# OPERATION钩子

当客户端请求OPERATION编译成的API时（对应9991端口），依次经历登录校验，授权校验，入参校验，参数注入，执行OPERATION，响应转换等步骤。

接下来，我们学习下，该流程图上的分支：OPERATION钩子。

它分为2大类：全局钩子和局部钩子。

全局钩子所有的API共用同一个，都可以修改请求头对象，即ClientRequest。全局钩子有3个，包括：网关钩子（beforeOriginRequest）、数据源前置钩子(onOriginRequest)、数据源后置钩子（onOriginResponse）。

局部钩子每个API单独启停，从是否改变请求逻辑的角度，也可以分为两类：

* 不改变流程的前后置钩子：
  * 前置钩子：preResolve、mutatingPreResolve
  * 后置钩子：postResolve、mutatingPostResolve
* 能改变流程：
  * 模拟钩子：mockResolve
  * 自定义钩子：customResolve

## HTTP请求流程图

<figure><img src="../.gitbook/assets/image (1).png" alt=""><figcaption><p>飞布服务请求流程</p></figcaption></figure>

## 全局钩子

### 网关钩子

`beforeOriginRequest` 钩子又名网关钩子，会拦截所有HTTP请求，可以改写请求内容，也可以终止后续流程。

```http
http://{serverAddress}/global/httpTransport/beforeOriginRequest

# Example:: http://localhost:9992/global/httpTransport/beforeOriginRequest

Content-Type: application/json
X-Request-Id: "83821325-9638-e1af-f27d-234624aa1824"

# JSON request
{
  "request": { // 与全局参数路径__wg.clientRequest格式一致
    "method": "POST",
    "requestURI":"/operations/Weather",
    "headers": {
      "Content-Type": "application/json; charset=utf-8"
    },
    "body": { "data": { "country": { "code": "DE", "name": "Germany", "capital": "Berlin" } } }
  },
  "operationName": "Weather",
  "operationType": "query"
}

# JSON response
{
  "op": "Weather",
  "hook": "beforeOriginRequest",
  "response": {
    "skip": false,
    "cancel": false,
    "request": { // 与全局参数路径__wg.clientRequest格式一致
      "statusCode": 200,
      "status": "200 OK",
      "method": "POST",
      "requestURI": "https://weather-api.fireboom.io/",
      "headers": {
        "content-type": "application/json; charset=utf-8"
      },
      "body": {
        "data": {
          "weather_getCityByName": {
            "weather": {
              "summary": { "title": "Clear", "description": "clear sky" },
              "temperature": { "actual": 290.45, "feelsLike": 289.23 }
            }
          }
        }
      }
    }
  }
}
```

### 数据源前置钩子

`onOriginRequest` 钩子又名数据源前置钩子，在请求每个数据源之前生效。可修改请求数据源时的请求头，也可以用来取消请求。常见用例是：动态授权，动态注入授权参数

```http
http://{serverAddress}/global/httpTransport/onOriginResponse

# Example:: http://localhost:9992/global/httpTransport/onOriginResponse

Content-Type: application/json
X-Request-Id: "83850325-9638-e5af-f27d-234624aa1824"

# JSON request
{
  "request": {
    "method": "POST", # 请求钩子的方法
    "requestURI": "https://weather-api.fireboom.io/", # 数据源的请求地址，可用来区分不同数据源
    "headers": {
      "Accept": "application/json",
      "Content-Type": "application/json",
      "X-Request-Id": "83850325-9638-e5af-f27d-234624aa1824"
    },
    "body": {
      "variables": {
        "capital": "Berlin"
      },
      "query": "query($capital: String!){weather_getCityByName: getCityByName(name: $capital){weather {summary {title description} temperature {actual feelsLike}}}}"
    }
  },
  "operationName": "Weather", # OPERATION 名称
  "operationType": "query", # OPERATION类型，QUERY、MUTATION、SUBSCRIPTION
  "__wg": { # 全局参数
    "clientRequest": { # 原始客户端请求，即请求9991端口的request对象
      "method": "GET",
      "requestURI": "/operations/Weather?code=DE",
      "headers": {
        "Accept": "application/json",
        "Content-Type": "application/json"
      }
    },
    "user": { # （可选）授权用户的信息
      "userID": "1",
      "roles": ["user"]
    }
  }
}

# JSON response
{
  "op": "Weather",
  "hook": "onOriginRequest",
  "response": {
    "skip": false, # 如果true，忽略该钩子的响应
    "cancel": false, # 如果true，取消请求
    "request": {
      "method": "POST",
      "requestURI": "https://weather-api.fireboom.com/",
      "headers": {
        "Accept": "application/json",
        "Content-Type": "application/json",
        "X-Request-Id": "83850325-9638-e5af-f27d-234624aa1824"
      },
      "body": {
        "variables": { "capital": "Berlin" },
        "query": "query($capital: String!){weather_getCityByName: getCityByName(name: $capital){weather {summary {title description} temperature {actual feelsLike}}}}"
      }
    }
  }
}
```

### 数据源后置钩子

`onOriginRequest` 钩子又名数据源后置钩子，在请求每个数据源之后生效。可修改数据源返回的数据，也可以用来取消请求。常见用例是：若数据源不支持json，可用该方法转换协议

```http
http://{serverAddress}/global/httpTransport/onOriginResponse

# Example:: http://localhost:9992/global/httpTransport/onOriginResponse

Content-Type: application/json
X-Request-Id: "83850325-9638-e5af-f27d-234624aa1824"

# JSON request
{
  "response": {
    "statusCode": 200,
    "status": "200 OK",
    "method": "POST",
    "requestURI": "https://countries.trevorblades.com/",
    "headers": {
      "Content-Type": "application/json; charset=utf-8"
    },
    "body": { "data": { "country": { "code": "DE", "name": "Germany", "capital": "Berlin" } } }
  },
  "operationName": "Weather",
  "operationType": "query",
  "__wg": {
    "clientRequest": {
      "method": "GET",
      "requestURI": "/operations/Weather?code=DE",
      "headers": {
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
        "Accept-Encoding": "gzip, deflate, br",
        "Accept-Language": "de-DE,de;q=0.9,en-DE;q=0.8,en;q=0.7,en-GB;q=0.6,en-US;q=0.5"
      }
    }
  }
}

# JSON response
{
  "op": "Weather",
  "hook": "onOriginResponse",
  "response": {
    "skip": false,
    "cancel": false,
    "response": {
      "statusCode": 200,
      "status": "200 OK",
      "method": "POST",
      "requestURI": "https://weather-api.fireboom.io/",
      "headers": {
        "access-control-allow-origin": "*",
        "content-type": "application/json; charset=utf-8",
        "date": "Mon, 01 May 2023 10:46:39 GMT"
      },
      "body": {
        "data": {
          "weather_getCityByName": {
            "weather": {
              "summary": { "title": "Clear", "description": "clear sky" },
              "temperature": { "actual": 290.45, "feelsLike": 289.23 }
            }
          }
        }
      }
    }
  }
}
```

## 局部钩子

与全局钩子不同，每个OPERTION都有对应的局部钩子，由开关单独控制。

### 前置钩子

前置钩子在 "执行OPERATION"前执行，可校验参数或修改输入参数。

#### 前置普通钩子

preResolve 钩子在参数注入后执行，能拿到请求入参，常用于入参校验。

```http
http://{serverAddress}/operation/{operation}/preResolve

# Example:: http://localhost:9992/operation/Weather/preResolve

Content-Type: application/json
X-Request-Id: "83850325-9638-e5af-f27d-234624aa1824"

# JSON request
{
  "__wg": {
    "clientRequest": {
      "method": "GET",
      "requestURI": "/operations/Weather?code=DE",
      "headers": {
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
        "Accept-Encoding": "gzip, deflate, br",
        "Accept-Language": "de-DE,de;q=0.9,en-DE;q=0.8,en;q=0.7,en-GB;q=0.6,en-US;q=0.5",
        "Cache-Control": "max-age=0",
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36"
      }
    },
    "user": {
      "userID": "1"
    }
  },
  "input": { "code": "DE" } # (可选)请求的输入参数
}

# JSON response
{
  "setClientRequestHeaders": ${__wg.clientRequest.headers} // 与全局参数__wg.clientRequest.headers格式保持一致
  "op": "Weather",
  "hook": "preResolve"
}
```

#### 前置修改入参钩子

mutatingPreResolve 钩子在参数注入后执行，能拿到请求入参，也能修改请求入参。

```http
http://{serverAddress}/operation/{operation}/mutatingPreResolve

# Example:: http://localhost:9992/operation/Weather/mutatingPreResolve

Content-Type: application/json
X-Request-Id: "83850325-9638-e5af-f27d-234624aa1824"

# JSON request
{
  "__wg": {
    "clientRequest": {
      "method": "GET",
      "requestURI": "/operations/Weather?code=DE",
      "headers": {
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
        "Accept-Encoding": "gzip, deflate, br",
        "Accept-Language": "de-DE,de;q=0.9,en-DE;q=0.8,en;q=0.7,en-GB;q=0.6,en-US;q=0.5",
        "Cache-Control": "max-age=0",
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36"
      }
    },
    "user": {
      "userID": "1"
    }
  },
  "input": { "code": "DE" }
}

# JSON response
{
  "op": "Weather",
  "hook": "mutatingPreResolve",
  "input": { "code": "US" } # 用来修改入参
}
```

### 后置钩子

后置钩子在 "执行OPERATION" 后执行，可触发自定义操作或修改响应结果。

#### 后置普通钩子

postResolve 钩子在响应转换后执行，能拿到请求入参和响应结果，常用于消息通知。

```http
http://{serverAddress}/operation/{operation}/postResolve

# Example:: http://localhost:9992/operation/Weather/postResolve

Content-Type: application/json
X-Request-Id: "83850325-9638-e5af-f27d-234624aa1824"

# JSON request
{
  "__wg": {
    "clientRequest": {
      "method": "GET",
      "requestURI": "/operations/Weather?code=DE",
      "headers": {
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
        "Accept-Encoding": "gzip, deflate, br",
        "Accept-Language": "de-DE,de;q=0.9,en-DE;q=0.8,en;q=0.7,en-GB;q=0.6,en-US;q=0.5",
        "Cache-Control": "max-age=0",
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36"
      }
    },
    "user": {
      "userID": "1"
    }
  },
  "input": { "code": "DE" }
}

# JSON response
{
  "op": "Weather",
  "hook": "postResolve"
}
```

#### 后置修改出参钩子

mutatingPostResolve 钩子在响应转换后执行，能拿到请求入参和响应结果，也能修改响应结果。

```http
http://{serverAddress}/operation/{operation}/mutatingPostResolve

# Example:: http://localhost:9992/operation/Weather/mutatingPostResolve

Content-Type: application/json
X-Request-Id: "83850325-9638-e5af-f27d-234624aa1824"

# JSON request
{
  "__wg": {
    "clientRequest": {
      "method": "GET",
      "requestURI": "/operations/Weather?code=DE",
      "headers": {
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
        "Accept-Encoding": "gzip, deflate, br",
        "Accept-Language": "de-DE,de;q=0.9,en-DE;q=0.8,en;q=0.7,en-GB;q=0.6,en-US;q=0.5",
        "Cache-Control": "max-age=0",
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36"
      }
    },
    "user": {
      "userID": "1"
    }
  },
  "input": { "code": "DE" },
  "response": {
    "data": {
      "weather": {
        "temperature": 10,
        "description": "Sunny"
      }
    }
  }
}

# JSON response
{
  "op": "Weather",
  "hook": "mutatingPostResolve",
  "response": {
    "data": { # 修改响应的结果
      "weather": {
        "temperature": 10,
        "description": "Sunny"
      }
    }
  }
}
```

{% hint style="info" %}
该钩子可以修改响应结构体，例如增加字段。当前仅node钩子支持，golang钩子暂不支持。
{% endhint %}

### 模拟钩子

mockResolve 钩子在前置钩子后执行，可以用来模拟操作的响应。开启后，将跳过查询数据源，详情看 [#http-qing-qiu-liu-cheng-tu](operation-gou-zi.md#http-qing-qiu-liu-cheng-tu "mention")。

```http
http://{serverAddress}/operation/{operation}/mockResolve

# Example:: http://localhost:9992/operation/Weather/mockResolve

Content-Type: application/json
X-Request-Id: "83850325-9638-e5af-f27d-234624aa1824"

# JSON request
{
  "__wg": {
    "clientRequest": {
      "method": "GET",
      "requestURI": "/operations/Weather?code=DE",
      "headers": {
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
        "Accept-Encoding": "gzip, deflate, br",
        "Accept-Language": "de-DE,de;q=0.9,en-DE;q=0.8,en;q=0.7,en-GB;q=0.6,en-US;q=0.5",
        "Cache-Control": "max-age=0",
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36"
      }
    },
    "user": {
      "userID": "1"
    }
  },
  "input": { "code": "DE" }
}

# JSON response
{
  "op": "Weather",
  "hook": "mockResolve",
  "response": {
    "data": { # 模拟的数据
      "weather": {
        "temperature": 10,
        "description": "Sunny"
      }
    }
  }
}
```

### 自定义钩子

customResolve 钩子在模拟钩子后执行。此钩子可用于用自定义OPERATION解析器替换默认解析器。

有两个逻辑：

* 返回结构体：跳过执行OPERATION的逻辑，返回结构体
* 返回NULL：继续执行OPERATION逻辑

<figure><img src="../.gitbook/assets/image (2).png" alt=""><figcaption><p>注意看customResolve逻辑</p></figcaption></figure>

```http
http://{serverAddress}/operation/{operation}/customResolve

# Example:: http://localhost:9992/operation/Weather/customResolve

Content-Type: application/json
X-Request-Id: "83850325-9638-e5af-f27d-234624aa1824"

# JSON request
{
  "__wg": {
    "clientRequest": {
      "method": "GET",
      "requestURI": "/operations/Weather?code=DE",
      "headers": {
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
        "Accept-Encoding": "gzip, deflate, br",
        "Accept-Language": "de-DE,de;q=0.9,en-DE;q=0.8,en;q=0.7,en-GB;q=0.6,en-US;q=0.5",
        "Cache-Control": "max-age=0",
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36"
      }
    },
    "user": {
      "userID": "1"
    }
  },
  "input": { "code": "DE" }
}

# JSON response
{
  "op": "Weather",
  "hook": "customResolve",
  "response": { # 若response不为空会中断后置钩子执行
    "data": {
      "weather": {
        "temperature": 10,
        "description": "Sunny"
      }
    }
  }
}
```
