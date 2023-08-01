# 服务端Operation

## HTTP请求流程

接着，我们复盘下上文提到的3中类型时序图。他们都能用该HTTP请求流程图表示。

<figure><img src="../.gitbook/assets/image (11).png" alt=""><figcaption></figcaption></figure>

即，客户端发起REST请求，飞布服务查找并解析对应OPERATION，根据指令调用不同的拦截器，根据命名空间拆分请求到不同数据源，或控制请求不同数据源的流程，如串行或并行。

并将请求结果组装后，以JSON的方式返回给客户端。

一般情况下，基于GraphQL的服务直接对外暴露gql 端点，operation一般写在客户端，而飞布反其道而行，把operation放到服务端。

这样做有如下好处：

1. 前端开发者无需感知GraphQL的存在，因此无需任何学习成本。
2. 飞布对外暴露REST API，可以复用HTTP基础设施，如CDN等。
3. OPERATION保存在服务端，攻击者无法触达OPERATION，保证了安全。
4. 无论客户端还是服务端OPERATION，都能利用GraphQL按需取用、类型系统的优势。

## GraphQL指令注解

GraphQL的指令系统实现更加复杂的业务逻辑，但如果将指令放到客户端，就无法保证安全。得益于服务端OPERATION的第三点优势，Fireboom充分发挥了GraphQL指令系统的优势。

<figure><img src="../.gitbook/assets/image (12).png" alt=""><figcaption></figcaption></figure>

飞布的指令面板用于可视化插入指令，除了跨源关联外，指令分为三大类：全局指令（包括角色和私有），入参指令（用于修饰OPERATION入参）、出参指令（用于修饰operation字段）。

* 全局指令：作用于Operation，用于定义整个operation的行为，例如设置权限或声明其为内部操作。
* 入参指令：作用于 入参变量，用于为参数设置默认值，默认值可能源于当前登录用户或固定函数（环境变量）
* 字段指令：作用于 输出字段，用于修改字段的行为，如transform拍平响应结构
* 跨源关联指令：用于实现多API的流程编排，详情见 [kua-yuan-guan-lian.md](../ji-chu-ke-shi-hua-kai-fa/api-gou-jian/kua-yuan-guan-lian.md "mention")

在飞布中，指令本质上用于实现HTTP的切面功能，与JAVA等的注解类似。
