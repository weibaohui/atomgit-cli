---
name: atg
description: Use when users need help with AtomGit CLI commands. Covers repo, issue, pr, and skills commands. Trigger when user mentions atg, gitcode, GitHub-like CLI, repository management, issue tracking, or pull requests.
---

# AtomGit CLI Skill

AtomGit (atomgit.com / gitcode.com) is a GitHub-like platform. The `gh` CLI does NOT support it, but the **atomgit CLI** does!

## Quick Start
 
# Create PR
``` bash
atg pr create -R owner/repo \
  --title "PR Title" \
  --body "Description" \
  --head feature-branch \
  --base main
```
 

## Skills Commands

```bash
# List available skills
atg skills list
```

## Common Workflows

### Create Pull Request
1. Create feature branch: `git checkout -b feature-branch`
2. Push: `git push -u origin feature-branch`
3. Create PR:
   ```bash
   ./atg pr create -R owner/repo \
     --title "Feature: description" \
     --body "## Summary\n- changes" \
     --head feature-branch \
     --base main
   ```

### View PR
```bash
# View PR details
./atg pr view -R owner/repo 123

# View with comments
./atg pr view -R owner/repo 123 --comments
```

### List PRs
```bash
./atg pr list -R owner/repo
./atg pr list -R owner/repo --state open  
```

### Merge PR
```bash
./atg pr merge -R owner/repo 123
```

## Repository Commands

```bash
# List repositories
atg repo list

# View repository details
atg repo view owner/repo

# Create a new repository
atg repo create name --public      # public repo
atg repo create name --private    # private repo

# Delete a repository
atg repo delete owner/repo
```

## Issue Commands

```bash
# List issues
atg issue list -R owner/repo

# Filter issues
atg issue list -R owner/repo --state open --label bug

# View issue details
atg issue view -R owner/repo 123

# Create an issue
atg issue create -R owner/repo -t "Title" -b "Description"
```

## Pull Request Commands

### List PRs
```bash
atg pr list -R owner/repo
atg pr list -R owner/repo --state open
```

### View PR
```bash
# View PR details
atg pr view -R owner/repo 123

# View with comments
atg pr view -R owner/repo 123 --comments
```

### Create PR
```bash
# Create PR (requires different head and base branches)
atg pr create -R owner/repo \
  --title "PR Title" \
  --body "PR Description" \
  --head feature-branch \
  --base main
```

### Merge PR
```bash
atg pr merge -R owner/repo 123
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

## When `gh` Fails

If `gh pr create` fails with:
> "none of the git remotes configured for this repository point to a known GitHub host"

This means the remote is GitCode, not GitHub. Use `atg` CLI instead.

## Quick Reference

| Task | GitHub (gh) | GitCode/AtomGit (atg) |
|------|------------|---------------------------|
| Create PR | `gh pr create` | `atg pr create` |
| View PR | `gh pr view` | `atg pr view` |
| List PRs | `gh pr list` | `atg pr list` |
| Merge PR | `gh pr merge` | `atg pr merge` |
| Clone | `git clone` | `git clone` |
| List Issues | `gh issue list` | `atg issue list` |
| Create Issue | `gh issue create` | `atg issue create` |
