# 使用API

学习了如何构建API后，接下来我们学习如何测试和使用API。

## 单API使用

点击“API管理”面板后的筛选按钮，或按住快捷键Ctrl+K，打开搜索面板，搜索对应API，并打开。

### 测试接口

在底部“输入”TAB中，输入参数，点击OPERATION编辑区上方工具栏的“测试”按钮，可在“响应”TAB中查看测试结果。

点击“测试”调用的是GraphQL端点，其执行格式为：

```bash
curl 'http://localhost:9123/app/main/graphql'
-H 'Accept: application/json'
--data-raw '{"query":"query MyQuery {\n bg_findFirstPost {\n authorId\n createdAt\n id\n published\n title\n auhor:User {\n email\n id\n name\n role\n }\n }\n}","variables":{},"operationName":"MyQuery"}'
--compressed
```

测试端点仅用于测试 GraphQL OPEARTION 到数据源的执行情况，未兼容指令。除跨源关联指令外，其他指令均不生效，如角色、响应转换和入参指令等。

{% hint style="info" %}
出于安全考虑，生产环境下，请不要暴露当前端口。
{% endhint %}

### 复制链接

对于已上线API，点击顶部的“链接复制ICON”，可获得对应的访问链接。

Query Operation：对应为GET请求，复制为普通URL，如下：

```
http://localhost:9991/operations/Goods/GetManyGoods
```

Mutation Operation：对应为POST请求，复制为curl，如下：

```bash
curl 'http://localhost:9991/operations/Goods/DeleteOneGoods' \
  -X POST  \
  -H 'Content-Type: application/json' \
  --data-raw '{"id":1}' \
  --compressed
```

Subscription Operation：对应为GET请求，复制为如下URL：

```
# 加上了?wg_sse=true
http://localhost:9991/operations/Sub?wg_sse=true
```

若Query Operation开启了实时查询，则复制为如下URL：

<pre><code># 加上了?wg_live=true后缀
<strong>http://localhost:9991/operations/Goods/GetManyGoods?wg_live=true
</strong></code></pre>

{% hint style="info" %}
复制连接中对应的域名：http://localhost:9991，需要前往设置->系统  API外网地址中修改。
{% endhint %}

## API预览

点击右上角"API预览ICON"，打开swagger文档页。

<figure><img src="../../.gitbook/assets/image (1) (5).png" alt=""><figcaption><p>预览页</p></figcaption></figure>

左侧为API列表，以文件夹作为分组名，未分组API位于Others分组，此外，还包含FileUpload分组，用于展示OSS对应的路由。

选择对应API，可以查看其Operation，输入入参后，点击“TRY”按钮，可进行测试。响应栏会展示响应状态码和响应结果。

若API需要授权才能访问，需要在右上角选择对应OIDC，进行登录。登录后可查看当前登录用户的基本信息。

此外，若想在POSTMAN等工具测试，点击“下载”按钮，可获取swagger文档。

## SDK生成

若想在客户端使用API，可点击状态栏“SDK模板”，下载对应模板，并配置生成路径，即可在项目中引用对应语言的客户端SDK。

更多详情，前往[sdk-sheng-cheng](../sdk-sheng-cheng/ "mention")。
