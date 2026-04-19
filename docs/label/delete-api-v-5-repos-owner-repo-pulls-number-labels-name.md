# 删除 Pull Request标签

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | DELETE |
| **Endpoint** | `https://api.atomgit.com/api/v5/repos/:owner/:repo/pulls/:number/labels/:name` |
| **文档链接** | https://docs.atomgit.com/docs/apis/delete-api-v-5-repos-owner-repo-pulls-number-labels-name |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| name | string | 是 | 路径参数 |
| number | string | 是 | 路径参数 |
| owner | string | 是 | 路径参数 |
| repo | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl -X DELETE "https://api.atomgit.com/api/v5/repos/:owner/:repo/pulls/:number/labels/:name" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
{"error_code":1000,"error_code_name":"PARAMETER_ERROR","error_message":"参数 :number 类型错误","trace_id":"6f281c7ed25121162f24ff44b5a0d4fd"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
