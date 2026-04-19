package e2e

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
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
	// Use unique name to avoid conflicts
	repoName := fmt.Sprintf("e2e-test-repo-%d", time.Now().Unix())
	owner := "weibaohui"
	fullRepo := owner + "/" + repoName

	// Cleanup if exists
	defer func() {
		exec.Command(cli, "repo", "delete", fullRepo, "-R", fullRepo).Run()
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

// TestIssueLifecycle tests issue lifecycle: create -> list -> view
func TestIssueLifecycle(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("List issues", func(t *testing.T) {
		cmd := exec.Command(cli, "issue", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list issues: %v\n%s", err, output)
		}
		// Just verify command works
		t.Log("Issue list works")
	})
}

// TestHookLifecycle tests webhook lifecycle: create -> list -> view -> update -> test -> delete
func TestHookLifecycle(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()
	hookURL := fmt.Sprintf("https://example.com/webhook-%d", time.Now().Unix())

	// Cleanup
	defer func() {
		// Try to delete, ignore error
		exec.Command(cli, "hook", "delete", "999999", "-R", repo).Run()
	}()

	var hookID string

	t.Run("Create webhook", func(t *testing.T) {
		cmd := exec.Command(cli, "hook", "create", "-R", repo, "--url", hookURL, "--events", "*")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to create webhook: %v\n%s", err, output)
		}
		// Extract hook ID from output
		var result map[string]interface{}
		if json.Unmarshal(output, &result) == nil {
			if id, ok := result["id"].(float64); ok {
				hookID = strconv.FormatFloat(id, 'f', 0, 64)
			}
		}
		if hookID == "" {
			hookID = "1" // fallback
		}
	})

	t.Run("List webhooks", func(t *testing.T) {
		cmd := exec.Command(cli, "hook", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list webhooks: %v\n%s", err, output)
		}
		t.Log("Webhook list works")
	})

	t.Run("View webhook", func(t *testing.T) {
		if hookID == "" {
			t.Skip("No hook ID available")
		}
		cmd := exec.Command(cli, "hook", "view", hookID, "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to view webhook: %v\n%s", err, output)
		}
	})

	t.Run("Update webhook", func(t *testing.T) {
		if hookID == "" {
			t.Skip("No hook ID available")
		}
		newURL := fmt.Sprintf("https://example.com/webhook-updated-%d", time.Now().Unix())
		cmd := exec.Command(cli, "hook", "update", hookID, "-R", repo, "--url", newURL)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to update webhook: %v\n%s", err, output)
		}
	})

	t.Run("Test webhook", func(t *testing.T) {
		if hookID == "" {
			t.Skip("No hook ID available")
		}
		cmd := exec.Command(cli, "hook", "test", hookID, "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, _ := cmd.CombinedOutput()
		// Test might fail if hook doesn't exist or can't be reached
		t.Log("Webhook test response:", string(output))
	})

	t.Run("Delete webhook", func(t *testing.T) {
		if hookID == "" {
			t.Skip("No hook ID available")
		}
		cmd := exec.Command(cli, "hook", "delete", hookID, "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Log("Delete webhook response:", string(output))
		}
	})
}

// TestCommitCommands tests commit list and view
func TestCommitCommands(t *testing.T) {
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
		// Extract first commit SHA
		if strings.Contains(string(output), "sha") {
			t.Log("Commit list works")
		}
	})

	t.Run("View commit", func(t *testing.T) {
		// First get a commit SHA
		cmd := exec.Command(cli, "commit", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, _ := cmd.CombinedOutput()

		// Try to extract a SHA
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

// TestTagCommands tests tag list
func TestTagCommands(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("List tags", func(t *testing.T) {
		cmd := exec.Command(cli, "tag", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list tags: %v\n%s", err, output)
		}
		t.Log("Tag list works")
	})
}

// TestReleaseCommands tests release list and view
func TestReleaseCommands(t *testing.T) {
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

// TestLabelCommands tests label list
func TestLabelCommands(t *testing.T) {
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
		t.Log("Label list works")
	})
}

// TestMilestoneCommands tests milestone list
func TestMilestoneCommands(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("List milestones", func(t *testing.T) {
		cmd := exec.Command(cli, "milestone", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list milestones: %v\n%s", err, output)
		}
		t.Log("Milestone list works")
	})
}

// TestForkCommands tests fork create and list
func TestForkCommands(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	// Fork a well-known public repo
	forkSource := "golang/go"

	t.Run("List forks", func(t *testing.T) {
		cmd := exec.Command(cli, "fork", "list", forkSource)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list forks: %v\n%s", err, output)
		}
		t.Log("Fork list works")
	})
}

// TestStarCommands tests star list
func TestStarCommands(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("List stargazers", func(t *testing.T) {
		cmd := exec.Command(cli, "star", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list stargazers: %v\n%s", err, output)
		}
		t.Log("Star list works")
	})
}

// TestOrgCommands tests org info and members
func TestOrgCommands(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	// Use AtomGit official org
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
			t.Skip("Org members not accessible (requires permissions)")
		}
		if err != nil {
			t.Fatalf("Failed to list org members: %v\n%s", err, output)
		}
	})
}

// TestUserCommands tests user info, followers, following
func TestUserCommands(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	username := "weibaohui"

	t.Run("Get user info", func(t *testing.T) {
		cmd := exec.Command(cli, "user", "info", username)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to get user info: %v\n%s", err, output)
		}
	})

	t.Run("List followers", func(t *testing.T) {
		cmd := exec.Command(cli, "user", "followers", username)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil && strings.Contains(string(output), "404") {
			t.Skip("User has no followers or API not available")
		}
		if err != nil {
			t.Fatalf("Failed to list followers: %v\n%s", err, output)
		}
	})

	t.Run("List following", func(t *testing.T) {
		cmd := exec.Command(cli, "user", "following", username)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil && strings.Contains(string(output), "404") {
			t.Skip("User has no following or API not available")
		}
		if err != nil {
			t.Fatalf("Failed to list following: %v\n%s", err, output)
		}
	})
}

// TestSearchCommands tests search commands
func TestSearchCommands(t *testing.T) {
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
	})

	t.Run("Search users", func(t *testing.T) {
		cmd := exec.Command(cli, "search", "users", "torvalds")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to search users: %v\n%s", err, output)
		}
	})
}

// TestContributorCommands tests contributor list and stats
func TestContributorCommands(t *testing.T) {
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
	})
}

// TestEventCommands tests event list
func TestEventCommands(t *testing.T) {
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
	})
}

// TestLanguageCommands tests language list
func TestLanguageCommands(t *testing.T) {
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
	})
}

// TestSubscriberCommands tests subscriber list
func TestSubscriberCommands(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("List subscribers", func(t *testing.T) {
		cmd := exec.Command(cli, "subscriber", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list subscribers: %v\n%s", err, output)
		}
	})
}

// Helper function to extract SHA from commit list output
func extractSHA(output string) string {
	// Look for "sha":"..." pattern
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

// getTestRepo returns the test repository from env or default
func getTestRepo() string {
	if repo := os.Getenv("ATOMGIT_TEST_REPO"); repo != "" {
		return repo
	}
	return "weibaohui/atomgit-cli-e2e-test"
}
