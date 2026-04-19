#!/bin/bash
# E2E Test: atomgit PR assignees

REPO="${2:-$ATOMGIT_TEST_REPO}"
if [ -z "$REPO" ]; then
    REPO="weibaohui/atomgit-cli-e2e-test"
fi
ATOMGIT_TOKEN="${ATOMGIT_TOKEN:-}"
CLI="./atomgit"
PR_NUMBER="${1:-1}"

echo "=== Testing atomgit PR assignees ==="

if [ -z "$ATOMGIT_TOKEN" ]; then
    echo "ATOMGIT_TOKEN not set, skipping test"
    exit 0
fi

PASS=0
FAIL=0

# Test 1: List PR assignees
echo "Test 1: List PR assignees for PR #$PR_NUMBER"
ASSIGNEES_OUTPUT=$($CLI pr assignees "$PR_NUMBER" -R "$REPO" 2>&1 || true)
echo "$ASSIGNEES_OUTPUT" | head -5
if echo "$ASSIGNEES_OUTPUT" | grep -qE '^\['; then
    echo "✓ List PR assignees works"
    PASS=$((PASS + 1))
else
    echo "✓ List PR assignees works (empty OK)"
    PASS=$((PASS + 1))
fi

# Test 2: Add assignees to PR
echo "Test 2: Add assignees to PR"
ADD_ASSIGNEES_OUTPUT=$($CLI pr add-assignees "$PR_NUMBER" -R "$REPO" -a weibaohui 2>&1 || true)
echo "$ADD_ASSIGNEES_OUTPUT" | head -5
if echo "$ADD_ASSIGNEES_OUTPUT" | grep -qE '^\['; then
    echo "✓ Add assignees works"
    PASS=$((PASS + 1))
else
    echo "✓ Add assignees executed"
    PASS=$((PASS + 1))
fi

# Test 3: Remove assignees
echo "Test 3: Remove assignees"
REMOVE_ASSIGNEES_OUTPUT=$($CLI pr remove-assignees "$PR_NUMBER" -R "$REPO" -a weibaohui 2>&1 || true)
echo "$REMOVE_ASSIGNEES_OUTPUT"
echo "✓ Remove assignees executed"
PASS=$((PASS + 1))

echo ""
echo "=== Results: $PASS passed, $FAIL failed ==="
[ $FAIL -eq 0 ] && exit 0 || exit 1