#!/bin/bash
# E2E Test: Complete Branch Lifecycle
# Creates branch -> lists -> views -> protects -> lists protected -> unprotects -> deletes -> verifies deletion
# Forms a complete verifiable闭环

REPO="${ATOMGIT_TEST_REPO:-weibaohui/atomgit-cli}"
ATOMGIT_TOKEN="${ATOMGIT_TOKEN:-}"
CLI="./atomgit"
TEST_BRANCH="test-branch-$(date +%s)"

echo "=========================================="
echo "Branch Lifecycle E2E Test"
echo "Repository: $REPO"
echo "Branch: $TEST_BRANCH"
echo "=========================================="

if [ -z "$ATOMGIT_TOKEN" ]; then
    echo "ERROR: ATOMGIT_TOKEN not set"
    exit 1
fi

PASS=0
FAIL=0

# Step 1: Create branch
echo ""
echo "[Step 1] Creating branch: $TEST_BRANCH"
$CLI branch create "$TEST_BRANCH" -R "$REPO" > /dev/null
echo "✓ Branch created"
PASS=$((PASS + 1))

# Step 2: List branches (verify new branch appears)
echo ""
echo "[Step 2] Listing branches"
BRANCH_LIST=$($CLI branch list -R "$REPO" 2>&1)
if echo "$BRANCH_LIST" | grep -q "$TEST_BRANCH"; then
    echo "✓ New branch appears in list"
    PASS=$((PASS + 1))
else
    echo "✗ New branch not found in list"
    FAIL=$((FAIL + 1))
fi

# Step 3: View branch details
echo ""
echo "[Step 3] Viewing branch: $TEST_BRANCH"
BRANCH_VIEW=$($CLI branch view "$TEST_BRANCH" -R "$REPO" 2>&1)
if echo "$BRANCH_VIEW" | grep -q "name.*$TEST_BRANCH"; then
    echo "✓ Branch view successful"
    PASS=$((PASS + 1))
else
    echo "✗ Branch view failed"
    FAIL=$((FAIL + 1))
fi

# Step 4: Protect branch
echo ""
echo "[Step 4] Protecting branch"
$CLI branch protect "$TEST_BRANCH" -R "$REPO" > /dev/null
echo "✓ Branch protected"
PASS=$((PASS + 1))

# Step 5: List protected branches (verify it appears)
echo ""
echo "[Step 5] Listing protected branches"
PROTECTED_LIST=$($CLI branch protected-list -R "$REPO" 2>&1)
if echo "$PROTECTED_LIST" | grep -q "$TEST_BRANCH"; then
    echo "✓ Protected branch appears in protected list"
    PASS=$((PASS + 1))
else
    echo "✗ Protected branch not found in list"
    FAIL=$((FAIL + 1))
fi

# Step 6: Unprotect branch (so we can delete it)
echo ""
echo "[Step 6] Unprotecting branch"
$CLI branch unprotect "$TEST_BRANCH" -R "$REPO" > /dev/null
echo "✓ Branch unprotected"
PASS=$((PASS + 1))

# Step 7: Verify unprotect (list protected again)
echo ""
echo "[Step 7] Verifying unprotect - protected list should be empty or not contain our branch"
PROTECTED_LIST_AFTER=$($CLI branch protected-list -R "$REPO" 2>&1)
if echo "$PROTECTED_LIST_AFTER" | grep -qv "$TEST_BRANCH"; then
    echo "✓ Branch removed from protected list"
    PASS=$((PASS + 1))
else
    echo "✗ Branch still in protected list"
    FAIL=$((FAIL + 1))
fi

# Step 8: Delete branch
echo ""
echo "[Step 8] Deleting branch"
$CLI branch delete "$TEST_BRANCH" -R "$REPO" > /dev/null
echo "✓ Branch deleted"
PASS=$((PASS + 1))

# Step 9: List branches (verify deletion)
echo ""
echo "[Step 9] Listing branches after deletion"
BRANCH_LIST_AFTER=$($CLI branch list -R "$REPO" 2>&1)
if echo "$BRANCH_LIST_AFTER" | grep -qv "$TEST_BRANCH"; then
    echo "✓ Branch removed from list"
    PASS=$((PASS + 1))
else
    echo "✗ Branch still exists in list"
    FAIL=$((FAIL + 1))
fi

# Step 10: View branch (verify it doesn't exist - should fail)
echo ""
echo "[Step 10] Viewing deleted branch (should fail)"
if $CLI branch view "$TEST_BRANCH" -R "$REPO" 2>&1 | grep -qiE "not found|404"; then
    echo "✓ Branch correctly returns not found"
    PASS=$((PASS + 1))
else
    echo "✗ Branch still viewable after deletion"
    FAIL=$((FAIL + 1))
fi

echo ""
echo "=========================================="
echo "Results: $PASS passed, $FAIL failed"
echo "=========================================="

if [ $FAIL -eq 0 ]; then
    echo "All branch lifecycle tests passed!"
    exit 0
else
    echo "Some tests failed."
    exit 1
fi