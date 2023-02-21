# 文件存储

飞布内置存储机制，任意S3兼容的供应商都能轻松接入飞布。配置S3供应商后，飞布将自动注册上传路由，飞布应用程序的用户可以轻松上传文件。

## 支持OIDC Provider

目前主流S3供应商如下：阿里云OSS、腾讯云COS、MINIO。

## 快速操作

### 基本配置

1. 在文件存储面板中点击“+”，进入S3新建页
2. 前往各S3 Provider的文档页，查看如何获取参数
3. 输入供应商名称及其他参数
4. 点击测试，若测试通过，点击保存，进入详情页
5. 点击顶部菜单栏的“<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAMAAAC7IEhfAAAAY1BMVEUAAADU1NRjZmxvcnePkZVgY2rPz9BfY2poaHTAwMOAhIxoa3FgYmpgYmlgY2nFxcbU1NRmZm+Ag427u73U1NRfYml/g4zt7e3k5eXW1te9vsCztLefoaWanKCOkJVydXtucXecDQKGAAAAFXRSTlMAzP336NDOiAvTz/rn2tjSph7Qs6d9epWLAAAAjElEQVQ4y+2T2Q6EIAxFK+A6mzMj4q7//5VaYngCG2N8cDkvNOlJSG9TuCq+XMQ3oiQ4p0jGsx+/fCIByDwrqRFzDYDn4BatYiw4Y1zEhBgIJjUsjJbED5eG19ctBtrr66rD9x05RYH9oVBKtViFTvGB7UZNlFg9N4n01/QwdDwrA0/mU0jtK/zDYRgBwgsrsPomQg4AAAAASUVORK5CYII=" alt="预览" data-size="line">”，前往API预览页
6. 在预览页顶部，选择OIDC供应商，点击前往登录，登录后可查看用户信息
7. 在FileUpload中选择当前上传路由，选择文件，并上传

### 文件管理

1. 点击文件存储列表的选项，进入文件管理页
2. 选择文件列表的文件，可对查看详情，或删除文件
3. 点击顶部的“上传”按钮，可上传文件
4. 也可点击顶部“文件夹”按钮，新建文件夹
5. 在顶部的搜索框，可搜索文件或文件夹

## 工作原理

飞布自动将S3供应商注册到特定端点中，用户可直接通过该端点上传文件。但，上传前需要登录。

具体规则，请前往API构建->API规范的[#wen-jian-shang-chuan](api-gou-jian/api-gui-fan.md#wen-jian-shang-chuan "mention") 查看。

## 客户端使用

todo
