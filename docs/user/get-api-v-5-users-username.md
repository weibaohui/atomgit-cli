# 获取一个用户

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/users/:username` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-users-username |

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
curl "https://api.atomgit.com/api/v5/users/:username?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":1003,"error_code_name":"NOT_EXIST","error_message":"用户不存在","trace_id":"da27494d4b8176b8fb814554ed20be17"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
