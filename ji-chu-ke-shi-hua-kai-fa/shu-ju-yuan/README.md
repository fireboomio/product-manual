# 数据源

飞布支持和整合了多种数据库和第三方 API（GraphQL/REST），提供更加灵活的数据访问方式。

本文档介绍如何在飞布中新建数据源。

![多数据源支持](https://www.fireboom.io/images/gif/01-01%E5%A4%9A%E6%95%B0%E6%8D%AE%E6%BA%90%E6%94%AF%E6%8C%81.gif)

## 支持数据源

飞布当前支持API、数据库、消息队列、[自定义](../../jin-jie-gou-zi-ji-zhi/graphql-gou-zi.md) 四类数据源类型，详情见下文。

每个数据源都需要设置不同的“**名称**”，用作命名空间。

<mark style="color:red;">若想进一步了解底层原理，请学习</mark> [chao-tu.md](../../he-xin-gai-nian/chao-tu.md "mention") <mark style="color:red;">！！！</mark>

## 新建数据源

{% embed url="https://www.bilibili.com/video/BV1fL411C72D" %}
04功能介绍 飞布如何新建数据源？
{% endembed %}

新建数据源步骤如下：

1.进入飞布主页面，点击左侧边栏“数据源”后的“+”号按钮

2.选择数据源类型

3.填写数据源配置信息，详情参阅“连接数据库”和“连接API”文档

4.点击“测试”按钮，检查当前配置信息是否正确

5.点击“保存”按钮

6.开启当前数据源

## 示例数据源

为方便开发者快速上手，飞布还内置了示例数据源，只需点击即可添加。

示例数据源托管在公开仓库，可前往[这里](https://github.com/fireboomio/files/blob/main/datasource.example.json)查看。
