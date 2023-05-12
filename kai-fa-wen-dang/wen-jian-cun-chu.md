# 文件存储

虽然Fireboom专注于结构化数据的处理，但文件/ Blob在构建应用程序中也占据着重要地位。

我们深入研究了各种解决方案，以方便用户上传图片、PDF文件等。但大多数解决方案需要在客户端、服务器端或两者都进行额外工作。

我们遵循“开发体验优先”的理念，添加了对S3兼容的支持。任意S3兼容的供应商都能轻松接入飞布。

配置S3供应商后，飞布将自动注册上传路由，飞布应用程序的用户可以轻松上传文件。

## 支持OIDC Provider

目前主流S3供应商如下：阿里云OSS、腾讯云COS、MINIO。

环境准备请参考：

[#wen-jian-cun-chu-s3](../huan-jing-zhun-bei.md#wen-jian-cun-chu-s3 "mention")

## 快速操作

### 基本配置

1. 在文件存储面板中点击“+”，进入S3新建页
2. 前往各S3 Provider的文档页，查看如何获取参数
3. 输入供应商名称及其他参数
4. 点击测试，若测试通过，点击保存，进入详情页
5. 点击顶部菜单栏的“<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAMAAAC7IEhfAAAAY1BMVEUAAADU1NRjZmxvcnePkZVgY2rPz9BfY2poaHTAwMOAhIxoa3FgYmpgYmlgY2nFxcbU1NRmZm+Ag427u73U1NRfYml/g4zt7e3k5eXW1te9vsCztLefoaWanKCOkJVydXtucXecDQKGAAAAFXRSTlMAzP336NDOiAvTz/rn2tjSph7Qs6d9epWLAAAAjElEQVQ4y+2T2Q6EIAxFK+A6mzMj4q7//5VaYngCG2N8cDkvNOlJSG9TuCq+XMQ3oiQ4p0jGsx+/fCIByDwrqRFzDYDn4BatYiw4Y1zEhBgIJjUsjJbED5eG19ctBtrr66rD9x05RYH9oVBKtViFTvGB7UZNlFg9N4n01/QwdDwrA0/mU0jtK/zDYRgBwgsrsPomQg4AAAAASUVORK5CYII=" alt="预览" data-size="line">”，前往API预览页
6. 在预览页顶部，选择OIDC供应商，点击前往登录，登录后可查看用户信息
7. 在FileUpload中选择当前上传路由，选择文件，并上传

需要注意的是：如果设置了profile，s3上传时就必须要设置profile

### 高级配置

1. 在文件存储面板中，选择列表项，右击点击“新建Profile”
2. 基本设置：
   * 最大尺寸：设置上传文件的最大尺寸
   * 匿名上传：默认情况下，需要身份验证才能上传文件。这里开启后，可支持匿名上传。
   * 文件数量：每次上传的最大文件数量
   * 文件类型：文件的媒体类型
   * 文件后缀：文件后缀
   * 附加信息：S3规范中的字段，用来描述本次上传的数据。
     1. 管理和追踪对象：元数据可以用来标识对象的创建日期、最后修改日期、对象大小、存储类别、对象所有者等信息，以方便管理员追踪和管理对象。
     2. 控制访问权限：元数据可以用来定义S3对象的访问权限，例如设置ACL（Access Control List）或Bucket策略，以决定哪些用户或群组可以读取或写入对象。
     3. 优化数据检索：元数据可以帮助加速数据的检索，例如在存储大量视频文件时，可以将元数据设置为视频的格式、码率、分辨率等信息，以便客户端快速定位和获取所需的视频文件。
3. 前置钩子：将在上传到OSS之前运行，允许执行任何必需的验证以及定义存储文件的路径。
4. 后置钩子：在上传完成或失败后运行，用来进行后置处理

### 文件管理

1. 点击文件存储列表的选项，进入文件管理页
2. 选择文件列表的文件，可对查看详情，或删除文件
3. 点击顶部的“上传”按钮，可上传文件
4. 也可点击顶部“文件夹”按钮，新建文件夹
5. 在顶部的搜索框，可搜索文件或文件夹

## 工作原理

飞布自动将S3供应商注册到特定端点中，用户可直接通过该端点上传文件。

具体规则，请前往API构建->API规范的[#wen-jian-shang-chuan](api-gou-jian/api-gui-fan.md#wen-jian-shang-chuan "mention") 查看。

若要使用高级功能，需要先配置Profile项。从文件上传生命周期角度来分析，配置项主要分为3个触点：上传校验、前置钩子以及后置钩子。

* 上传校验：文件上传前的一些校验，例如文件大小、文件类型、文件数量等
* 前置钩子：文件通过校验，但未上传到OSS前的时刻
* 后置钩子：文件上传到OSS后的时刻

## 客户端使用

使用文件上传，只需要按照文件上传的接口规范，构建一个POST文件上传请求，即可将文件上传至S3 bucket中。这适用于任何可以使用HTTP请求发送FormData的环境。

具体使用，请参考：

[#wen-jian-shang-chuan](sdk-sheng-cheng/wei-xin-xiao-cheng-xu-sdk.md#wen-jian-shang-chuan "mention")



