#!/bin/bash
set -e

DOCS_DIR="$(dirname "$0")"

files=$(find "$DOCS_DIR" -name "*.md" -type f ! -name "apis_list.md" ! -name "scrape*.sh" ! -name "scrape*.py" ! -name "scrape*.cjs" ! -name "index.md" | sort)

total=$(echo "$files" | wc -l | tr -d ' ')
echo "Fixing curl commands in $total files..."

for file in $files; do
    fname=$(basename "$file")

    # Extract method and endpoint from the file
    method=$(grep -oE 'GET|POST|PUT|PATCH|DELETE' "$file" | head -1)
    endpoint=$(grep -oE 'https://api\.atomgit\.com/api/v[0-9]+/[^`"]+' "$file" | head -1)

    if [ -z "$method" ] || [ -z "$endpoint" ]; then
        echo "SKIP $fname: no method or endpoint"
        continue
    fi

    # Build proper curl command
    if [ "$method" = "GET" ]; then
        curl_cmd="curl \"${endpoint}?access_token=\$ATOMGIT_TOKEN\" \\"
        curl_cmd="$curl_cmd\n  -H \"Authorization: token \$ATOMGIT_TOKEN\""
    else
        curl_cmd="curl -X $method \"${endpoint}\" \\"
        curl_cmd="$curl_cmd\n  -H \"Authorization: token \$ATOMGIT_TOKEN\" \\"
        curl_cmd="$curl_cmd\n  -H \"Content-Type: application/json\" \\"
        curl_cmd="$curl_cmd\n  -d '{\"key\": \"value\"}'"
    fi

    # Replace the curl example
    awk -v new_cmd="$curl_cmd" '
        /^```bash$/ { in_code=1; print; print new_cmd; next }
        in_code && /^```$/ { in_code=0; next }
        !in_code { print }
    ' "$file" > "$file.tmp" && mv "$file.tmp" "$file"

    echo "FIX $fname"
done

echo "Done!"