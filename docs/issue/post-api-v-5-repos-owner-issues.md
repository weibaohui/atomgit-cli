# 创建 Issue

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | POST |
| **Endpoint** | `https://api.atomgit.com/api/v5/repos/:owner/issues` |
| **文档链接** | https://docs.atomgit.com/docs/apis/post-api-v-5-repos-owner-issues |
| **CLI命令** | 未实现 |

## Path Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| owner | string | 是 | 仓库所属空间地址(组织或个人的地址path) |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 (使用环境变量 `$ATOMGIT_TOKEN`) |

## Body Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| repo | string | 是 | 仓库路径 |
| title | string | 是 | Issue标题 |
| body | string | 是 | Issue描述 |
| assignee | string | 否 | Issue负责人的username，多个用英文逗号隔开 |
| milestone | integer | 否 | 里程碑序号 |
| labels | string | 否 | 用逗号分开的标签，名称要求长度在 2-20 之间且非特殊字符。如: bug,performance |
| security_hole | string | 否 | 是否是私有issue(默认为false) |
| template_path | string | 否 | issue模板路径 |
| issue_type | string | 否 | issue类型（企业版支持） |
| issue_severity | string | 否 | issue优先级（企业版支持） |
| custom_fields | object[] | 否 | 自定义字段 |

## 请求示例

```bash
curl -X POST "https://api.atomgit.com/api/v5/repos/:owner/issues" \
  -H "Authorization: token $ATOMGIT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "repo": "owner/repo",
    "title": "Bug: Login issue",
    "body": "用户无法登录系统",
    "assignee": "zhangsan",
    "labels": "bug",
    "milestone": 1
  }'
```

## 响应示例

```json
{
  "id": 123,
  "number": 1,
  "html_url": "https://atomgit.com/owner/repo/issues/1",
  "state": "open",
  "title": "Bug: Login issue",
  "body": "用户无法登录系统",
  "user": {
    "id": 1,
    "login": "zhangsan",
    "avatar_url": "https://..."
  },
  "assignee": null,
  "labels": [],
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

## 相关 CLI 命令

目前未实现，可使用 `atomgit api` 命令直接调用：

```bash
atomgit api POST /repos/:owner/issues --body '{"repo":"owner/repo","title":"Bug","body":"desc"}'
```
