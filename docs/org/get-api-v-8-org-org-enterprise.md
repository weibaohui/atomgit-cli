# 获取组织关联的企业

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v8/org/:org/enterprise` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-8-org-org-enterprise |

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
curl "https://api.atomgit.com/api/v8/org/:org/enterprise?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":400,"error_code_name":"Unknown","error_message":"org: :org not found","trace_id":"8474ececf27f5e4209eadcc57fe56b9a"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
