# 获取某个issue下的操作日志

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/repos/:owner/issues/:number/operate_logs` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-repos-owner-issues-number-operate-logs |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| number | string | 是 | 路径参数 |
| owner | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/repos/:owner/issues/:number/operate_logs?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":1000,"error_code_name":"PARAMETER_ERROR","error_message":"参数 :number 类型错误","trace_id":"3e1de44eca8bfd93cd75cadf83f3832b"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
