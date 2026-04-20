# AtomGit CLI 项目

## 项目概述

AtomGit CLI 是一个 GitHub-like CLI 工具，用于操作 AtomGit (atomgit.com / gitcode.com) 平台。

## 代码提交流程

**重要：所有代码更改必须通过 PR 合并，禁止直接提交到 main 分支。**

### 标准流程

1. **新建分支**
   ```bash
   git checkout -b <branch-name>
   # 例如: git checkout -b feat/add-feature
   ```

2. **编写代码并提交**
   ```bash
   git add .
   git commit -m "描述信息"
   ```

3. **推送分支**
   ```bash
   git push -u origin <branch-name>
   ```

4. **创建 PR**
   ```bash
   # 先构建
   go build -o atomgit-cli .

   # 使用 atomgit CLI 创建 PR
   ./atomgit-cli pr create -R owner/repo \
     --title "PR 标题" \
     --body "## Summary\n- 变更内容" \
     --head <branch-name> \
     --base main
   ```

5. **切换回 main**
   ```bash
   git checkout main
   ```

## 使用 atomgit CLI

**注意：`gh` CLI 不支持 GitCode/AtomGit，必须使用 `atomgit` CLI。**

### 构建
```bash
go build -o atomgit-cli .
```

### 常用命令

```bash
# PR 操作
./atomgit-cli pr list -R owner/repo
./atomgit-cli pr view -R owner/repo 123 --comments
./atomgit-cli pr create -R owner/repo --title "..." --body "..." --head feature --base main
./atomgit-cli pr merge -R owner/repo 123

# Issue 操作
./atomgit-cli issue list -R owner/repo
./atomgit-cli issue view -R owner/repo 123
./atomgit-cli issue create -R owner/repo -t "Title" -b "Description"

# Repo 操作
./atomgit-cli repo list
./atomgit-cli repo view owner/repo
./atomgit-cli repo create name --public
./atomgit-cli repo delete owner/repo

# Skills 操作
./atomgit-cli skills list
./atomgit-cli skills install --claude  # 安装到 ~/.claude/skills
```

## Skills

项目内置 skills 位于 `cmd/skills/atomgit-cli/SKILL.md`。

安装到 Claude Code:
```bash
./atomgit-cli skills install --claude
```
