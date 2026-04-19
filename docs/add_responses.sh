#!/bin/bash
set -e

DOCS_DIR="$(dirname "$0")"

files=$(find "$DOCS_DIR" -name "*.md" -type f ! -name "apis_list.md" ! -name "scrape*.sh" ! -name "scrape*.py" ! -name "scrape*.cjs" | sort)

total=$(echo "$files" | wc -l | tr -d ' ')
echo "Processing $total files..."

count=0
for file in $files; do
    count=$((count + 1))
    fname=$(basename "$file")
    echo "[$count/$total] $fname"

    # Extract curl command - get everything between ```bash and ```
    # Remove backslash-n and backslash-backslash-n sequences
    curl_cmd=$(sed -n '/```bash/,/```/p' "$file" | sed '1d;$d' | sed 's/\\n//g' | tr -d '\')

    if [ -z "$curl_cmd" ]; then
        echo "  SKIP: no curl cmd"
        continue
    fi

    # Replace $ATOMGIT_TOKEN with actual token
    curl_cmd=$(echo "$curl_cmd" | sed "s/\\\$ATOMGIT_TOKEN/${ATOMGIT_TOKEN}/g")

    # Execute curl
    response=$(eval "$curl_cmd" 2>/dev/null) || {
        echo "  WARN: curl failed"
        continue
    }

    if [ -z "$response" ]; then
        echo "  WARN: empty response"
        continue
    fi

    # Format JSON if valid
    if echo "$response" | jq . > /dev/null 2>&1; then
        response=$(echo "$response" | jq -c)
    fi

    # Escape for awk/sed
    response_esc=$(echo "$response" | sed 's/\\/\\\\/g; s/"/\\"/g; s/$/\\n/g' | tr -d '\n')
    response_esc="${response_esc%??}" # remove trailing \n

    # Add response section
    if grep -q "## 响应示例" "$file"; then
        # Replace existing
        awk -v resp="$response_esc" '
            BEGIN { in_resp=0 }
            /^## 响应示例$/ { found=1; print; next }
            found && /^```json$/ { in_resp=1; print; next }
            in_resp && /^```$/ { in_resp=0; print; next }
            !in_resp { print }
        ' "$file" > "$file.tmp" && mv "$file.tmp" "$file"
    else
        # Insert before "## 相关 CLI"
        awk -v resp="$response_esc" '
            /^## 相关 CLI 命令/ { print ""; print "## 响应示例"; print ""; print "```json"; print resp; print "```"; print "" }
            {print}
        ' "$file" > "$file.tmp" && mv "$file.tmp" "$file"
    fi

    echo "  OK"
done

echo ""
echo "Done!"