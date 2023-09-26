# RBAC

首先，我们学习RBAC基于角色的访问控制。

{% embed url="https://www.bilibili.com/video/BV1UL411C73g" %}
03功能介绍 如何用飞布实现API授权和访问控制？
{% endembed %}

## RBAC介绍

> RBAC，全称Role-based access control，即基于角色的权限控制，通过角色关联用户，角色关联权限的方式间接赋予用户权限。

当使用 RBAC 时，可以授予用户一个或多个角色，每个角色具有一个或多个权限。用户通过角色间接拥有权限，说白了就是给权限分个组，用户直接绑定分组，间接拥有分组的所有权限。相比直接为用户分配权限的ACL模式，RBAC实现了更灵活的访问控制。

<figure><img src="../../../.gitbook/assets/image (1) (1) (1) (1).png" alt=""><figcaption></figcaption></figure>

例如，用户张三的角色是销售经理，销售经理的权限有3个：客户列表、添加客户、删除客户。因此，用户张三就通过继承销售经理的角色，变相拥有了3个权限。

## RBAC0：授权流程

RBAC模型有多种变形，可分为RBAC0、RBAC1、RBAC2、RBAC3，但其核心都是RBAC0。

<figure><img src="../../../.gitbook/assets/image (1) (1) (1) (1) (2).png" alt=""><figcaption><p>RBAC0控制图</p></figcaption></figure>

如图是RBAC0的控制图，由四部分构成：

* 用户（User）
* 角色（Role）
* 会话（Session）
* 许可（Pemission），包括“操作”和“控制对象”

许可被赋予角色，而不是用户，当一个角色被指定给一个用户时，此用户就拥有了该角色所包含的许可。

用户与角色是多对多的关系，角色与许可是多对多的关系。这两个关系都是存储在数据库中的。

会话是动态的概念，用户必须通过会话才可以设置角色，是当前用户与激活的角色之间的映射关系。可以认为，每次登录都会创建一个会话，在登录时，为用户激活角色。

每个用户只能有1个会话，因此用户与会话是一对一的关系。会话时可以根据业务需求，激活对应的角色，所以会话与角色是一对多的关系。

总结下，实现权限控制核心是3步：

1. 用户绑定角色
2. 角色绑定权限
3. 创建会话激活角色

## 用户绑定角色

用户绑定角色指的是建立用户和角色间的关联关系，一般用数据库存储。

下面是示例数据模型：

```prisma
model admin_role {
  id     Int    @id @default(autoincrement())
  code   String @unique(map: "uniqueKey") @db.VarChar(20) # 角色码，对应Fireboom中的角色
  remark String @db.VarChar(40) # 角色描述
}

model admin_role_user {
  id      Int @id @default(autoincrement())
  role_id Int
  user_id Int

  @@unique([role_id, user_id], map: "role_user")
}

model admin_user {
  id            Int         @id @default(autoincrement())
  created_at    DateTime?   @db.DateTime(0)
  name          String      @db.VarChar(32)
  avatar        String?     @db.VarChar(255)
  phone         String?     @db.Char(13)
  password_salt String?     @db.VarChar(100)
  password      String?     @db.VarChar(100)
  country_code  String?     @db.VarChar(6)
  password_type String?     @db.VarChar(100)
  user_id       String?     @db.VarChar(255)
}
```

包含用户表`admin_user`和角色表`admin_role`，两表用`admin_role_user`关联。

为用户绑定角色，可用如下OPERATION实现：

{% code title="System/User/ConnectRole.graphql" %}
```graphql
# userId 用户ID，roleId 角色ID
mutation MyQuery($userId: Int!, $roleId: Int!) {
  data: main_createOneadmin_role_user(data: {role_id: $roleId, user_id: $userId}) {
    id
  }
}
```
{% endcode %}

```bash
curl 'http://localhost:9991/operations/System/User/ConnectRole' \
  -X POST  \
  -H 'Content-Type: application/json' \
  --data-raw '{"userId":1,"roleId":1}' \
  --compressed
```

## 角色绑定权限

Fireboom通过`@rbac`指令实现了角色绑定权限，详情见 [jie-kou-quan-xian-kong-zhi.md](../jie-kou-quan-xian-kong-zhi.md "mention")。

## 激活角色

Fireboom基于OIDC协议实现了 [shen-fen-yan-zheng](../../shen-fen-yan-zheng/ "mention") ，但OIDC中不包含角色相关的约定。用户通过OIDC流程登录后，claims中不包含`roles`字段。

因此，需要有个地方为用户动态注入`roles`字段，即用户第一次登录时，根据用户ID或email去特定数据源（可能是自有数据库或者其他数据源）查找其关联的角色，并绑定到roles字段上。

借助 [shen-fen-yan-zheng-gou-zi.md](../../../jin-jie-gou-zi-ji-zhi/shen-fen-yan-zheng-gou-zi.md "mention")，可实现上述功能，常用下面两个钩子：

* `mutatingPostAuth`：在用户授权登录后触发，可以**设置**用户信息，包括用户角色
* `revalidateAuth` ：在用户主动更新用户信息时触发，能**更新**缓存中的用户信息，包括用户角色

<figure><img src="../../../.gitbook/assets/image (56).png" alt=""><figcaption></figcaption></figure>

### 编写钩子

1. 在身份验证面板中点击“<img src="http://localhost:9123/assets/workbench/panel-role.png" alt="头像" data-size="line">”，进入“角色管理”TAB
2. 根据业务需求添加角色，系统默认内置admin和user角色（请确保必须有1个角色）
3. 切换到“身份鉴权”TAB，开启`mutatingPostAuthentication` 和 `revalidateAuthentication`
4. 编写按照下述用例，编写钩子，启动钩子

{% hint style="info" %}
步骤2中的角色列表必须包含步骤3中激活的角色！
{% endhint %}

设置/更新用户角色示例代码：

{% tabs %}
{% tab title="golang" %}
```go
func MutatingPostAuthentication(hook *base.AuthenticationHookRequest) (*plugins.AuthenticationResponse, error) {
	user := hook.User
	// 设置用户的角色
	user.Roles = []string{"user", "asssitent"}

	return &plugins.AuthenticationResponse{User: user, Status: "ok"}, nil
}
```

```go
func Revalidate(hook *base.AuthenticationHookRequest) (*plugins.AuthenticationResponse, error) {
	fmt.Println("Revalidate", hook.User.UserId)
	user := hook.User
	// 更新用户的角色
	user.Roles = []string{"system"}

	return &plugins.AuthenticationResponse{
		Status: "ok",
		User:   user,
	}, nil
}
```
{% endtab %}
{% endtabs %}

详情，请前往 [shen-fen-yan-zheng-gou-zi.md](../../../jin-jie-gou-zi-ji-zhi/shen-fen-yan-zheng-gou-zi.md "mention")

### 钩子测试

1. 点击顶部菜单栏的“<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAMAAAC7IEhfAAAAY1BMVEUAAADU1NRjZmxvcnePkZVgY2rPz9BfY2poaHTAwMOAhIxoa3FgYmpgYmlgY2nFxcbU1NRmZm+Ag427u73U1NRfYml/g4zt7e3k5eXW1te9vsCztLefoaWanKCOkJVydXtucXecDQKGAAAAFXRSTlMAzP336NDOiAvTz/rn2tjSph7Qs6d9epWLAAAAjElEQVQ4y+2T2Q6EIAxFK+A6mzMj4q7//5VaYngCG2N8cDkvNOlJSG9TuCq+XMQ3oiQ4p0jGsx+/fCIByDwrqRFzDYDn4BatYiw4Y1zEhBgIJjUsjJbED5eG19ctBtrr66rD9x05RYH9oVBKtViFTvGB7UZNlFg9N4n01/QwdDwrA0/mU0jtK/zDYRgBwgsrsPomQg4AAAAASUVORK5CYII=" alt="预览" data-size="line">”，前往API预览页，选择当前API
2. 在预览页顶部，选择OIDC供应商，点击前往登录
3. 登录后可查看用户信息，可以看到当前登录用户`roles`字段包含钩子中赋予的角色
