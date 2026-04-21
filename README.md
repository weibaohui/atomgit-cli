# AtomGit CLI

类 GitHub CLI (gh) 的 AtomGit 平台命令行工具，使用 Go 构建。命令行简称为 `atg`。

## 功能特性

- **仓库管理** - 创建、删除、查看、搜索仓库
- **Issue 管理** - 列表、查看、评论
- **Pull Request** - 列表、查看、审查
- **代码浏览** - 分支、提交、标签、发布版本
- **用户交互** - 关注、收藏、叉取
- **组织功能** - 组织成员、仓库管理
- **开发者工具** - Webhook、搜索、API 调用

## 安装

### Make 方式（推荐）

```bash
git clone https://atomgit.com/weibaohui/atomgit-cli.git
cd atomgit-cli
make install
```

### 源码编译

```bash
git clone https://atomgit.com/weibaohui/atomgit-cli.git
cd atomgit-cli
make build
# 或手动编译
go build -o atg .
```

### 安装 Skill

```bash
# 安装到 ~/.claude/skills（Claude Code skills 目录）
atg skills install --claude

# 安装到 ~/.agents/skills
atg skills install
```

安装后 Claude Code 会自动加载 skill，在与 AtomGit/AtomGit 平台交互时会获得帮助提示。

### 使用

```bash
# 登录
atg auth login -t YOUR_TOKEN

# 仓库操作
atg repo create my-project --public
atg repo list
atg repo view owner/repo

# Issue 操作
atg issue list -R owner/repo
atg issue view 123 -R owner/repo

# Pull Request
atg pr list -R owner/repo
atg pr view 456 -R owner/repo

# 代码浏览
atg branch list -R owner/repo
atg commit list -R owner/repo
atg release list -R owner/repo

# 搜索
atg search repos golang
atg search users username

# 直接调用 API
atg api /api/v1/user
```

## 命令参考

| 命令 | 说明 |
|------|------|
| `atg auth` | 认证 (login/logout/status) |
| `atg repo` | 仓库管理 |
| `atg issue` | Issue 管理 |
| `atg pr` | Pull Request |
| `atg branch` | 分支操作 |
| `atg commit` | 提交历史 |
| `atg release` | 发布版本 |
| `atg tag` | 标签管理 |
| `atg fork` | 叉取仓库 |
| `atg star` | 收藏仓库 |
| `atg user` | 用户信息 |
| `atg org` | 组织管理 |
| `atg search` | 搜索 |
| `atg api` | 直接调用 API |
| `atg skills` | 技能管理 |

## 配置

配置文件: `~/.config/atomgit-cli/config.toml`

```toml
base_url = "https://api.atomgit.com"
token = "your-token"
```

环境变量:

- `AMGIT_TOKEN` - 访问令牌
- `AMGIT_BASE_URL` - API 地址 (默认: `https://api.atomgit.com`)

## 开发

```bash
# 依赖安装
go mod download

# 构建
make build        # 构建到 ./bin/atg
make install      # 构建并安装到 ~/bin/atg
make clean        # 清理

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
