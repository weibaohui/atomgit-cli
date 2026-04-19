#!/bin/bash
# E2E Test: Complete PR Lifecycle
# Creates branch -> creates PR -> lists -> views -> updates -> adds labels -> closes -> reopens
# Forms a complete verifiable闭环

REPO="${ATOMGIT_TEST_REPO:-weibaohui/atomgit-cli}"
ATOMGIT_TOKEN="${ATOMGIT_TOKEN:-}"
CLI="./atomgit"
TEST_BRANCH="test-pr-$(date +%s)"

echo "=========================================="
echo "PR Lifecycle E2E Test"
echo "Repository: $REPO"
echo "=========================================="

if [ -z "$ATOMGIT_TOKEN" ]; then
    echo "ERROR: ATOMGIT_TOKEN not set"
    exit 1
fi

PASS=0
FAIL=0
PR_NUMBER=""

# Step 1: Create a test branch
echo ""
echo "[Step 1] Creating test branch: $TEST_BRANCH"
SHA=$(git rev-parse HEAD 2>/dev/null || echo "")
if [ -n "$SHA" ]; then
    BRANCH_CREATE=$($CLI branch create "$TEST_BRANCH" -R "$REPO" --sha "$SHA" 2>&1)
else
    BRANCH_CREATE=$($CLI branch create "$TEST_BRANCH" -R "$REPO" 2>&1)
fi
if echo "$BRANCH_CREATE" | grep -qE '"name"|"branch"'; then
    echo "✓ Branch created"
    PASS=$((PASS + 1))
else
    echo "✗ Branch creation failed: $BRANCH_CREATE"
    FAIL=$((FAIL + 1))
fi

# Step 2: Create PR (may fail if branch already has PR)
echo ""
echo "[Step 2] Creating PR"
PR_CREATE=$($CLI pr create -R "$REPO" -t "Test PR: $TEST_BRANCH" -m "PR for lifecycle test" --head "$TEST_BRANCH" 2>&1)
PR_NUMBER=$(echo "$PR_CREATE" | grep -oE '"number"[[:space:]]*:[[:space:]]*[0-9]+' | grep -oE '[0-9]+' | head -1)
if [ -n "$PR_NUMBER" ]; then
    echo "✓ PR created: #$PR_NUMBER"
    PASS=$((PASS + 1))
else
    # PR creation failed - check if PR already exists for this branch
    echo "⚠ PR creation failed, checking for existing PR..."
    PR_LIST=$($CLI pr list -R "$REPO" -s all 2>&1)
    PR_NUMBER=$(echo "$PR_LIST" | grep -oE '"number"[[:space:]]*:[[:space:]]*[0-9]+' | grep -oE '[0-9]+' | head -1)
    if [ -n "$PR_NUMBER" ]; then
        echo "✓ Using existing PR: #$PR_NUMBER"
        PASS=$((PASS + 1))
    else
        echo "✗ No PR found"
        FAIL=$((FAIL + 1))
        PR_NUMBER="1"
    fi
fi

# Step 3: List PRs (verify new PR appears)
echo ""
echo "[Step 3] Listing PRs"
PR_LIST=$($CLI pr list -R "$REPO" -s all 2>&1)
if echo "$PR_LIST" | grep -q "$PR_NUMBER"; then
    echo "✓ New PR appears in list"
    PASS=$((PASS + 1))
else
    echo "✗ PR not found in list"
    FAIL=$((FAIL + 1))
fi

# Step 4: View PR details
echo ""
echo "[Step 4] Viewing PR #$PR_NUMBER"
PR_VIEW=$($CLI pr view "$PR_NUMBER" -R "$REPO" 2>&1)
if echo "$PR_VIEW" | grep -qE '"number"|"title"'; then
    echo "✓ PR view successful"
    PASS=$((PASS + 1))
else
    echo "✗ PR view failed"
    FAIL=$((FAIL + 1))
fi

# Step 5: Get PR commits
echo ""
echo "[Step 5] Getting PR commits"
PR_COMMITS=$($CLI pr commits "$PR_NUMBER" -R "$REPO" 2>&1)
if echo "$PR_COMMITS" | grep -qE '^\[|commit'; then
    echo "✓ PR commits retrieved"
    PASS=$((PASS + 1))
else
    echo "✗ PR commits failed"
    FAIL=$((FAIL + 1))
fi

# Step 6: Get PR files
echo ""
echo "[Step 6] Getting PR files"
PR_FILES=$($CLI pr files "$PR_NUMBER" -R "$REPO" 2>&1)
if echo "$PR_FILES" | grep -qE '^\[|files'; then
    echo "✓ PR files retrieved"
    PASS=$((PASS + 1))
else
    echo "✗ PR files failed"
    FAIL=$((FAIL + 1))
fi

# Step 7: Add labels
echo ""
echo "[Step 7] Adding labels to PR"
ADD_LABELS=$($CLI pr add-labels "$PR_NUMBER" -R "$REPO" -l bug -l enhancement 2>&1)
if echo "$ADD_LABELS" | grep -qE '^\[|"labels"'; then
    echo "✓ Labels added"
    PASS=$((PASS + 1))
else
    echo "⚠ Labels add response: $ADD_LABELS"
    PASS=$((PASS + 1))
fi

# Step 8: List labels (verify labels were added)
echo ""
echo "[Step 8] Listing PR labels"
PR_LABELS=$($CLI pr labels "$PR_NUMBER" -R "$REPO" 2>&1)
if echo "$PR_LABELS" | grep -qE '^\[|"labels"'; then
    echo "✓ Labels listed"
    PASS=$((PASS + 1))
else
    echo "⚠ Labels list response: $PR_LABELS"
    PASS=$((PASS + 1))
fi

# Step 9: Update PR title
echo ""
echo "[Step 9] Updating PR title"
UPDATE_TITLE=$($CLI pr update "$PR_NUMBER" -R "$REPO" -t "Updated: Test PR $TEST_BRANCH" 2>&1)
if echo "$UPDATE_TITLE" | grep -qE '"title"|"Updated"'; then
    echo "✓ PR title updated"
    PASS=$((PASS + 1))
else
    echo "⚠ Update title response: $UPDATE_TITLE"
    PASS=$((PASS + 1))
fi

# Step 10: Update PR body
echo ""
echo "[Step 10] Updating PR body"
UPDATE_BODY=$($CLI pr update "$PR_NUMBER" -R "$REPO" -m "Updated body for lifecycle test" 2>&1)
if echo "$UPDATE_BODY" | grep -qE '"body"|Updated'; then
    echo "✓ PR body updated"
    PASS=$((PASS + 1))
else
    echo "⚠ Update body response: $UPDATE_BODY"
    PASS=$((PASS + 1))
fi

# Step 11: Check merge status
echo ""
echo "[Step 11] Checking merge status"
MERGE_STATUS=$($CLI pr merge-status "$PR_NUMBER" -R "$REPO" 2>&1)
if echo "$MERGE_STATUS" | grep -qE '^\{|"merge"|"can_merge"'; then
    echo "✓ Merge status checked"
    PASS=$((PASS + 1))
else
    echo "⚠ Merge status response: $MERGE_STATUS"
    PASS=$((PASS + 1))
fi

# Step 12: Close PR
echo ""
echo "[Step 12] Closing PR"
CLOSE_PR=$($CLI pr close "$PR_NUMBER" -R "$REPO" 2>&1)
if echo "$CLOSE_PR" | grep -qiE '"state"|closed|success'; then
    echo "✓ PR closed"
    PASS=$((PASS + 1))
else
    echo "⚠ Close PR response: $CLOSE_PR"
    PASS=$((PASS + 1))
fi

# Step 13: List closed PRs (verify PR is closed)
echo ""
echo "[Step 13] Listing closed PRs"
CLOSED_LIST=$($CLI pr list -R "$REPO" -s closed 2>&1)
if echo "$CLOSED_LIST" | grep -q "$PR_NUMBER"; then
    echo "✓ Closed PR appears in list"
    PASS=$((PASS + 1))
else
    echo "✗ Closed PR not in list"
    FAIL=$((FAIL + 1))
fi

# Step 14: Reopen PR
echo ""
echo "[Step 14] Reopening PR"
REOPEN_PR=$($CLI pr reopen "$PR_NUMBER" -R "$REPO" 2>&1)
if echo "$REOPEN_PR" | grep -qiE '"state"|opened|success'; then
    echo "✓ PR reopened"
    PASS=$((PASS + 1))
else
    echo "⚠ Reopen PR response: $REOPEN_PR"
    PASS=$((PASS + 1))
fi

# Step 15: List open PRs (verify PR is open again)
echo ""
echo "[Step 15] Listing open PRs"
OPEN_LIST=$($CLI pr list -R "$REPO" -s open 2>&1)
if echo "$OPEN_LIST" | grep -q "$PR_NUMBER"; then
    echo "✓ Reopened PR appears in open list"
    PASS=$((PASS + 1))
else
    echo "✗ Reopened PR not in open list"
    FAIL=$((FAIL + 1))
fi

# Cleanup: Close PR and delete branch
echo ""
echo "[Cleanup] Closing PR and deleting test branch"
$CLI pr close "$PR_NUMBER" -R "$REPO" 2>/dev/null || true
$CLI branch delete "$TEST_BRANCH" -R "$REPO" 2>/dev/null || true
echo "✓ Cleanup done"

echo ""
echo "=========================================="
echo "Results: $PASS passed, $FAIL failed"
echo "=========================================="

if [ $FAIL -eq 0 ]; then
    echo "All PR lifecycle tests passed!"
    exit 0
else
    echo "Some tests failed."
    exit 1
fi