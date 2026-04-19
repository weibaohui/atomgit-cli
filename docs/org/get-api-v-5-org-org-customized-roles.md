# 获取组织自定义角色

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/orgs/:org/customized_roles` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-org-org-customized-roles |

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
curl "https://api.atomgit.com/api/v5/orgs/:org/customized_roles?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":400,"error_code_name":"Unknown","error_message":"org not found","trace_id":"51a56c52232c08ad4fe44dc22bd64ac5"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
