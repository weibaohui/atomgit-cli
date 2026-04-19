package e2e

import (
	"os"
	"os/exec"
	"testing"
)

// TestMilestoneList tests milestone list
func TestMilestoneList(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("List milestones", func(t *testing.T) {
		cmd := exec.Command(cli, "milestone", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list milestones: %v\n%s", err, output)
		}
		t.Log("Milestone list works")
	})
}
