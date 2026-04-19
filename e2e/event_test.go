package e2e

import (
	"os"
	"os/exec"
	"testing"
)

// TestEventList tests event list
func TestEventList(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("List events", func(t *testing.T) {
		cmd := exec.Command(cli, "event", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list events: %v\n%s", err, output)
		}
		t.Log("Event list works")
	})
}
