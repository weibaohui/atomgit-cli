# 创建一个仓库的Tag

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | POST |
| **Endpoint** | `https://api.atomgit.com/api/v5/repos/:owner/:repo/tags` |
| **文档链接** | https://docs.atomgit.com/docs/apis/post-api-v-5-repos-owner-repo-tags |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| owner | string | 是 | 路径参数 |
| repo | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl -X POST "https://api.atomgit.com/api/v5/repos/:owner/:repo/tags" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
{"error_code":1000,"error_code_name":"PARAMETER_ERROR","error_message":"tag_name not blank.","trace_id":"6f50723fac77ea8b270a2f9cffe6e3e5"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
