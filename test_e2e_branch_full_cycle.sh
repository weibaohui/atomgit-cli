#!/bin/bash
# E2E Test: Complete Branch Lifecycle
# Tests: create -> list -> view -> protect -> list-protected -> unprotect -> list-protected -> delete -> list -> view (should fail)
# Usage: ATOMGIT_TEST_REPO="owner/repo" ./test_e2e_branch_full_cycle.sh

set -e

REPO="${ATOMGIT_TEST_REPO:-weibaohui/atomgit-cli-e2e-test}"
BRANCH_NAME="test-branch-$(date +%s)"
CLI="./atomgit"

log_info() { echo -e "\033[0;34m[INFO]\033[0m $1"; }
log_ok() { echo -e "\033[0;32m[OK]\033[0m $1"; }
log_fail() { echo -e "\033[0;31m[FAIL]\033[0m $1"; }

echo "=========================================="
echo "Branch Lifecycle E2E Test"
echo "Repository: $REPO"
echo "Branch: $BRANCH_NAME"
echo "=========================================="

# Build CLI if not exists
if [ ! -f "$CLI" ]; then
    log_info "Building atomgit CLI..."
    go build -o "$CLI" .
fi

# Check token
if [ -z "$ATOMGIT_TOKEN" ]; then
    log_fail "ATOMGIT_TOKEN not set"
    exit 1
fi

# ========================================
# Step 1: Create branch
# ========================================
log_info "Step 1: Create branch"
$CLI branch create "$BRANCH_NAME" -R "$REPO" > /dev/null
log_ok "Branch created: $BRANCH_NAME"

# ========================================
# Step 2: List branches (verify it appears)
# ========================================
log_info "Step 2: List branches"
LIST_OUTPUT=$($CLI branch list -R "$REPO")
if echo "$LIST_OUTPUT" | grep -q "\"name\": \"$BRANCH_NAME\""; then
    log_ok "Branch appears in list"
else
    log_fail "Branch not found in list"
    exit 1
fi

# ========================================
# Step 3: View branch (verify it exists)
# ========================================
log_info "Step 3: View branch"
VIEW_OUTPUT=$($CLI branch view "$BRANCH_NAME" -R "$REPO")
if echo "$VIEW_OUTPUT" | grep -q "\"name\": \"$BRANCH_NAME\""; then
    log_ok "Branch view successful"
else
    log_fail "Branch view failed"
    exit 1
fi

# ========================================
# Step 4: Protect branch
# ========================================
log_info "Step 4: Protect branch"
$CLI branch protect "$BRANCH_NAME" -R "$REPO" > /dev/null
log_ok "Branch protected"

# ========================================
# Step 5: List protected branches (verify protection)
# ========================================
log_info "Step 5: List protected branches"
PROTECTED_LIST=$($CLI branch protected-list -R "$REPO")
if echo "$PROTECTED_LIST" | grep -q "\"name\": \"$BRANCH_NAME\""; then
    log_ok "Branch appears in protected list"
else
    log_fail "Branch not found in protected list"
    exit 1
fi

# ========================================
# Step 6: Unprotect branch
# ========================================
log_info "Step 6: Unprotect branch"
$CLI branch unprotect "$BRANCH_NAME" -R "$REPO" > /dev/null
log_ok "Branch unprotected"

# ========================================
# Step 7: List protected branches (verify it's gone)
# ========================================
log_info "Step 7: List protected branches (should be empty)"
PROTECTED_LIST_AFTER=$($CLI branch protected-list -R "$REPO")
if echo "$PROTECTED_LIST_AFTER" | grep -q "\"name\": \"$BRANCH_NAME\""; then
    log_fail "Branch still in protected list after unprotect"
    exit 1
else
    log_ok "Branch removed from protected list"
fi

# ========================================
# Step 8: Delete branch
# ========================================
log_info "Step 8: Delete branch"
$CLI branch delete "$BRANCH_NAME" -R "$REPO" > /dev/null
log_ok "Branch deleted"

# ========================================
# Step 9: List branches (verify it's gone)
# ========================================
log_info "Step 9: List branches (branch should be gone)"
LIST_AFTER_DELETE=$($CLI branch list -R "$REPO")
if echo "$LIST_AFTER_DELETE" | grep -q "\"name\": \"$BRANCH_NAME\""; then
    log_fail "Branch still in list after delete"
    exit 1
else
    log_ok "Branch removed from list"
fi

# ========================================
# Step 10: View branch (should fail)
# ========================================
log_info "Step 10: View deleted branch (should fail)"
VIEW_AFTER_DELETE=$($CLI branch view "$BRANCH_NAME" -R "$REPO" 2>&1 || true)
if echo "$VIEW_AFTER_DELETE" | grep -q "404\|not found\|Not Found"; then
    log_ok "Branch correctly returns 404 after deletion"
else
    log_fail "Branch view did not fail as expected"
    exit 1
fi

echo ""
echo "=========================================="
echo "All Branch Lifecycle Tests Passed!"
echo "=========================================="
echo ""
echo "Verified complete lifecycle:"
echo "  1. create ✓"
echo "  2. list (verify appears) ✓"
echo "  3. view ✓"
echo "  4. protect ✓"
echo "  5. list-protected (verify protection) ✓"
echo "  6. unprotect ✓"
echo "  7. list-protected (verify removal) ✓"
echo "  8. delete ✓"
echo "  9. list (verify deletion) ✓"
echo "  10. view (should fail) ✓"
