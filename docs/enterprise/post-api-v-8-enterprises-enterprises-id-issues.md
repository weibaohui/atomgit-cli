# 获取企业Issue列表

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | POST |
| **Endpoint** | `https://api.atomgit.com/api/v8/enterprises/:enterprise_id/issues` |
| **文档链接** | https://docs.atomgit.com/docs/apis/post-api-v-8-enterprises-enterprises-id-issues |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| enterprise_id | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl -X POST "https://api.atomgit.com/api/v8/enterprises/:enterprise_id/issues" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
{"error_code":400,"error_code_name":"BAD_REQUEST","error_message":"参数类型错误","trace_id":"f9e6a704b5d6766e134a55d811c67d44"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
