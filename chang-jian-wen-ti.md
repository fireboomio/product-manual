# 常见问题



{% embed url="https://github.com/fireboomio/product-manual/discussions/categories/q-a" %}
飞布QA汇总
{% endembed %}

SDK显示不了列表？SDK不能下载？首页通知无法展示？ 我们使用 `Github` 作为我们的数据仓库，如果你无法访问 Github，那么 `Fireboom` 中的一些功能可能受到影响，我们建议你优先处理 `Github` 访问的网络问题。或者使用下面的方法来尝试解决（推荐方案1）

1. 参考 https://github1s.com/521xueweihan/GitHub520 中的方法修改本地host文件
2. 参考 https://doc.fastgit.org/zh-cn/node.html 中的反代地址，在 `Fireboom` 中添加环境变量（“设置”->"环境变量"）GITHUB\_PROXY\_URL=hub.fastgit.xyz GITHUB\_RAW\_PROXY\_URL=raw.fastgit.org 如果反代后仍无法使用，请检查反代服务是否正常，如果异常请另外寻找其它反代服务



faq添加一个 数据库用户名密码含特殊字符导致无法连接成功？参考 https://developer.mozilla.org/en-US/docs/Glossary/percent-encoding 将特殊字符编码后再次尝试

