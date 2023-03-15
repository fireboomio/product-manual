# 微信小程序SDK

## 初始化
在项目入口文件中引入 SDK，并配置后端地址。
```ts
import sdk from './sdk'
sdk.setBaseUrl('http://127.0.0.1:9991')
```

## API 请求
SDK 提供了查询、实时查询和变更三种 API 请求方式。
### 查询接口
sdk.query 下挂载了所有查询接口，调用方式如下：
#### 参数
- data?: ParamType: 查询参数，类型根据 API 定义自动生成

#### 返回值
- ReturnType:  根据 API 定义自动生成
```ts
const resp = await sdk.query.GetUserList(data)
```

### 实时查询接口
将API配置为实时查询后，除了可以通过 query 调用，同时也会在 sdk.liveQuery 下生成实时查询接口，调用方式如下：
#### 参数
- callback: (ReturnType) => void: 回调函数，每次数据更新时会调用此函数
- data?: ParamType: 查询参数，类型根据 API 定义自动生成
#### 返回值
无
```ts
sdk.liveQuery.GetUserList(callback, data)
```
### 变更接口
sdk.mutation 下挂载了所有变更接口，调用方式如下：
#### 参数
- data?: ParamType: 查询参数，类型根据 API 定义自动生成

#### 返回值
- ReturnType: 根据 API 定义自动生成
```ts
const resp = await sdk.mutation.DeleteManyUser(data)
```

## 文件上传
通过 sdk.upload 可以调用文件存储的上传功能
#### 参数
- profileName: string: 后台中文件存储配置的名称
- tempPath: string: 待上传的文件临时地址，该地址可以通过 wx.chooseMedia 等 API 获取
- filaName: string: 上传后的文件名，

#### 返回值
- string: 返回上传后文件的访问地址

```ts
const uploadUrl = await sdk.upload(profileName, tempPath, filaName)
```

## 身份验证
因为小程序没有原生的 cookie 支持，因此小程序中仅提供隐式模式（基于 cookie）登录。

目前只对接了 authing.cn 的身份验证服务，后续会考虑对接更多的身份验证服务。

### 初始化
初始化 auth 模块
#### 参数
以下参数来自会被透传至 Authing SDK，具体逻辑可以参考 Authing SDK 文档。
- appId: string: authing 应用 appId
- host: string: authing 应用域名
- userPoolId: string: authing 用户池id
- poolName: string: authing 用户池唯一标志

#### 返回值
- boolean: 当前是否已登录

```ts
import auth from "./sdk/auth";
const success = await auth.init({
  appId: 'xxx',
  host: 'https://xxx.authing.cn',
  userPoolId: 'xxx',
  poolName: 'xxx',
})
const success = await auth.login()
```
### 登录
触发登录
#### 参数
无

#### 返回值
- boolean: 是否登录成功

```ts
import auth from "./sdk/auth";
const success = await auth.login()
```
