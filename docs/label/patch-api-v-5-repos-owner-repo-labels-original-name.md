# 更新一个仓库的任务标签

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | PATCH |
| **Endpoint** | `https://api.atomgit.com/api/v5/repos/:owner/:repo/labels/:original_name` |
| **文档链接** | https://docs.atomgit.com/docs/apis/patch-api-v-5-repos-owner-repo-labels-original-name |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| original_name | string | 是 | 路径参数 |
| owner | string | 是 | 路径参数 |
| repo | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl -X PATCH "https://api.atomgit.com/api/v5/repos/:owner/:repo/labels/:original_name" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
{"error_code":404,"error_code_name":"error","error_message":"404, token not found","trace_id":"be8cf3019dcf4677b962b160d27cf777"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
