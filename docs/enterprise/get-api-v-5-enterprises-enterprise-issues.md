# 获取某个企业的所有Issues

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/enterprises/:enterprise/issues` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-enterprises-enterprise-issues |

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
curl "https://api.atomgit.com/api/v5/enterprises/:enterprise/issues?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":404,"error_code_name":"UN_KNOW","error_message":"404 Group Not Found","trace_id":"c75ef68d2221c5cdb7545e3dd44deafe"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
