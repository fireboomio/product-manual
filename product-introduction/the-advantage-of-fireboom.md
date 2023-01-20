---
description: 相比于传统开发、前端低代码开发，飞布有哪些优势
---

# 飞布的优势

本文主要阐述飞布的核心优势有哪些，产品的能力如何。

<figure><img src="../.gitbook/assets/image (17).png" alt=""><figcaption><p>飞布同时满足易用性和开放性</p></figcaption></figure>

## 相对于传统开发

### 飞布学习成本低

采用传统方式开发API，需要掌握一门后端开发语言，而这通常需要1-2年的专业训练。使用飞布，开发者只需要30分钟的简单学习，就能上手构建出生产级API。

传统开发框架都基于特定开发语言实现，开发者必须选择自己技术栈范围内的开发框架，例如只会php的开发者就不能选用java生态的积木框架。而飞布底层基于Graphql，这本身是一种协议，类似json或sql，开发者很容易就能掌握其规则。此外，飞布通过可视化界面封装了Graphql细节，即使不会graphql也能构建满足绝大多数场景的API。

### 飞布开发效率高

为了提高研发效率，目前市场上有很多API开发框架。这些框架极大的提升了生产力，例如基于PHP语言的fastadmin框架，相对于ThinkPHP效率提升了50%。但对飞布而言，这仍然不够。飞布可以在2分钟内，完成用框架2小时才能完成的开发工作。

### 飞布多语言兼容

为了实现和传统开发一样的“PRO CODE”能力，飞布实现了基于HTTP协议的钩子机制，开发者可以用任意喜欢的后端编程语言，如Java、Golang、PHP、Node.js等，扩展API，实现自定义业务逻辑。目前飞布已经官方支持了Node.js SDK，Golang SDK正在开发中。

<figure><img src="../.gitbook/assets/image (1).png" alt=""><figcaption><p>飞布在线IDE</p></figcaption></figure>

### 飞布提升了前后端协作效率

传统开发过程中，前后端协作往往需要通过文档进行，但文档往往无法及时维护，联调时接连翻车，导致前后端开发者之间产生矛盾，更甚者大打出手。此外，研发团队中往往后端强势，前端为了适配后端接口，有时不得不使出各种“黑魔法”，降低了代码可维护性，延长了工期。最后，即使有一致的swagger文档，前端仍要根据文档，用最传统的方式对接接口，该过程往往没有语法提醒，效率较低。一旦接口字段变更，又需要人工调整，繁琐低效。

依赖于graphql强大的类型系统，飞布可实时将API生成swagger文档和各客户端语言的SDK，如ts、flutter等。后端无需写文档，更不用手工维持一致性；前端直接调用对应SDK，并根据语法提醒，对接接口。一旦接口变更，SDK也实时更新，对于强类型语言的SDK，语法报错将帮助前端开发者定位到变更处。

<figure><img src="../.gitbook/assets/image (9).png" alt=""><figcaption><p>API文档和SDK</p></figcaption></figure>

### 飞布更能响应需求变更

很多优秀的传统开发框架都内置了根据表单配置快速生成代码的功能。在项目初始化时，代码生成可以发挥很大功效，提升研发效率。但遇到需求变更后，例如单表需要拆分为主子表，该功能就无法工作了。因为，开发者要在生成的代码中实现自定义逻辑，这时重新生成代码就可能覆盖原有逻辑，风险极大。最好的方式是前往代码，找到所有变更的部分，手工修改。

依赖于graphql的内省特性，飞布可自动内省数据库表结构，检测并提醒哪些API非法。开发者根据界面提醒，稍微删减字段，就可实现功能变更。

<figure><img src="../.gitbook/assets/image (6).png" alt=""><figcaption><p>非法API检测</p></figcaption></figure>

### 飞布开发体验更好

传统开发的所有功能都需要用编码的方式实现，需要深入了解一门编程语言。而飞布提供了可视化的界面，封装了技术细节，提供了友好的交互体验，开发者只需要勾选就能生成API。同时，飞布集成了日志面板和问题面板，可实时查看系统日志和错误信息。此外，飞布还内置了数据建模、数据预览、ER图等功能，开发者无需打开其他工具，如navicat，就能实现全流程的API开发工作。

飞布还内置了webcontainer，nodejs开发者，无需在本机安装node环境，即可进行钩子开发。

<figure><img src="../.gitbook/assets/image (14).png" alt=""><figcaption><p>飞布核心功能</p></figcaption></figure>

## 相对于前端低代码

### 飞布适用场景更广

前端低代码平台大多针对中后台开发场景，无法适用移动端开发场景，如小程序开发、APP开发等。飞布不仅能完成中后台开发，而且适用于移动端、BI大屏，甚至WEB3等。

### 飞布灵活性更高

前端低代码平台大多基于表单驱动，无法完成复杂需求，例如创建文章后发送一封邮件、从数据库查询设备列表时同时获取设备在线状态、更新点赞数量后，实时广播给所有订阅的客户端等。依赖于graphql的指令系统，飞布实现了复杂的数据编排能力，例如跨数据源关联，而前端低代码平台不具备该能力。飞布的钩子机制，让飞布具备了“PRO CODE”能力，能够实现任意复杂场景的case。飞布的服务端订阅能力，让飞布天生适用于实时或准实时通讯场景。

## 相对于后端低代码

飞布本身属于“后端低代码”，但国内并无类似产品。目前国外类似产品也较少，其中hasura为佼佼者。

<figure><img src="../.gitbook/assets/image (18).png" alt=""><figcaption><p>飞布VS hasura</p></figcaption></figure>

### 飞布更适用于本土开发者

Hasura主要市场在印度，产品界面和文档以英文为主，主要兼容pgsql，不兼容国内常用的mysql，对外暴露API为graphql端点，而不是国内常用的REST API。而飞布拥有中文社区，兼容多个数据源，包括mysql，pgsql、sqlite等，同时尽管飞布以graphql为核心，但考虑到国内开发者的习惯，对外只暴露REST API。

### 飞布权限系统更灵活

Hasura基于数据表实现权限体系，能够实现API控制和数据控制。该方案对数据库特别适用，但不适用第三方API。飞布充分利用了graphql的指令系统，设计了一套同时兼容数据库和第三方API的权限系统，包括基于RBAC的接口权限和基于OIDC的数据权限。

### 飞布指令系统更强大

Hasura采用了直接暴露graphql端点的方式对外提供服务，开发者必须在客户端编写operation，出于安全的目的该方式无法利用graphql的指令系统。而飞布将graphql作为中间层，开发者编写的operation放在了服务端，无需担心安全，可利用graphql的指令系统，实现任意功能的逻辑。例如：参数校验指令`jsonSchema`、当前时间注入指令`injectCurrentDateTime`、响应转换指令`transform`等。&#x20;

<figure><img src="../.gitbook/assets/image (11).png" alt=""><figcaption><p>指令枚举</p></figcaption></figure>

### 飞布支持内省REST API

Hasura支持REST API，但使用体验很差，用户必须手工编写graphql类型定义和operation函数签名。飞布不仅支持自动内省数据库和graphql API，还基于OAS规范，自动内省REST API，提升了接入REST API的效率，降低了学习成本。

<figure><img src="../.gitbook/assets/image.png" alt=""><figcaption><p>REST API配置</p></figcaption></figure>

