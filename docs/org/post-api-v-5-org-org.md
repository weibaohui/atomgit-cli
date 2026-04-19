# 创建组织

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | POST |
| **Endpoint** | `https://api.atomgit.com/api/v5/orgs` |
| **文档链接** | https://docs.atomgit.com/docs/apis/post-api-v-5-org-org |

## Path Parameters

| - | - | - | - |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl -X POST "https://api.atomgit.com/api/v5/orgs" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
{"error_code":404,"error_code_name":"error","error_message":"404, token not found","trace_id":"36491e607340453c3dc8bd596d02b910"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
