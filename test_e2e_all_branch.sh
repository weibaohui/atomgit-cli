#!/bin/bash
# E2E Test: All Branch Commands
# Usage: ./test_e2e_all_branch.sh

set -e

echo "=========================================="
echo "Running All Branch E2E Tests"
echo "=========================================="

# Build once
go build -o ./atomgit .

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
