# 更新仓库设置

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | PATCH |
| **Endpoint** | `https://api.atomgit.com/api/v5/repos/:owner/:repo` |
| **文档链接** | https://docs.atomgit.com/docs/apis/patch-api-v-5-repos-owner-repo |

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
curl -X PATCH "https://api.atomgit.com/api/v5/repos/:owner/:repo" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
{"error_code":404,"error_code_name":"error","error_message":"404, token not found","trace_id":"51046d040633f63f746e33729df056af"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
