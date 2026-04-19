# 获取企业issue状态

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/enterprises/:enterprise/issue/statuses` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-enterprises-enterprise-issue-statuses |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| enterprise | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/enterprises/:enterprise/issue/statuses?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":400,"error_code_name":"Unknown","error_message":"该组织不存在","trace_id":"a36a5b41075b6c9e0c1f315e3f6c22fe"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
