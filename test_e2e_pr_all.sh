#!/bin/bash
# E2E Test: atomgit PR Commands - Master Test Runner

echo "============================================"
echo "atomgit PR Commands E2E Test Suite"
echo "============================================"
echo ""

PASS=0
FAIL=0
TOTAL=0

run_test() {
    local name="$1"
    local script="$2"
    local args="${3:-}"

    echo "--------------------------------------------"
    echo "Running: $name"
    echo "--------------------------------------------"
    TOTAL=$((TOTAL + 1))

    if [ -n "$args" ]; then
        ./"$script" "$args" 2>&1
    else
        ./"$script" 2>&1
    fi

    if [ $? -eq 0 ]; then
        echo "✓ $name PASSED"
        PASS=$((PASS + 1))
    else
        echo "✗ $name FAILED"
        FAIL=$((FAIL + 1))
    fi
    echo ""
}

# Check if token is set
if [ -z "$ATOMGIT_TOKEN" ]; then
    echo "WARNING: ATOMGIT_TOKEN not set, tests may fail"
    echo ""
fi

# Get a real PR number for tests that need it
echo "Getting PR number for targeted tests..."
PR_NUMBER=$(./atomgit pr list -R weibaohui/atomgit-cli -s all -L 1 2>/dev/null | grep -o '"number":[0-9]*' | head -1 | cut -d: -f2 || echo "1")
echo "Using PR number: $PR_NUMBER"
echo ""

# Run tests
run_test "PR View and List" "test_e2e_pr_view.sh"
run_test "PR Create" "test_e2e_pr_create.sh"
run_test "PR Labels" "test_e2e_pr_labels.sh" "$PR_NUMBER"
run_test "PR Assignees" "test_e2e_pr_assignees.sh" "$PR_NUMBER"
run_test "PR Reviewers and Testers" "test_e2e_pr_reviewers.sh" "$PR_NUMBER"
run_test "PR Merge/Update/Close" "test_e2e_pr_merge.sh" "$PR_NUMBER"

echo "============================================"
echo "FINAL RESULTS"
echo "============================================"
echo "Tests run: $TOTAL"
echo "Passed: $PASS"
echo "Failed: $FAIL"
echo ""

if [ $FAIL -eq 0 ]; then
    echo "All PR E2E tests passed!"
    exit 0
else
    echo "Some PR E2E tests failed."
    exit 1
fi