package e2e

import (
	"os"
	"os/exec"
	"testing"
)

// TestStarLifecycle tests star: list
func TestStarLifecycle(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("List stargazers", func(t *testing.T) {
		cmd := exec.Command(cli, "star", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list stargazers: %v\n%s", err, output)
		}
	})
}
