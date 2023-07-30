---
description: 从开发到生产的部署迁移流程
---

# 发布到生产

在实际的开发过程中，我们一般会将开发流程拆分为多个迭代，每个迭代完成不同的功能，本章将介绍如何在多轮迭代下从开发到生产的迁移过程。

### 第1个迭代

第一个迭代我们做最小化需求，我们定义了一个简单的数据结构

```prisma
model Todo {
  id        Int      @id @default(autoincrement())
  title     String
  completed Boolean  @default(false)
  createdAt DateTime
}
```

然后使用fb的批量创建，创建出我们的增删改查 API 。

然后这个迭代就完成了，我们参考下一节的手动部署给部署到线上服务器，使用`pm2 start ./fireboom -- start` 启动 Fireboom 服务，于是就得到一个线上 API 地址 `http://your.comany.com:9991` 。

业务上线后我们使用了一段时间，数据库中已存在了一部分数据，我们希望后续迭代过程中仍然保留这些数据。

### 第2个迭代

之前迭代1的时候我们忘记添加了更新时间和删除时间

