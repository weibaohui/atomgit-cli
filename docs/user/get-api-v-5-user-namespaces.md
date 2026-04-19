# 列出授权用户所有的 Namespace

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/user/namespaces` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-user-namespaces |

## Path Parameters

| - | - | - | - |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/user/namespaces?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
[{"id":1591596,"path":"Cangjie","name":"Cangjie","html_url":"https://atomgit.com/Cangjie","type":"group"},{"id":2196715,"path":"weibaohui","name":"weibaohui","html_url":"https://atomgit.com/weibaohui","type":"user"}]
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
