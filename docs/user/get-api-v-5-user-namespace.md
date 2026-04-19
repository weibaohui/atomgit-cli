# 获取授权用户的一个 Namespace

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/user/namespace` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-user-namespace |

## Path Parameters

| - | - | - | - |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/user/namespace?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":400,"error_code_name":"BAD_REQUEST","error_message":"Required request parameter 'path' for method parameter type String is not present","trace_id":"872fbe5939f05ad972830675cc3025ec"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
