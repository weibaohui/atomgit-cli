package e2e

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

// TestBranchLifecycle tests the complete branch lifecycle
func TestBranchLifecycle(t *testing.T) {
	repo := os.Getenv("ATOMGIT_TEST_REPO")
	if repo == "" {
		repo = "weibaohui/atomgit-e2e"
	}

	cli := getCLIPath()
	testBranch := fmt.Sprintf("test-branch-%d", time.Now().Unix())

	// Step 1: Create branch
	t.Log("[Step 1] Creating branch:", testBranch)
	cmd := exec.Command(cli, "branch", "create", testBranch, "-R", repo)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to create branch: %v\nOutput: %s", err, output)
	}
	if !strings.Contains(string(output), testBranch) {
		t.Fatalf("Branch creation response missing branch name")
	}
	t.Log("✓ Branch created")

	// Step 2: List branches (verify new branch appears)
	t.Log("[Step 2] Listing branches")
	cmd = exec.Command(cli, "branch", "list", "-R", repo)
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to list branches: %v\nOutput: %s", err, output)
	}
	if !strings.Contains(string(output), testBranch) {
		t.Fatalf("New branch not found in list")
	}
	t.Log("✓ New branch appears in list")

	// Step 3: View branch details
	t.Log("[Step 3] Viewing branch:", testBranch)
	cmd = exec.Command(cli, "branch", "view", testBranch, "-R", repo)
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to view branch: %v\nOutput: %s", err, output)
	}
	if !strings.Contains(string(output), testBranch) {
		t.Fatalf("Branch view response missing branch name")
	}
	t.Log("✓ Branch view successful")

	// Step 4: Protect branch
	t.Log("[Step 4] Protecting branch")
	cmd = exec.Command(cli, "branch", "protect", testBranch, "-R", repo)
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to protect branch: %v\nOutput: %s", err, output)
	}
	t.Log("✓ Branch protected")

	// Step 5: List protected branches (verify it appears)
	t.Log("[Step 5] Listing protected branches")
	cmd = exec.Command(cli, "branch", "protected-list", "-R", repo)
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to list protected branches: %v\nOutput: %s", err, output)
	}
	if !strings.Contains(string(output), testBranch) {
		t.Fatalf("Protected branch not found in protected list")
	}
	t.Log("✓ Protected branch appears in protected list")

	// Step 6: Unprotect branch
	t.Log("[Step 6] Unprotecting branch")
	cmd = exec.Command(cli, "branch", "unprotect", testBranch, "-R", repo)
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to unprotect branch: %v\nOutput: %s", err, output)
	}
	t.Log("✓ Branch unprotected")

	// Step 7: Verify unprotect
	t.Log("[Step 7] Verifying unprotect")
	cmd = exec.Command(cli, "branch", "protected-list", "-R", repo)
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to list protected branches: %v\nOutput: %s", err, output)
	}
	if strings.Contains(string(output), testBranch) {
		t.Fatalf("Branch still in protected list after unprotect")
	}
	t.Log("✓ Branch removed from protected list")

	// Step 8: Delete branch
	t.Log("[Step 8] Deleting branch")
	cmd = exec.Command(cli, "branch", "delete", testBranch, "-R", repo)
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to delete branch: %v\nOutput: %s", err, output)
	}
	t.Log("✓ Branch deleted")

	// Step 9: List branches (verify deletion)
	t.Log("[Step 9] Listing branches after deletion")
	cmd = exec.Command(cli, "branch", "list", "-R", repo)
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to list branches: %v\nOutput: %s", err, output)
	}
	if strings.Contains(string(output), testBranch) {
		t.Fatalf("Branch still exists in list after deletion")
	}
	t.Log("✓ Branch removed from list")

	// Step 10: View deleted branch (should fail)
	t.Log("[Step 10] Viewing deleted branch (should fail)")
	cmd = exec.Command(cli, "branch", "view", testBranch, "-R", repo)
	output, err = cmd.CombinedOutput()
	if err == nil || (!strings.Contains(string(output), "not found") && !strings.Contains(string(output), "404")) {
		t.Fatalf("Deleted branch should return not found error")
	}
	t.Log("✓ Branch correctly returns not found")

	t.Log("=== Branch lifecycle test PASSED ===")
}
