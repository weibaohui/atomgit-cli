# 设置Issue关联的分支

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | PUT |
| **Endpoint** | `https://api.atomgit.com/api/v5/repos/:owner/:repo/issues/:number/related_branches` |
| **文档链接** | https://docs.atomgit.com/docs/apis/put-api-v-5-repos-owner-repo-issues-number-related-branches |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| number | string | 是 | 路径参数 |
| owner | string | 是 | 路径参数 |
| repo | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl -X PUT "https://api.atomgit.com/api/v5/repos/:owner/:repo/issues/:number/related_branches" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
{"error_code":1000,"error_code_name":"PARAMETER_ERROR","error_message":"参数 :number 类型错误","trace_id":"064a2db377a4d3fc6f0ebe20ab29b6d9"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
