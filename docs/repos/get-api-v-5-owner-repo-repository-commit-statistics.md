# 获取代码量贡献

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/:owner/:repo/repository/commit_statistics` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-owner-repo-repository-commit-statistics |

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
curl "https://api.atomgit.com/api/v5/:owner/:repo/repository/commit_statistics?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
{"error_code":400,"error_code_name":"BAD_REQUEST","error_message":"Request body parsing error, please check if the header content-type:null matches","trace_id":"34fbca2b283a1d4856770b3d60099ba3"}
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
