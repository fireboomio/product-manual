# 高级配置：profile

飞布的文件上传还支持高级配置。主要实现3个功能：匿名登录、文件限制以及文件钩子。

* 匿名登录很简单，只需要开启即可。后续上传文件，无需登录。
* 文件限制，支持最大文件尺寸，文件类型、文件后缀等的限制。需要注意到的是，文件数量 需要由客户端限制，该参数用于生成客户端SDK时，限制数量。
* 文件钩子，允许用户在飞布上传文件到OSS之前和之后编写自定义逻辑，修改上传行为。

相关配置：

<table><thead><tr><th width="153">名称</th><th width="188.33333333333331">值</th><th>备注</th></tr></thead><tbody><tr><td>匿名登录</td><td>否</td><td>默认情况下，需要身份验证才能上传文件。这里开启后，可支持匿名上传。</td></tr><tr><td>最大尺寸</td><td>10 M</td><td>设置上传文件的最大尺寸</td></tr><tr><td> 文件数量限制</td><td>1</td><td>生成客户端SDK时会用到</td></tr><tr><td>文件类型</td><td>image/*,video/*</td><td>文件的媒体类型</td></tr><tr><td>文件后缀</td><td>gif,png</td><td>文件后缀</td></tr><tr><td>META</td><td>{ "$schema": "http://json-schema.org/draft-07/schema#", "type": "object", "properties": { "postId": { "type": "string" } }, "required": [ "postId" ] }</td><td>文件的元数据，和钩子配置使用，值为JSON SCHEMA</td></tr><tr><td>前置钩子</td><td>否</td><td>将在上传到OSS之前运行，允许执行任何必需的验证以及定义存储文件的路径。</td></tr><tr><td>后置钩子</td><td>否</td><td>在上传完成或失败后运行，用来进行后置处理</td></tr></tbody></table>

操作步骤：

1. 在文件存储面板中，选择列表项，右击点击“新建Profile”
2. 重命名Profile名称，并填写基本设置（meta和钩子后续讲解）

<figure><img src="../../.gitbook/assets/image (4) (1) (2).png" alt=""><figcaption></figcaption></figure>

使用profile，需要在上传文件的请求头中增加X-Upload-Profile字段，该字段为枚举类型，从profile中选择。

{% hint style="info" %}
X-Upload-Profile不设置时，默认走普通上传。
{% endhint %}
