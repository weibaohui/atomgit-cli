package e2e

import (
	"os"
	"os/exec"
	"testing"
)

// TestTagLifecycle tests tag: list
func TestTagLifecycle(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("List tags", func(t *testing.T) {
		cmd := exec.Command(cli, "tag", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list tags: %v\n%s", err, output)
		}
	})
}
