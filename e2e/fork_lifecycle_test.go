package e2e

import (
	"os"
	"os/exec"
	"testing"
)

// TestForkLifecycle tests fork: list
func TestForkLifecycle(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	forkSource := "golang/go"

	t.Run("List forks", func(t *testing.T) {
		cmd := exec.Command(cli, "fork", "list", forkSource)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list forks: %v\n%s", err, output)
		}
	})
}
