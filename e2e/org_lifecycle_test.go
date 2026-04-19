package e2e

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// TestOrgLifecycle tests org: info -> members
func TestOrgLifecycle(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	org := "AtomGit"

	t.Run("Get org info", func(t *testing.T) {
		cmd := exec.Command(cli, "org", "info", org)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil && strings.Contains(string(output), "404") {
			t.Skip("Org not found")
		}
		if err != nil {
			t.Fatalf("Failed to get org info: %v\n%s", err, output)
		}
	})

	t.Run("List org members", func(t *testing.T) {
		cmd := exec.Command(cli, "org", "members", org)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil && (strings.Contains(string(output), "404") || strings.Contains(string(output), "403")) {
			t.Skip("Org members not accessible")
		}
		if err != nil {
			t.Fatalf("Failed to list org members: %v\n%s", err, output)
		}
	})
}
