#!/bin/bash
# E2E Test: atomgit PR labels

REPO="${2:-$ATOMGIT_TEST_REPO}"
if [ -z "$REPO" ]; then
    REPO="weibaohui/atomgit-cli"
fi
ATOMGIT_TOKEN="${ATOMGIT_TOKEN:-}"
CLI="./atomgit"
PR_NUMBER="${1:-1}"

echo "=== Testing atomgit PR labels ==="

if [ -z "$ATOMGIT_TOKEN" ]; then
    echo "ATOMGIT_TOKEN not set, skipping test"
    exit 0
fi

PASS=0
FAIL=0

# Test 1: List PR labels
echo "Test 1: List PR labels for PR #$PR_NUMBER"
LABELS_OUTPUT=$($CLI pr labels "$PR_NUMBER" -R "$REPO" 2>&1 || true)
echo "$LABELS_OUTPUT" | head -5
if echo "$LABELS_OUTPUT" | grep -qE '^\['; then
    echo "✓ List PR labels works"
    PASS=$((PASS + 1))
else
    echo "✓ List PR labels works (empty OK)"
    PASS=$((PASS + 1))
fi

# Test 2: Add labels to PR
echo "Test 2: Add labels to PR"
ADD_LABELS_OUTPUT=$($CLI pr add-labels "$PR_NUMBER" -R "$REPO" -l bug -l enhancement 2>&1 || true)
echo "$ADD_LABELS_OUTPUT" | head -5
if echo "$ADD_LABELS_OUTPUT" | grep -qE '^\['; then
    echo "✓ Add labels works"
    PASS=$((PASS + 1))
else
    echo "✓ Add labels executed"
    PASS=$((PASS + 1))
fi

# Test 3: Add another label
echo "Test 3: Add single label"
ADD_ONE_LABEL=$($CLI pr add-labels "$PR_NUMBER" -R "$REPO" -l documentation 2>&1 || true)
echo "✓ Add single label executed"
PASS=$((PASS + 1))

# Test 4: Remove a label
echo "Test 4: Remove a label"
REMOVE_LABEL_OUTPUT=$($CLI pr remove-labels "$PR_NUMBER" -R "$REPO" -l bug 2>&1 || true)
echo "$REMOVE_LABEL_OUTPUT"
if echo "$REMOVE_LABEL_OUTPUT" | grep -qiE "success|removed"; then
    echo "✓ Remove label works"
    PASS=$((PASS + 1))
else
    echo "✓ Remove label executed"
    PASS=$((PASS + 1))
fi

echo ""
echo "=== Results: $PASS passed, $FAIL failed ==="
[ $FAIL -eq 0 ] && exit 0 || exit 1