package e2e

import (
	"os"
	"os/exec"
	"testing"
)

// TestReleaseList tests release list
func TestReleaseList(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("List releases", func(t *testing.T) {
		cmd := exec.Command(cli, "release", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list releases: %v\n%s", err, output)
		}
		t.Log("Release list works")
	})
}
