# 更新授权用户所管理的组织资料

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | PATCH |
| **Endpoint** | `https://api.atomgit.com/api/v5/orgs/:org` |
| **文档链接** | https://docs.atomgit.com/docs/apis/patch-api-v-5-orgs-org |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| org | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl -X PATCH "https://api.atomgit.com/api/v5/orgs/:org" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
{"error_code":404,"error_code_name":"error","error_message":"404, token not found","trace_id":"b9cd6761dba6d6858952b220afbd8887"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
