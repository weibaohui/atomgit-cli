package e2e

import (
	"os"
	"os/exec"
	"testing"
)

// TestLabelLifecycle tests label: list
func TestLabelLifecycle(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("List labels", func(t *testing.T) {
		cmd := exec.Command(cli, "label", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list labels: %v\n%s", err, output)
		}
	})
}
