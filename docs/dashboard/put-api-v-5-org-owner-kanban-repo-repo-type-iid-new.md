# 更新Issue或者Pull Request关联的看板

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | PUT |
| **Endpoint** | `https://api.atomgit.com/api/v5/org/:owner/kanban/repo/:repo/:type/:iid` |
| **文档链接** | https://docs.atomgit.com/docs/apis/put-api-v-5-org-owner-kanban-repo-repo-type-iid-new |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| iid | string | 是 | 路径参数 |
| owner | string | 是 | 路径参数 |
| repo | string | 是 | 路径参数 |
| type | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl -X PUT "https://api.atomgit.com/api/v5/org/:owner/kanban/repo/:repo/:type/:iid" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
{"error_code":1000,"error_code_name":"PARAMETER_ERROR","error_message":"must not be blank","trace_id":"f6e9cc885c9a728dc3d9157d2e72cf43"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
