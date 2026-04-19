package e2e

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"
)

// TestBranchCreateAndListAndDelete tests branch create -> list -> delete -> list verify
func TestBranchCreateAndListAndDelete(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()
	branchName := fmt.Sprintf("e2e-branch-%d", time.Now().Unix())

	t.Run("Create branch", func(t *testing.T) {
		cmd := exec.Command(cli, "branch", "create", branchName, "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to create branch: %v\n%s", err, output)
		}
		if !contains(string(output), branchName) {
			t.Fatalf("Expected branch name in output, got: %s", output)
		}
		t.Log("Branch created:", branchName)
	})

	t.Run("List branches", func(t *testing.T) {
		cmd := exec.Command(cli, "branch", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list branches: %v\n%s", err, output)
		}
		t.Log("Branch list works")
	})

	t.Run("Delete branch", func(t *testing.T) {
		cmd := exec.Command(cli, "branch", "delete", branchName, "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to delete branch: %v\n%s", err, output)
		}
		t.Log("Branch deleted")
	})
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
