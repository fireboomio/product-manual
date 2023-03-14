---
description: 使用飞布开发微信小程序
---

# AI魔法师实战

### 视频课程

{% embed url="https://www.bilibili.com/video/BV1Wg4y147KK/" %}

{% embed url="https://www.bilibili.com/video/BV1fb411f7G7" %}

### 体验码

![](../assets/AIGC-qrcode.jpeg)

仓库地址：[https://github.com/fireboomio/case-ai-art](https://github.com/fireboomio/case-ai-art)

### 1. 项目需求 <a href="#hslit" id="hslit"></a>

[https://lanhuapp.com/link/#/invite?sid=lX0tcrqd](https://lanhuapp.com/link/#/invite?sid=lX0tcrqd)

### 2. 项目初始化 <a href="#pddqo" id="pddqo"></a>

#### 2.1 初始化目录 <a href="#fw5q7" id="fw5q7"></a>

```bash
mkdir -p fb-art
cd fb-art
curl -fsSL https://www.fireboom.io/install.sh | bash -s server
```

#### 2.2. 前端代码生成 <a href="#qjiph" id="qjiph"></a>

我们使用 codefun 快速生成前端代码

[https://ide.code.fun/projects/63aafc28d109e600121b3472/pages](https://ide.code.fun/projects/63aafc28d109e600121b3472/pages)

稍做整理放到 miniapp 目录

或者使用已经整理好的代码

```bash
git clone -b miniapp https://github.com/fireboomio/case-ai-art.git miniapp
cd miniapp
```

### 3. 需求拆解生成数据库模型 <a href="#xwgds" id="xwgds"></a>

```prisma
// 用户
model AppUser {
  id             String   @id
  nickname       String
  avatar         String
  phone          String?
  provider       String
  providerId     String
  // 积分
  points         Int      @default(0)
  createdAt      DateTime @default(now())

  // 积分记录
  pointRecords PointRecord[]
  // 点赞记录
  likeRecords  LikeRecord[]

  // 创作
  creations Creation[]

  // 邀请我的人的ID
  inviteById   String?
  // 被谁邀请
  inviteBy     AppUser?  @relation("InviteHistory", fields: [inviteById], references: [id])
  // 我邀请的
  invitedUsers AppUser[] @relation("InviteHistory")
}

// 创作过程
model Creation {
  id        Int      @id @default(autoincrement())
  authorId  String
  author    AppUser  @relation(fields: [authorId], references: [id])
  // 创作时的参数 json 序列化
  args      String
  createdAt DateTime @default(now())

  artWorks ArtWork[]
}

// 作品
model ArtWork {
  id            Int          @id @default(autoincrement())
  creation      Creation     @relation(fields: [creationId], references: [id])
  creationId    Int
  // 作品存储地址
  url           String
  // 是否是草稿
  isDraft       Boolean      @default(true)
  // 是否发布到画廊
  published     Boolean      @default(false)
  // 发布时间
  publishAt     DateTime?
  // 点赞数
  likeCount     Int          @default(0)
  // 点赞记录
  likeRecords   LikeRecord[]
  // 推荐指数
  recommendRate Int          @default(0)
  // 被分享次数
  sharedCount   Int          @default(0)
}

// 点赞记录
model LikeRecord {
  id        Int      @id @default(autoincrement())
  artWorkId Int
  artWork   ArtWork  @relation(fields: [artWorkId], references: [id])
  userId    String
  user      AppUser  @relation(fields: [userId], references: [id])
  createdAt DateTime @default(now())
}

// 积分获得方式
enum PointWays {
  // 分享给好友
  ShareToFriend
  // 分享到群
  ShareToGroup
  // 观看广告
  WatchAD
  // 邀请好友加入
  Invite
  // 画图使用消耗
  Draw
}

// 积分获得记录
model PointRecord {
  id        Int       @id @default(autoincrement())
  userId    String
  user      AppUser   @relation(fields: [userId], references: [id])
  // 获得积分方式
  way       PointWays
  // 获得积分数
  point     Int
  createdAt DateTime  @default(now())
}
```

### 4. Fireboom接口创建 <a href="#ztx77" id="ztx77"></a>

#### 4.1 身份鉴权 - OIDC登录 <a href="#hecnk" id="hecnk"></a>

Authing 配置

```
clientId 63ae88490f3dff9d1a651cea
secert e8094233545dec4b3b7fc848eb04e306
issuer sail.authing.cn/oidc
```

业务表用户创建

```typescript
import { AuthenticationHookRequest } from 'fireboom-wundersdk/server'

export default async function postAuthentication(hook: AuthenticationHookRequest) : Promise<void>{
  if (hook.user?.userId && hook.user.providerId !== 'admin') {
    const { provider, providerId, userId } = hook.user
    const { nickname, picture } = hook.user.idToken!
    const resp = await hook.internalClient.queries.FindOneAppUser({
      input: {
        id: userId
      }
    })
    const existedUser = resp?.data?.data
    if (!existedUser) {
      await hook.internalClient.mutations.CreateOneAppUser({
        input: {
          data: {
            id: userId,
            provider,
            providerId,
            nickname,
            avatar: picture,
            phone: '',
            points: 100
          }
        }
      })
    }
  }
}
```

需要前置2个查询

```graphql
query FindOneAppUser($id: String!) {
  data: art_findFirstAppUser(where: {id: {equals: $id}}) {
    id
    avatar
    inviteById
    nickname
    phone
    points
    provider
    providerId
    createdAt
  }
}
```

```graphql
mutation CreateOneAppUser($data: art_AppUserCreateInput!) @internalOperation {
  data: art_createOneAppUser(
    data: $data
  ) {
    id
  }
}
```

登录后就需要获取用户信息 `fromClaim` 功能

```graphql
query GetUserinfo($userId: String @fromClaim(name: USERID)) {
  data: art_findFirstAppUser(where: {id: {equals: $userId}}) {
    avatar
    id
    nickname
    phone
    points
  }
}
```

修改用户基本信息接口

```graphql
mutation UpdateUserBaseinfo( $nickname: String, $avatar: String,$id: String! @fromClaim(name: USERID)) {
  data: art_updateOneAppUser(
    data: {nickname: {set: $nickname}, avatar: {set: $avatar}}
    where: {id: $id}
  ) {
    id
    nickname
    avatar
  }
}
```

#### 4.2 创作接口 - 自定义数据源能力 <a href="#patox" id="patox"></a>

翻译

```typescript
import fetch from '@web-std/fetch'
import { type FastifyBaseLogger } from 'fastify/types/logger'

export async function translate(str: string, logger: FastifyBaseLogger) {
  try {
    const resp = await fetch("https://www2.deepl.com/jsonrpc?method=LMT_handle_jobs", {
      "headers": {
        "accept": "*/*",
        "accept-language": "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7,ja;q=0.6",
        "cache-control": "no-cache",
        "content-type": "application/json",
        "pragma": "no-cache",
        "sec-ch-ua": "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Google Chrome\";v=\"108\"",
        "sec-ch-ua-mobile": "?0",
        "sec-ch-ua-platform": "\"macOS\"",
        "sec-fetch-dest": "empty",
        "sec-fetch-mode": "cors",
        "sec-fetch-site": "same-site",
        "cookie": "dapUid=30d87eb9-d7bc-4a37-85a1-31efc18da077; dl_session=fa.ee3ef862-c610-4c93-a99f-435934073958; userRecommendations=RC-3.2; releaseGroups=340.DF-2477.2.3_399.RC-3.2.8_469.DM-541.2.2_470.DM-542.2.2_471.DM-547.2.2_475.DM-544.2.2_604.DM-595.2.2_605.DM-585.2.3_606.DWFA-165.2.3_612.DM-538.2.2_622.B2B-158.2.1_633.DM-695.2.2_769.B2B-127.2.2_778.DM-705.1.1_863.DM-601.2.2_866.DM-592.2.2_867.DM-684.2.4_972.B2B-138.2.3_975.DM-609.1.2_1084.TG-1207.2.3_1085.TC-432.2.2_1086.TC-104.1.6_1090.DWFA-345.2.2_1092.SEO-44.2.2_1119.B2B-251.2.2_1202.DF-2381.1.2_1205.DAL-176.2.1_1207.DWFA-96.2.2_1219.DAL-136.1.2_1224.DAL-186.2.2_1245.SEO-113.2.2_1330.DF-3073.2.1_220.DF-1925.1.9_1338.AAEXP-724.1.1_1328.DWFA-285.1.1_774.DWFA-212.2.2_1342.AAEXP-728.1.1_1347.AAEXP-733.1.1_1354.AAEXP-740.1.1_1336.AAEXP-722.1.1_1339.AAEXP-725.2.1_1341.AAEXP-727.1.1_1337.AAEXP-723.2.1_1352.AAEXP-738.1.1_1353.AAEXP-739.1.1_1349.AAEXP-735.1.1_1350.AAEXP-736.1.1_1344.AAEXP-730.1.1_1351.AAEXP-737.1.1_1345.AAEXP-731.1.1_1327.DWFA-391.1.1_1343.AAEXP-729.2.1_1348.AAEXP-734.1.1_1340.AAEXP-726.2.1_1346.AAEXP-732.1.1_1335.AAEXP-721.2.1; LMTBID=v2|94536b41-e487-4a2f-9615-29ca9ba751b4|490dd240be839f71bae7772125afb228; userCountry=CN; privacySettings=%7B%22v%22%3A%221%22%2C%22t%22%3A1672790400%2C%22m%22%3A%22LAX%22%2C%22consent%22%3A%5B%22NECESSARY%22%2C%22PERFORMANCE%22%2C%22COMFORT%22%5D%7D; dapVn=2; dapSid=%7B%22sid%22%3A%22939602dd-b714-4278-b4e6-e0153929242c%22%2C%22lastUpdate%22%3A1672799309%7D",
        "Referer": "https://www.deepl.com/",
        "Referrer-Policy": "strict-origin-when-cross-origin"
      },
      "body": `{\"jsonrpc\":\"2.0\",\"method\": \"LMT_handle_jobs\",\"params\":{\"jobs\":[{\"kind\":\"default\",\"sentences\":[{\"text\":\"${str}\",\"id\":0,\"prefix\":\"\"}],\"raw_en_context_before\":[],\"raw_en_context_after\":[],\"preferred_num_beams\":4,\"quality\":\"fast\"}],\"lang\":{\"preference\":{\"weight\":{\"EN\":10},\"default\":\"default\"},\"source_lang_user_selected\":\"ZH\",\"target_lang\":\"EN\"},\"priority\":-1,\"commonJobParams\":{\"regionalVariant\":\"en-US\",\"mode\":\"translate\",\"browserType\":1},\"timestamp\":${+new Date()}},\"id\":${Math.ceil(Math.random() * 99999999)}}`,
      "method": "POST"
    });
    const res = await resp.json()
    if (res.result) {
      return res.result.translations[0].beams[0].sentences[0].text
    }
    return str
  } catch (error) {
    logger.error(error)
    return str
  }
}
```

AI 生成图片 - 自定义数据源

```typescript
import { GraphQLObjectType, GraphQLSchema, GraphQLString, GraphQLNonNull, GraphQLID } from 'graphql'
import { FastifyBaseLogger } from 'fastify/types/logger'
import { InternalClient } from 'fireboom-wundersdk/server'
import { Mutations, Queries } from 'generated/fireboom.internal.client';
import { createClient } from 'generated/client'
import { translate } from './translate'

export default new GraphQLSchema({
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
          args: {
            type: new GraphQLNonNull(GraphQLString)
          },
          userId: {
            type: new GraphQLNonNull(GraphQLString)
          }
        },
        type: new GraphQLObjectType({
          fields: {
            id: {
              type: GraphQLID,
            },
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
          if (json.prompt && /[\u4e00-\u9fa5]/.test(json.prompt)) {
            json.prompt = await translate(json.prompt, log)
          }
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
            // const data = await fetch('https://stablediffusionapi.com/api/v3/dreambooth', {
            //   method: 'post',
            //   headers: {
            //     'Content-Type': 'application/json'
            //   },
            //   body: JSON.stringify({
            //     key: 'EXKt6qfISnSLeFYlLpw9gqHirkCT8hWIpYWllNfuvCUHXBIYk8UoIkUkX8xk',
            //     samples: '1',
            //     num_inference_steps: '30',
            //     guidance_scale: 7.5,
            //     ...json
            //   })
            // }).then(resp => resp.json())
            const data = {
              status: 'success',
              messege: '',
              output: ['https://stable-diffusion-api.s3.amazonaws.com/generations/dd011494-4bd7-4257-9eb6-dc2269b1288c-0.png']
            }
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

记录创作过程

```graphql
mutation CreateOneCreation($args: String!, $userId: String! @fromClaim(name: USERID)) @internalOperation {
  data: art_createOneCreation(
    data: {args: $args, AppUser: {connect: {id: $userId}}}
  ) {
    id
  }
}
```

根据输入创建草稿

```graphql
mutation CreateOneDraft($url: String!, $creationId: Int!) @internalOperation {
  data: art_createOneArtWork(
    data: {url: $url, Creation: {connect: {id: $creationId}}, published: false, isDraft: true, likeCount: 0, recommendRate: 0, sharedCount: 0}
  ) {
    id
    url
  }
}
```

#### 4.3 开始创作 - API快速创建 <a href="#gvsbl" id="gvsbl"></a>

**4.3.1 创作**

发起创作

```graphql
mutation GeneratePictureWithAI($args: String!, $userId: String! @fromClaim(name: USERID)) {
  data: ai_GeneratePictureWithAI(args: $args, userId: $userId) {
    id
    url
  }
}
```

增积分

```graphql
mutation IncreasePoints($id: String!, $points: Int!) @internalOperation {
  data: art_updateOneAppUser(data: {points: {increment: $points}}, where: {id: $id}) {
    id
  }
}
```

减积分

```graphql
mutation DecreasePoints($id: String!, $points: Int!) @internalOperation {
  data: art_updateOneAppUser(data: {points: {decrement: $points}}, where: {id: $id}) {
    id
  }
}
```

创作消耗积分

```graphql
mutation UsePoints($id: String!) {
  data: art_createOnePointRecord(
    data: {
        way: Draw,
        AppUser: {connect: {id: $id}},
        point: -1
      }
  ) {
    id
  }
}
```

使用`customResolve`钩子判断积分是否足够

```typescript
import { HookRequestWithInput } from 'generated/fireboom.hooks'
import { InjectedUsePointsInput, UsePointsResponse } from 'generated/models'

export default async function customResolve(hook: HookRequestWithInput<InjectedUsePointsInput>)
  //: Promise<void | UsePointsResponse>{ // 取消注释以使用严格的返回类型
  : Promise<any>{
  // TODO: 在此处添加代码
  const resp = await hook.internalClient.queries.FindOneAppUser({
    input: {
      id: hook.input?.id
    }
  })
  // 积分不足时，跳过后续流程直接返回
  const points = resp?.data?.data?.points ?? 0
  if (points < 1) {
    return {
      data: {},
      errors: [{
        message: '积分不足'
      }]
    }
  }
  // 修改用户积分
  await hook.internalClient.mutations.DecreasePoints({
    input: {
      points: 1,
      id: hook.input?.id
    }
  })
}
```

**4.3.2 草稿转画夹**

```graphql
mutation MoveDraftToAlbum($id: Int!) {
  data: art_updateOneArtWork(
    data: {isDraft: {set: false}}
    where: {id: $id}
  ) {
    id
  }
}
```

**4.3.3 画夹转画廊**

```graphql
mutation PublishMyArtWork($id: Int!, $date: DateTime @injectCurrentDateTime(format: ISO8601)) {
  data: art_updateOneArtWork(
    data: {publishAt: {set: $date}, published: {set: true}}
    where: {id: $id}
  ) {
    id
    publishAt
    published
  }
}
```

**4.3.4 画廊列表详情 分页**

最新

```graphql
query GetWorksByNewest($skip: Int = 0, $take: Int = 10) {
  data: art_findManyArtWork(
    where: {isDraft: { equals: false}, published: {equals: true}}
    orderBy: [{ publishAt: desc }]
    skip: $skip
    take: $take
  ) {
    id
    likeCount
    sharedCount
    url
  }
}
```

最热

```graphql
query GetWorksByHot($skip: Int = 0, $take: Int = 10) {
  data: art_findManyArtWork(
    where: {isDraft: { equals: false}, published: {equals: true}}
    orderBy: [{likeCount: desc}, { sharedCount: desc }, {recommendRate:desc}]
    skip: $skip
    take: $take
  ) {
    id
    likeCount
    sharedCount
    url
  }
}
```

推荐

```graphql
query GetWorksByRecommend($skip: Int = 0, $take: Int = 10) {
  data: art_findManyArtWork(
    where: {isDraft: { equals: false}, published: {equals: true}}
    orderBy: [{recommendRate: desc}, {likeCount: desc}]
    skip: $skip
    take: $take
  ) {
    id
    likeCount
    recommendRate
    sharedCount
    url
  }
}
```

作品详情

```graphql
query GetArtWorkDetail($id: Int!, $userId: String! @fromClaim(name: USERID)) {
  data: art_findFirstArtWork(where: {id: {equals: $id}}) {
    id
    likeCount
    sharedCount
    url
    args: Creation @transform(get: "args") {
      args
    }
    user: Creation @transform(get: "AppUser") {
      AppUser {
        avatar
        id
        nickname
      }
    }
    isDraft
    published
    likeRecords: LikeRecord(
      where: {AppUser: {is: {id: {equals: $userId}}}, ArtWork: {is: {id: {equals: $id}}}}
    ) {
      id
    }
  }
}
```

**4.3.5 点赞**

增加点赞数

```graphql
mutation IncreaseArtWorkLikeCount($id: Int!) @internalOperation {
  data: art_updateOneArtWork(data: {likeCount: {increment: 1}}, where: {id: $id}) {
    id
  }
}
```

创建点赞记录

```graphql
mutation LikeOneArtWork($userId: String! @fromClaim(name: USERID), $artWorkId: Int!) {
  data: art_createOneLikeRecord(
    data: {ArtWork: {connect: {id: $artWorkId}}, AppUser: {connect: {id: $userId}}}
  ) {
    id
  }
}
```

点赞后增加点赞数钩子

```typescript
import { HookRequestWithInput, HookRequestWithResponse } from 'generated/fireboom.hooks'
import { InjectedLikeOneArtWorkInput, LikeOneArtWorkResponse } from 'generated/models'

export default async function postResolve(hook: HookRequestWithInput<InjectedLikeOneArtWorkInput> & HookRequestWithResponse<LikeOneArtWorkResponse>) : Promise<void>{
	// 增加like数
   await hook.internalClient.mutations.IncreaseArtWorkLikeCount({ input: {id: hook.input.artWorkId }})
}
```

减少点赞数

```graphql
mutation DecreaseArtWorkLikeCount($id: Int!) @internalOperation {
  data: art_updateOneArtWork(data: {likeCount: {decrement: 1}}, where: {id: $id}) {
    id
  }
}
```

取消点赞记录

```graphql
mutation UnlikeOneArtWork($userId: String! @fromClaim(name: USERID), $artWorkId: Int!) {
  data: art_deleteManyLikeRecord(
    where: {AppUser: {is: {id: {equals: $userId}}}, ArtWork: {is: {id: {equals: $artWorkId}}}}
  ) {
    count
  }
}
```

取消点赞钩子

```typescript
import { HookRequestWithInput, HookRequestWithResponse } from 'generated/fireboom.hooks'
import { InjectedUnlikeOneArtWorkInput, UnlikeOneArtWorkResponse } from 'generated/models'

export default async function postResolve(hook: HookRequestWithInput<InjectedUnlikeOneArtWorkInput> & HookRequestWithResponse<UnlikeOneArtWorkResponse>) : Promise<void>{
	// like - 1
  await hook.internalClient.mutations.DecreaseArtWorkLikeCount({ input: {id: hook.input.artWorkId }})
}
```

**4.3.6 个人画展**

* 我的草稿

```graphql
query GetMyDrafts($userId: String! @fromClaim(name: USERID), $creationId: Int! @internal, $skip: Int = 0, $take: Int = 10) {
  data: art_findManyCreation(
    orderBy: {createdAt: desc}
    skip: $skip
    take: $take
    where: {AppUser: {is: {id: {equals: $userId}}}}
  ) {
    id @export(as: "creationId")
    count: _count {
      ArtWork
    }
    artWork: _join @transform(get: "art_findFirstArtWork") {
      art_findFirstArtWork(where: {creationId: {equals: $creationId}}) {
        id
        url
      }
    }
  }
}
```

草稿集

```graphql
query GetMyDraftItems($draftId: Int!, $userId: String! @fromClaim(name: USERID)) {
  data: art_findManyArtWork(
    where: {creationId: {equals: $draftId}, isDraft: {equals: true}, Creation: {is: {authorId: {equals: $userId}}}}
  ) {
    id
    url
  }
}
```

* 我的画夹

```graphql
query GetMyAlbum($skip: Int = 0, $take: Int = 10, $userId: String! @fromClaim(name: USERID)) {
  data: art_findManyArtWork(
    where: {isDraft: {equals: false}, Creation: {is: {AppUser: {is: {id: {equals: $userId}}}}}}
    skip: $skip
    take: $take
  ) {
    id
    likeCount
    publishAt
    published
    recommendRate
    sharedCount
    url
  }
}
```

* 我的画廊

```graphql
query GetMyPublicAlbum($skip: Int = 0, $take: Int = 10, $userId: String! @fromClaim(name: USERID)) {
  data: art_findManyArtWork(
    where: {published: {equals: true}, Creation: {is: {AppUser: {is: {id: {equals: $userId}}}}}}
    skip: $skip
    take: $take
  ) {
    id
    likeCount
    sharedCount
    url
  }
}
```

* 我的点赞

```graphql
query GetMyLiked($skip: Int = 0, $take: Int = 10, $userId: String! @fromClaim(name: USERID)) {
  data: art_findManyLikeRecord(
    skip: $skip
    take: $take
    where: {AppUser: {is: {id: {equals: $userId}}}}
  ) {
    artWork: ArtWork {
      id
      url
      sharedCount
      likeCount
    }
  }
}
```

**4.3.7 个人资料**

获取我的积分

```graphql
query GetMyPointRecords($way: [art_PointRecord_way], $userId: String! @fromClaim(name: USERID), $timeStart: DateTime!) {
  data: art_findManyPointRecord( where: {
    way: {in: $way},
    userId: {equals: $userId},
    createdAt: {gte: $timeStart}
  }) {
    id,
    way
  }
}
```

**4.3.8 分享**

* 创建积分记录

```graphql
mutation CreateOnePointRecord($way:art_PointRecord_way!, $point: Int! = 0, $userId: String!) {
  data: art_createOnePointRecord(
    data: {
      way: $way,
      AppUser: {connect: {id: $userId}},
      point: $point
    }
  ) {
    id
  }
}
```

使用 `mutatingPreResolve`钩子设置每种积分方式的分值

```typescript
import { HookRequestWithInput } from 'generated/fireboom.hooks'
import { InjectedCreateOnePointRecordInput } from 'generated/models'

export default async function mutatingPreResolve(hook: HookRequestWithInput<InjectedCreateOnePointRecordInput>) : Promise<InjectedCreateOnePointRecordInput>{
  switch (hook.input.way) {
    case 'ShareToFriend':
      hook.input.point = 2
      break
    case 'ShareToGroup':
      hook.input.point = 2
      break
    case 'WatchAD':
      hook.input.point = 4
      break
    case 'Invite':
      hook.input.point = 10
      break
    default:
      hook.input.point = 0
  }
  return hook.input
}
```

使用`customResolve`钩子判断今日积分增加次数是否用完

```typescript
import { HookRequestWithInput } from 'generated/fireboom.hooks'
import { InjectedCreateOnePointRecordInput, CreateOnePointRecordResponse } from 'generated/models'

export default async function customResolve(hook: HookRequestWithInput<InjectedCreateOnePointRecordInput>)
  : Promise<void | CreateOnePointRecordResponse> {
  // 根据用户查询今日分享次数，超出直接200返回
  const now = new Date(Math.floor(Date.now() / 86400000) * 86400000)
  const resp = await hook.internalClient.queries.GetMyPointRecords({
    input: {
      way: [hook.input.way],
      userId: hook.input.userId,
      timeStart: now.toISOString()
    }
  })
  const limitTimes = {
    ShareToFriend: 2,
    ShareToGroup: 2,
    WatchAD: 10
  }[hook.input.way] ?? 0
  const currentTimes = resp.data?.data?.length??0
  if ((currentTimes >= limitTimes) && limitTimes) {
    return {
      data:{}
    }
  }
}
```

使用`postResolve`钩子创建完记录后增加用户积分

```typescript
import { HookRequestWithInput, HookRequestWithResponse } from 'generated/fireboom.hooks'
import { InjectedCreateOnePointRecordInput, CreateOnePointRecordResponse } from 'generated/models'

export default async function postResolve(hook: HookRequestWithInput<InjectedCreateOnePointRecordInput> & HookRequestWithResponse<CreateOnePointRecordResponse>): Promise<void> {
  await hook.internalClient.mutations.IncreasePoints({
    input: {
      id: hook.input.userId,
      points: hook.input.point
    }
  })
}
```

* 增加积分接口

```graphql
mutation IncreasePoints($id: String!, $points: Int!) @internalOperation {
  data: art_updateOneAppUser(data: {points: {increment: $points}}, where: {id: $id}) {
    id
  }
}
```

* 增加作品分享次数

```graphql
mutation IncreaseArtWorkShareCount($id: Int!) @internalOperation {
  data: art_updateOneArtWork(
    data: {sharedCount: {increment: 1}}
    where: {id: $id}
  ) {
    id
  }
}
```

* 记录我的邀请者接口

```graphql
mutation RecoredMyInviter($inviterId: String!, $userId: String! @fromClaim(name: USERID)) {
  data: art_updateOneAppUser(
    data: {AppUser: {connect: {id: $inviterId}}}
    where: {id: $userId}
  ) {
    id
  }
}
```

使用`customResolve`钩子在记录前判断邀请人状态

```typescript
import { createClient } from 'generated/client'
import {HookRequestWithInput} from 'generated/fireboom.hooks'
import {InjectedRecordMyInviterInput, RecordMyInviterResponse} from 'generated/models'
export default async function customResolve(hook: HookRequestWithInput<InjectedRecordMyInviterInput>)
  //: Promise<void | RecordMyInviterResponse>{ // 取消注释以使用严格的返回类型
  : Promise<any> {
  // TODO: 在此处添加代码
  const currentUser = await hook.internalClient.queries.GetOneAppUser({
    input: {
      id: hook.input.userId
    }
  })
  if (!currentUser) {
    return {
      data: {}, errors: [{
        message: '用户不存在'
      }]
    }
  }
  if (currentUser.data?.data?.AppUser?.id) {
    return {
      data: {}, errors: [{
        message: '已有邀请人'
      }]
    }
  }

  const client = createClient()
  // 增加魔法值
  await client.mutate({
    operationName: 'CreateOnePointRecord',
    input: {
      way:'Invite',
      point: 0,
      userId: hook.input.userId
    }
  })
  await client.mutate({
    operationName: 'CreateOnePointRecord',
    input: {
      way:'Invite',
      point: 0,
      userId: hook.input.inviterId
    }
  })
}
```

使用批量新建创建接口`GetOneAppUser`，勾选AppUser的选项（如下图）

![](https://cdn.nlark.com/yuque/0/2023/png/26463209/1678086006540-734c3ddb-0895-46b8-b9df-244b86a9cab5.png)

最后，需要将sdk/miniapp目录下的文件进行复制，替换到小程序的sdk目录下

![](https://cdn.nlark.com/yuque/0/2023/png/26463209/1678093602704-af69ee84-becc-4d7d-b59a-c25ee19d3505.png)

#### 4.4 管理后台接口 <a href="#wtd1d" id="wtd1d"></a>

**4.4.1 权限控制**

在飞布控制台设置页面——安全——重定向URL中，增加 \[ localhost:4321/oidc/callback ]

Authing 配置

```
clientId 63ae88490f3dff9d1a651cea
secert e8094233545dec4b3b7fc848eb04e306
issuer sail.authing.cn/oidc
```

**4.4.2 用户管理**

AppUser表，批量创建 —— 详情接口 & 分页查询接口

**4.4.3 创作列表**

ArtWork表，批量创建 —— 删除接口 & 分页查询接口 (**需要勾选下面的 Creation-AppUser 关联字段**)

**4.4.4 统计数据**

* 获取总用户和总作品

```graphql
query GetCount {
  user: art_aggregateAppUser {
    _count {
      _all
    }
  }
  creation: art_aggregateCreation {
    _count {
      _all
    }
  }
}
```

* 获取今日作品数量

```graphql
query GetTodayCount($date: DateTime!) {
  creation: art_aggregateCreation(where: {createdAt: {gte: $date}}) {
    _count {
      _all
    }
  }
}
```

* 每日作品数量——查询接口

每日作品数量——自定义数据源

```typescript
import {GraphQLFloat, GraphQLObjectType, GraphQLSchema, GraphQLString, GraphQLList } from 'graphql';
import { prisma } from 'generated/prisma'

export default new GraphQLSchema({
  query: new GraphQLObjectType({
    name: 'Query',
    fields: {
      DailyCreation: {
            type: new GraphQLList(new GraphQLObjectType({
                name: 'DailyCreation',
                fields: {
                  date: {
                        type: GraphQLString,
                    },
                    total: {
                        type: GraphQLFloat
                    }
                },
            })),
            resolve() {
                return prisma.art.queryRaw(`SELECT DATE_FORMAT(createdAt,'%Y-%m-%d') date, COUNT(1) total from Creation GROUP BY date`, {});
            },
        },
    },
  }),
})
```

```graphql
query QueryStatistic {
  data: statistics_DailyCreation {
    date
    total
  }
}
```

### 5. 上线部署 <a href="#nrlpb" id="nrlpb"></a>

[参考手动部署](../bu-shu-yun-wei/shou-dong-bu-shu.md)
