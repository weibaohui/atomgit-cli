# 获取授权用户的全部邮箱

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/emails` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-emails |

## Path Parameters

| - | - | - | - |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/emails?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
[{"email":"weibaohui@yeah.net","state":"confirmed"}]
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
