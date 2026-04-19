# Commits对比

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/repos/:owner/:repo/compare/:base...:head` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-repos-owner-repo-compare-base-head |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| base | string | 是 | 路径参数 |
| head | string | 是 | 路径参数 |
| owner | string | 是 | 路径参数 |
| repo | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/repos/:owner/:repo/compare/:base...:head?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":404,"error_code_name":"UN_KNOW","error_message":"Project not found:%3Aowner%2F:repo","trace_id":"08f53f7531a7e199fc5886545ab7bfb2"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
