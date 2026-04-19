# AtomGit CLI (Go版本) 设计文档

## 1. 项目概述

**项目名称:** atomgit-cli (amg)
**项目类型:** 命令行工具 (Go)
**核心功能:** GitHub-like CLI for AtomGit platform
**目标用户:** 开发者

## 2. 技术栈

| 功能 | Go 库 | 说明 |
|------|-------|------|
| CLI框架 | github.com/spf13/cobra@latest | 知名开源CLI框架 |
| 配置管理 | github.com/spf13/viper@latest | 配合cobra使用 |
| HTTP客户端 | net/http (内置) | Go标准库 |
| Spinner | github.com/briandowns/spinner@latest | ora的Go实现 |
| 配置文件 | toml | Viper原生支持 |

## 3. 项目结构

```
go/
├── cmd/
│   ├── root.go        # 根命令
│   ├── auth.go        # auth命令组 (login/logout/status/token)
│   ├── repo.go        # repo命令组 (create/delete/list/view)
│   ├── issue.go       # issue命令组 (list/view)
│   ├── pr.go          # pr命令组 (list/view)
│   └── api.go         # api命令
├── internal/
│   ├── config/
│   │   └── config.go  # 配置管理 (Viper)
│   └── httpclient/
│       └── client.go  # HTTP客户端
├── main.go            # 入口
├── go.mod
└── go.sum
```

## 4. 命令设计

### 4.1 根命令 (root)
- 全局选项: 无
- 子命令: auth, repo, issue, pr, api

### 4.2 auth 命令组
| 子命令 | 功能 |
|--------|------|
| login | 登录 (交互式或 -t token) |
| logout | 登出 |
| status | 显示认证状态 |
| token | 显示token (脱敏) |

### 4.3 repo 命令组
| 子命令 | 功能 |
|--------|------|
| create \<name\> | 创建仓库 |
| delete \<repo\> | 删除仓库 |
| list | 列出仓库 |
| view [owner/repo] | 查看仓库详情 |

### 4.4 issue 命令组
| 子命令 | 功能 |
|--------|------|
| list | 列出issues |
| view \<number\> | 查看issue详情 |

### 4.5 pr 命令组
| 子命令 | 功能 |
|--------|------|
| list | 列出PRs |
| view \<number\> | 查看PR详情 |

### 4.6 api 命令
- `api <endpoint>` - 发起API请求
- 选项: `-X method`, `-d data`

## 5. 配置设计

### 5.1 配置文件路径
- Linux/Mac: `~/.config/atomgit-cli/config.toml`
- Windows: `%APPDATA%/atomgit-cli/config.toml`

### 5.2 配置项
```toml
base_url = "https://api.atomgit.com"
token = ""
```

### 5.3 环境变量
- `ATOMGIT_BASE_URL` - API基础URL
- `ATOMGIT_TOKEN` - 访问令牌

## 6. API 设计

### 6.1 基础URL
- 默认: `https://api.atomgit.com`
- 自动添加 `/api/v5` 前缀

### 6.2 认证
- Header: `Authorization: Bearer <token>`
- User-Agent: `atomgit-cli/0.1.0`

## 7. 构建

```bash
go build -o atomgit .
```

## 8. 错误处理

- HTTP错误返回格式: `HTTP {status}: {statusText}\n{body}`
- 配置错误: 使用viper默认值降级
- 网络错误: 返回具体错误信息
