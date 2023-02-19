# 授权与访问控制

飞布基于RBAC规范，结合GraphQL的注解能力，实现了API的的授权与访问控制。



## 快速操作

### 基本设置

1. 在身份验证面板中点击“\
   <img src="http://localhost:9123/assets/workbench/panel-role.png" alt="头像" data-size="line">”，进入“角色管理”TAB
2. 根据业务需求添加角色，系统默认内置admin和user角色（请确保必须有1个角色）
3. 切换到“身份鉴权”TAB，在auth目录下选择`mutatingPostAuthentication`文件
4. 编写钩子脚本或选择预制脚本，详情前往钩子章节

{% hint style="info" %}
由于OIDC规范中没有用户角色的相关声明，用户通过OIDC流程登录后，claim中不包含roles字段。而RBAC要求用户绑定角色，通过角色匹配接口权限。因此，需要有个地方为用户动态注入roles字段，即用户第一次登录时，根据用户ID或email去特定数据源（可能是自有数据库或者其他数据源）查找其关联的角色，并绑定到roles字段上。考虑到灵活扩展，钩子是实现该功能的最佳场所。
{% endhint %}

### API设置



1. 前往API管理面板，选择需要设置权限的API
2. 在GraphQL编辑区的工具栏中点击“@角色”，选择匹配模式并添加角色
3. 点击顶部菜单栏的“<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAMAAAC7IEhfAAAAY1BMVEUAAADU1NRjZmxvcnePkZVgY2rPz9BfY2poaHTAwMOAhIxoa3FgYmpgYmlgY2nFxcbU1NRmZm+Ag427u73U1NRfYml/g4zt7e3k5eXW1te9vsCztLefoaWanKCOkJVydXtucXecDQKGAAAAFXRSTlMAzP336NDOiAvTz/rn2tjSph7Qs6d9epWLAAAAjElEQVQ4y+2T2Q6EIAxFK+A6mzMj4q7//5VaYngCG2N8cDkvNOlJSG9TuCq+XMQ3oiQ4p0jGsx+/fCIByDwrqRFzDYDn4BatYiw4Y1zEhBgIJjUsjJbED5eG19ctBtrr66rD9x05RYH9oVBKtViFTvGB7UZNlFg9N4n01/QwdDwrA0/mU0jtK/zDYRgBwgsrsPomQg4AAAAASUVORK5CYII=" alt="预览" data-size="line">”，前往API预览页
4. ~~进入详情页，点击右上角“测试”按钮，登录后将能查看到当前用户信息~~

## 授权





## 访问控制



用户的角色何时绑定？

