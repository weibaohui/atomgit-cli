#!/bin/bash
# E2E Test: atomgit PR Commands - Master Test Runner
# Uses repository: weibaohui/atomgit-cli

set -e

REPO="${ATOMGIT_TEST_REPO:-weibaohui/atomgit-cli}"
ATOMGIT_TOKEN="${ATOMGIT_TOKEN:-}"

echo "============================================"
echo "atomgit PR Commands E2E Test Suite"
echo "Repository: $REPO"
echo "============================================"
echo ""

if [ -z "$ATOMGIT_TOKEN" ]; then
    echo "ERROR: ATOMGIT_TOKEN not set"
    exit 1
fi

# Run PR lifecycle test (complete闭环)
echo "Running PR Lifecycle Test..."
echo "--------------------------------------------"
bash test_e2e_pr_lifecycle.sh

echo ""
echo "============================================"
echo "All PR E2E tests completed!"
echo "============================================"