# 修改企业成员权限

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | PUT |
| **Endpoint** | `https://api.atomgit.com/api/v8/enterprises/:enterprise/members/:username` |
| **文档链接** | https://docs.atomgit.com/docs/apis/put-api-v-8-enterprises-enterprise-members-username |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| enterprise | string | 是 | 路径参数 |
| username | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl -X PUT "https://api.atomgit.com/api/v8/enterprises/:enterprise/members/:username" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
{"error_code":1000,"error_code_name":"PARAMETER_ERROR","error_message":"must not be blank","trace_id":"68692b0c30f81dc05f8b0d79a4426cf0"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
