package e2e

import (
	"os"
	"os/exec"
	"testing"
)

// TestLanguageList tests language list
func TestLanguageList(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("List languages", func(t *testing.T) {
		cmd := exec.Command(cli, "language", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list languages: %v\n%s", err, output)
		}
		t.Log("Language list works")
	})
}
