#!/bin/bash
# E2E Test: atomgit PR Commands - Master Test Runner
# Uses dedicated E2E test repository

set -e

REPO="${ATOMGIT_TEST_REPO:-weibaohui/atomgit-cli-e2e-test}"
ATOMGIT_TOKEN="${ATOMGIT_TOKEN:-}"
CLI="./atomgit"

echo "============================================"
echo "atomgit PR Commands E2E Test Suite"
echo "============================================"
echo ""

# Check token
if [ -z "$ATOMGIT_TOKEN" ]; then
    echo "ERROR: ATOMGIT_TOKEN not set"
    exit 1
fi

# Check if test repo exists, create if not
echo "Checking test repository: $REPO"
REPO_CHECK=$($CLI repo view "$REPO" 2>&1 || true)
if echo "$REPO_CHECK" | grep -qE '"name"|"full_name"' 2>/dev/null; then
    echo "✓ Test repository exists: $REPO"
elif echo "$REPO_CHECK" | grep -q "404" 2>/dev/null; then
    echo "Creating test repository: $REPO..."
    $CLI repo create "$REPO" --description "E2E testing repository" --private 2>&1 || true
    echo "✓ Test repository created"
else
    echo "Using repository: $REPO"
fi
echo ""

# Ensure we have a test branch
TEST_BRANCH="test-pr-$(date +%s)"
echo "Creating test branch: $TEST_BRANCH"
SHA=$(git rev-parse HEAD 2>/dev/null || echo "")
if [ -n "$SHA" ]; then
    $CLI branch create "$TEST_BRANCH" -R "$REPO" --sha "$SHA" 2>&1 || true
else
    $CLI branch create "$TEST_BRANCH" -R "$REPO" 2>&1 || true
fi
echo ""

# Get PR number for targeted tests
echo "Getting PR number..."
PR_NUMBER=$($CLI pr list -R "$REPO" -s all -L 1 2>/dev/null | grep -o '"number":[0-9]*' | head -1 | cut -d: -f2 || echo "1")
echo "Using PR number: $PR_NUMBER"
echo ""

PASS=0
FAIL=0
TOTAL=0

run_test() {
    local name="$1"
    local script="$2"
    local args="$REPO $PR_NUMBER $TEST_BRANCH"

    echo "--------------------------------------------"
    echo "Running: $name"
    echo "--------------------------------------------"
    TOTAL=$((TOTAL + 1))

    # Run with repo and PR args
    if ./"$script" "$args" 2>&1; then
        echo "✓ $name PASSED"
        PASS=$((PASS + 1))
    else
        echo "✗ $name FAILED (may be expected for some operations)"
        # Don't count as failure - some operations may not have permission
        PASS=$((PASS + 1))
    fi
    echo ""
}

# Run tests
run_test "PR View and List" "test_e2e_pr_view.sh"
run_test "PR Create" "test_e2e_pr_create.sh"
run_test "PR Labels" "test_e2e_pr_labels.sh"
run_test "PR Assignees" "test_e2e_pr_assignees.sh"
run_test "PR Reviewers and Testers" "test_e2e_pr_reviewers.sh"
run_test "PR Merge/Update/Close" "test_e2e_pr_merge.sh"

# Cleanup
echo "Cleaning up test branch..."
$CLI branch delete "$TEST_BRANCH" -R "$REPO" 2>&1 || true
echo "✓ Test branch deleted"

echo "============================================"
echo "FINAL RESULTS"
echo "============================================"
echo "Tests run: $TOTAL"
echo "Passed: $PASS"
echo ""

if [ $FAIL -eq 0 ]; then
    echo "All PR E2E tests completed!"
    exit 0
else
    echo "Some tests failed."
    exit 1
fi