# 获取企业Issue自定义字段列表

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v8/enterprises/:enterprise_id/issue_extend_field` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-8-enterprises-enterprises-id-issue-extend-field |

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
curl "https://api.atomgit.com/api/v8/enterprises/:enterprise_id/issue_extend_field?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":400,"error_code_name":"BAD_REQUEST","error_message":"参数类型错误","trace_id":"dfa9f1d16abd379daf282684b895ff9a"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
