package e2e

import (
	"os"
	"os/exec"
	"testing"
)

// TestCommitListAndView tests commit list -> view
func TestCommitListAndView(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("List commits", func(t *testing.T) {
		cmd := exec.Command(cli, "commit", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list commits: %v\n%s", err, output)
		}
		if contains(string(output), "sha") {
			t.Log("Commit list works")
		}
	})

	t.Run("View commit", func(t *testing.T) {
		// First get a commit SHA
		cmd := exec.Command(cli, "commit", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, _ := cmd.CombinedOutput()

		sha := extractSHA(string(output))
		if sha == "" {
			t.Skip("No commit SHA available")
		}

		cmd = exec.Command(cli, "commit", "view", sha, "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to view commit: %v\n%s", err, output)
		}
		t.Log("Commit view works")
	})
}
