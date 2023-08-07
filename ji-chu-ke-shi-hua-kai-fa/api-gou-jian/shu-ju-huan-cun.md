# 数据缓存

`stale-while-revalidate` ，简称SWR，是 HTTP 的一种缓存失效策略。

对应到 HTTP 响应头  `cache-control` 的属性值，如图所示：

<figure><img src="../../.gitbook/assets/image (50).png" alt=""><figcaption></figcaption></figure>

这种策略首先从缓存中返回数据（过期的），同时发送 fetch 请求（重新验证），最后得到最新数据。

假设，我们设置：

* max-age=600,
* stale-while-revalidate=30

1，浏览器收到该响应后，会将结果缓存起来，请求的结果在600s内都是新鲜的（stale的反义词）。如果在600s内发起了相同请求，则直接返回缓存。

2，若 600s\~630s（max-age+ stale-while-revalidate）内发起了相同的请求 ，虽然缓存过期了，但依旧使用该缓存结果；另外同时进行异步 revalidate，即重新验证，相当于浏览器在背后自己发一个请求，响应结果留作下次使用。

3，在 630s 秒后，发起相同请求时，则将和第一次请求一样，从服务器获取响应结果，并且浏览器并把响应结果缓存起来。
