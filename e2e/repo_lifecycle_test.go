package e2e

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

// TestRepoLifecycle tests repository lifecycle: create -> list -> view -> delete
func TestRepoLifecycle(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repoName := fmt.Sprintf("e2e-test-repo-%d", time.Now().Unix())
	owner := "weibaohui"
	fullRepo := owner + "/" + repoName

	// Cleanup if exists
	defer func() {
		exec.Command(cli, "repo", "delete", fullRepo, "-y").Run()
	}()

	t.Run("Create repository", func(t *testing.T) {
		cmd := exec.Command(cli, "repo", "create", repoName, "--description", "E2E test repo", "--private")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to create repo: %v\n%s", err, output)
		}
		if !strings.Contains(string(output), repoName) {
			t.Fatalf("Expected repo name in output, got: %s", output)
		}
	})

	t.Run("List repository", func(t *testing.T) {
		cmd := exec.Command(cli, "repo", "list")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list repos: %v\n%s", err, output)
		}
		if !strings.Contains(string(output), repoName) {
			t.Fatalf("Expected repo in list, got: %s", output)
		}
	})

	t.Run("View repository", func(t *testing.T) {
		cmd := exec.Command(cli, "repo", "view", fullRepo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to view repo: %v\n%s", err, output)
		}
		if !strings.Contains(string(output), repoName) {
			t.Fatalf("Expected repo name in output, got: %s", output)
		}
	})

	t.Run("Delete repository", func(t *testing.T) {
		cmd := exec.Command(cli, "repo", "delete", fullRepo, "-y")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to delete repo: %v\n%s", err, output)
		}
	})
}
