# 开放API

上面，我们学习了如何为接口绑定角色，以及如何为用户激活角色。

假设要实现一个带有RBAC的管理后台，有如下功能：

* 用户管理：用户的CRUD，及角色分配
* 角色管理：角色CRUD，及权限分配
* 菜单管理：菜单的增删改查，及权限关联

表面上看，这些都是数据库层面的关联关系，用OEPRTION构建API，供前端调用即可。但只在数据库中建立关系，如角色和API的关系，并无法影响Fireboom中的接口访问状态。

因此，我们需要在建立数据库关联的同时，将关联同步到飞布存储中，如角色权限关联关系，如何变成rbac指令描述。

<figure><img src="../../.gitbook/assets/image (4) (5).png" alt=""><figcaption></figcaption></figure>

我们当然可以在数据库建立关联后，使用`@rbac`指令手动在飞布控制台逐一为API绑定角色，但实在繁琐，且容易出错。

为解决上述问题，飞布提供了开放API。

开放API，本质上是一个内置的 rest 数据源，名字叫做`system`，由9123端口提供服务。

```http
http://localhost:9123
```

它有5个接口：

* 获取所有角色：/api/v1/role/all
* 获取角色的所有接口：/api/v1/role/apis
* 获取所有开启的API：/api/v1/operateApi/opened
* 为角色绑定API：/api/v1/role/bindApi
* 为角色解绑API：/api/v1/role/bindApi

## 获取所有角色

获取Fireboom系统中存储的所有角色。

其，REST 定义：

```json
{
    "/api/v1/role/all": {
        "get": {
            "operationId": "getAllRoles",
            "responses": {
                "200": {
                    "description": "",
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "array",
                                "items": {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "string" // 角色编码
                                        },
                                        "remark": {
                                            "type": "string" // 角色描述
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
```

使用方式如下：

```graphql
query MyQuery {
  system_getAllRoles {
    code
    remark
  }
} 
```

## 获取角色绑定的接口

根据角色获取其绑定的接口。

其，REST 定义：

```json
{
    "/api/v1/role/apis": {
        "get": {
            "operationId": "getRoleBindApis",
            "parameters": [
                {
                    "name": "code",
                    "in": "query",
                    "required": true,
                    "schema": {
                        "type": "string"
                    }
                },
                {
                    "name": "roleType",
                    "in": "query",
                    "required": false,
                    "schema": {
                        "type": "string",
                        "default": "requireMatchAny",
                        "enum": [
                            "requireMatchAll",
                            "requireMatchAny",
                            "denyMatchAll",
                            "denyMatchAny"
                        ]
                    }
                }
            ],
            "responses": {
                "200": {
                    "description": "",
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "array",
                                "items": {
                                    "type": "object",
                                    "properties": {
                                        "content": {
                                            "description": "内容",
                                            "type": "string"
                                        },
                                        "createTime": {
                                            "type": "string"
                                        },
                                        "deleteTime": {
                                            "type": "string"
                                        },
                                        "enabled": {
                                            "description": "开关 true开 false关",
                                            "type": "boolean"
                                        },
                                        "id": {
                                            "type": "integer"
                                        },
                                        "illegal": {
                                            "description": "是否非法 true 非法 false 合法",
                                            "type": "boolean"
                                        },
                                        "isPublic": {
                                            "description": "状态 true 公有 false 私有",
                                            "type": "boolean"
                                        },
                                        "liveQuery": {
                                            "description": "是否实时 true 是 false 否",
                                            "type": "boolean"
                                        },
                                        "method": {
                                            "description": "方法类型 GET、POST、PUT、DELETE",
                                            "type": "string"
                                        },
                                        "operationType": {
                                            "description": "请求类型 queries,mutations,subscriptions",
                                            "type": "string"
                                        },
                                        "remark": {
                                            "description": "说明",
                                            "type": "string"
                                        },
                                        "restUrl": {
                                            "description": "方法类型 GET、POST、PUT、DELETE",
                                            "type": "string"
                                        },
                                        "roleType": {
                                            "type": "string"
                                        },
                                        "roles": {
                                            "type": "string"
                                        },
                                        "title": {
                                            "description": "路径",
                                            "type": "string"
                                        },
                                        "updateTime": {
                                            "type": "string"
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
```

使用方式如下：

```graphql
# 获取角色绑定类型为requireMatchAny的所有API列表
query MyQuery($code: String!) {
  system_getRoleBindApis(code: $code, roleType: requireMatchAny) {
    createTime
    deleteTime
    enabled
    id
    illegal
    isPublic
    liveQuery
    method
    operationType
    remark
    restUrl
    roleType
    roles
    title
    updateTime
  }
}
```

## 查询所有开启的接口

获取Fireboom中所有已上线的接口。为角色绑定API时，需要知道API的ID，可通过该接口获取。

其，REST 定义：

```json
{
    "/api/v1/operateApi/opened": {
        "get": {
            "operationId": "getAllApis",
            "responses": {
                "200": {
                    "description": "",
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "array",
                                "items": {
                                    "type": "object",
                                    "properties": {
                                        "content": {
                                            "description": "内容",
                                            "type": "string"
                                        },
                                        "createTime": {
                                            "type": "string"
                                        },
                                        "deleteTime": {
                                            "type": "string"
                                        },
                                        "enabled": {
                                            "description": "开关 true开 false关",
                                            "type": "boolean"
                                        },
                                        "id": {
                                            "type": "integer"
                                        },
                                        "illegal": {
                                            "description": "是否非法 true 非法 false 合法",
                                            "type": "boolean"
                                        },
                                        "isPublic": {
                                            "description": "状态 true 公有 false 私有",
                                            "type": "boolean"
                                        },
                                        "liveQuery": {
                                            "description": "是否实时 true 是 false 否",
                                            "type": "boolean"
                                        },
                                        "method": {
                                            "description": "方法类型 GET、POST、PUT、DELETE",
                                            "type": "string"
                                        },
                                        "operationType": {
                                            "description": "请求类型 queries,mutations,subscriptions",
                                            "type": "string"
                                        },
                                        "remark": {
                                            "description": "说明",
                                            "type": "string"
                                        },
                                        "restUrl": {
                                            "description": "方法类型 GET、POST、PUT、DELETE",
                                            "type": "string"
                                        },
                                        "roleType": {
                                            "type": "string"
                                        },
                                        "roles": {
                                            "type": "string"
                                        },
                                        "title": {
                                            "description": "路径",
                                            "type": "string"
                                        },
                                        "updateTime": {
                                            "type": "string"
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
```

使用方式如下：

```graphql
query MyQuery {
  system_getAllApis {
    createTime
    deleteTime
    id
    enabled
    illegal
    isPublic
    liveQuery
    method
    operationType
    remark
    restUrl
    roleType
    roles
    title
    updateTime
    content
  }
}
```

## 角色绑定API

该接口比较特殊，能实现两个功能：

* 同步角色列表到飞布中
* 为某个角色<mark style="color:orange;">增量</mark>绑定API：只影响指定的API，未指定API不会发生变化

其，REST 定义：

```json
{
    "/api/v1/role/bindApi": {
        "post": {
            "operationId": "bindRoleApis",
            "requestBody": {
                "content": {
                    "application/json": {
                        "schema": {
                            "required": [
                                "roleCode",
                                "apis",
                                "allRoles"
                            ],
                            "type": "object",
                            "properties": {
                                "roleType": {
                                    "type": "string",
                                    "default": "requireMatchAny",
                                    "enum": [
                                        "requireMatchAll",
                                        "requireMatchAny",
                                        "denyMatchAll",
                                        "denyMatchAny"
                                    ]
                                },
                                "roleCode": {
                                    "type": "string"
                                },
                                "apis": {
                                    "type": "array",
                                    "items": {
                                        "type": "number"
                                    }
                                },
                                "allRoles": {
                                    "type": "array",
                                    "items": {
                                        "type": "string"
                                    }
                                }
                            }
                        }
                    }
                }
            },
            "responses": {
                "200": {
                    "description": "successful operations",
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "properties": {
                                    "count": {
                                        "type": "number"
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
```

使用方式如下：

```graphql
mutation MyQuery($roleCode: String!,  $apis: [Int]!,  $allRoles: [String]!) {
  system_bindRoleApis(
    POSTApiV1RoleBindApiInput: {apis: $apis, allRoles: $allRoles, roleCode: $roleCode, roleType: requireMatchAny}
  ) {
    count
  }
} 
```

其入参为角色`roleCode`，以及要绑定的api列表`apis`，为int数组，即接口的ID列表，对应 [#cha-xun-suo-you-kai-qi-de-jie-kou](kai-fang-api.md#cha-xun-suo-you-kai-qi-de-jie-kou "mention")的返回值 `id`。

此外还有一个入参是所有的角色列表`allRoles`。之所以这样设计是为了把数据库中的角色同步到飞布的角色列表中。和 [#huo-qu-suo-you-jiao-se](kai-fang-api.md#huo-qu-suo-you-jiao-se "mention")接口相反。

## 角色解绑API

为角色增量解绑API。

```json
{
    "/api/v1/role/unbindApi": {
        "post": {
            "operationId": "unBindRoleApis",
            "requestBody": {
                "content": {
                    "application/json": {
                        "schema": {
                            "required": [
                                "roleCode",
                                "apis"
                            ],
                            "type": "object",
                            "properties": {
                                "roleType": {
                                    "type": "string",
                                    "default": "requireMatchAny",
                                    "enum": [
                                        "requireMatchAll",
                                        "requireMatchAny",
                                        "denyMatchAll",
                                        "denyMatchAny"
                                    ]
                                },
                                "roleCode": {
                                    "type": "string"
                                },
                                "apis": {
                                    "type": "array",
                                    "items": {
                                        "type": "number"
                                    }
                                }
                            }
                        }
                    }
                }
            },
            "responses": {
                "200": {
                    "description": "successful operations",
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "properties": {
                                    "count": {
                                        "type": "number"
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
```

使用方式如下：

```graphql
mutation MyQuery($roleCode: String!, $apis: [Int]!) {
  system_unBindRoleApis(
    POSTApiV1RoleUnbindApiInput: {roleCode: $roleCode, apis: $apis}
  ) {
    count
  }
}
```

## 总结

总结下，实现管理后台相对比较复杂，需要自行编写业务逻辑，调用数据存储和开放API，保持两者同步。

{% hint style="info" %}
[#jiao-se-bang-ding-api](kai-fang-api.md#jiao-se-bang-ding-api "mention") [#jiao-se-jie-bang-api](kai-fang-api.md#jiao-se-jie-bang-api "mention")两个接口为变更操作，执行后需手动触发编译后，Fireboom中才能生效。
{% endhint %}

当前，飞布提供了完整的后台管理示例，您可以参考代码实现自己的管理后台。

详情查看：[fireboom-admin](../../zui-jia-shi-jian/fireboom-admin/ "mention")
