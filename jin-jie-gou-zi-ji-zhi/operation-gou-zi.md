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

<figure><img src="../.gitbook/assets/image (1) (1).png" alt=""><figcaption><p>飞布服务请求流程</p></figcaption></figure>

为了方便理解，我们采用如下OPERATION说明情况。

{% code title="Weather.graphql" %}
```graphql
query MyQuery($capital: String!) {
  weather_getCityByName(name: $capital) {
    weather {
      summary {
        title
        description
      }
      temperature {
        actual
        feelsLike
      }
    }
  }
}
```
{% endcode %}

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
  "request": { # 客户端请求对象，即请求9991端口的参数
    "method": "GET", 
    "requestURI":"/operations/Weather?capital=DE",
    "headers": {
      "Content-Type": "application/json; charset=utf-8"
    }，
    "body":null
  },
  "operationName": "Weather",
  "operationType": "query"
}

# JSON response
{
  "op": "Weather",
  "hook": "beforeOriginRequest",
  "response": {
    "skip": false, # 如果true，忽略该钩子的响应
    "cancel": false, # 如果true，取消请求
    "request": { # 修改后的客户端请求对象
      "method": "GET",
      "requestURI":"/operations/Weather?capital=DE",
      "headers": {
        "content-type": "application/json; charset=utf-8"
      },
      "body":null
    }
  }
}
```

{% tabs %}
{% tab title="nodejs" %}
暂未支持
{% endtab %}

{% tab title="golang" %}
```go
func BeforeOriginRequest(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientRequest, error) {
	// 实现OPEN API能力
	if key, ok := body.Request.Headers["key1"]; ok {
		realJwt := getAccessTokenByOpenKEY(key)
		body.Request.Headers["Authorization"] = "Bearer " + realJwt
		hook.Logger().Infof("match real JWT [%s] for key [%s]", realJwt, key)
	}
	// body.Request对应上面的response.request
	return body.Request, nil
}
```
{% endtab %}
{% endtabs %}

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
    "body": { # 请求数据源的body
      "variables": {
        "capital": "beijing"
      },
      "query": "query($capital: String!){weather_getCityByName: getCityByName(name: $capital){weather {summary {title description} temperature {actual feelsLike}}}}"
    }
  },
  "operationName": "Weather", # OPERATION 名称
  "operationType": "query", # OPERATION类型，QUERY、MUTATION、SUBSCRIPTION
  "__wg": { # 全局参数
    "clientRequest": { # 原始客户端请求，即请求9991端口的request对象
      "method": "GET",
      "requestURI": "/operations/Weather?code=beijing",
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
    "request": { # 修改访问数据源的请求对象
      "method": "POST",
      "requestURI": "https://weather-api.fireboom.com/",
      "headers": {
        "Accept": "application/json",
        "Content-Type": "application/json",
        "X-Request-Id": "83850325-9638-e5af-f27d-234624aa1824"
      },
      "body": {
        "variables": { "capital": "beijing" },
        "query": "query($capital: String!){weather_getCityByName: getCityByName(name: $capital){weather {summary {title description} temperature {actual feelsLike}}}}"
      }
    }
  }
}
```

{% tabs %}
{% tab title="nodejs" %}

{% endtab %}

{% tab title="golang" %}
```go
func OnOriginRequest(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientRequest, error) {
	fmt.Println("OnOriginRequest")
	// 修改请求数据源的请求头
	if body.Request.RequestURI == "https://api.openweathermap.org/data/2.5/weather?appid=322110335eaafeea8b31d8263910ac70&q=beijing" {
		body.Request.Headers["test1"] = "sss1"
	} else if body.Request.RequestURI == "http://localhost:58688/" {
		body.Request.Headers["test2"] = "sss2"
	} else if body.Request.RequestURI == "https://fireboom-gql.ansoncode.repl.co/graphql" {
		body.Request.Headers["test3"] = "sss3"
	} else {
		fmt.Println("未支持数据源...", body.Request.RequestURI)
	}
	fmt.Println("数据源:", body.Request.RequestURI)
	// body.Request对应上面JSON response的response.request
	return body.Request, nil
}
```
{% endtab %}
{% endtabs %}

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
    "requestURI": "https://weather-api.fireboom.com/",
    "headers": {
      "Content-Type": "application/json; charset=utf-8"
    },
    "body": {
      "data": { # 从数据源获取的数据
            "weather_getCityByName": {
              "weather": {
                "summary": { "title": "Clear", "description": "clear sky" },
                "temperature": { "actual": 290.45, "feelsLike": 289.23 }
            }
          }
        }
    }
  },
  "operationName": "Weather",
  "operationType": "query",
  "__wg": {
    "clientRequest": {
      "method": "GET",
      "requestURI": "/operations/Weather?code=beijing",
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
        "data": { # 钩子修改后的数据
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



{% tabs %}
{% tab title="nodejs" %}

{% endtab %}

{% tab title="golang" %}
```go
func OnOriginResponse(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientResponse, error) {
	fmt.Println("OnOriginResponse")
	if body.Response.RequestURI == "https://api.openweathermap.org/data/2.5/weather?appid=322110335eaafeea8b31d8263910ac70&q=beijing" {
		fmt.Println("协议转换，将xml数据转成JSON")
	} else if body.Response.RequestURI == "http://localhost:58688/" {

	} else if body.Response.RequestURI == "https://fireboom-gql.ansoncode.repl.co/graphql" {

	} else {
		fmt.Println("未支持数据源...", body.Response.RequestURI)
	}
	// body.Response对应response.response
	return body.Response, nil
}
```
{% endtab %}
{% endtabs %}

## 局部钩子

与全局钩子不同，每个OPERTION都有对应的局部钩子，由开关单独控制。

### 前置钩子

前置钩子在 "执行OPERATION"前执行，可校验参数或修改输入参数。

#### 前置普通钩子

preResolve 钩子在参数注入后执行，能拿到请求入参，常用于入参校验。

<pre class="language-http"><code class="lang-http">http://{serverAddress}/operation/{operation}/preResolve

# Example:: http://localhost:9992/operation/Weather/preResolve

Content-Type: application/json
X-Request-Id: "83850325-9638-e5af-f27d-234624aa1824"

# JSON request
{
  "__wg": {
    "clientRequest": {
      "method": "GET",
      "requestURI": "/operations/Weather?code=beijing",
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
  "input": { "capital": "beijing" } # (可选)请求的输入参数
}

# JSON response
{
<strong>  "setClientRequestHeaders": { # 设置值后，可传递给后续的钩子
</strong>        "Accept": "text/html"
  },
  "op": "Weather",
  "hook": "preResolve"
}
</code></pre>

{% tabs %}
{% tab title="nodejs" %}

{% endtab %}

{% tab title="golang" %}
```go
func PreResolve(hook *base.HookRequest, body generated.WeatherBody) (res generated.WeatherBody, err error) {
	hook.Logger().Info("PreResolve")

	hook.Logger().Info("请求参数是：", body.Input.Capital)
	if body.Input.Capital != "beijing" {
		msg := fmt.Sprintf("入参必须是beijing，当前入参是：%s", body.Input.Capital)
		return body, errors.New(msg)
	}
	return body, nil
}
```
{% endtab %}
{% endtabs %}

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
      "requestURI": "/operations/Weather?code=beijing",
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
  "input": { "capital": "beijing" }
}

# JSON response
{
  "op": "Weather",
  "hook": "mutatingPreResolve",
  "input": { "capital": "修改为固定值" } # 用来修改入参
}
```

{% tabs %}
{% tab title="nodejs" %}

{% endtab %}

{% tab title="golang" %}
```go
func MutatingPreResolve(hook *base.HookRequest, body generated.WeatherBody) (res generated.WeatherBody, err error) {
	hook.Logger().Info("MutatingPreResolve")
	body.Input.Capital = "修改为固定值"
	return body, nil
}
```
{% endtab %}
{% endtabs %}

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
      "requestURI": "/operations/Weather?code=beijing",
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
  "input": { "capital": "beijing" },
   "response": {
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

# JSON response
{
  "op": "Weather",
  "hook": "postResolve"
}
```

{% tabs %}
{% tab title="nodejs" %}

{% endtab %}

{% tab title="golang" %}
```go
func PostResolve(hook *base.HookRequest, body generated.WeatherBody) (res generated.WeatherBody, err error) {
	hook.Logger().Info("PostResolve")
	// 可以拿到入参和响应
	hook.Logger().Info("发送一封邮件：", body.Input, body.Response.Data)
	return body, nil
}
```
{% endtab %}
{% endtabs %}

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
      "requestURI": "/operations/Weather?code=beijing",
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
  "input": { "capital": "beijing" },
  "response": {
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

# JSON response
{
  "op": "Weather",
  "hook": "mutatingPostResolve",
  "response": {
    "data": { # 修改响应的结果（响应结构可以变化！！！）
      "weather_getCityByName": {
        "weather": {
          "summary": { "title": "Clear", "description": "修改响应值" },
          "temperature": { "actual": 290.45, "feelsLike": 289.23 }
        }
      }
    }
  }
}
```

{% tabs %}
{% tab title="nodejs" %}

{% endtab %}

{% tab title="golang" %}
```go
func MutatingPostResolve(hook *base.HookRequest, body generated.WeatherBody) (res generated.WeatherBody, err error) {
	hook.Logger().Info("MutatingPostResolve")
	if body.Input.Capital == "beijing" {
		body.Response.Data.Weather_getCityByName.Weather.Summary.Description = "修改响应值"
		return body, nil
	}
	return body, nil
}
```
{% endtab %}
{% endtabs %}

{% hint style="info" %}
nodejs钩子可以修改响应形状，golang钩子暂未支持。
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
      "requestURI": "/operations/Weather?code=beijing",
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
  "input": { "capital": "beijing" }
}

# JSON response
{
  "op": "Weather",
  "hook": "mockResolve",
  "response": {
    "data": { # 模拟的数据
      "weather_getCityByName": {
        "weather": {
          "summary": { "title": "Clear", "description": "mockdata" }
        }
      }
    }
  }
}
```

{% tabs %}
{% tab title="nodejs" %}

{% endtab %}

{% tab title="golang" %}
```go
func MockResolve(hook *base.HookRequest, body generated.WeatherBody) (res generated.WeatherBody, err error) {
	hook.Logger().Info("MockResolve")
	body.Response = &base.OperationBodyResponse[generated.WeatherResponseData]{
		Data: generated.WeatherResponseData{
			Weather_getCityByName: generated.WeatherResponseData_weather_getCityByName{
				Weather: generated.WeatherResponseData_weather_getCityByName_weather{
					Summary: generated.WeatherResponseData_weather_getCityByName_weather_summary{
						Description: "mock data",
						Title:       "Clear",
					},
				},
			},
		},
	}
	return body, nil
}
```
{% endtab %}
{% endtabs %}

### 自定义钩子

customResolve 钩子在模拟钩子后执行。此钩子可用于用自定义OPERATION解析器替换默认解析器。

有两个逻辑：

* 返回结构体：跳过执行OPERATION的逻辑，返回结构体
* 返回NULL：继续执行OPERATION逻辑

<figure><img src="../.gitbook/assets/image (2) (1).png" alt=""><figcaption><p>注意看customResolve逻辑</p></figcaption></figure>

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
      "requestURI": "/operations/Weather?code=beijing",
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
  "input": { "capital": "beijing" }
}

# JSON response
{
  "op": "Weather",
  "hook": "customResolve",
  "response": { # 若response不为空会中断后置钩子执行
    "data": {
      "weather_getCityByName": {
        "weather": {
          "summary": { "title": "Clear", "description": "custom data" }
        }
      }
    }
  }
}
```



{% tabs %}
{% tab title="nodejs" %}

{% endtab %}

{% tab title="golang" %}
```go
func CustomResolve(hook *base.HookRequest, body generated.WeatherBody) (res generated.WeatherBody, err error) {
	hook.Logger().Info("CustomResolve")
	// 1，如果返回null,则继续接下来的逻辑
	if body.Input.Capital == "beijing" {
		hook.Logger().Info("写一些自定义逻辑")
		// 继续后面的流程
		return nil, nil
	}
	// 2，类似mock钩子
	body.Response = &base.OperationBodyResponse[generated.WeatherResponseData]{
		Data: generated.WeatherResponseData{
			Weather_getCityByName: generated.WeatherResponseData_weather_getCityByName{
				Weather: generated.WeatherResponseData_weather_getCityByName_weather{
					Summary: generated.WeatherResponseData_weather_getCityByName_weather_summary{
						Description: "custom data",
						Title:       "Clear",
					},
				},
			},
		},
	}
	return body, nil
}
```
{% endtab %}
{% endtabs %}
