# 获取组织成员详情

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/orgs/:org/members/:username` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-orgs-org-members-username |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| org | string | 是 | 路径参数 |
| username | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/orgs/:org/members/:username?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":404,"error_code_name":"NOT_FOUND","error_message":"用户 :username 未找到","trace_id":"8c0fe5b63fb7b0e2bcb4667ff1364ea7"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
