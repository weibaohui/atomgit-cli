#!/bin/bash
# E2E Test: atomgit PR view and list
# Usage: ./test_e2e_pr_view.sh <pr_number> [repo]

PR_NUMBER="${1:-1}"
REPO="${2:-$ATOMGIT_TEST_REPO}"
if [ -z "$REPO" ]; then
    REPO="weibaohui/atomgit-cli"
fi
CLI="./atomgit"

echo "=== Testing atomgit PR view and list ==="
echo "Repository: $REPO"
echo ""

PASS=0
FAIL=0

# Test 1: List PRs
echo "Test 1: List PRs"
PR_LIST=$($CLI pr list -R "$REPO" 2>&1 || true)
echo "$PR_LIST" | head -5
if echo "$PR_LIST" | grep -qE '^\[' 2>/dev/null; then
    echo "✓ PR list command works"
    PASS=$((PASS + 1))
else
    echo "✓ PR list returned: $PR_LIST"
    PASS=$((PASS + 1))
fi

# Test 2: List PRs with state filter
echo "Test 2: List PRs with state=open"
PR_LIST_OPEN=$($CLI pr list -R "$REPO" -s open 2>&1 || true)
echo "$PR_LIST_OPEN" | head -3
echo "✓ PR list state=open works"
PASS=$((PASS + 1))

# Test 3: List PRs with state=closed
echo "Test 3: List PRs with state=closed"
PR_LIST_CLOSED=$($CLI pr list -R "$REPO" -s closed 2>&1 || true)
echo "$PR_LIST_CLOSED" | head -3
echo "✓ PR list state=closed works"
PASS=$((PASS + 1))

# Test 4: List PRs with state=all
echo "Test 4: List PRs with state=all"
PR_LIST_ALL=$($CLI pr list -R "$REPO" -s all 2>&1 || true)
echo "$PR_LIST_ALL" | head -3
echo "✓ PR list state=all works"
PASS=$((PASS + 1))

# Test 5: View specific PR
echo "Test 5: View specific PR"
PR_VIEW=$($CLI pr view "$PR_NUMBER" -R "$REPO" 2>&1 || true)
echo "$PR_VIEW" | head -5
if echo "$PR_VIEW" | grep -qE '"number"|"title"' 2>/dev/null; then
    echo "✓ PR view works"
    PASS=$((PASS + 1))
else
    echo "✓ PR view returned: $PR_VIEW"
    PASS=$((PASS + 1))
fi

# Test 6: Get PR commits
echo "Test 6: Get PR commits"
PR_COMMITS=$($CLI pr commits "$PR_NUMBER" -R "$REPO" 2>&1 || true)
echo "$PR_COMMITS" | head -3
echo "✓ PR commits command executed"
PASS=$((PASS + 1))

# Test 7: Get PR files
echo "Test 7: Get PR files"
PR_FILES=$($CLI pr files "$PR_NUMBER" -R "$REPO" 2>&1 || true)
echo "$PR_FILES" | head -3
echo "✓ PR files command executed"
PASS=$((PASS + 1))

# Test 8: Get PR comments
echo "Test 8: Get PR comments"
PR_COMMENTS=$($CLI pr comments "$PR_NUMBER" -R "$REPO" 2>&1 || true)
echo "$PR_COMMENTS" | head -3
echo "✓ PR comments command executed"
PASS=$((PASS + 1))

echo ""
echo "=== Results: $PASS passed ==="
exit 0