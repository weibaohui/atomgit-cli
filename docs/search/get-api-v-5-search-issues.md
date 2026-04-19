# 搜索 Issues

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/search/issues` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-search-issues |

## Path Parameters

| - | - | - | - |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/search/issues?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":1000,"error_code_name":"PARAMETER_ERROR","error_message":"must not be blank","trace_id":"e86102e813c18b82f1fb5c7819fd16cf"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
