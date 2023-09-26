# 请求时序图

## 单数据源

<figure><img src="../.gitbook/assets/image (8) (1) (2).png" alt=""><figcaption><p>Fireboom单数据源请求时序图</p></figcaption></figure>

0 定义接口

```graphql
query MyQuery($id: Int = 1) {
  todo_findUniqueTodo(where: {id: $id}) {
    completed
    createdAt
    id
    title
  }
}
```

1-2飞布启动时

9-10飞布关闭时

3 请求

```bash
curl 'http://localhost:9991/operations/Single?id=1'
```

4 findUniqueTodo

```graphql
query MyQuery {
  findUniqueTodo(where: { id: 1 }) {
    completed
    createdAt
    id
    title
  }
}
```

5 SQL req

```sql
SELECT `completed`,`createdAt`,`id`,`title` FROM Todo WHERE id=1
```

6 SQL res

![](<../.gitbook/assets/image (14).png>)

7 Query res(json)

```json
{
  "data": {
    "findUniqueTodo": {
      "completed": false,
      "createdAt": "2022-12-31T16:00:01.000Z",
      "id": 1,
      "title": "This is Fireboom"
    }
  }
}
```

8 响应

```json
{
    "todo_findUniqueTodo": {
        "completed": true,
        "createdAt": "2022-12-31T16:00:01.000Z",
        "id": 1,
        "title": "Hello world"
    }
}
```

## 多数据源（并行）

<figure><img src="../.gitbook/assets/image (9) (1).png" alt=""><figcaption><p>Fireboom多数据源请求时序图</p></figcaption></figure>

0 定义接口

```graphql
query MyQuery($id: Int = 1) {
  todo_findUniqueTodo(where: { id: $id }) {
    completed
    createdAt
    id
    title
  }
  system_getAllRoles {
    code
    remark
  }
} 
```

6 getAllRoles req

```bash
curl 'http://localhost:9123/api/v1/role/all'
```

7 getAllRoles res

```json
[
  { "code": "admin", "remark": "" },
  { "code": "user", "remark": "" }
]
```

8 响应

```json
{
    "todo_findUniqueTodo": {
        "completed": true,
        "createdAt": "2022-12-31T16:00:01.000Z",
        "id": 1,
        "title": "Hello world"
    },
    "system_getAllRoles": [
        {
            "code": "admin",
            "remark": ""
        },
        {
            "code": "user",
            "remark": ""
        }
    ]
}
```

## 多数据源（串行-跨源关联）

最后，我们介绍下跨源特性，看下它的时序图。

<figure><img src="../.gitbook/assets/image (10) (3).png" alt=""><figcaption></figcaption></figure>

跨源关联本质上是一种流程编排，将通常情况下并行的请求，改造成串行。

即先执行第一个接口，然后将第一个接口的返回结果作为入参传递给第二个接口。

详情请查看 [kua-yuan-guan-lian.md](../ji-chu-ke-shi-hua-kai-fa/api-gou-jian/kua-yuan-guan-lian.md "mention")
