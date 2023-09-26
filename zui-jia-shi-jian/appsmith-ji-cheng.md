# appsmith集成

Appsmith 是一个开源的前端低代码平台，帮助快速构建内部工具。例如，仪表盘、工具箱、CMS以及其它能帮团队完成特定任务的工具

### appsmith特性

* 拖拽构建UI：像设计PPT一样，在网格画布上拖拽预制组件构建UI，无需掌握Html+CSS
* 简化前后端集成：支持多种数据源，快速为UI绑定数据
* JS扩展逻辑：支持用JS处理数据，定制复杂工作流

### 集成fireboom

appsmith 从前端切入，擅长界面构建，利用 sql 能够处理简单数据。但对于复杂业务逻辑，需要依赖API处理。**而API构建是Fireboom的强项，因此两者结合将进一步扩展appsmith的适用范围。**不仅适用内部系统开发，还能扩展到更多领域。

#### appsmith引入js库

appsmith支持[引入UMD格式的js库](https://docs.appsmith.com/core-concepts/writing-code/ext-libraries)，导入方式如下：

1. 点击Libraries->新建 icon
2. 输入js库url，点击 install
3. 使用库

<figure><img src="../.gitbook/assets/image (25).png" alt=""><figcaption><p>appsmith引入js库</p></figcaption></figure>

#### Fireboom 生成JS-SDK

Fireboom支持生成UMD格式的客户端SDK—— [js-sdk.md](../shi-yong-bu-shu-shang-xian/sdk-sheng-cheng/js-sdk.md "mention")。

1. 浏览模板市场，下载JavaScript client
2. 打开`js-client`，文件将生成到`./generated-sdk/js-client`
3. 查看js，访问：[http://localhost:9123/generated-sdk/js-client/index.umd.min.js](http://localhost:9123/generated-sdk/js-client/index.umd.min.js)
4. 将上述链接填入appsmith的Add Js Libraries-> Library URL中，安装即可

<figure><img src="../.gitbook/assets/image (41).png" alt=""><figcaption></figcaption></figure>

#### appsmith中使用Js-SDK

appsmith中有如下方式使用sdk，如：

**1，js object中使用**

```javascript
export default {
	async create(){
		let res=await FBClient.mutation.CreateT({input:{name:"sss",des:"des"}})
		console.log(res)
	},
	async get(){
		let res=await FBClient.query.GetT({input:{id:1}})
		console.log(res)
	}
}
```

**2，在组件中直接调用**

<figure><img src="../.gitbook/assets/image (57).png" alt=""><figcaption></figcaption></figure>

FBClient更多用法，请参考 [js-sdk.md](../shi-yong-bu-shu-shang-xian/sdk-sheng-cheng/js-sdk.md "mention")

### 总结

使用该方法可以将fireboom和appsmith集成在一起，充分发挥两者各自的优势。未来，希望将fireboom作为数据源直接集成到appsmith中，进一步提升开发体验。



参考：

* [appsmith简明教程](https://bar9vnf09af.feishu.cn/mindnotes/Bi5nbgcZtmACZfnAMN7c796lnfd)
