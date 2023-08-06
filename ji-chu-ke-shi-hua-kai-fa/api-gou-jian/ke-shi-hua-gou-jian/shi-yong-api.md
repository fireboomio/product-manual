# 使用API

学习了如何构建API后，接下来我们学习如何测试和使用API。

## API预览

点击右上角"API预览ICON"，打开swagger文档页。

<figure><img src="../../../.gitbook/assets/image (1) (5).png" alt=""><figcaption><p>预览页</p></figcaption></figure>

左侧为API列表，以文件夹作为分组名，未分组API位于Others分组，此外，还包含FileUpload分组，用于展示OSS对应的路由。

选择对应API，可以查看其Operation，输入入参后，点击“TRY”按钮，可进行测试。响应栏会展示响应状态码和响应结果。

若API需要授权才能访问，需要在右上角选择对应OIDC，进行登录。登录后可查看当前登录用户的基本信息。

此外，若想在POSTMAN等工具测试，点击“下载”按钮，可获取swagger文档。

## SDK生成

若想在客户端使用API，可点击状态栏“SDK模板”，下载对应模板，并配置生成路径，即可在项目中引用对应语言的客户端SDK。

更多详情，前往[sdk-sheng-cheng](../../../shi-yong-bu-shu-shang-xian/sdk-sheng-cheng/ "mention")。
