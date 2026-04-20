# AtomGit CLI 项目

## 项目概述

AtomGit CLI 是一个 GitHub-like CLI 工具，用于操作 AtomGit (atomgit.com / gitcode.com) 平台。

## 代码提交流程

**重要：所有代码更改必须通过 PR 合并，禁止直接提交到 main 分支。**
 

## 使用 atomgit CLI

**注意：`gh` CLI 不支持 GitCode/AtomGit，必须使用 `atomgit` CLI。**
 
### 常用命令

```bash
# PR 操作
./atomgit pr list -R owner/repo
./atomgit pr view -R owner/repo 123 --comments
./atomgit pr create -R owner/repo --title "..." --body "..." --head feature --base main
./atomgit pr merge -R owner/repo 123

# Issue 操作
./atomgit issue list -R owner/repo
./atomgit issue view -R owner/repo 123
./atomgit issue create -R owner/repo -t "Title" -b "Description"

# Repo 操作
./atomgit repo list
./atomgit repo view owner/repo
./atomgit repo create name --public
./atomgit repo delete owner/repo

# Skills 操作
./atomgit skills list
./atomgit skills install --claude  # 安装到 ~/.claude/skills
```

## Skills

项目内置 skills 位于 `cmd/skills/atomgit/SKILL.md`。

安装到 Claude Code:
```bash
./atomgit skills install --claude
```
