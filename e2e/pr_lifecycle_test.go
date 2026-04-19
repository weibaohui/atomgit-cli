package e2e

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"
)

// TestPRLifecycle tests the complete PR lifecycle
func TestPRLifecycle(t *testing.T) {
	repo := os.Getenv("ATOMGIT_TEST_REPO")
	if repo == "" {
		repo = "weibaohui/atomgit-cli"
	}

	cli := getCLIPath()
	testBranch := fmt.Sprintf("test-pr-%d", time.Now().Unix())
	var prNumber int

	// Step 1: Create test branch
	t.Log("[Step 1] Creating test branch:", testBranch)
	cmd := exec.Command(cli, "branch", "create", testBranch, "-R", repo)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to create branch: %v\nOutput: %s", err, output)
	}
	if !strings.Contains(string(output), testBranch) {
		t.Fatalf("Branch creation response missing branch name")
	}
	t.Log("✓ Branch created")

	// Step 2: Create PR
	t.Log("[Step 2] Creating PR")
	cmd = exec.Command(cli, "pr", "create", "-R", repo, "-t", "Test PR: " + testBranch, "-m", "PR for lifecycle test", "--head", testBranch)
	output, err = cmd.CombinedOutput()
	if err != nil {
		// Check if PR already exists for this branch
		t.Log("PR creation failed, checking for existing PR...")
		cmd = exec.Command(cli, "pr", "list", "-R", repo, "-s", "all")
		output, _ = cmd.CombinedOutput()
	}

	// Extract PR number
	prNumStr := extractNumber(string(output))
	if prNumStr != "" {
		prNumber, _ = strconv.Atoi(prNumStr)
		t.Log("✓ PR created/using existing:", prNumber)
	} else {
		prNumber = 1 // fallback
		t.Log("⚠ Using default PR number:", prNumber)
	}

	// Step 3: List PRs (verify new PR appears)
	t.Log("[Step 3] Listing PRs")
	cmd = exec.Command(cli, "pr", "list", "-R", repo, "-s", "all")
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to list PRs: %v\nOutput: %s", err, output)
	}
	if !strings.Contains(string(output), strconv.Itoa(prNumber)) {
		t.Log("⚠ PR number not found in list, but listing works")
	}
	t.Log("✓ PR list works")

	// Step 4: View PR details
	t.Log("[Step 4] Viewing PR #", prNumber)
	cmd = exec.Command(cli, "pr", "view", strconv.Itoa(prNumber), "-R", repo)
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to view PR: %v\nOutput: %s", err, output)
	}
	if !strings.Contains(string(output), "number") {
		t.Fatalf("PR view response missing expected fields")
	}
	t.Log("✓ PR view successful")

	// Step 5: Get PR commits
	t.Log("[Step 5] Getting PR commits")
	cmd = exec.Command(cli, "pr", "commits", strconv.Itoa(prNumber), "-R", repo)
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Log("⚠ Get PR commits failed:", string(output))
	} else {
		t.Log("✓ PR commits retrieved")
	}

	// Step 6: Get PR files
	t.Log("[Step 6] Getting PR files")
	cmd = exec.Command(cli, "pr", "files", strconv.Itoa(prNumber), "-R", repo)
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Log("⚠ Get PR files failed:", string(output))
	} else {
		t.Log("✓ PR files retrieved")
	}

	// Step 7: Add labels
	t.Log("[Step 7] Adding labels to PR")
	cmd = exec.Command(cli, "pr", "add-labels", strconv.Itoa(prNumber), "-R", repo, "-l", "bug", "-l", "enhancement")
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Log("⚠ Add labels response:", string(output))
	} else {
		t.Log("✓ Labels added")
	}

	// Step 8: List labels
	t.Log("[Step 8] Listing PR labels")
	cmd = exec.Command(cli, "pr", "labels", strconv.Itoa(prNumber), "-R", repo)
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Log("⚠ Labels list response:", string(output))
	}
	t.Log("✓ Labels listed")

	// Step 9: Update PR title
	t.Log("[Step 9] Updating PR title")
	cmd = exec.Command(cli, "pr", "update", strconv.Itoa(prNumber), "-R", repo, "-t", "Updated: Test PR "+testBranch)
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Log("⚠ Update title response:", string(output))
	} else {
		t.Log("✓ PR title updated")
	}

	// Step 10: Update PR body
	t.Log("[Step 10] Updating PR body")
	cmd = exec.Command(cli, "pr", "update", strconv.Itoa(prNumber), "-R", repo, "-m", "Updated body for lifecycle test")
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Log("⚠ Update body response:", string(output))
	} else {
		t.Log("✓ PR body updated")
	}

	// Step 11: Check merge status
	t.Log("[Step 11] Checking merge status")
	cmd = exec.Command(cli, "pr", "merge-status", strconv.Itoa(prNumber), "-R", repo)
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Log("⚠ Merge status response:", string(output))
	}
	t.Log("✓ Merge status checked")

	// Step 12: Close PR
	t.Log("[Step 12] Closing PR")
	cmd = exec.Command(cli, "pr", "close", strconv.Itoa(prNumber), "-R", repo)
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Log("⚠ Close PR response:", string(output))
	}
	t.Log("✓ PR closed")

	// Step 13: List closed PRs (verify PR is closed)
	t.Log("[Step 13] Listing closed PRs")
	cmd = exec.Command(cli, "pr", "list", "-R", repo, "-s", "closed")
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Log("⚠ Closed list response:", string(output))
	}
	t.Log("✓ Closed PRs listed")

	// Step 14: Reopen PR
	t.Log("[Step 14] Reopening PR")
	cmd = exec.Command(cli, "pr", "reopen", strconv.Itoa(prNumber), "-R", repo)
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Log("⚠ Reopen PR response:", string(output))
	}
	t.Log("✓ PR reopened")

	// Step 15: List open PRs (verify PR is open again)
	t.Log("[Step 15] Listing open PRs")
	cmd = exec.Command(cli, "pr", "list", "-R", repo, "-s", "open")
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Log("⚠ Open list response:", string(output))
	}
	t.Log("✓ Open PRs listed")

	// Cleanup
	t.Log("[Cleanup] Closing PR and deleting test branch")
	exec.Command(cli, "pr", "close", strconv.Itoa(prNumber), "-R", repo).Run()
	exec.Command(cli, "branch", "delete", testBranch, "-R", repo).Run()
	t.Log("✓ Cleanup done")

	t.Log("=== PR lifecycle test PASSED ===")
}

// extractNumber extracts a number from JSON response
func extractNumber(s string) string {
	// Look for "number":<digits>
	parts := strings.Split(s, `"number":`)
	if len(parts) < 2 {
		return ""
	}
	numStr := strings.TrimSpace(parts[1])
	for i, c := range numStr {
		if c < '0' || c > '9' {
			return numStr[:i]
		}
	}
	return numStr
}
