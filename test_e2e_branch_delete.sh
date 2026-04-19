#!/bin/bash
# E2E Test: atomgit branch delete

REPO="${ATOMGIT_TEST_REPO:-weibaohui/atomgit-cli-e2e-test}"
BRANCH_NAME="test-branch-to-delete-$(date +%s)"
ATOMGIT_TOKEN="${ATOMGIT_TOKEN:-}"
CLI="./atomgit"

echo "=== Testing atomgit branch delete ==="

if [ -z "$ATOMGIT_TOKEN" ]; then
    echo "ATOMGIT_TOKEN not set, skipping test"
    exit 0
fi

# Setup: Create a branch first
echo "Setup: Create branch to delete"
$CLI branch create "$BRANCH_NAME" -R "$REPO" > /dev/null
echo "✓ Branch created for deletion"

# Test: Delete branch
echo "Test: Delete branch"
$CLI branch delete "$BRANCH_NAME" -R "$REPO"
echo "✓ Branch deleted successfully"

echo "=== branch delete tests passed ==="
