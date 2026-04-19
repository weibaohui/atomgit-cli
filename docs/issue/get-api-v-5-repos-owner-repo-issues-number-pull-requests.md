# 获取 issue 关联的 pull requests

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/repos/:owner/:repo/issues/:number/pull_requests` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-repos-owner-repo-issues-number-pull-requests |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| number | string | 是 | 路径参数 |
| owner | string | 是 | 路径参数 |
| repo | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/repos/:owner/:repo/issues/:number/pull_requests?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":1000,"error_code_name":"PARAMETER_ERROR","error_message":"参数 :number 类型错误","trace_id":"bb34e9f7604211895d81edb77022906e"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
