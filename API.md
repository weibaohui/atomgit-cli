# AtomGit CLI API 文档

## 概述

AtomGit CLI 提供了与 AtomGit 平台交互的命令行工具，类似于 GitHub CLI (gh)。

**基础URL:** `https://api.atomgit.com/api/v5`

**认证:** 通过 `ATOMGIT_TOKEN` 环境变量或 `atomgit auth login` 命令

---

## 命令列表

### 1. auth - 认证

```bash
atomgit auth login [-t <token>]     # 登录
atomgit auth logout                # 登出
atomgit auth status                # 查看认证状态
atomgit auth token                # 显示token (脱敏)
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| login | 登录 | - |
| logout | 登出 | - |
| status | 查看认证状态 | GET /user |
| token | 显示token | - |

---

### 2. repo - 仓库

```bash
atomgit repo create <name> [-d <desc>] [--public|--private]  # 创建仓库
atomgit repo delete <owner/repo> [--yes]                      # 删除仓库
atomgit repo list                                            # 列出仓库
atomgit repo view <owner/repo>                               # 查看仓库详情
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| create | 创建仓库 | POST /user/repos |
| delete | 删除仓库 | DELETE /repos/{owner}/{repo} |
| list | 列出仓库 | GET /user/repos |
| view | 查看仓库详情 | GET /repos/{owner}/{repo} |

---

### 3. issue - 问题

```bash
atomgit issue list -R <owner/repo> [-s <state>] [-L <limit>]  # 列出问题
atomgit issue view <number> -R <owner/repo>                   # 查看问题
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| list | 列出问题 | GET /repos/{owner}/{repo}/issues |
| view | 查看问题详情 | GET /repos/{owner}/{repo}/issues/{number} |

**参数:**
- `-R, --repo`: 仓库路径 (必需)
- `-s, --state`: 状态过滤 (open, closed, all)
- `-L, --limit`: 返回数量限制

---

### 4. pr - 拉取请求

```bash
atomgit pr list -R <owner/repo> [-s <state>] [-L <limit>]  # 列出PR
atomgit pr view <number> -R <owner/repo>                   # 查看PR
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| list | 列出PR | GET /repos/{owner}/{repo}/pulls |
| view | 查看PR详情 | GET /repos/{owner}/{repo}/pulls/{number} |

---

### 5. branch - 分支

```bash
atomgit branch list -R <owner/repo>   # 列出分支
atomgit branch view <branch> -R <owner/repo>  # 查看分支
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| list | 列出分支 | GET /repos/{owner}/{repo}/branches |
| view | 查看分支详情 | GET /repos/{owner}/{repo}/branches/{branch} |

---

### 6. commit - 提交

```bash
atomgit commit list -R <owner/repo> [--sha <sha>] [--path <path>]  # 列出提交
atomgit commit view <sha> -R <owner/repo>                         # 查看提交
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| list | 列出提交 | GET /repos/{owner}/{repo}/commits |
| view | 查看提交详情 | GET /repos/{owner}/{repo}/commits/{sha} |

**参数:**
- `-R, --repo`: 仓库路径 (必需)
- `--sha`: SHA或分支名
- `--path`: 路径过滤

---

### 7. tag - 标签

```bash
atomgit tag list -R <owner/repo>   # 列出标签
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| list | 列出标签 | GET /repos/{owner}/{repo}/tags |

---

### 8. release - 发布

```bash
atomgit release list -R <owner/repo>   # 列出发布
atomgit release view <tag> -R <owner/repo>  # 查看发布
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| list | 列出发布 | GET /repos/{owner}/{repo}/releases |
| view | 查看发布详情 | GET /repos/{owner}/{repo}/releases/{tag} |

---

### 9. label - 标签

```bash
atomgit label list -R <owner/repo>   # 列出标签
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| list | 列出标签 | GET /repos/{owner}/{repo}/labels |

---

### 10. milestone - 里程碑

```bash
atomgit milestone list -R <owner/repo>   # 列出里程碑
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| list | 列出里程碑 | GET /repos/{owner}/{repo}/milestones |

---

### 11. event - 事件

```bash
atomgit event list -R <owner/repo>   # 列出事件
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| list | 列出事件 | GET /repos/{owner}/{repo}/events |

---

### 12. user - 用户

```bash
atomgit user info [username]          # 获取用户信息
atomgit user followers [username]    # 列出粉丝
atomgit user following [username]    # 列出关注
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| info | 用户信息 | GET /user 或 GET /users/{username} |
| followers | 粉丝列表 | GET /user/followers 或 GET /users/{username}/followers |
| following | 关注列表 | GET /user/following 或 GET /users/{username}/following |

---

### 13. org - 组织

```bash
atomgit org info <org>      # 组织信息
atomgit org members <org>   # 成员列表
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| info | 组织信息 | GET /orgs/{org} |
| members | 成员列表 | GET /orgs/{org}/members |

---

### 14. search - 搜索

```bash
atomgit search repos <query> [-L <limit>]   # 搜索仓库
atomgit search users <query> [-L <limit>]   # 搜索用户
atomgit search code <query> [-L <limit>]     # 搜索代码
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| repos | 搜索仓库 | GET /search/repositories?q={query} |
| users | 搜索用户 | GET /search/users?q={query} |
| code | 搜索代码 | GET /search/code?q={query} |

---

### 15. fork - Fork

```bash
atomgit fork create <repo>   # Fork仓库
atomgit fork list <repo>    # 列出Forks
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| create | Fork仓库 | POST /repos/{owner}/{repo}/forks |
| list | 列出Forks | GET /repos/{owner}/{repo}/forks |

---

### 16. star - Star

```bash
atomgit star list -R <owner/repo>   # 列出Stargazers
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| list | 列出Stargazers | GET /repos/{owner}/{repo}/stargazers |

---

### 17. subscriber - 订阅者

```bash
atomgit subscriber list -R <owner/repo>   # 列出订阅者
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| list | 列出订阅者 | GET /repos/{owner}/{repo}/subscribers |

---

### 18. hook - Webhook

```bash
atomgit hook list -R <owner/repo>                    # 列出webhooks
atomgit hook create -R <owner/repo> --url <url>     # 创建webhook
atomgit hook view <id> -R <owner/repo>              # 查看webhook
atomgit hook update <id> -R <owner/repo> --url <url> # 更新webhook
atomgit hook delete <id> -R <owner/repo>            # 删除webhook
atomgit hook test <id> -R <owner/repo>              # 测试webhook
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| list | 列出webhooks | GET /repos/{owner}/{repo}/hooks |
| create | 创建webhook | POST /repos/{owner}/{repo}/hooks |
| view | 查看webhook | GET /repos/{owner}/{repo}/hooks/{id} |
| update | 更新webhook | PATCH /repos/{owner}/{repo}/hooks/{id} |
| delete | 删除webhook | DELETE /repos/{owner}/{repo}/hooks/{id} |
| test | 测试webhook | POST /repos/{owner}/{repo}/hooks/{id}/tests |

---

### 19. language - 语言

```bash
atomgit language list -R <owner/repo>   # 列出编程语言
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| list | 列出语言 | GET /repos/{owner}/{repo}/languages |

---

### 20. contributor - 贡献者

```bash
atomgit contributor list -R <owner/repo>   # 列出贡献者
atomgit contributor stats -R <owner/repo>    # 贡献者统计
```

| 子命令 | 说明 | API端点 |
|--------|------|---------|
| list | 贡献者列表 | GET /repos/{owner}/{repo}/contributors |
| stats | 贡献者统计 | GET /repos/{owner}/{repo}/contributors-statistic |

---

### 21. api - 原始API请求

```bash
atomgit api <endpoint> [-X <method>] [-d <json>]
```

直接调用 AtomGit API，用于测试或未封装的API。

**参数:**
- `-X, --method`: HTTP方法 (默认 GET)
- `-d, --data`: JSON格式的请求体

**示例:**
```bash
atomgit api /user
atomgit api /repos/weibaohui/atomgit-cli -X GET
atomgit api /user/repos -X POST -d '{"name":"test","private":true}'
```

---

## 环境变量

| 变量 | 说明 |
|------|------|
| ATOMGIT_TOKEN | 访问令牌 |
| ATOMGIT_BASE_URL | API基础URL (默认: https://api.atomgit.com) |

---

## 配置文件

配置存储在: `~/.config/atomgit-cli/config.toml`

---

## 安装

```bash
go install github.com/weibaohui/atomgit-cli@latest
```

设置别名:
```bash
alias ag='atomgit'
```
