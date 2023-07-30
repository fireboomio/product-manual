# 新建API

了解飞布的API规范后，接下来我们学习下如何新建一条API。

新建API可能需要同时操作 API管理面板、超图Schema面板、编辑器GraphQL面板、概览面板、钩子面板以及工具栏。

为了更加直观的演示，我们提供了如下视频，供您快速上手：

{% embed url="https://www.bilibili.com/video/BV1MG4y1j7rf/" %}
\[飞布入门课程] 01-API创建
{% endembed %}



接下来，我们在分别介绍各面板的功能及用途。

下面内容只是简单概述，建议结合[视频](https://space.bilibili.com/3493080529373820/channel/collectiondetail?sid=1267673)进行学习。

## API管理面板

主要用途是新建或管理API，参考VSCODE的文件目录实现。

* API新建：点击右上角“+”或下方“新建”，可创建API
* API列表：展示所有API，不同状态说明如下
  * 方法：POST对应MUTATION，GET对应QUERY和SUBSCRIPTION
  * 实时：GET标识右上角的闪电符表示当前API为QUERY的实时查询或SUBSCRIPTION
  * 内部：API名称后的内部表示当前OPERATION仅供内部调用，不对外公开
  * 上线：未上线API用灰色表示表示
  * 非法：“非法”标识当前API的OPERATION有异常，无法正常使用
* 全局设置：应用于所有API的全局设置，主要包含授权配置、缓存配置、实时配置
* 批量新建：进入批量新建页，了解更多[前往查看](ding-yue.md)
* 端点测试：进入GraphQL测试页，用于探索超图的GraphQL端点
* 批量操作：按住shift键多选API，右击可进行批量操作，包括上下线、删除等

## 超图Schema面板

该面板是超图Schema的可视化展示，基于[GraphiQL Explorer](https://github.com/OneGraph/graphiql-explorer)项目二次开发，主要包含如下功能：

{% embed url="https://www.bilibili.com/video/BV1fx4y1N7wJ/" %}
05功能介绍-如何使用飞布超图面板构建API?
{% endembed %}

* 搜索：下拉选择命名空间，或输入函数名搜索所需方法
* 筛选：支持查询QUERY、变更MUTATION、订阅SUBSCRIPTION的筛选
* 勾选：展开对应方法，选择所需字段，并设置过滤条件
  * 选择字段：勾选方法下的蓝色字段，也可以展开折叠的蓝色字段，勾选嵌套字段（蓝色字段对应 GraphQL的 [字段](https://graphql.cn/learn/queries/#fields) ）
  * 过滤条件：勾选紫色字段，设置过滤条件。（紫色字段对应GraphQL的 [参数](https://graphql.cn/learn/queries/#arguments) ）
    * 默认值：在蓝色字段后的输入框中设置默认值
    * 函数入参：点击蓝色字段和输入框之间的`$`符，可以将过滤条件设置为函数入参

{% hint style="info" %}
\_join字段的用法比较特殊，详情见下文“[跨源关联](kua-yuan-guan-lian.md)”
{% endhint %}

## 编辑器GraphQL

GraphQL编辑器主要用来查看、修改以及测试 OPERATION ，基于 [GraphiQL](https://graphql-dotnet.github.io/docs/getting-started/graphiql/) 项目二次开发，具体功能如下：

<figure><img src="../../.gitbook/assets/image (3).png" alt=""><figcaption><p>Graphql编辑器</p></figcaption></figure>

* 编辑：手动修改OPERATION，支持语法提醒和自动补全
* 输入：输入OPERATION的入参
  * 可视化输入：用表格的形式展示入参，提供类型校验，和录入组件，如日期录入
  * <mark style="color:red;">源码输入</mark>（**推荐**）：在源码视图输入JSON，支持语法校验和提醒
* 响应：以JSON形式展示请求结果

接下来，介绍工具栏。

工具栏主要为GraphQL指令的可视化封装。用于帮助开发者快速掌握飞布指令的使用方法。

GraphQL指令分为三类：全局指令、入参指令、字段指令。

**全局指令：**作用于QUERY | MUTATION | SUBSCRIPTION，包括@rbac和@internalOperation。

**入参指令：**作用于GraphQL的入参，包括@fromClaim、@jsonSchema、@hooksVariable、@injectGeneratedUUID、@injectGeneratedUUID、@injectEnvironmentVariable、@internal。

**字段指令：**作用于GraphQL 字段，包括@transform。

详情见 [API指令](api-zhi-ling.md)。

## 概览面板

概览面板以可视化的形式展示当前 OPERATION 的运行机制，主要包括：HTTP流程图、内部调用流程图、订阅流程图。

详情见，[API运行机制](api-yun-hang-ji-zhi.md)。

## 设置面板

与全局设置对应，启用独立配置后，可对当前OPERATION单独设置。

* 开启授权：设置当前OPERATION的权限，开启后用户登录才能访问
* 缓存配置：仅对QUERY生效
  * 查询缓存：开启后启用缓存，提升服务性能
  * 最大时长：缓存有效时长，过期后，缓存失效
  * 重校验时长：缓存的客户端重校验时长
* 实时配置：仅对QUERY生效
  * 实时查询：开启后服务器轮询生效，可定时将消息推送至客户端
  * 轮询间隔：服务器轮询间隔，越短响应越快

## 钩子面板

将概览面板中的钩子单独提取出来，简化钩子展示。点击可展开钩子面板。

## API工具栏

编辑器顶部是API工具栏，主要功能如下：

<figure><img src="../../.gitbook/assets/image (8).png" alt=""><figcaption><p>API工具栏</p></figcaption></figure>

* 重命名：重命名API
* 状态：展示当前OPERATION是否被保存
* 克隆：克隆当前OPERATION，包括钩子
* 复制：复制当前OPERATION的API地址
  * QUERY  OPEARTION：对应GET请求，复制URL，粘贴到浏览器中可直接访问
    * 若开启实时查询，则转换为SSE推送，通过后缀\&wg\_live=true指定是否启用
  * MUTATION OPERATION：对应POST 请求，复制CURL链接，在terminal中可直接执行
  * SUBSCRIPTION OPERATION：对应GET请求（SSE推送），通过后缀 \&wg\_sse=true制定是否启用
* 保存：保存当前OPEARTION
* 开关：上下线当前API

{% hint style="info" %}
关于API和OPERATION用词说明：

在某种情况下，两者等价，因为OPEARTION可以构建为API。

但API的范围更广，它不仅包含OPERATION，还包含钩子。
{% endhint %}

