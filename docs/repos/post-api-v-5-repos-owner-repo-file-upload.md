# 上传文件

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | POST |
| **Endpoint** | `https://api.atomgit.com/api/v5/repos/:owner/:repo/file/upload` |
| **文档链接** | https://docs.atomgit.com/docs/apis/post-api-v-5-repos-owner-repo-file-upload |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| owner | string | 是 | 路径参数 |
| repo | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl -X POST "https://api.atomgit.com/api/v5/repos/:owner/:repo/file/upload" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
{"error_code":400,"error_code_name":"BAD_REQUEST","error_message":"Content type 'application/json' not supported","trace_id":"edadf84c97a643a8a496972f676ed8d3"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
