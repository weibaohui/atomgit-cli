# 列出用户 watch 了的仓库

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/users/:username/subscriptions` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-user-username-subscriptions |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| username | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/users/:username/subscriptions?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":404,"error_code_name":"NOT_FOUND","error_message":"user not found","trace_id":"53d3c83b80acafede0230a48cbb4feb3"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
