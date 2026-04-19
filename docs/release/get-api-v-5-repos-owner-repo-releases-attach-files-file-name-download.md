# 下载仓库release附件

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/repos/:owner/:repo/releases/:tag/attach_files/:file_name/download` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-repos-owner-repo-releases-attach-files-file-name-download |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| file_name | string | 是 | 路径参数 |
| owner | string | 是 | 路径参数 |
| repo | string | 是 | 路径参数 |
| tag | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/repos/:owner/:repo/releases/:tag/attach_files/:file_name/download?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":403,"error_code_name":"FORBIDDEN","error_message":"Forbidden","trace_id":"6fa63f6d1a81a3542a2aaeb5d9dfa657"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
