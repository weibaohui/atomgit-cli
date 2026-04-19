# 获取用户个人动态

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/users/:username/events` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-users-username-events |

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
curl "https://api.atomgit.com/api/v5/users/:username/events?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":404,"error_code_name":"NOT_FOUND","error_message":"Not Found","trace_id":"16f808a857de7e495972fe10a466f256"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
