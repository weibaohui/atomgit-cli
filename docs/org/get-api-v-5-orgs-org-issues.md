# 获取当前用户某个组织的Issues

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/orgs/:org/issues` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-orgs-org-issues |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| org | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/orgs/:org/issues?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":404,"error_code_name":"UN_KNOW","error_message":"404 Group Not Found","trace_id":"199f44e0f4cdebb9d5fd3432a34a8c96"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
