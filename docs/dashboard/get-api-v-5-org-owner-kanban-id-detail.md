# 获取单个看板详情

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/org/:owner/kanban/:kanban_id/detail` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-org-owner-kanban-id-detail |

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
curl "https://api.atomgit.com/api/v5/org/:owner/kanban/:kanban_id/detail?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":400,"error_code_name":"BAD_REQUEST","error_message":"参数类型错误","trace_id":"bc37eadbdb446419d8fc14ec61c35180"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
