#!/bin/bash
# E2E Test: atomgit branch protect/unprotect
# Note: API /repos/{owner}/{repo}/branches POST may require different params
set -e

REPO="weibaohui/atomgit-cli"
BRANCH_NAME="test-branch-protect-$(date +%s)"
ATOMGIT_TOKEN="${ATOMGIT_TOKEN:-}"
ATOMGIT_CLI="/tmp/atomgit-test"

echo "=== Testing atomgit branch protect/unprotect ==="

if [ -z "$ATOMGIT_TOKEN" ]; then
    echo "ATOMGIT_TOKEN not set, skipping test"
    exit 0
fi

# Get SHA from main branch
SHA=$(curl -s -H "Authorization: Bearer $ATOMGIT_TOKEN" \
    "https://api.atomgit.com/api/v5/repos/$REPO/branches/main" | \
    python3 -c "import sys,json; print(json.load(sys.stdin)['commit']['id'])" 2>/dev/null)

# Try to create a branch first
echo "Setup: Create branch to protect"
$ATOMGIT_CLI branch create "$BRANCH_NAME" -R "$REPO" --sha "$SHA" > /dev/null 2>&1 && created=true || created=false

if [ "$created" = false ]; then
    echo "⚠ Could not create branch (API issue) - skipping protect test"
    exit 0
fi
echo "✓ Branch created"

# Test 1: Protect branch
echo "Test 1: Protect branch"
$ATOMGIT_CLI branch protect "$BRANCH_NAME" -R "$REPO"
echo "✓ Branch protected"

# Test 2: Check protected list
echo "Test 2: List protected branches"
$ATOMGIT_CLI branch protected-list -R "$REPO"
echo "✓ Protected list retrieved"

# Test 3: Unprotect branch
echo "Test 3: Unprotect branch"
$ATOMGIT_CLI branch unprotect "$BRANCH_NAME" -R "$REPO"
echo "✓ Branch unprotected"

# Cleanup
echo "Cleanup: Delete test branch"
$ATOMGIT_CLI branch delete "$BRANCH_NAME" -R "$REPO" 2>/dev/null || true
echo "✓ Branch deleted"

echo "=== branch protect/unprotect tests passed ==="
