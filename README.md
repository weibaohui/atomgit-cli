# AtomGit CLI

类 GitHub CLI (gh) 的 AtomGit 平台命令行工具，使用 Go 构建。命令行简称为 `amc`。

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
go build -o amc .
```

### 安装 Skill（Claude Code 用户）

```bash
# 安装到 ~/.claude/skills（Claude Code skills 目录）
amc skills install --claude

# 安装到 ~/.agents/skills
amc skills install
```

安装后 Claude Code 会自动加载 skill，在与 AtomGit/AtomGit 平台交互时会获得帮助提示。

### 使用

```bash
# 登录
amc auth login -t YOUR_TOKEN

# 仓库操作
amc repo create my-project --public
amc repo list
amc repo view owner/repo

# Issue 操作
amc issue list -R owner/repo
amc issue view 123 -R owner/repo

# Pull Request
amc pr list -R owner/repo
amc pr view 456 -R owner/repo

# 代码浏览
amc branch list -R owner/repo
amc commit list -R owner/repo
amc release list -R owner/repo

# 搜索
amc search repos golang
amc search users username

# 直接调用 API
amc api /api/v1/user
```

## 命令参考

| 命令 | 说明 |
|------|------|
| `amc auth` | 认证 (login/logout/status) |
| `amc repo` | 仓库管理 |
| `amc issue` | Issue 管理 |
| `amc pr` | Pull Request |
| `amc branch` | 分支操作 |
| `amc commit` | 提交历史 |
| `amc release` | 发布版本 |
| `amc tag` | 标签管理 |
| `amc fork` | 叉取仓库 |
| `amc star` | 收藏仓库 |
| `amc user` | 用户信息 |
| `amc org` | 组织管理 |
| `amc search` | 搜索 |
| `amc api` | 直接调用 API |
| `amc skills` | 技能管理 |

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
make build        # 构建到 ./bin/amc
make install      # 构建并安装到 ~/bin/amc
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
