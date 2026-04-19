package e2e

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// TestCommitLifecycle tests commit: list -> view
func TestCommitLifecycle(t *testing.T) {
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
	})
}

// extractSHA extracts SHA from commit list output
func extractSHA(output string) string {
	parts := strings.Split(output, `"sha":"`)
	if len(parts) < 2 {
		return ""
	}
	sha := parts[1]
	for i, c := range sha {
		if c == '"' {
			return sha[:i]
		}
	}
	return ""
}
