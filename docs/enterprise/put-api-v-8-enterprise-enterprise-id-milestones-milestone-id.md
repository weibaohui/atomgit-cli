# 修改企业里程碑

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | PUT |
| **Endpoint** | `https://api.atomgit.com/api/v8/enterprises/:enterprise/milestones/:milestone_id` |
| **文档链接** | https://docs.atomgit.com/docs/apis/put-api-v-8-enterprise-enterprise-id-milestones-milestone-id |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| enterprise | string | 是 | 路径参数 |
| milestone_id | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl -X PUT "https://api.atomgit.com/api/v8/enterprises/:enterprise/milestones/:milestone_id" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
{"error_code":404,"error_code_name":"error","error_message":"404, token not found","trace_id":"ceedac1c5b688779d2de51bb0d723a78"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
