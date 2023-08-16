# 授权与访问控制（废弃）

飞布基于RBAC规范，结合GraphQL的注解能力，实现了API的的授权与访问控制。

{% embed url="https://www.bilibili.com/video/BV1UL411C73g" %}
03功能介绍 如何用飞布实现API授权和访问控制？
{% endembed %}

## 快速操作

### 基本设置

1. 在身份验证面板中点击“<img src="http://localhost:9123/assets/workbench/panel-role.png" alt="头像" data-size="line">”，进入“角色管理”TAB
2. 根据业务需求添加角色，系统默认内置admin和user角色（请确保必须有1个角色）
3. 切换到“身份鉴权”TAB，在auth目录下选择`mutatingPostAuthentication`文件
4. 编写钩子脚本或选择预制脚本，这里设置所有用户拥有`user`权限，启动钩子，详情前往[钩子章节](../../../jin-jie-gou-zi-ji-zhi/gou-zi-ji-zhi.md)

{% tabs %}
{% tab title="TS" %}
<pre class="language-typescript" data-title="custom-ts/auth/mutatingPostAuthentication.ts"><code class="lang-typescript"><strong>import { AuthenticationHookRequest, AuthenticationResponse } from 'fireboom-wundersdk/server'
</strong>import type { User,Role } from "generated/claims"
export default async function mutatingPostAuthentication(hook: AuthenticationHookRequest) : Promise&#x3C;AuthenticationResponse&#x3C;User>>{
 let roles=getRolesByUid(hook.user.userId!)
 return {
          user: {...hook.user,roles},
          status: 'ok',
        }
}
function getRolesByUid( uid:string ) :Role[]{
 //TODO：根据用户uid获取用户角色

 return ["user"]
}
</code></pre>
{% endtab %}

{% tab title="Go" %}
{% code title="custom-go/auth/mutatingPostAuthentication.go" %}
```go
package auth

import (
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
)

func MutatingPostAuthentication(hook *base.AuthenticationHookRequest) (*plugins.AuthenticationResponse, error) {

	hook.User.Roles = []string{
		"user",
	}
	res := plugins.AuthenticationResponse{
		Status: "ok",
		User:   *hook.User,
	}
	return &res, nil
}

```
{% endcode %}
{% endtab %}
{% endtabs %}

### API设置

1. 前往API管理面板，选择需要设置权限的API
2. 在GraphQL编辑区的工具栏中点击“@角色”，选择匹配模式并添加角色，如添加admin角色
3. 点击顶部菜单栏的“<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAMAAAC7IEhfAAAAY1BMVEUAAADU1NRjZmxvcnePkZVgY2rPz9BfY2poaHTAwMOAhIxoa3FgYmpgYmlgY2nFxcbU1NRmZm+Ag427u73U1NRfYml/g4zt7e3k5eXW1te9vsCztLefoaWanKCOkJVydXtucXecDQKGAAAAFXRSTlMAzP336NDOiAvTz/rn2tjSph7Qs6d9epWLAAAAjElEQVQ4y+2T2Q6EIAxFK+A6mzMj4q7//5VaYngCG2N8cDkvNOlJSG9TuCq+XMQ3oiQ4p0jGsx+/fCIByDwrqRFzDYDn4BatYiw4Y1zEhBgIJjUsjJbED5eG19ctBtrr66rD9x05RYH9oVBKtViFTvGB7UZNlFg9N4n01/QwdDwrA0/mU0jtK/zDYRgBwgsrsPomQg4AAAAASUVORK5CYII=" alt="预览" data-size="line">”，前往API预览页，选择当前API
4. 在预览页顶部，选择OIDC供应商，点击前往登录
5. 登录后可查看用户信息，可以看到当前登录用户roles字段包含user角色
6. 输入参数，测试接口，发现接口返回401
7. 重复步骤2，添加admin角色
8. 重复步骤6，可以看到接口执行成功

## 工作原理

### RBAC模型

RBAC（Role-Based Access Control）即：基于角色的权限控制。通过角色关联用户，角色关联权限的方式间接赋予用户权限。

![RBAC](https://image.woshipm.com/wp-files/2018/07/Tv1YwLlngzOs6oQN9UG0.png)

从模型中可以看出，我们要实现两个关键工作：

* 角色绑定权限：快速操作->API设置中介绍了该过程。
* 用户绑定角色：快速操作->基本设置中介绍了该过程。

因此，在调用API接口时，只需要判断用户是否拥有当前API的权限即可。

### 角色注入

由于OIDC规范中没有用户角色的相关声明，用户通过OIDC流程登录后，claim中不包含roles字段。而RBAC要求用户绑定角色，通过角色匹配接口权限。

因此，需要有个地方为用户动态注入roles字段，即用户第一次登录时，根据用户ID或email去特定数据源（可能是自有数据库或者其他数据源）查找其关联的角色，并绑定到roles字段上。考虑到灵活扩展，钩子是实现该功能的最佳场所。具体代码，参考 [上文](shou-quan-yu-fang-wen-kong-zhi.md#ji-ben-she-zhi)。

### 匹配规则

飞布支持四种匹配模式，我们以集合的概念进行讲解。

首先，需要知道三个集合域：

* 全部角色：角色管理列表中配置的角色

<figure><img src="../../../.gitbook/assets/image (10) (2).png" alt=""><figcaption><p>全部角色</p></figcaption></figure>

* 用户角色：用户拥有的角色，对应[mutatingPostAuthentication](shou-quan-yu-fang-wen-kong-zhi.md#ts)钩子中为用户设置的角色
* API角色：API RBAC指令上所设置的角色

```graphql
query GetOneTodo($id: Int!) @rbac(requireMatchAny: [admin]) {
  data: todo_findFirstTodo(where: {id: {equals: $id}}) {
    id
    title
    completed
    createdAt
  }
}
```

然后，了解四种集合关系：

* requireMatchAll：全部匹配，用户角色包含API角色时，可访问
* requireMatchAny：任意匹配，用户角色与API角色有交集时，可访问
* denyMatchAll：非全部匹配，当任意匹配或互斥匹配时，可访问
* denyMatchAny：互斥匹配，用户角色与API角色互斥时，可访问

最后，以全部角色作为全集，结合四种关系，看用户角色和API角色的交并补情况，确定当前用户是否能访问当前API。

## 下一步

飞布提供了完整的后台管理示例，您可以[参考代码](https://github.com/fireboomio/fb-admin)实现自己的管理后台。
