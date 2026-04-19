# 创建企业里程碑

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | POST |
| **Endpoint** | `https://api.atomgit.com/api/v8/enterprises/:enterprise/milestones` |
| **文档链接** | https://docs.atomgit.com/docs/apis/post-api-v-8-enterprise-enterprise-id-milestones |

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
curl -X POST "https://api.atomgit.com/api/v8/enterprises/:enterprise/milestones" \\n  -H "Authorization: token $ATOMGIT_TOKEN" \\n  -H "Content-Type: application/json" \\n  -d '{"key": "value"}'
```


## 响应示例

```json
{"error_code":1000,"error_code_name":"PARAMETER_ERROR","error_message":"title can not be blank","trace_id":"c5009a3c3deb1cd2f16a0766566c7804"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
