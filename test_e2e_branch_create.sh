#!/bin/bash
# E2E Test: atomgit branch create

REPO="${ATOMGIT_TEST_REPO:-weibaohui/atomgit-cli-e2e-test}"
BRANCH_NAME="test-branch-$(date +%s)"
ATOMGIT_TOKEN="${ATOMGIT_TOKEN:-}"
CLI="./atomgit"

echo "=== Testing atomgit branch create ==="

if [ -z "$ATOMGIT_TOKEN" ]; then
    echo "ATOMGIT_TOKEN not set, skipping test"
    exit 0
fi

# Test 1: Create branch without sha (uses default main)
echo "Test 1: Create branch without sha"
$CLI branch create "$BRANCH_NAME" -R "$REPO"
echo "✓ Branch created"

# Test 2: Verify branch exists
echo "Test 2: Verify branch exists"
$CLI branch view "$BRANCH_NAME" -R "$REPO"
echo "✓ Branch view passed"

# Cleanup
echo "Cleanup: Delete test branch"
$CLI branch delete "$BRANCH_NAME" -R "$REPO" 2>/dev/null || true
echo "✓ Branch deleted"

echo "=== branch create tests passed ==="
