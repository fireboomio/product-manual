# 自定义数据源

当 Fireboom 默认提供的 GraphQL 查询变更功能不能满足需求时，你可以使用自定义数据源来补充。目前自定义数据源支持使用 `NodeJs` 进行开发，在此之前你需要先[准备环境](../../../huan-jing-zhun-bei/#nodejs)。

在 Fireboom 控制台点击`数据源`面板的`+`号，选择`自定义` -> `node.js`，输入新增数据源的名称，点击确定，此时会进入编辑模式。

默认填充的是示例代码，并没有进行保存，你可以在此修改代码后点击保存，或者点击`选择模板`按钮，选择一个模板插入并保存。

![自定义数据源面板](../../../.gitbook/assets/custom-ds-editor.png)

下面是我们 [AI魔法精灵案例](../../../zui-jia-shi-jian/xiao-cheng-xu-shi-zhan.md) 中的一段自定义数据源代码

```ts
// NodeJs 18以上可以不用引入该库
import fetch from '@web-std/fetch'
import { GraphQLObjectType, GraphQLSchema, GraphQLString, GraphQLNonNull, GraphQLID } from 'graphql'
import { type FastifyBaseLogger } from 'fastify/types/logger'
import { InternalClient } from 'fireboom-wundersdk/server'
import { type Mutations, type Queries } from 'generated/fireboom.internal.client';
import { createClient } from 'generated/client'

export default new GraphQLSchema({
  // 这里定义了一个保留的query，没有会报错
  query: new GraphQLObjectType({
    name: 'Query',
    fields: {
      _dummy: { type: GraphQLString }
    }
  }),
  mutation: new GraphQLObjectType<{
    args: string,
    userId: string
  }, {
    wundergraph: {
      log: FastifyBaseLogger,
      internalClient: InternalClient<Queries, Mutations>
    }
  }>({
    name: 'Mutation',
    fields: {
      GeneratePictureWithAI: {
        args: {
          // 必填的args
          args: {
            type: new GraphQLNonNull(GraphQLString)
          },
          // 必填的userId
          userId: {
            type: new GraphQLNonNull(GraphQLString)
          }
        },
        type: new GraphQLObjectType({
          fields: {
            // 返回字段 id
            id: {
              type: GraphQLID,
            },
            // 返回字段 url
            url: {
              type: GraphQLString
            }
          },
          name: 'data'
        }),
        async resolve(_, input, ctx) {
          const { log, internalClient } = ctx.wundergraph
          const userId = input.userId
          const json = JSON.parse(input.args)
          try {
            const client = createClient()
            // 消耗积分
            const resp = await client.mutate({
              operationName: 'UsePoints',
              input: {id: userId }
            })
            if (resp.error) {
            	throw resp.error
            }
            // 在自定义数据源中发起请求，实现自定义业务
            const data = await fetch('https://stablediffusionapi.com/api/v3/dreambooth', {
              method: 'post',
              headers: {
                'Content-Type': 'application/json'
              },
              body: JSON.stringify({
                key: 'xxx',
                samples: '1',
                num_inference_steps: '30',
                guidance_scale: 7.5,
                ...json
              })
            }).then(resp => resp.json())
            if (data.status === 'error') {
              throw new Error(data.messege)
            }
            if (data.status === 'success') {
              const resp1 = await internalClient.mutations.CreateOneCreation({ input: {args: input.args, userId }})
              if (!resp1.errors) {
                const resp2 = await internalClient.mutations.CreateOneDraft({ input: {
                  creationId: resp1.data!.data!.id!,
                  url: data.output[0]
                }})
                if (!resp2.errors) {
                  return {
                    url: data.output[0],
                    id: resp2.data!.data!.id!
                  }
                }
              }
            }
            return data
          } catch (error) {
            throw error
          }
        }
      }
    }
  })
})
```

自定义数据源创建完成后点击保存并打开开关，然后在创建API面板中就可以选择该数据源了。
