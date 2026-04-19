package e2e

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// TestContributorListAndStats tests contributor list and stats
func TestContributorListAndStats(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("List contributors", func(t *testing.T) {
		cmd := exec.Command(cli, "contributor", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list contributors: %v\n%s", err, output)
		}
		t.Log("Contributor list works")
	})

	t.Run("Contributor stats", func(t *testing.T) {
		cmd := exec.Command(cli, "contributor", "stats", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil && strings.Contains(string(output), "404") {
			t.Skip("Contributor stats not available for this repo")
		}
		if err != nil {
			t.Fatalf("Failed to get contributor stats: %v\n%s", err, output)
		}
		t.Log("Contributor stats works")
	})
}
