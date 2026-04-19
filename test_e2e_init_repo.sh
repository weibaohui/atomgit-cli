#!/bin/bash
# Initialize E2E Test Repository
# This script creates a dedicated repository for E2E testing

set -e

CLI="./atomgit"
TEST_REPO="${ATOMGIT_TEST_REPO:-atomgit-cli-e2e-test}"
TEST_BRANCH="test-branch-$(date +%s)"

log_info() { echo -e "\033[0;32m[INFO]\033[0m $1"; }
log_warn() { echo -e "\033[1;33m[WARN]\033[0m $1"; }
log_error() { echo -e "\033[0;31m[ERROR]\033[0m $1"; }

# Check token
if [ -z "$ATOMGIT_TOKEN" ]; then
    log_error "ATOMGIT_TOKEN environment variable is not set"
    exit 1
fi

log_info "Initializing E2E test repository: $TEST_REPO"

# Check if repo already exists
EXISTING=$($CLI repo view "$TEST_REPO" 2>&1 || true)
if echo "$EXISTING" | grep -qE '"name"' || echo "$EXISTING" | grep -q "Already exists"; then
    log_info "Repository $TEST_REPO already exists"
else
    # Create repository
    log_info "Creating repository..."
    CREATE_OUTPUT=$($CLI repo create "$TEST_REPO" --description "E2E testing repository for atomgit-cli" --private 2>&1 || true)
    if echo "$CREATE_OUTPUT" | grep -qE '"name"|Created|成功'; then
        log_info "Repository created successfully"
    else
        log_warn "Repository creation returned: $CREATE_OUTPUT"
    fi
fi

# Create a test file
log_info "Creating test file..."
echo "# E2E Test Repository
Created at: $(date)
Test branch: $TEST_BRANCH
" > /tmp/test-e2e-$TEST_BRANCH.md

# Create initial commit via API or file creation
# For now, just ensure the repo is accessible
VERIFICATION=$($CLI repo list 2>&1 || true)
if echo "$VERIFICATION" | grep -q "$TEST_REPO"; then
    log_info "Repository $TEST_REPO is accessible"
    echo ""
    log_info "=========================================="
    log_info "E2E Test Repository Ready!"
    log_info "=========================================="
    echo ""
    echo "Test Repository: $TEST_REPO"
    echo "Test Branch: $TEST_BRANCH"
    echo ""
    echo "Use this repository for PR E2E tests:"
    echo "  export ATOMGIT_TEST_REPO=\"$TEST_REPO\""
    echo "  ./test_e2e_pr_all.sh"
else
    log_error "Failed to verify repository access"
    exit 1
fi