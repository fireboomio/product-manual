---
description: 待完善
---

# 数据建模

数据建模用来进行数据库设计，具有完备的数据表设计能力，类似[Navicat](https://navicat.com.cn/products#navicat)数据库开发工具。它本质上是对Prisma Schema的可视化封装。

## 控制台打开

<figure><img src="../../../.gitbook/assets/image (20).png" alt=""><figcaption><p>数据建模</p></figcaption></figure>

1，切换到“数据建模”页签

2，选择数据源，例如`todo`数据库

3，点击“![](<../../../.gitbook/assets/image (2).png>)”切换到数据建模功能页

## 模型设计

数据建模支持两种模式：普通视图和源码视图，分别适用于新手开发者和熟悉Prisma的开发者。点击右上角的图标![](<../../../.gitbook/assets/image (13).png>)和![](<../../../.gitbook/assets/image (21).png>)，可切换两种视图。

### 普通视图

#### 新增表

1. 点击数据建模右侧的“+”，双击输入表名
2. 点击右侧面板顶部的“+”，输入字段名称和类型
3. 点击字段行后的“?”，设置字段为数组或是否为空
4. 点击字段行后的“@”，为字段增加描述
5. 点击顶部的“@”，为表增加属性，前往[参考](https://www.prisma.io/docs/concepts/components/prisma-schema/data-model#defining-attributes)
6. 点击顶部的“迁移”按钮，保存修改

> 迁移本质上调用了 prisma db pull 命令，该方式不可用于<mark style="color:orange;">生产环境</mark>，详情查看[prisma文档](https://www.prisma.io/docs/concepts/components/prisma-migrate/db-push)。
>
> <img src="https://website-v9.vercel.app/illustrations/home-page/hasslefree-migrations.svg" alt="" data-size="original">

#### 删除表

有两个方式可以删除表，一是选中表后，点击顶部的“删除”按钮，二是在左侧列表右击，点击“删除”。

### 源码视图

源码视图展示prisma schema源文件，同时支持语法提醒和高亮展示，你可以用它实现任意形式的数据建模。

<figure><img src="../../../.gitbook/assets/image (8).png" alt=""><figcaption><p>源码视图</p></figcaption></figure>







