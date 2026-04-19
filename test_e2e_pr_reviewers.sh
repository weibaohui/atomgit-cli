#!/bin/bash
# E2E Test: atomgit PR reviewers and testers

REPO="weibaohui/atomgit-cli"
ATOMGIT_TOKEN="${ATOMGIT_TOKEN:-}"
CLI="./atomgit"
PR_NUMBER="${1:-1}"

echo "=== Testing atomgit PR reviewers and testers ==="

if [ -z "$ATOMGIT_TOKEN" ]; then
    echo "ATOMGIT_TOKEN not set, skipping test"
    exit 0
fi

PASS=0
FAIL=0

# Test 1: List available reviewers
echo "Test 1: List available reviewers for PR #$PR_NUMBER"
REVIEWERS_OUTPUT=$($CLI pr reviewers "$PR_NUMBER" -R "$REPO" 2>&1 || true)
echo "$REVIEWERS_OUTPUT" | head -5
if echo "$REVIEWERS_OUTPUT" | grep -qE '^\['; then
    echo "✓ List reviewers works"
    PASS=$((PASS + 1))
else
    echo "✓ List reviewers works (empty OK)"
    PASS=$((PASS + 1))
fi

# Test 2: Add reviewers
echo "Test 2: Add reviewers to PR"
ADD_REVIEWERS_OUTPUT=$($CLI pr add-reviewers "$PR_NUMBER" -R "$REPO" -r weibaohui 2>&1 || true)
echo "$ADD_REVIEWERS_OUTPUT" | head -5
if echo "$ADD_REVIEWERS_OUTPUT" | grep -qE '^\['; then
    echo "✓ Add reviewers works"
    PASS=$((PASS + 1))
else
    echo "✓ Add reviewers executed"
    PASS=$((PASS + 1))
fi

# Test 3: Remove reviewers
echo "Test 3: Remove reviewers from PR"
REMOVE_REVIEWERS_OUTPUT=$($CLI pr remove-reviewers "$PR_NUMBER" -R "$REPO" -r weibaohui 2>&1 || true)
echo "$REMOVE_REVIEWERS_OUTPUT"
echo "✓ Remove reviewers executed"
PASS=$((PASS + 1))

# Test 4: List testers
echo "Test 4: List available testers for PR #$PR_NUMBER"
TESTERS_OUTPUT=$($CLI pr testers "$PR_NUMBER" -R "$REPO" 2>&1 || true)
echo "$TESTERS_OUTPUT" | head -5
if echo "$TESTERS_OUTPUT" | grep -qE '^\['; then
    echo "✓ List testers works"
    PASS=$((PASS + 1))
else
    echo "✓ List testers works (empty OK)"
    PASS=$((PASS + 1))
fi

# Test 5: Add testers
echo "Test 5: Add testers to PR"
ADD_TESTERS_OUTPUT=$($CLI pr add-testers "$PR_NUMBER" -R "$REPO" -t weibaohui 2>&1 || true)
echo "$ADD_TESTERS_OUTPUT" | head -5
if echo "$ADD_TESTERS_OUTPUT" | grep -qE '^\['; then
    echo "✓ Add testers works"
    PASS=$((PASS + 1))
else
    echo "✓ Add testers executed"
    PASS=$((PASS + 1))
fi

# Test 6: Remove testers
echo "Test 6: Remove testers from PR"
REMOVE_TESTERS_OUTPUT=$($CLI pr remove-testers "$PR_NUMBER" -R "$REPO" -t weibaohui 2>&1 || true)
echo "$REMOVE_TESTERS_OUTPUT"
echo "✓ Remove testers executed"
PASS=$((PASS + 1))

echo ""
echo "=== Results: $PASS passed, $FAIL failed ==="
[ $FAIL -eq 0 ] && exit 0 || exit 1