#!/bin/bash
# Script to scrape GitCode API documentation using playwright-cli
# Usage: ./scrape_apis.sh

set -e

SESSION="atomgit-scrape"
DOCS_DIR="$(cd "$(dirname "$0")" && pwd)"
RESULTS_DIR="$DOCS_DIR/scrape_results"
API_LIST="$DOCS_DIR/apis_list.md"

# Create directories
mkdir -p "$RESULTS_DIR"
mkdir -p "$DOCS_DIR/repos"
mkdir -p "$DOCS_DIR/branch"
mkdir -p "$DOCS_DIR/issue"
mkdir -p "$DOCS_DIR/pull_request"
mkdir -p "$DOCS_DIR/commit"
mkdir -p "$DOCS_DIR/tag"
mkdir -p "$DOCS_DIR/label"
mkdir -p "$DOCS_DIR/milestone"
mkdir -p "$DOCS_DIR/user"
mkdir -p "$DOCS_DIR/org"
mkdir -p "$DOCS_DIR/hook"
mkdir -p "$DOCS_DIR/release"
mkdir -p "$DOCS_DIR/enterprise"
mkdir -p "$DOCS_DIR/dashboard"
mkdir -p "$DOCS_DIR/ai"
mkdir -p "$DOCS_DIR/search"

# Extract URLs from apis_list.md
extract_urls() {
    grep -oE 'https://docs\.atomgit\.com/docs/apis/[a-zA-Z0-9_-]+' "$API_LIST" | sort -u
}

# Get category from URL
get_category() {
    local url="$1"
    local url_lower=$(echo "$url" | tr '[:upper:]' '[:lower:]')

    case "$url_lower" in
        *kanban*) echo "dashboard" ;;
        *similarity*|*audio-transcriptions*|*detect-yolo*|*video-generate*|*video-status*|*audio-classification*) echo "ai" ;;
        *search-issues*|*search-repositories*) echo "search" ;;
        *hooks*) echo "hook" ;;
        *releases*|*release*) echo "release" ;;
        *enterprises*|*enterprise-enterprise*) echo "enterprise" ;;
        *orgs-org*|*org-org*|*org-owner*) echo "org" ;;
        *milestones*) echo "milestone" ;;
        *labels*) echo "label" ;;
        *tags*|*protected-tags*) echo "tag" ;;
        *commits*|*comments*|*commit-statistics*) echo "commit" ;;
        *pulls*|*merge-requests*) echo "pull_request" ;;
        *issues*) echo "issue" ;;
        *users*|*user-*|*emails*) echo "user" ;;
        *repos*|*fork*) echo "repos" ;;
        *branch*) echo "branch" ;;
        *) echo "repos" ;;
    esac
}

# Extract filename from URL
get_filename() {
    local url="$1"
    echo "$url" | sed 's|https://docs.atomgit.com/docs/apis/||' | sed 's|/||g'
}

# Open page and get title
scrape_page() {
    local url="$1"
    local category=$(get_category "$url")
    local filename=$(get_filename "$url")
    local output_file="$DOCS_DIR/$category/${filename}.md"

    echo "Scraping: $url -> $category/${filename}.md"

    # Open page
    playwright-cli -s="$SESSION" open "$url" > /dev/null 2>&1
    sleep 1

    # Get page content
    local content=$(playwright-cli -s="$SESSION" eval "document.body.innerText" 2>/dev/null)

    # Get raw HTML for more details
    local html=$(playwright-cli -s="$SESSION" eval "document.body.innerHTML" 2>/dev/null)

    echo "$content"
    echo "---HTML_START---"
    echo "$html"
    echo "---HTML_END---"
}

export -f scrape_page
export -f get_category
export -f get_filename
export SESSION
export DOCS_DIR

# Main loop
echo "Starting API scraping..."
URLS=$(extract_urls)
TOTAL=$(echo "$URLS" | wc -l)
COUNT=0

for url in $URLS; do
    COUNT=$((COUNT + 1))
    echo "[$COUNT/$TOTAL] $url"

    category=$(get_category "$url")
    filename=$(get_filename "$url")
    output_file="$DOCS_DIR/$category/${filename}.md"

    # Skip if already exists
    if [ -f "$output_file" ]; then
        echo "  -> Already exists, skipping"
        continue
    fi

    # Open and scrape
    playwright-cli -s="$SESSION" open "$url" > /dev/null 2>&1
    sleep 0.5

    # Get content
    playwright-cli -s="$SESSION" snapshot > /dev/null 2>&1

    echo "  -> Saved to $category/${filename}.md"
done

echo "Scraping complete!"
playwright-cli -s="$SESSION" close > /dev/null 2>&1 || true