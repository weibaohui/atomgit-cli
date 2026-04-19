#!/bin/bash
# E2E Test: atomgit branch list/view
set -e

REPO="weibaohui/atomgit-cli"
ATOMGIT_TOKEN="${ATOMGIT_TOKEN:-}"
ATOMGIT_CLI="/tmp/atomgit-test"

echo "=== Testing atomgit branch list/view ==="

if [ -z "$ATOMGIT_TOKEN" ]; then
    echo "ATOMGIT_TOKEN not set, skipping test"
    exit 0
fi

# Test 1: List branches
echo "Test 1: List branches"
$ATOMGIT_CLI branch list -R "$REPO"
echo "✓ Branch list retrieved"

# Test 2: View main branch
echo "Test 2: View main branch"
$ATOMGIT_CLI branch view main -R "$REPO"
echo "✓ Branch view passed"

echo "=== branch list/view tests passed ==="
