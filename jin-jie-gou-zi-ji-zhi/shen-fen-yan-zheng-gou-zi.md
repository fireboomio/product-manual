# 身份验证钩子

为了定制OIDC身份认证流程，飞布提供了3种类型的钩子：认证钩子、重校验钩子、登出钩子。

## 认证钩子

首先，我们学习登录过程中，需要用到的钩子——认证钩子。

它在用户完成OIDC供应商授权后触发，若想重复触发，需要先退出客户端和飞布服务的登录态，无需退出客户端和oidc供应商的登录态！

认证钩子有2个，先执行postAuth，在执行mutatingpost。

<figure><img src="../.gitbook/assets/image.png" alt=""><figcaption></figcaption></figure>

### postAuth

`postAuthentication` 钩子又名认证前置钩子，在身份验证已经验证并且用户已经通过身份验证之后执行。

主要用例：

* 同步OIDC用户至自己的数据库

```http
http://{serverAddress}/authentication/postAuthentication

Example:: http://localhost:9992/authentication/postAuthenthication

Content-Type: application/json
X-Request-Id: "83821325-9638-e1af-f27d-234624aa1824"

# JSON request
{
  "__wg": {
    "clientRequest": {},
    "user": {
      "access_token": "<secret>", // 访问令牌
      "id_token": "<secret>" // ID令牌
    }
  }
}

# JSON response
no response
```



{% tabs %}
{% tab title="golang" %}
```go
func PostAuthentication(hook *base.AuthenticationHookRequest) error {
	user := hook.User
	// 将用户同步到数据库
	upsertRes, err := plugins.ExecuteInternalRequestMutations[generated.Lesson0305__UpsertUserInput, generated.Lesson0305__UpsertUserResponseData](hook.InternalClient, generated.Lesson0305__UpsertUser, generated.Lesson0305__UpsertUserInput{
		UserId: user.UserId,
		Create: &generated.Todo_UserCreateInput{
			ProviderId: user.ProviderId,
			UserId:     user.UserId,
			Name:       user.Name,
			Nickname:   user.NickName,
			Email:      "E" + user.Email,
		},
		Update: &generated.Todo_UserUpdateInput{
			Name: &generated.Todo_StringFieldUpdateOperationsInput{
				Set: user.Name,
			},
			Nickname: &generated.Todo_StringFieldUpdateOperationsInput{
				Set: user.NickName,
			},
			Email: &generated.Todo_StringFieldUpdateOperationsInput{
				Set: "E" + user.Email,
			},
		},
	})
	if err != nil {
		return err
	}
	fmt.Println(upsertRes)
	return nil
}
```
{% endtab %}

{% tab title="nodejs" %}

{% endtab %}
{% endtabs %}

### mutatingPostAuth

`mutatingPostAuthentication` 钩子又名认证后置钩子，钩子在用户登录并验证身份后执行，可以用来在用户对象返回给客户端之前改变它。

主要用例：

* 为用户绑定角色
* 中断认证流程

```http
http://{serverAddress}/authentication/mutatingPostAuthentication

Example:: http://localhost:9992/authentication/mutatingPostAuthentication

Content-Type: application/json
X-Request-Id: "83821325-9638-e1af-f27d-234624aa1824"

# JSON request
{
  "__wg": {
    "clientRequest": {},
    "user": {
      "access_token": "<secret>", // 访问令牌
      "id_token": "<secret>" // ID令牌
    }
  }
}

# JSON response
{
  "hook": "mutatingPostAuthentication",
  "response": {
    "status": "ok", // 枚举值，ok或deny， deny表示拒绝登录
    "message": "not ok message", // 可选，为deny时填写
    "user": { // 可选，为ok时填写
      "userID": "1",
      "roles":["admin"] // 为用户修改角色
    }
  }
}
```



{% tabs %}
{% tab title="golang" %}
```go
func MutatingPostAuthentication(hook *base.AuthenticationHookRequest) (*plugins.AuthenticationResponse, error) {
	user := hook.User
	user.Roles = []string{"user", "asssitent"}
	// 修改用户的角色
	return &plugins.AuthenticationResponse{User: user, Status: "ok"}, nil
}
```
{% endtab %}

{% tab title="nodejs" %}

{% endtab %}
{% endtabs %}



两者的区别是，前者无法修改user对象，后者可以修改user或终止流程。其返回参数包含，user\status\以及message。

值得注意的是，这两个钩子的`user`，是从oidc拿到的原始user。而其他钩子的user，可能经由`mutatingpost`修改过了。详情可以参考时序图！

### 隐式模式

除了授权码模式，隐式模式也支持认证钩子。下面是其时序图。

<figure><img src="../.gitbook/assets/image (2).png" alt=""><figcaption></figcaption></figure>

当用户使用`access_token`访问任意OPERATION生成的接口时，都会走该流程。

第一次使用时，先去authing上获取用户信息，然后执行post钩子，接着执行mutatingpost钩子，并缓存用户信息。后续使用时，直接用缓存，不走钩子。

此外，若请求OPERATION 接口时，加上了revalidate后缀，则不用缓存，必走钩子。

总结下，第一次访问时，认为是登录，后续有了登录态，就不走登录流程了。

## 重校验钩子

`revalidateAuthentication` 钩子又名重校验钩子，引擎重新验证用户身份时执行。

主要用例：

* 在不退出登录的情况下，更新用户的信息，如角色等

<figure><img src="../.gitbook/assets/image (1).png" alt=""><figcaption></figcaption></figure>

OIDC 授权码模式和隐式模式调用该构造的方式不同。

* token模式（隐式模式）：携带access token请求任意OPERATION接口且携带revalidate字段，例如：operations/op?revalidate=true，同时调用3个钩子，包括revalidate
* cookie模式（授权码模式）：携带cookie，请求auth/user端点且携带revalidate字段，例如：/auth/user?revalidate=true

详情请参考：授权码模式->[#deng-lu-liu-cheng](../ji-chu-ke-shi-hua-kai-fa/shen-fen-yan-zheng/shou-quan-ma-mo-shi/#deng-lu-liu-cheng "mention")

```http
http://{serverAddress}/authentication/revalidateAuthentication

Example:: http://localhost:9992/authentication/revalidateAuthentication

Content-Type: application/json
X-Request-Id: "83821325-9638-e1af-f27d-234624aa1824"

# JSON request
{
  "__wg": {
    "clientRequest": {},
    "user": {
      "access_token": "<secret>", // 访问令牌
      "id_token": "<secret>" // ID令牌
    }
  }
}

# JSON response
{
  "hook": "mutatingPostAuthentication",
  "response": {
    "status": "ok", // 枚举值，ok或deny， deny表示拒绝登录
    "message": "not ok message", // 可选，为deny时填写
    "user": { // 可选，为ok时填写
      "userID": "1",
      "roles":["admin"] // 为用户修改角色
    }
  }
}
```



{% tabs %}
{% tab title="golang" %}
```go
func Revalidate(hook *base.AuthenticationHookRequest) (*plugins.AuthenticationResponse, error) {
	fmt.Println("Revalidate", hook.User.UserId)
	user := hook.User
	user.Roles = []string{"system"}

	return &plugins.AuthenticationResponse{
		Status: "ok",
		User:   user,
	}, nil
}
```
{% endtab %}

{% tab title="nodejs" %}

{% endtab %}
{% endtabs %}

## 登出钩子

`postLogout` 钩子又名登出钩子，在用户注销后执行。该钩子处于身份验证生命周期的最后一环，主要用于做些收尾操作。

主要用例：

* 记录用户登出时间

<figure><img src="../.gitbook/assets/image (3).png" alt=""><figcaption></figcaption></figure>

客户端请求退出登录端点时，会触发退出登录钩子。

```http
http://{serverAddress}/authentication/postLogout

Example:: http://localhost:9992/authentication/postLogout

Content-Type: application/json
X-Request-Id: "83821325-9638-e1af-f27d-234624aa1824"

# JSON request
{
  "__wg": {
    "clientRequest": {},
    "user": {
      "access_token": "<secret>", // 访问令牌
      "id_token": "<secret>" // ID令牌
    }
  }
}

# JSON response
no response
```



{% tabs %}
{% tab title="golang" %}
```go
func PostLogout(hook *base.AuthenticationHookRequest) error {

	fmt.Println(hook.User)
	return nil
}
```
{% endtab %}

{% tab title="nodejs" %}

{% endtab %}
{% endtabs %}
