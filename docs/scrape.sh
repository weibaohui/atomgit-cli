#!/bin/bash
set -e

BASE_URL="https://docs.atomgit.com/docs/apis"
API_BASE="https://api.atomgit.com"
OUTPUT_DIR="$(dirname "$0")"

# Create category directories
mkdir -p "$OUTPUT_DIR"/{issue,pull_request,branch,repos,user,search,hook,release,commit,tag,label,milestone,org,enterprise,dashboard,ai}

get_category() {
    local url="$1"
    if echo "$url" | grep -q "kanban"; then echo "dashboard"
    elif echo "$url" | grep -qE "(similarity|audio-transcriptions|detect-yolo|video-generate|video-status|audio-classification)"; then echo "ai"
    elif echo "$url" | grep -qE "search-issues|search-repositories"; then echo "search"
    elif echo "$url" | grep -q "hooks"; then echo "hook"
    elif echo "$url" | grep -qE "releases|release"; then echo "release"
    elif echo "$url" | grep -qE "enterprises|enterprise-enterprise"; then echo "enterprise"
    elif echo "$url" | grep -qE "orgs-org|org-org|org-owner"; then echo "org"
    elif echo "$url" | grep -q "milestones"; then echo "milestone"
    elif echo "$url" | grep -q "labels"; then echo "label"
    elif echo "$url" | grep -qE "tags|protected-tags"; then echo "tag"
    elif echo "$url" | grep -qE "commits|comments"; then echo "commit"
    elif echo "$url" | grep -qE "pulls|merge-requests"; then echo "pull_request"
    elif echo "$url" | grep -q "issues"; then echo "issue"
    elif echo "$url" | grep -qE "users|user-|emails"; then echo "user"
    elif echo "$url" | grep -qE "repos|fork"; then echo "repos"
    elif echo "$url" | grep -q "branch"; then echo "branch"
    else echo "repos"; fi
}

scrape_api() {
    local url="$1"
    local api_path="${url#$BASE_URL/}"
    local category=$(get_category "$url")
    local filename="$OUTPUT_DIR/$category/$api_path.md"
    local html

    echo "Scraping: $api_path"

    # Fetch page
    html=$(curl -sL "$url" 2>/dev/null) || return 1

    # Extract title
    local title=$(echo "$html" | grep -oE '<h1[^>]*>[^<]+</h1>' | head -1 | sed 's/<[^>]*>//g')
    [ -z "$title" ] && title="$api_path"

    # Extract method
    local method=$(echo "$html" | grep -oE '(GET|POST|PUT|PATCH|DELETE)' | head -1)
    [ -z "$method" ] && method="GET"

    # Extract endpoint
    local endpoint=$(echo "$html" | grep -oE '/api/v[0-9]+/[^"'"'"'<>]+' | head -1)
    [ -z "$endpoint" ] && endpoint="/unknown"

    # Extract path params (things like :owner, :repo, :number)
    local path_params=$(echo "$endpoint" | grep -oE ':[a-z_]+' | sort -u)

    # Build query string
    local query_str="access_token=\$ATOMGIT_TOKEN"

    # Build curl command
    local curl_cmd
    if [ "$method" = "GET" ]; then
        curl_cmd="curl \"${API_BASE}${endpoint}?${query_str}\" \\\n  -H \"Authorization: token \$ATOMGIT_TOKEN\""
    else
        curl_cmd="curl -X $method \"${API_BASE}${endpoint}\" \\\n  -H \"Authorization: token \$ATOMGIT_TOKEN\" \\\n  -H \"Content-Type: application/json\" \\\n  -d '{\"key\": \"value\"}'"
    fi

    # Write markdown
    cat > "$filename" << EOF
# $title

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | $method |
| **Endpoint** | \`${API_BASE}${endpoint}\` |
| **文档链接** | $url |

## Path Parameters

$([ -n "$path_params" ] && echo "| 参数名 | 类型 | 必填 | 描述 |" && echo "|--------|------|------|------|" && echo "$path_params" | while read p; do echo "| ${p#:} | string | 是 | 路径参数 |"; done || echo "| - | - | - | - |")

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

\`\`\`bash
$curl_cmd
\`\`\`

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
EOF

    echo "  -> $filename"
}

# Read URLs from apis_list.md and scrape each
total=$(grep -oE 'https://docs\.atomgit\.com/docs/apis/[a-zA-Z0-9_-]+' "$OUTPUT_DIR/apis_list.md" | wc -l)
echo "Found $total APIs to scrape"

count=0
grep -oE 'https://docs\.atomgit\.com/docs/apis/[a-zA-Z0-9_-]+' "$OUTPUT_DIR/apis_list.md" | sort -u | while read url; do
    count=$((count + 1))
    echo "[$count/$total] $(basename $url)"
    scrape_api "$url" || echo "  FAILED: $url"
done

echo ""
echo "Done! Check go/docs/*/*.md"