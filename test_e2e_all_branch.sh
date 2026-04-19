#!/bin/bash
# E2E Test: All Branch Commands
# Usage: ./test_e2e_all_branch.sh
# Requires: ATOMGIT_TOKEN, ATOMGIT_TEST_REPO (optional, defaults to weibaohui/atomgit-cli-e2e-test)

set -e

REPO="${ATOMGIT_TEST_REPO:-weibaohui/atomgit-cli}"

echo "=========================================="
echo "Running All Branch E2E Tests"
echo "Repository: $REPO"
echo "=========================================="

# Build CLI
if [ ! -f "./atomgit" ]; then
    echo "Building atomgit CLI..."
    go build -o ./atomgit .
fi

# Run each test file
for test_file in test_e2e_branch_list_view.sh test_e2e_branch_create.sh test_e2e_branch_delete.sh test_e2e_branch_protect.sh; do
    echo ""
    echo "Running: $test_file"
    echo "------------------------------------------"
    bash "$test_file"
done

echo ""
echo "=========================================="
echo "All Branch E2E Tests Passed!"
echo "=========================================="
