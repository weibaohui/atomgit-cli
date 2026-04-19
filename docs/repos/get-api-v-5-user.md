# 获取授权用户的资料

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/user` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-user |

## Path Parameters

| - | - | - | - |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/user?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"avatar_url":"https://cdn-img.gitcode.com/ec/fa/e2cfa4bd75bb785f72e0b35a6455666d2e1bed31e5c6b1a8f611a0acf56e3e4f.png","followers_url":"https://api.atomgit.com/api/v5/users/weibaohui/followers","html_url":"https://atomgit.com/weibaohui","id":"66b0d2441cdfc50d45c0e886","login":"weibaohui","name":"weibaohui","type":"User","url":"https://api.atomgit.com/api/v5/user/weibaohui","bio":"","blog":"https://github.com/weibaohui","company":"","email":"weibaohui@yeah.net","followers":0,"following":3,"top_languages":["Go","C#","TSX","TypeScript","HTML+Razor"],"location":"","github_account":"https://github.com/weibaohui","website":"https://github.com/weibaohui","description":""}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
