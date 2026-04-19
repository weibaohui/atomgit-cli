# 列出授权用户 watch 了的仓库

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/user/subscriptions` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-user-subscriptions |

## Path Parameters

| - | - | - | - |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/user/subscriptions?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
[{"id":4165259,"full_name":"weibaohui/k8m","human_name":"weibaohui / k8m","url":"https://api.atomgit.com/api/v5/repos/weibaohui/k8m","namespace":{"id":2196715,"type":"user","name":"weibaohui","path":"weibaohui","html_url":"https://atomgit.com/weibaohui"},"path":"k8m","name":"k8m","description":"一款轻量级、跨平台的 Mini Kubernetes AI Dashboard，支持大模型+智能体+MCP(支持设置操作权限)，集成多集群管理、智能分析、实时异常检测等功能，支持多架构并可单文件部署，助力高效集群管理与运维优化。","status":"开始","ssh_url_to_repo":"git@gitcode.com:weibaohui/k8m.git","http_url_to_repo":"https://atomgit.com/weibaohui/k8m.git","web_url":"https://atomgit.com/weibaohui/k8m","created_at":"2024-10-11T12:11:07.317+08:00","updated_at":"2025-05-08T18:13:06.790+08:00","homepage":"https://atomgit.com/weibaohui/k8m","members":["weibaohui"],"forks_count":2,"stargazers_count":5,"relation":"master","permission":{"pull":true,"push":true,"admin":true},"internal":false,"open_issues_count":1,"has_issue":true,"has_issues":true,"watchers_count":2,"enterprise":{"id":2196715,"path":"weibaohui","html_url":"https://atomgit.com/weibaohui","type":"user"},"default_branch":"main","fork":false,"pushed_at":"2026-04-19T02:31:01.047+08:00","owner":{"id":"898013","login":"weibaohui","name":"weibaohui","avatar_url":"https://cdn-img.gitcode.com/ec/fa/e2cfa4bd75bb785f72e0b35a6455666d2e1bed31e5c6b1a8f611a0acf56e3e4f.png","html_url":"https://atomgit.com/weibaohui","type":"User","url":"https://api.atomgit.com/api/v5/users/weibaohui"},"issue_template_source":"project","language":"Go","has_wiki":false,"pull_requests_enabled":true,"project_creator":"weibaohui","private":false,"public":true}]
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
