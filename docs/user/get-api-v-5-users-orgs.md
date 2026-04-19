# 列出授权用户所属的组织

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/users/orgs` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-users-orgs |

## Path Parameters

| - | - | - | - |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/users/orgs?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
[{"avatar_url":"https://cdn-img.gitcode.com/cf/bf/349c8fbf998f96f60e10d8918239dfe678f9e78cdc4d07701efdd591ebbed7cb.jpg?time1715738758513","description":"仓颉编程语言是一款面向全场景智能的新一代编程语言，主打智能化、全场景、高性能、强安全。","id":1591596,"login":"Cangjie","path":"Cangjie","name":"Cangjie","url":"https://atomgit.com/Cangjie","my_role":{"id":1554521,"access_level":15,"source_id":1591596,"source_type":"Namespace","user_id":898013,"notification_level":3,"created_at":"8/24/24 8:47 AM","updated_at":"9/12/24 1:43 AM","created_by_id":501265,"limited":false}}]
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
