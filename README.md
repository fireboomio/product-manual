# 序言

> 欢迎阅读 飞布 的技术产品文档。本文档旨在为您提供对我们产品背后技术及其功能的全面了解。我们的目标是为您提供充分利用 飞布 的所有功能和特性所需的知识。

从双创时代初期开始的创业之路，一程风雨洗练，让我们积累了技术经验和久经考验的团队，也让我们明白一款提升开发效率，降低开发成本的工具对开发者而言，有着怎样的意义，更让我们决心把这款理想中的工具照进现实。从实践经验和行业风向中，我们看到了aPaaS背后的无限潜力。我们相信，一款足够灵活，足够开放的aPaaS平台，将让目前开发者面临的许多难题迎刃而解，让新手开发者和资深开发者都能从中受益。

于是，怀着为开发者服务的信念，带着对低代码前景的信心，我们投入了飞布的开发。项目从2022年初启动，历经一年的打磨，为中国本土开发者量身打造的可视化API开发平台——飞布API 终于和开发者见面了。无数次头脑风暴，无数次深夜debug，无数次反复打磨，飞布的每一处细节，都凝聚着整个团队的智慧和心血。

飞布的LOGO是一团小小的火苗。有道是众人拾柴火焰高，全体成员的共同努力，让我们有足够的信心，将飞布带到各位开发者面前。希望更多优秀的开发者通过体验和反馈，参与到飞布产品的功能评估和性能完善中来，一起让飞布这朵小火苗越烧越旺。

{% embed url="https://www.bilibili.com/video/BV148411T7WY" %}

## 为什么要做飞布？

在研发过程中，不同角色有不同的困扰：

如果你是后端开发者，想做有趣的产品，却只能在日复一日的增查改删中蹉跎了岁月？

如果你是前端开发者，想成为全栈开发者，却苦于后端技术门槛高、上手难？

如果你是创业团队，想用最快的速度、最低的成本做出产品，却只能眼睁睁看着本就拮据的资源在毫无进展的日子里消耗殆尽？

这些问题，我们都曾经历过。项目发起人曾在本科期间尝试创业，在产品研发上花费了大量心血，但产品还没上线，就因现金流不足而失败；读研期间，凭借积累的技术经验和团队，拉来原班人马，成立了外包团队，服务了众多创业者，却发现做的大多数项目都是增删改查。而外包企业的主要成本就是人力成本。

一言以蔽之，公司需要降本增效，开发者不想做CRUD Boy。

因此，我们非常了解一款能够提高开发者的开发效率和开发体验，节约公司的人力成本和时间成本的工具意味着什么。

## 飞布不是什么？

提到飞布，就不得不提提低代码平台。广义上将，飞布也属于低代码。尽管飞布具有低代码开发平台的所有优势，如可视化、低成本、高效率、稳定性等，但飞布与市面上的低代码平台完全不同。

因为，用飞布可以构建低代码，但，用低代码构建不出飞布！

目前国内市场上所有低代码平台均从前端切入，通过可视化拖拽构建应用，主要解决前端页面复现和数据绑定的问题。

尽管这些工具提供了简单的数据处理机制，如连接数据源写SQL等，但该方式不够灵活，遇到稍微复杂的逻辑就无法处理，只适用于中后台系统的开发。

此外，这些产品无法实现APP或小程序等移动端产品的开发，应用场景受限，对于独立开发者和外包企业而言无疑是一大硬伤：开发人员需要的是趁手的工具，而不是看似一劳永逸，实则处处掣肘的鸡肋。

我更倾向于称这些低代码平台为“前端低代码”。

目前国内暂时还没有像飞布这样专注API开发的平台，但在国外的低代码平台起步较早，积累了一批较为成功的开发者工具，比如Hasura、Supabase、Nhost等等。其中Supabase是一家比较年轻的企业，近两年的发展势头也非常可观，而且它也是在充分汲取Firebase、Hasura等产品的成功经验的基础上发展起来的，它的经验对即将问世的飞布而言参考价值也更高。

但对于国内的开发者而言，使用国外低代码平台存在语言不通、支付方式不同、数据安全问题等诸多不便。

而且，国外开发者的逻辑和需求与国内开发者也有差异，国外开发者用起来得心应手的工具，对国内开发者而言却可能只是另一种口味的鸡肋。

因此，受海外Hasura、Supabase、Firebase等产品的启发，我们决心打造适合中国开发者的低代码平台。怀揣着“<mark style="color:orange;">极致开发体验，飞速布署应用</mark>”的愿景，飞布项目于2022年初开始启动，历经一年打磨，于2023年2月10日发布内测版本。

## 飞布是什么？

飞布是中国版Hasura。它从后端切入，专注API开发，是开发体验优先的可视化API开发平台，前后端开发者都能使用飞布构建生产级WEB API，从而让<mark style="color:orange;">”前端变全栈，后端不搬砖“</mark>。

详情见[README (1).md](<README (1).md> "mention")

{% embed url="https://www.bilibili.com/video/BV1w24y1U7fx" %}
入门课程导读
{% endembed %}

## 联系我们

无论您是开发人员还是技术团队管理人员，本文档都将为您提供所需的信息，以有效地使用和集成飞布到您的工作流程中。文档涵盖了产品的所有方面，从其架构和设计，到其功能和特性，再到如何自定义和与其他系统集成。

我们相信本文档中包含的信息将对您的飞布之旅极为宝贵。如果您有任何问题或疑虑，请随时与我们的支持团队联系。您也可以前往[飞布官网](https://www.fireboom.io/)，获取更多信息。

![](<.gitbook/assets/image (12) (1).png>)

谢谢您选择飞布。我们很高兴成为您技术之旅的一部分。

最诚挚的问候，飞布团队。
