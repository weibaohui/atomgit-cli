#!/bin/bash
set -e

BASE_URL="https://docs.atomgit.com"
API_BASE="https://api.atomgit.com/api/v5"
OUTPUT_DIR="/Users/weibh/projects/go/atomgit_cli/docs"

# Define API categories and their endpoints
declare -A APIS=(
  # Issue APIs
  ["issue/post-api-v-5-repos-owner-issues"]="post-api-v-5-repos-owner-issues|创建Issue|POST|/repos/:owner/issues"
  ["issue/patch-api-v-5-repos-owner-issues-number"]="patch-api-v-5-repos-owner-issues-number|更新Issue|PATCH|/repos/:owner/issues/:number"
  ["issue/get-api-v-5-repos-owner-repo-issues-number"]="get-api-v-5-repos-owner-repo-issues-number|获取仓库的某个Issue|GET|/repos/:owner/:repo/issues/:number"
  ["issue/get-api-v-5-repos-owner-repo-issues"]="get-api-v-5-repos-owner-repo-issues|获取仓库所有issues|GET|/repos/:owner/:repo/issues"
  ["issue/get-api-v-5-repos-owner-repo-issues-number-comments"]="get-api-v-5-repos-owner-repo-issues-number-comments|获取仓库某个Issue所有的评论|GET|/repos/:owner/:repo/issues/:number/comments"
  ["issue/post-api-v-5-repos-owner-repo-issues-number-comments"]="post-api-v-5-repos-owner-repo-issues-number-comments|创建Issue评论|POST|/repos/:owner/:repo/issues/:number/comments"
  
  # PR APIs
  ["pull_request/get-api-v-5-repos-owner-repo-pulls"]="get-api-v-5-repos-owner-repo-pulls|获取Pull Request列表|GET|/repos/:owner/:repo/pulls"
  ["pull_request/get-api-v-5-repos-owner-repo-pulls-number"]="get-api-v-5-repos-owner-repo-pulls-number|获取单个Pull Request|GET|/repos/:owner/:repo/pulls/:number"
  ["pull_request/post-api-v-5-repos-owner-repo-pulls"]="post-api-v-5-repos-owner-repo-pulls|创建Pull Request|POST|/repos/:owner/:repo/pulls"
  
  # Branch APIs  
  ["branch/get-api-v-5-repos-owner-repo-branches"]="get-api-v-5-repos-owner-repo-branches|获取分支列表|GET|/repos/:owner/:repo/branches"
  ["branch/post-api-v-5-repos-owner-repo-branches"]="post-api-v-5-repos-owner-repo-branches|创建分支|POST|/repos/:owner/:repo/branches"
  ["branch/delete-api-v-5-repos-owner-repo-branches-name"]="delete-api-v-5-repos-owner-repo-branches-name|删除分支|DELETE|/repos/:owner/:repo/branches/:name"
  
  # Repo APIs
  ["repos/get-api-v-5-repos-owner-repo"]="get-api-v-5-repos-owner-repo|获取仓库详情|GET|/repos/:owner/:repo"
  ["repos/post-api-v-5-user-repos"]="post-api-v-5-user-repos|创建仓库|POST|/user/repos"
  ["repos/delete-api-v-5-repos-owner-repo"]="delete-api-v-5-repos-owner-repo|删除仓库|DELETE|/repos/:owner/:repo"
  
  # User APIs
  ["user/get-api-v-5-user"]="get-api-v-5-user|获取当前用户|GET|/user"
  ["user/get-api-v-5-users-username"]="get-api-v-5-users-username|获取用户信息|GET|/users/:username"
  
  # Search APIs
  ["search/get-api-v-5-search-repositories"]="get-api-v-5-search-repositories|搜索仓库|GET|/search/repositories"
  ["search/get-api-v-5-search-users"]="get-api-v-5-search-users|搜索用户|GET|/search/users"
  
  # Hook APIs
  ["hook/get-api-v-5-repos-owner-repo-hooks"]="get-api-v-5-repos-owner-repo-hooks|获取Webhooks列表|GET|/repos/:owner/:repo/hooks"
  ["hook/post-api-v-5-repos-owner-repo-hooks"]="post-api-v-5-repos-owner-repo-hooks|创建Webhook|POST|/repos/:owner/:repo/hooks"
  ["hook/delete-api-v-5-repos-owner-repo-hooks-id"]="delete-api-v-5-repos-owner-repo-hooks-id|删除Webhook|DELETE|/repos/:owner/:repo/hooks/:id"
  
  # Release APIs
  ["release/get-api-v-5-repos-owner-repo-releases"]="get-api-v-5-repos-owner-repo-releases|获取Releases列表|GET|/repos/:owner/:repo/releases"
  ["release/get-api-v-5-repos-owner-repo-releases-tag"]="get-api-v-5-repos-owner-repo-releases-tag|获取单个Release|GET|/repos/:owner/:repo/releases/:tag"
)

# Create directories
mkdir -p "$OUTPUT_DIR"/{issue,pull_request,branch,repos,user,search,hook,release,commit,tag,label,milestone,org,enterprise,dashboard,ai}

# Function to create API doc
create_api_doc() {
  local key=$1
  local api_path=$2
  local name=$3
  local method=$4
  local endpoint=$5
  local category=$(echo $key | cut -d'/' -f1)
  local filename="$OUTPUT_DIR/$category/$api_path.md"
  
  echo "Creating: $filename"
  
  cat > "$filename" << EOF
# $name

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | $method |
| **Endpoint** | \`$API_BASE$endpoint\` |
| **文档链接** | $BASE_URL/docs/apis/$api_path |
| **CLI命令** | 参见下方 |

## 路径参数

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| - | - | - | 详见下方请求示例 |

## 查询参数

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

\`\`\`bash
curl -X $method "$API_BASE$endpoint" \\
  -H "Authorization: token \$ATOMGIT_TOKEN" \\
  -H "Content-Type: application/json"
\`\`\`

## 响应格式

\`\`\`json
{
  // 详见官方文档
}
\`\`\`

## 相关链接

- [官方文档]($BASE_URL/docs/apis/$api_path)
EOF
}

# Process APIs
for key in "${!APIS[@]}"; do
  IFS='|' read -r api_path name method endpoint <<< "${APIS[$key]}"
  create_api_doc "$key" "$api_path" "$name" "$method" "$endpoint"
done

echo "Done! Created $(ls -1 $OUTPUT_DIR/*/*.md 2>/dev/null | wc -l) files"
