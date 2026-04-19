package e2e

import (
	"os"
	"os/exec"
	"testing"
)

// TestSearchReposAndUsers tests search repos and users
func TestSearchReposAndUsers(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()

	t.Run("Search repos", func(t *testing.T) {
		cmd := exec.Command(cli, "search", "repos", "golang")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to search repos: %v\n%s", err, output)
		}
		t.Log("Search repos works")
	})

	t.Run("Search users", func(t *testing.T) {
		cmd := exec.Command(cli, "search", "users", "torvalds")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to search users: %v\n%s", err, output)
		}
		t.Log("Search users works")
	})
}
