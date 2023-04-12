---
description: 二锅头
---

# 组合式API

Fireboom 的核心概念之一是"Operation"。 Operation 可以是持久化的 GraphQL Operation，也可以是 TypeScript Operation， 它们都具有三种形式：Query,Mutation和Subscription. 此外，还有使用服务器端轮询的LiveQuery Operation。

编写 TypeScript 操作非常简单，因为我们正在使用 TypeScript 作为钩子服务语言。与此相反，配置 GraphQL 操作有点冗长，因为 GraphQL 不提供泛型，而且难以在多个操作间复用。

出于这个原因，我们使用 Typescript 来组合所有的 Operation，同时提供更丰富的自定义能力，后面简称 ts-operation。

## 快速上手

要使用 ts-operation，具体步骤如下：

1. 你需要在`custom-ts`目录新建`operations`目录。
2. 然后在该目录下新建 ts 文件，ts文件的具体编写方式见下文。
3. 编写完毕后，执行`npm run build-operations`命令生成配置。
4. 最后启动钩子服务即可！

## ts构建方式

### 查询请求

例如想要一个 `/operations/users/get`路由，那么应该建立`custom-ts/operations/users/get.ts`文件。然后在该文件中编辑

```ts
import { createOperation, z } from 'generated/fireboom.factory'

export default createOperation.query({
	input: z.object({
		id: z.string(),
	}),
	handler: async ({ input }) => {
    // 这里可以执行异步逻辑
		return {
			id: input.id,
			name: '张三',
      age: 22,
			avatar: 'https://i.pravatar.cc/300'
		}
	}
})
```

上面的示例合成的API请求为 `http://localhost:9991/operations/users/get`，method 为 get，请求参数为`id`，返回值结构为

```json
{
  "data": {
    "id": "1",
    "name": "张三",
    "age": 22,
    "avatar": "https://i.pravatar.cc/300"
  }
}
```

### 服务端轮询

如果需要支持服务端轮询，只需要将代码稍作修改

```diff
import { createOperation, z } from 'generated/fireboom.factory'

export default createOperation.query({
+ live: {
+		enable: true,
+		pollingIntervalSeconds: 2
+	},
	input: z.object({
		id: z.string(),
	}),
	handler: async ({ input }) => {
    // 这里可以执行异步逻辑
		return {
			id: input.id,
			name: '张三',
      age: 22,
			avatar: 'https://i.pravatar.cc/300'
		}
	}
})
```

### 变更请求

如果希望创建 mutation 的请求，只需要将`createOperation.query`改为`createOperation.mutation`，下面是一个示例

```ts
import { createOperation, z } from 'generated/fireboom.factory'

export default createOperation.mutation({
  // 参数格式
	input: z.object({
		id: z.string(),
		name: z.string(),
		age: z.optional(z.number()),
    avatar: z.optional(z.string())
	}),
	handler: async ({ input }) => {
    // 这里可以执行异步变更逻辑
		return {
			...input
		}
	}
})
```

### &#x20;订阅请求

你也可以使用 ts-operation 创建 subscription 类型的 operation，下面是一个示例

```ts
import { createOperation, z } from 'generated/fireboom.factory'

export default createOperation.subscription({
  input: z.object({
    // 参数
		id: z.string()
	}),
	handler: async function* ({ input }) {
		try {
			// 在这里你可以创建你自己的订阅，比如连接一个队列服务或者数据流服务
      // 这里演示隔1s发送一次消息，共10次
			for (let i = 0; i < 10; i++) {
        // 通过迭代器进行返回
				yield {
					id: input.id,
					name: '张三',
					age: 22,
			    avatar: 'https://i.pravatar.cc/300'
					time: new Date().toISOString()
				}
				// 模拟一些延迟
				await new Promise((resolve) => setTimeout(resolve, 1000))
			}
		} finally {
      // 在这里处理一些销毁操作，比如断开连接
			console.log('client disconnected')
		}
	}
})
```

### 内部调用

在 ts-operation 中你不仅可以从零开发业务，也可以直接使用`internalClient`直接调用已经创建的 Operation，下面是一个示例

```graphql
# 定义一个internalOperation GetOneUser
query MyQuery($id: Int!) @internalOperation {
  data: art_findOneUser(where: {id: {equals: $id}}) {
    id
    avatar
    createdAt
  }
}
```

然后在 ts-operation 调用

```ts
import { createOperation, z } from 'generated/fireboom.factory'

export default createOperation.query({
	input: z.object({
		id: z.string(),
	}),
	handler: async ({ input, log, internalClient }) => {
    // 这里可以调用 internalClient 进行操作
    const resp = await internalClient.queries.GetOneUser({
			input: 1
		})
    log.info('response', resp)
    if (resp.errors) {
      return {
        id: input.id,
        name: '张三',
        age: 22,
        avatar: 'https://i.pravatar.cc/300'
      }
    }
		return resp.data!.data!
	}
})
```

### 内部访问（开发中）&#x20;

如果你希望添加的 ts-operation 不可对外访问，仅作为内部接口使用，那么只需要在接口目录的最外层加一个 `internal` 即可，例如上面的`custom-ts/operations/users/get.ts`路径改为`custom-ts/operations/internal/users/get.ts`，就会变成仅内部可访问的接口。
