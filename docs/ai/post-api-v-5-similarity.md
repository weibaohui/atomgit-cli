# 句子相似度

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | POST |
| **Endpoint** | `https://api.atomgit.com/api/v5/similarity/` |
| **文档链接** | https://docs.atomgit.com/docs/apis/post-api-v-5-similarity |

## Path Parameters

| - | - | - | - |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl -X POST "https://api.atomgit.com/api/v5/similarity/" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
