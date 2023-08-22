# 跨域访问

在浏览器中，当一个网页的代码向不同源（协议、域名、端口）的服务器发送请求时，就会发生跨域请求。浏览器出于安全考虑，限制了跨域请求的执行，以防止恶意代码获取用户的敏感信息。

跨域问题本质上是浏览器限制，而不是服务端限制。可以查看Network，请求能够正确响应，response返回的值也是正确的，但浏览器不展示。

通常有如下措施可解决跨域访问，例如：

* 客户端浏览器解除跨域限制（理论上可以但是不现实）
* 发送JSONP请求替代XHR请求（并不能适用所有的请求方式，不推荐）
* 修改服务器端的设置（推荐）

推荐第三种方式，在服务器端设置跨域参数，可前往 <mark style="color:purple;">设置->跨域</mark> 设置，接下来我们逐一介绍参数。

## 允许源

Access-Control-Allow-Origin：指定允许访问的源（域名、协议、端口），使用通配符 `*` 表示允许所有源进行访问，也可以指定具体的源。

例如：Access-Control-Allow-Origin: \* 或 Access-Control-Allow-Origin: [http://example.com](http://example.com/)

当浏览器发出跨站请求时，服务器会校验当前请求是不是来自被允许的站点。服务端根据浏览器请求首部字段 `Origin` 判断。

**Origin组成**：协议+域名+端口号

## 允许方法

Access-Control-Allow-Methods：指定允许的HTTP请求方法，多个方法使用逗号分隔。常见的方法包括GET、POST、PUT、DELETE等。

例如：Access-Control-Allow-Methods: GET, POST

## 允许头

Access-Control-Allow-Headers：指定允许的自定义请求头，多个头部字段使用逗号分隔。默认为 \* ，表示不限制。

## 暴露头

Access-Control-Expose-Headers：指定响应中允许客户端访问的响应头，多个头部字段使用逗号分隔。

例如：Access-Control-Expose-Headers: Authorization

## 跨域时间

Access-Control-Max-Age：指定预检请求（OPTIONS请求）的有效期，单位为秒，默认 120 秒 。

在有效期内，浏览器无需再发送预检请求。例如：Access-Control-Max-Age: 3600

## 允许 Credentials

Access-Control-Allow-Credentials：指定是否允许发送Cookie等凭证信息。如果设置为true，则表示允许发送凭证信息；如果设置为false，则不允许发送凭证信息。

例如：Access-Control-Allow-Credentials: true
