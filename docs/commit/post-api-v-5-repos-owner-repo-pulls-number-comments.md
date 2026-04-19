# 提交pull request 评论

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | POST |
| **Endpoint** | `https://api.atomgit.com/api/v5/repos/:owner/:repo/pulls/:number/comments` |
| **文档链接** | https://docs.atomgit.com/docs/apis/post-api-v-5-repos-owner-repo-pulls-number-comments |

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
curl -X POST "https://api.atomgit.com/api/v5/repos/:owner/:repo/pulls/:number/comments" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
{"error_code":1000,"error_code_name":"PARAMETER_ERROR","error_message":"must not be null","trace_id":"ecc39099276f9a7c96e0a94c11b3d8f5"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
