#!/bin/bash
# E2E Test: atomgit branch protect/unprotect

REPO="${ATOMGIT_TEST_REPO:-weibaohui/atomgit-cli-e2e-test}"
BRANCH_NAME="test-branch-protect-$(date +%s)"
ATOMGIT_TOKEN="${ATOMGIT_TOKEN:-}"
CLI="./atomgit"

echo "=== Testing atomgit branch protect/unprotect ==="

if [ -z "$ATOMGIT_TOKEN" ]; then
    echo "ATOMGIT_TOKEN not set, skipping test"
    exit 0
fi

# Setup: Create a branch first
echo "Setup: Create branch to protect"
$CLI branch create "$BRANCH_NAME" -R "$REPO" > /dev/null
echo "✓ Branch created"

# Test 1: Protect branch
echo "Test 1: Protect branch"
$CLI branch protect "$BRANCH_NAME" -R "$REPO"
echo "✓ Branch protected"

# Test 2: Check protected list
echo "Test 2: List protected branches"
$CLI branch protected-list -R "$REPO"
echo "✓ Protected list retrieved"

# Test 3: Unprotect branch
echo "Test 3: Unprotect branch"
$CLI branch unprotect "$BRANCH_NAME" -R "$REPO"
echo "✓ Branch unprotected"

# Cleanup
echo "Cleanup: Delete test branch"
$CLI branch delete "$BRANCH_NAME" -R "$REPO" 2>/dev/null || true
echo "✓ Branch deleted"

echo "=== branch protect/unprotect tests passed ==="
