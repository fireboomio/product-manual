# HTTP请求流程指令

飞布底层基于GraphQL协议构建，因此能充分发挥其指令系统的优势，扩展功能。

我们知道，GraphQL指令按作用位置分为3类：全局指令、入参指令、字段指令。详情见 [#graphql-zhi-ling-zhu-jie](../../he-xin-gai-nian/fu-wu-duan-operation.md#graphql-zhi-ling-zhu-jie "mention")

但，按照影响的流程也可以分为3类：HTTP请求流程、[内部调用流程](../../jin-jie-gou-zi-ji-zhi/nei-bu-tiao-yong.md#nei-bu-operation)、[跨源关联](kua-yuan-guan-lian.md#kua-yuan-guan-lian)。

本节重点介绍：HTTP请求流程指令。

其相关指令，按顺序可分为5类：登录校验、授权校验、入参校验、参数注入、响应转换！每一类都能完成特定的功能，体现在HTTP流程图上，对应一个个节点。

## HTTP请求流程

在介绍指令前，我们先学习OPERATION编译为REST API后的HTTP请求流程。

<figure><img src="../../.gitbook/assets/image (45).png" alt=""><figcaption></figcaption></figure>

1，登录校验：当客户端发起请求时，首先要检查用户是否登录

2，授权校验：接着检查用户拥有的角色是否匹配API需要的角色

3，入参校验：接着检查请求参数是否合法，如字符串入参是否满足正则表达式

4，注入参数：接着在服务端为某些参数设置值，包括环境变量、当前时间、UUID、Claims等

5，执行OPEARTION：去对应数据源获取数据，如数据库、REST API等，例如 如果是数据库，则可以把该过程类比为传统API开发框架的ORM层

6，响应转换：拍平上述过程拿到的响应，即将复杂的结构体变成扁平化结构体，方便前端使用

7，返回响应：返回结构体到客户端

{% hint style="info" %}
上述过程，未提及钩子对结果的影响，详情前往 [gou-zi-ji-zhi.md](../../jin-jie-gou-zi-ji-zhi/gou-zi-ji-zhi.md "mention")
{% endhint %}

## 登录校验指令

`@fromClaim`指令，又名登录校验指令，结合OIDC协议，实现了API数据权限控制。

{% code overflow="wrap" %}
```graphql
query GetOnetodo($uid: Int! @fromClaim(name: USERID) # 注入当前登录用户的ID) {
  data: todo_findFirsttodo(where: {user_id: {equals: $uid}}) {
    id
    title
    user_id
  }
}
```
{% endcode %}

当访问用`@fromClaim`指令修饰的接口时，引擎从当前登录用户会话的Claims中获取用户的基本信息，例如邮箱、UID等，并注入到OPERATION的入参中，保证本次请求只能获取或操作登录用户拥有的数据，从而实现数据权限控制。可注入字段具体包含USERID、EMAIL、EMAIL\_VERIFIED、NAME、 NICKNAME、 LOCATION、 PROVIDER。

{% embed url="https://www.bilibili.com/video/BV1Pk4y1b7je/" %}
08功能介绍-飞布如何限制API数据权限?
{% endembed %}

## 授权校验指令

`@rbac`指令，又名授权校验指令，实现了API接口的RBAC控制，对应到流程图上，为授权校验节点。它本质上是全局指令，作用于整个OPERATION上，修改了OPERATION的整体行为。

{% code overflow="wrap" %}
```graphql
query GetOnetodo($uid: Int!) @rbac(requireMatchAll: [admin]) # 拥有admin角色用户才能访问 {
  data: todo_findFirsttodo(where: {user_id: {equals: $uid}}) {
    id
    title
    user_id
  }
}
```
{% endcode %}

详情见 [授权与访问控制](../yan-zheng-he-shou-quan/shou-quan-yu-fang-wen-kong-zhi/)。

## 入参校验指令

`@jsonSchema`指令，又名入参校验指令，实现了API入参校验，体现在节点上如图所示。

```graphql
query GetOnetodo($uid: Int! @jsonSchema(pattern: "^ [0-9]*$")# 正则表达式校验入参 ) {
  data: todo_findFirsttodo(where: {user_id: {equals: $uid}}) {
    id
    title
    user_id
  }
}
```

入参校验指令支持正则表达式，可实现常用的入参合法性校验。

{% embed url="https://www.bilibili.com/video/BV1Ns4y1J7sQ/?vd_source=4be85d63cfdf7c8dbfaee9fce8d56792" %}
09功能介绍-飞布如何实现API入参校验？
{% endembed %}

主要包含如下校验方式：

* 正则校验：正则表达式校验参数，用法`@jsonSchema(pattern: "这里是正则表达式")`
* 通用校验：使用内置的规则校验入参，包含EMAIL和DOMAIN，用法`@jsonSchema(`commonPattern`:` EMAIL`)`
* 长度校验：针对字符串，校验其长度`minLength`和`maxLength`
* 大小校验：针对数字，校验其大小`minimum`和`maximum`
* 数组校验：针对数组，校验数组的尺寸`minItems`和`maxItems`

## 注入参数指令

注入参数由服务端注入，用于在服务端设置参数的值，不允许客户端修改，以保证数据安全。

注入参数修饰的入参，在编译为接口时，会自行去除，即编译的REST API的入参中不包含对应字段。

{% embed url="https://www.bilibili.com/video/BV1eM411p7j7/" %}
07功能介绍-飞布如何设置UpdateAt为当前时间？
{% endembed %}

### **injectGeneratedUUID**

`@injectGeneratedUUID` 指令，在服务端注入UUID，仅能修饰`string`字段。

```graphql
query myQuery ($uuid: String! @injectGeneratedUUID) {}
```

### **injectCurrentDatetime**

`@injectCurrentDatetime` 指令，服务端自动注入Datetime，仅能修饰`date`字段。包含两种使用方式：

* 内置格式：系统内置了多种日期规范，包括：

```graphql
query myQuery ($updatedAt: DateTime! @injectCurrentDateTime(format: ISO8601)) {}
```

```
# 其他格式枚举
  ISO8601：2006-01-02T15:04:05-0700
  ANSIC：Mon Jan _2 15:04:05 2006
  UnixDate：Mon Jan _2 15:04:05 MST 2006
  RubyDate：Mon Jan 02 15:04:05 -0700 2006
  RFC822：02 Jan 06 15:04 MST
  RFC822Z：02 Jan 06 15:04 -0700
  RFC850：Monday, 02-Jan-06 15:04:05 MST
  RFC1123：Mon, 02 Jan 2006 15:04:05 MST
  RFC1123Z：Mon, 02 Jan 2006 15:04:05 -0700
  RFC3339：2006-01-02T15:04:05Z07:00
  RFC3339Nano：2006-01-02T15:04:05.999999999Z07:00
  Kitchen：3:04PM
  Stamp：Jan _2 15:04:05
  StampMilli：Jan _2 15:04:05.000
  StampMicro：Jan _2 15:04:05.000000
  StampNano：Jan _2 15:04:05.000000000
```

* 自定义格式：采用符合golang规范的日期自定义格式，用法

```graphql
query myQuery ($updatedAt: DateTime! @injectCurrentDateTime(customFormat: "符合golang规范的日期")) {}
```

{% hint style="info" %}
不同数据库的datetime支持的格式不一样，可根据报错提醒，确定当前数据库支持的日期格式。
{% endhint %}

### **injectEnvironmentVariable**

`@injectEnvironmentVariable` 指令，服务端自动注入环境变量。环境变量可前往 设置-> 环境变量 进行配置。

```graphql
query myQuery ($applicationID: String! @injectEnvironmentVariable(name: "AUTH_APP_ID")) {}
```

## 响应转换指令

`@transform`指令，又名响应转换指令，用于修改响应的结构，主要是拍平嵌套较深的响应。它作用于选择集的对象字段上，也是字段指令。

{% embed url="https://www.bilibili.com/video/BV1Co4y1p7NN/?spm_id_from=333.788&vd_source=480e53a0bff92bc5368b03f2b1305865" %}
10功能介绍-飞布如何实现响应转换？
{% endembed %}

### 拍平对象

API所需的结构与数据库对应字段的层级不一致，通过该指令进行映射。

```graphql
query GettodoList {
  total: todo_aggregatetodo @transform(get: "_count.id") # 将_count.id值赋值给total字段 {
    _count {
      id
    }
  }
}
```

```json
# 转换前返回结果
{
    "total":{
        "_count":{
            "id":10
        }
    }
}
# 转换后返回结果
{
    "total":10
}
```

本质上是提取json结构的某个嵌套字段，然后赋值给上级字段。

### 对象数组到普通数组

此外，transfrom指令目前已支持直接将对象数组中的某个字段提取为普通数组。

在下方示例中，`presetList`字段原先的返回值类型是对象数组，通过`@transform(get: "[].I18n.Preset.code")`指令将`I18nItem`对象中的`I18n.Preset.code`提取为普通数组。

```graphql
query GetManyLearningLanguage {
  data: freetalk_findManyLearningLanguage(orderBy: {sort: desc}) {
    id
    name
    azure
    presetList: I18nItem @transform(get: "[].I18n.Preset.code")
   #将I18n.Preset.code值转换为普通数组并赋值给presetList字段 
   {
      I18n {
        Preset {
          code
        }
      }
    }
  }
} 
```

<figure><img src="../../.gitbook/assets/image (10) (1).png" alt=""><figcaption><p>对象数组转换前返回的结果</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/737eaec541ccda943bd6484b0e4cf5a.jpg" alt=""><figcaption><p>对象数组转换后返回的结果</p></figcaption></figure>
