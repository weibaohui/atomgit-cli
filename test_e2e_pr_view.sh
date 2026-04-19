#!/bin/bash
# E2E Test: atomgit PR view and list

REPO="weibaohui/atomgit-cli"
ATOMGIT_TOKEN="${ATOMGIT_TOKEN:-}"
CLI="./atomgit"

echo "=== Testing atomgit PR view and list ==="

if [ -z "$ATOMGIT_TOKEN" ]; then
    echo "ATOMGIT_TOKEN not set, skipping test"
    exit 0
fi

PASS=0
FAIL=0

check() {
    local desc="$1"
    local result="$2"
    if [ "$result" -eq 0 ]; then
        echo "✓ $desc"
        ((PASS++))
    else
        echo "✗ $desc"
        ((FAIL++))
    fi
}

# Test 1: List PRs
echo "Test 1: List PRs"
PR_LIST=$($CLI pr list -R "$REPO" 2>&1 || true)
echo "$PR_LIST" | head -5
if echo "$PR_LIST" | grep -qE '^\['; then
    check "PR list command works" 0
else
    check "PR list command works" 1
fi

# Test 2: List PRs with state filter
echo "Test 2: List PRs with state=open"
PR_LIST_OPEN=$($CLI pr list -R "$REPO" -s open 2>&1 || true)
echo "$PR_LIST_OPEN" | head -3
if echo "$PR_LIST_OPEN" | grep -qE '^\['; then
    check "PR list state=open works" 0
else
    check "PR list state=open works" 1
fi

# Test 3: List PRs with state=closed
echo "Test 3: List PRs with state=closed"
PR_LIST_CLOSED=$($CLI pr list -R "$REPO" -s closed 2>&1 || true)
echo "$PR_LIST_CLOSED" | head -3
if echo "$PR_LIST_CLOSED" | grep -qE '^\['; then
    check "PR list state=closed works" 0
else
    check "PR list state=closed works" 1
fi

# Test 4: List PRs with state=all
echo "Test 4: List PRs with state=all"
PR_LIST_ALL=$($CLI pr list -R "$REPO" -s all 2>&1 || true)
echo "$PR_LIST_ALL" | head -3
if echo "$PR_LIST_ALL" | grep -qE '^\['; then
    check "PR list state=all works" 0
else
    check "PR list state=all works" 1
fi

echo ""
echo "=== Results: $PASS passed, $FAIL failed ==="
[ $FAIL -eq 0 ] && exit 0 || exit 1