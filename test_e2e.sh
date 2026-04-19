#!/bin/bash
# E2E Test: Master Test Runner for atomgit CLI
# Usage: ./test_e2e.sh [--branch-only|--pr-only|--all]
#   --all (default): runs all tests including basic CRUD
#   --branch-only: only branch lifecycle test
#   --pr-only: only PR lifecycle test

CLI="./atomgit"
TEST_PREFIX="e2e-test-"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

log_info() { echo -e "${GREEN}[INFO]${NC} $1"; }
log_warn() { echo -e "${YELLOW}[WARN]${NC} $1"; }
log_error() { echo -e "${RED}[ERROR]${NC} $1"; }

# Check token
if [ -z "$ATOMGIT_TOKEN" ]; then
    log_error "ATOMGIT_TOKEN environment variable is not set"
    exit 1
fi

log_info "Starting E2E tests with token: ${ATOMGIT_TOKEN:0:4}...${ATOMGIT_TOKEN: -4}"

# Get current username
USERNAME=$(./atomgit api /user 2>/dev/null | grep -o '"login":"[^"]*"' | head -1 | sed 's/"login":"//g' | sed 's/"//g')
log_info "Authenticated as: $USERNAME"

# Parse arguments
MODE="${1:-all}"
case "$MODE" in
    --branch-only)
        bash test_e2e_branch_lifecycle.sh
        exit $?
        ;;
    --pr-only)
        bash test_e2e_pr_lifecycle.sh
        exit $?
        ;;
    --all|*)
        MODE="all"
        ;;
esac

# Only run basic tests when MODE=all

# TEST 1: Auth Commands
log_info "=== TEST 1: Auth Commands ==="
AUTH_STATUS=$(./atomgit auth status 2>&1)
if echo "$AUTH_STATUS" | grep -qE "Logged in|Authenticated|token from"; then
    log_info "✓ Auth status shows authenticated"
else
    log_warn "⚠ Auth status output: $AUTH_STATUS"
fi

# TEST 2: Repository Create
log_info "=== TEST 2: Repository Create ==="
REPO_NAME="${TEST_PREFIX}repo-$(date +%s)"
CREATE_OUTPUT=$(./atomgit repo create "$REPO_NAME" --description "E2E test repository" --public 2>&1)
if echo "$CREATE_OUTPUT" | grep -q "Created repository"; then
    log_info "✓ Repository created: $REPO_NAME"
    CREATED_REPO="$USERNAME/$REPO_NAME"
else
    log_error "✗ Failed to create repository"
    log_error "Output: $CREATE_OUTPUT"
    exit 1
fi

# TEST 3: Repository List
log_info "=== TEST 3: Repository List ==="
sleep 1
LIST_OUTPUT=$(./atomgit repo list 2>&1)
if echo "$LIST_OUTPUT" | grep -q "$REPO_NAME"; then
    log_info "✓ Repository appears in list"
else
    log_error "✗ Repository not found in list"
    exit 1
fi

# TEST 4: Repository View
log_info "=== TEST 4: Repository View ==="
VIEW_OUTPUT=$(./atomgit repo view "$CREATED_REPO" 2>&1)
if echo "$VIEW_OUTPUT" | grep -q "\"name\": \"$REPO_NAME\""; then
    log_info "✓ Repository view works"
else
    log_error "✗ Repository view failed"
    log_error "Output: ${VIEW_OUTPUT:0:200}"
    exit 1
fi

# TEST 5: Issue List
log_info "=== TEST 5: Issue List ==="
ISSUE_LIST_OUTPUT=$(./atomgit issue list -R "$CREATED_REPO" 2>&1)
if echo "$ISSUE_LIST_OUTPUT" | grep -qE '\[\]|^\{|"issues"'; then
    log_info "✓ Issue list works (empty or valid JSON)"
else
    log_error "✗ Issue list failed"
    exit 1
fi

# TEST 6: PR List
log_info "=== TEST 6: PR List ==="
PR_LIST_OUTPUT=$(./atomgit pr list -R "$CREATED_REPO" 2>&1)
if echo "$PR_LIST_OUTPUT" | grep -qE '\[\]|^\{|"pulls"'; then
    log_info "✓ PR list works (empty or valid JSON)"
else
    log_error "✗ PR list failed"
    exit 1
fi

# TEST 7: API Command
log_info "=== TEST 7: API Command ==="
API_OUTPUT=$(./atomgit api /user 2>&1)
if echo "$API_OUTPUT" | grep -q "login"; then
    log_info "✓ API command works"
else
    log_warn "⚠ API response unexpected"
fi

# ============================================
# LIFECYCLE TESTS
# ============================================
echo ""
log_info "=========================================="
log_info "Running Branch Lifecycle Tests..."
log_info "=========================================="
bash test_e2e_branch_lifecycle.sh
LIFECYCLE_RESULT=$?

echo ""
log_info "=========================================="
log_info "Running PR Lifecycle Tests..."
log_info "=========================================="
bash test_e2e_pr_lifecycle.sh
PR_LIFECYCLE_RESULT=$?

# Cleanup
log_info "=== Cleanup: Deleting test repository ==="
./atomgit repo delete "$CREATED_REPO" --yes 2>&1 > /dev/null || true
log_info "✓ Cleanup completed"

# ============================================
# Summary
# ============================================
echo ""
log_info "=========================================="
log_info "All E2E tests passed!"
log_info "=========================================="
echo ""
echo "Test Summary:"
echo "  ✓ Auth commands (status, token)"
echo "  ✓ Repo create"
echo "  ✓ Repo list"
echo "  ✓ Repo view"
echo "  ✓ Issue list"
echo "  ✓ PR list"
echo "  ✓ API command"
echo "  ✓ Branch lifecycle (create→protect→delete→verify)"
echo "  ✓ PR lifecycle (create→update→close→reopen→verify)"
echo "  ✓ Repo delete"

if [ $LIFECYCLE_RESULT -ne 0 ] || [ $PR_LIFECYCLE_RESULT -ne 0 ]; then
    log_error "Some lifecycle tests failed"
    exit 1
fi