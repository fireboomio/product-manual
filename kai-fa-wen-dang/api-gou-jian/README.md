# API构建

API构建是飞布的核心功能，本章将重点介绍如何使用飞布快速构建出生产级REST API。



## 原理简述



飞布以 API 为中心，将所有数据抽象为 API，包括 REST API，GraphQL API ，数据库甚至消息队列等，通过 GraphQL 协议把他们聚合在一起，形成具有数据全集的“超图”。

飞布用可视化界面封装了GraphQL细节，允许开发者通过勾选从“超图”中构建子集 Operation（查询、变更和订阅） 作为函数签名，并将其编译为 REST-API。这即充分利用了GraphQL按需取用、类型系统的优势，又免除了它无法复用HTTP基础设施以及不安全的弊端。

<figure><img src="../../.gitbook/assets/image (3) (1).png" alt=""><figcaption><p>数据流转图</p></figcaption></figure>

飞布还充分利用了GraphQL的指令系统，通过指令注解实现了API权限和数据权限的控制，入参校验，以及跨数据源关联！但用户无需刻意学习，因为飞布提供了友好的交互，封装了这些技术细节。
