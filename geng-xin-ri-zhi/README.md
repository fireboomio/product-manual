# 更新日志

<details>

<summary>V2.0.0  2023/09/04</summary>

「新增功能」&#x20;

1、新增fromClaim参数custom，从User.CustomClaims中获取数据，实现了任意数据的注入&#x20;

2、新增@transaction参数，允许mutation执行事务&#x20;

3、底层存储升级，目录结构更加清晰&#x20;

4、引擎重构，大幅缩小配置文件大小和编译时间&#x20;

5、增量编译，新版本将全量编译改为了增量编译，极大提升了OPERATION的编译速度！

6、支持了function钩子，且还解决了钩子的循环依赖问题&#x20;

7、新增prisma数据源：支持虚拟外键、视图

「缺陷修复」 修复其他已知问题，提升了飞布的稳定性

「优化」 界面交互优化，提升产品易用性

</details>

详情见： [v2.0-geng-xin-shuo-ming.md](v2.0-geng-xin-shuo-ming.md "mention")
