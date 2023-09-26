# Js SDK

`js-client`是基于javascript语言UMD风格的Fireboom API 封装，可以在浏览器环境中使用。

### 引用SDK

#### 配置基本请求头

在index.html文件中增加baseURL配置，自定义访问域名。

{% code title="index.html" lineNumbers="true" %}
```html
<script src="./index.umd.min.js"></script>
```
{% endcode %}

#### 自定义请求域名

```javascript
FBClient.setBaseURL('https://xxx.cc.com')
```

#### 自定义请求头

```javascript
FBClient.setExtraHeaders({ 'X-CUSTOM-HEADER': 'VALUE' })
```

### 查询和变更

#### 查询

```javascript
FBClient.query.GetT({input:{id:1}}).then(console.log)
```

#### 变更

```javascript
FBClient.mutation.CreateT({input:{name:"sss",des:"des"}}).then(console.log)
```

### 实时

#### 实时查询

```typescript
FBClient.subscription.GetT({input:{id:1},liveQuery: true}).then(async res => {
  for await (const v of res) {
    console.log(v)
  }
})
```

#### 订阅

```typescript
FBClient.subscription.GetT({input:{id:1}}).then(async res => {
  for await (const v of res) {
    console.log(v)
  }
})
```

#### 订阅一次

使用`subscribeOnce`运行订阅，这将直接返回订阅响应，而不会推流。适用于SSR场景。

```typescript
FBClient.subscription.GetT({input:{id:1},subscribeOnce: true}, ({ data, error }) => {
  console.log(data)
})
```

### 文件上传

```html
<input type="file" id="fileInput" name="fileInput" >
<script>
    const fileInput = document.getElementById('fileInput');
    fileInput.addEventListener('change', (event) => {
        const files = event.target.files
        if (files == null) return
        // 调用上传函数
        FBClient.uploadFiles({
            provider: 'tengxunyun',
            files: files,
            profile: 'avatar', // （可选）高级配置
            directory:"sss",     // （可选）上传目录
            meta:{              // （可选）meta信息
                postId:"sss"
            }
        })
    });
</script>
```

### 身份验证

身份认证包含两种模式：授权码模式（基于cookie）和隐式模式（基于token）。

#### 授权码模式

**登录**

```typescript
FBClient.login('auth0');
```

**获取用户**

```typescript
FBClient.fetchUser();
```

**退出登录**

```typescript
FBClient.logout({
  logoutOpenidConnectProvider: true,
});
```

#### 隐式模式

**获取Token**

详情见 [#yin-shi-mo-shi](./#yin-shi-mo-shi "mention") ->获取Token

**使用Token**

{% code title="index.js" %}
```javascript
FBClient.setAuthorizationToken('<access_token>')
```
{% endcode %}

