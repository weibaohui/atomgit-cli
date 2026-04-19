# 获取企业某个Issue所有标签

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/enterprises/:enterprise/issues/:issue_id/labels` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-enterprises-enterprise-issues-issue-id-labels |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| enterprise | string | 是 | 路径参数 |
| issue_id | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/enterprises/:enterprise/issues/:issue_id/labels?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":400,"error_code_name":"BAD_REQUEST","error_message":"参数类型错误","trace_id":"23083706c27cedb747a00c276759dbdc"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
