---
name: atomgit-cli
description: Use when users need help with AtomGit CLI commands. Covers repo, issue, pr, and skills commands. Trigger when user mentions atomgit, gitcode, GitHub-like CLI, repository management, issue tracking, or pull requests.
---

# AtomGit CLI Skill

AtomGit (atomgit.com / gitcode.com) is a GitHub-like platform. The `gh` CLI does NOT support it, but the **atomgit CLI** does!

## Quick Start

```bash
# Build atomgit CLI
go build -o atomgit-cli .

# Create PR
./atomgit-cli pr create -R owner/repo \
  --title "PR Title" \
  --body "Description" \
  --head feature-branch \
  --base main
```

## Installation

```bash
# Install to Claude Code skills (~/.claude/skills/{name})
atomgit skills install --claude

# Install to default path (~/.agents/skills)
atomgit skills install

# Install to custom path
atomgit skills install --path /custom/path
```

## Skills Commands

```bash
# List available skills
atomgit skills list
```

## Common Workflows

### Create Pull Request
1. Create feature branch: `git checkout -b feature-branch`
2. Push: `git push -u origin feature-branch`
3. Create PR:
   ```bash
   ./atomgit-cli pr create -R owner/repo \
     --title "Feature: description" \
     --body "## Summary\n- changes" \
     --head feature-branch \
     --base main
   ```

### View PR
```bash
# View PR details
./atomgit-cli pr view -R owner/repo 123

# View with comments
./atomgit-cli pr view -R owner/repo 123 --comments
```

### List PRs
```bash
./atomgit-cli pr list -R owner/repo
./atomgit-cli pr list -R owner/repo --state open
```

### Merge PR
```bash
./atomgit-cli pr merge -R owner/repo 123
```

## Repository Commands

```bash
# List repositories
atomgit repo list

# View repository details
atomgit repo view owner/repo

# Create a new repository
atomgit repo create name --public      # public repo
atomgit repo create name --private    # private repo

# Delete a repository
atomgit repo delete owner/repo
```

## Issue Commands

```bash
# List issues
atomgit issue list -R owner/repo

# Filter issues
atomgit issue list -R owner/repo --state open --label bug

# View issue details
atomgit issue view -R owner/repo 123

# Create an issue
atomgit issue create -R owner/repo -t "Title" -b "Description"
```

## Pull Request Commands

### List PRs
```bash
atomgit pr list -R owner/repo
atomgit pr list -R owner/repo --state open
```

### View PR
```bash
# View PR details
atomgit pr view -R owner/repo 123

# View with comments
atomgit pr view -R owner/repo 123 --comments
```

### Create PR
```bash
# Create PR (requires different head and base branches)
atomgit pr create -R owner/repo \
  --title "PR Title" \
  --body "PR Description" \
  --head feature-branch \
  --base main
```

### Merge PR
```bash
atomgit pr merge -R owner/repo 123
```

## Common Options

```
-R, --repo owner/repo    Specify repository
-L, --limit 20           Limit number of results
--state open|closed      Filter by state
--format table|json      Output format
```

## Supported Platforms

| Platform | URL Format |
|----------|-----------|
| GitCode | `https://gitcode.com/{owner}/{repo}` |
| AtomGit | `https://atomgit.com/{owner}/{repo}` |
| Git Clone | `https://atomgit.com/{owner}/{repo}.git` |

## When `gh` Fails

If `gh pr create` fails with:
> "none of the git remotes configured for this repository point to a known GitHub host"

This means the remote is GitCode, not GitHub. Use `atomgit` CLI instead.

## Quick Reference

| Task | GitHub (gh) | GitCode/AtomGit (atomgit) |
|------|------------|---------------------------|
| Create PR | `gh pr create` | `atomgit pr create` |
| View PR | `gh pr view` | `atomgit pr view` |
| List PRs | `gh pr list` | `atomgit pr list` |
| Merge PR | `gh pr merge` | `atomgit pr merge` |
| Clone | `git clone` | `git clone` |
| List Issues | `gh issue list` | `atomgit issue list` |
| Create Issue | `gh issue create` | `atomgit issue create` |
