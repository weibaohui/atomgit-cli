# 获取企业里程碑详情

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v8/enterprises/:enterprise/milestones/:milestone_id` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-8-enterprise-enterprise-id-milestones-milestone-id |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| enterprise | string | 是 | 路径参数 |
| milestone_id | string | 是 | 路径参数 |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v8/enterprises/:enterprise/milestones/:milestone_id?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":403,"error_code_name":"UN_KNOW","error_message":"CH.00000403 apig token has not permission to request url","trace_id":"777fb6c373f0de082b221fe4bddb52ad"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
