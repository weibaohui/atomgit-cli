# 删除企业成员

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | DELETE |
| **Endpoint** | `https://api.atomgit.com/api/v8/enterprises/:enterprise/members/:username` |
| **文档链接** | https://docs.atomgit.com/docs/apis/delete-api-v-8-enterprises-enterprise-members-usernames |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| enterprise | string | 是 | 路径参数 |
| username | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl -X DELETE "https://api.atomgit.com/api/v8/enterprises/:enterprise/members/:username" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
{"error_code":404,"error_code_name":"error","error_message":"404, token not found","trace_id":"1311a059f1dce3fc2c4d0bd7cf04fec0"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
