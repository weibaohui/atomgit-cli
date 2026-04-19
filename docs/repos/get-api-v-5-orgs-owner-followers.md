# 列出指定组织的所有关注者

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/orgs/:owner/followers` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-orgs-owner-followers |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| owner | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/orgs/:owner/followers?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":404,"error_code_name":"UN_KNOW","error_message":"404 Namespace Not Found","trace_id":"9e647d5ccfe5df27080397706dbb212d"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
