# 新建保护分支规则

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | PUT |
| **Endpoint** | `https://api.atomgit.com/api/v5/repos/:owner/:repo/branches/setting/new` |
| **文档链接** | https://docs.atomgit.com/docs/apis/put-api-v-5-repos-owner-repo-branches-setting-new |

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
curl -X PUT "https://api.atomgit.com/api/v5/repos/:owner/:repo/branches/setting/new" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
{"error_code":1000,"error_code_name":"PARAMETER_ERROR","error_message":"must not be blank","trace_id":"e12c96bbe16c002d6df9df3547d0a074"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
