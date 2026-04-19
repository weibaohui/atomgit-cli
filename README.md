# AtomGit CLI

类 GitHub CLI (gh) 的 AtomGit 平台命令行工具，使用 Go 构建。

## 功能特性

- **仓库管理** - 创建、删除、查看、搜索仓库
- **Issue 管理** - 列表、查看、评论
- **Pull Request** - 列表、查看、审查
- **代码浏览** - 分支、提交、标签、发布版本
- **用户交互** - 关注、收藏、叉取
- **组织功能** - 组织成员、仓库管理
- **开发者工具** - Webhook、搜索、API 调用

## 安装

### 源码编译

```bash
git clone https://github.com/weibaohui/atomgit-cli.git
cd atomgit-cli
go build -o atomgit .
sudo mv atomgit /usr/local/bin/
```

### 使用

```bash
# 登录
atomgit auth login -t YOUR_TOKEN

# 仓库操作
atomgit repo create my-project --public
atomgit repo list
atomgit repo view owner/repo

# Issue 操作
atomgit issue list -R owner/repo
atomgit issue view 123 -R owner/repo

# Pull Request
atomgit pr list -R owner/repo
atomgit pr view 456 -R owner/repo

# 代码浏览
atomgit branch list -R owner/repo
atomgit commit list -R owner/repo
atomgit release list -R owner/repo

# 搜索
atomgit search repos golang
atomgit search users username

# 直接调用 API
atomgit api /api/v1/user
```

## 命令参考

| 命令 | 说明 |
|------|------|
| `atomgit auth` | 认证 (login/logout/status) |
| `atomgit repo` | 仓库管理 |
| `atomgit issue` | Issue 管理 |
| `atomgit pr` | Pull Request |
| `atomgit branch` | 分支操作 |
| `atomgit commit` | 提交历史 |
| `atomgit release` | 发布版本 |
| `atomgit tag` | 标签管理 |
| `atomgit fork` | 叉取仓库 |
| `atomgit star` | 收藏仓库 |
| `atomgit user` | 用户信息 |
| `atomgit org` | 组织管理 |
| `atomgit search` | 搜索 |
| `atomgit api` | 直接调用 API |

## 配置

配置文件: `~/.config/atomgit-cli/config.toml`

```toml
base_url = "https://api.atomgit.com"
token = "your-token"
```

环境变量:

- `ATOMGIT_TOKEN` - 访问令牌
- `ATOMGIT_BASE_URL` - API 地址 (默认: `https://api.atomgit.com`)

## 开发

```bash
# 依赖安装
go mod download

# 构建
go build -o atomgit .

# 测试
go test ./...
```

## 项目结构

```
cmd/              # 命令实现
├── root.go       # 根命令
├── auth.go       # 认证
├── repo.go       # 仓库
├── issue.go      # Issue
├── pr.go         # Pull Request
├── branch.go     # 分支
├── commit.go     # 提交
├── release.go    # 发布
├── fork.go       # 叉取
├── star.go       # 收藏
├── search.go     # 搜索
├── api.go        # API 调用
└── ...
```

## 技术栈

- Go 1.21+
- [cobra](https://github.com/spf13/cobra) - CLI 框架
- [viper](https://github.com/spf13/viper) - 配置管理
- [spinner](https://github.com/briandowns/spinner) - 加载动画

## 许可证

MIT
