# S3配置及使用

首先，我们学习如何配置S3，这里以腾讯云为例。

{% embed url="https://www.bilibili.com/video/BV1hk4y1Y7jM/?share_source=copy_web&vd_source=f8709d15baaa835ea2d0bb3bcc6857da" %}
16功能介绍-飞布如何实现文件上传？
{% endembed %}

### S3配置

S3配置主要有：

<table data-header-hidden><thead><tr><th width="129">字段</th><th width="367">值</th><th>备注</th></tr></thead><tbody><tr><td>名称</td><td>tengxunyun</td><td>文件存储的名称</td></tr><tr><td>服务地址</td><td>cos.ap-nanjing.myqcloud.com</td><td>S3服务访问地址</td></tr><tr><td>APP ID</td><td>AKID0zoR4VmsnWFsIVVIsFPM6htvlPo0yw43</td><td></td></tr><tr><td>APP Secret</td><td>********************************</td><td>可以用环境变量存储</td></tr><tr><td>区域</td><td>ap-nanjing</td><td></td></tr><tr><td>桶名称</td><td>test-1314985928</td><td>bucket名称</td></tr><tr><td>开启SSL</td><td>是</td><td>开启后用HTTPS访问</td></tr></tbody></table>

使用步骤：

1. 在文件存储面板中点击“+”，进入S3新建页
2. 前往各S3 Provider的文档页，查看如何获取参数（前往查看[教程](../../huan-jing-zhun-bei/wen-jian-cun-chu-s3.md#teng-xun-yun)）
3. 输入供应商名称及其他参数
4. 点击测试，若测试通过，点击保存，进入详情页

### S3使用

配置S3供应商后，飞布将注册上传路由，路由规则为：

```
http://localhost:9991/S3/文件存储名称/upload?directory=xxx
// 文件存储名称，不是存储桶名称。
```

用户可通过该路由，上传文件至指定目录，目录由directory字段指定。需要注意的是，使用该路由上传文件时，必须要登录。

<figure><img src="../../.gitbook/assets/image (13) (4).png" alt=""><figcaption></figcaption></figure>

上传文件后，返回文件的相对地址，为：

<pre class="language-json"><code class="lang-json"><strong>[
</strong>    {
        "fileKey":"directory/hash.jpg"
        //目录+文件的hash值（xxHash是一种非常快速的哈希算法）
    }
]
</code></pre>

有两种方式访问上述文件：

* 标准方式：https://桶名称.服务地址/fileKey，例如：<mark style="color:orange;">https://</mark><mark style="color:blue;">test-1314985928</mark>.<mark style="color:red;">cos.ap-nanjing.myqcloud.com</mark><mark style="color:purple;">/aaaa/logotest.png</mark>
* 其他方式：https://服务地址/桶名称/fileKey，例如：<mark style="color:orange;">https://</mark><mark style="color:red;">cos.ap-nanjing.myqcloud.com</mark>/<mark style="color:blue;">test-1314985928</mark><mark style="color:purple;">/aaaa/logotest.png</mark>

对于私有桶，还需要追加临时签名才能访问：

```
https://桶名称.服务地址/fileKey
?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=LTAI5tGAiNbpnDb4ZghQ7MaG%2F20230720%2Foss-cn-beijing%2Fs3%2Faws4_request&X-Amz-Date=20230720T072030Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&response-content-disposition=attachment%3B%20filename%3D%22%22&X-Amz-Signature=9c09d9ba4b03180bfc524ff986b3ed8d1b785f665c4fc1e8bf99ef96160568a2
```

查看临时签名的生成方式，请前往这里\~

测试步骤：

1. 点击顶部菜单栏的“<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAMAAAC7IEhfAAAAY1BMVEUAAADU1NRjZmxvcnePkZVgY2rPz9BfY2poaHTAwMOAhIxoa3FgYmpgYmlgY2nFxcbU1NRmZm+Ag427u73U1NRfYml/g4zt7e3k5eXW1te9vsCztLefoaWanKCOkJVydXtucXecDQKGAAAAFXRSTlMAzP336NDOiAvTz/rn2tjSph7Qs6d9epWLAAAAjElEQVQ4y+2T2Q6EIAxFK+A6mzMj4q7//5VaYngCG2N8cDkvNOlJSG9TuCq+XMQ3oiQ4p0jGsx+/fCIByDwrqRFzDYDn4BatYiw4Y1zEhBgIJjUsjJbED5eG19ctBtrr66rD9x05RYH9oVBKtViFTvGB7UZNlFg9N4n01/QwdDwrA0/mU0jtK/zDYRgBwgsrsPomQg4AAAAASUVORK5CYII=" alt="预览" data-size="line">”，前往API预览页
2. 在FileUpload中选择上传路由，设置上传目录directory的值，选择文件，点击”TRY“，返回数组，fileKey
   1. 若上传时返回401错误，请登录后重试
   2. 在预览页顶部，选择OIDC供应商，点击前往登录
3. 拼接目录，访问文件



~~需要注意的是：如果设置了profile，s3上传时就必须要设置profile~~

