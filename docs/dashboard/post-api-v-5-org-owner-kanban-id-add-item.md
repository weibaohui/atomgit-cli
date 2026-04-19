# 添加Issue或者Pull Request到看板

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | POST |
| **Endpoint** | `https://api.atomgit.com/api/v5/org/:owner/kanban/:kanban_id/add_item` |
| **文档链接** | https://docs.atomgit.com/docs/apis/post-api-v-5-org-owner-kanban-id-add-item |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| kanban_id | string | 是 | 路径参数 |
| owner | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl -X POST "https://api.atomgit.com/api/v5/org/:owner/kanban/:kanban_id/add_item" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
{"error_code":400,"error_code_name":"BAD_REQUEST","error_message":"参数类型错误","trace_id":"6d284e23ce0edc1073b0ea1775d748ca"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
