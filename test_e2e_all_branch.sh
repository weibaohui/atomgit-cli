#!/bin/bash
# E2E Test: Branch Lifecycle Test Runner
# Forms a complete verifiable闭环: create -> list -> view -> protect -> list protected -> unprotect -> delete -> verify

set -e

REPO="${ATOMGIT_TEST_REPO:-weibaohui/atomgit-cli}"
ATOMGIT_TOKEN="${ATOMGIT_TOKEN:-}"

echo "=========================================="
echo "Branch Lifecycle E2E Test Suite"
echo "Repository: $REPO"
echo "=========================================="
echo ""

if [ -z "$ATOMGIT_TOKEN" ]; then
    echo "ERROR: ATOMGIT_TOKEN not set"
    exit 1
fi

# Build CLI if needed
if [ ! -f "./atomgit" ]; then
    echo "Building atomgit CLI..."
    go build -o ./atomgit .
fi

# Run the lifecycle test
bash test_e2e_branch_lifecycle.sh

echo ""
echo "=========================================="
echo "Branch E2E tests completed!"
echo "=========================================="
