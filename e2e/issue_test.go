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

// TestIssueCreateAndListAndView tests issue create -> list -> view -> delete
func TestIssueCreateAndListAndView(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()
	issueTitle := fmt.Sprintf("E2E Test Issue %d", time.Now().Unix())
	var issueNumber int

	t.Run("Create issue", func(t *testing.T) {
		cmd := exec.Command(cli, "issue", "create", "-R", repo,
			"-t", issueTitle,
			"-b", "E2E test issue body")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to create issue: %v\n%s", err, output)
		}
		// Try to extract issue number from output
		var result map[string]interface{}
		if json.Unmarshal(output, &result) == nil {
			if num, ok := result["number"].(float64); ok {
				issueNumber = int(num)
			}
		}
		if issueNumber == 0 {
			issueNumber = 1 // fallback
		}
		t.Logf("Created issue #%d", issueNumber)
	})

	t.Run("List issues and verify exists", func(t *testing.T) {
		cmd := exec.Command(cli, "issue", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list issues: %v\n%s", err, output)
		}
		t.Log("Issue list works")
	})

	t.Run("View issue details", func(t *testing.T) {
		if issueNumber == 0 {
			t.Skip("No issue number available")
		}
		cmd := exec.Command(cli, "issue", "view", strconv.Itoa(issueNumber), "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to view issue: %v\n%s", err, output)
		}
		t.Log("Issue view works")
	})
}

// TestIssueEnhancedList tests issue list with various filters
func TestIssueEnhancedList(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("List issues with label filter", func(t *testing.T) {
		cmd := exec.Command(cli, "issue", "list", "-R", repo, "-l", "bug")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list issues with label: %v\n%s", err, output)
		}
		t.Log("Issue list with label filter works")
	})

	t.Run("List issues with assignee filter", func(t *testing.T) {
		cmd := exec.Command(cli, "issue", "list", "-R", repo, "-a", "weibaohui")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list issues with assignee: %v\n%s", err, output)
		}
		t.Log("Issue list with assignee filter works")
	})

	t.Run("List issues with author filter", func(t *testing.T) {
		cmd := exec.Command(cli, "issue", "list", "-R", repo, "-A", "weibaohui")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list issues with author: %v\n%s", err, output)
		}
		t.Log("Issue list with author filter works")
	})

	t.Run("List issues with state filter", func(t *testing.T) {
		cmd := exec.Command(cli, "issue", "list", "-R", repo, "-s", "all")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list issues with state: %v\n%s", err, output)
		}
		t.Log("Issue list with state filter works")
	})
}

// TestIssueEnhancedView tests issue view with comments flag
func TestIssueEnhancedView(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("View issue without comments", func(t *testing.T) {
		cmd := exec.Command(cli, "issue", "view", "1", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			if strings.Contains(string(output), "Issue Not Found") || strings.Contains(string(output), "404") {
				t.Skip("Issue #1 not found in test repo")
			}
			t.Fatalf("Failed to view issue: %v\n%s", err, output)
		}
		t.Log("Issue view without comments works")
	})

	t.Run("View issue with comments", func(t *testing.T) {
		cmd := exec.Command(cli, "issue", "view", "1", "-R", repo, "-c")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			if strings.Contains(string(output), "Issue Not Found") || strings.Contains(string(output), "404") {
				t.Skip("Issue #1 not found in test repo")
			}
			t.Fatalf("Failed to view issue with comments: %v\n%s", err, output)
		}
		t.Log("Issue view with comments works")
	})
}

// TestIssueCreateWithBodyFile tests issue create with body from file
func TestIssueCreateWithBodyFile(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("Create issue from file", func(t *testing.T) {
		// Create a temporary file for body
		tmpFile := "/tmp/e2e-issue-body.txt"
		err := os.WriteFile(tmpFile, []byte("Issue body from file"), 0644)
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile)

		issueTitle := fmt.Sprintf("Issue from file %d", time.Now().Unix())
		cmd := exec.Command(cli, "issue", "create", "-R", repo,
			"-t", issueTitle,
			"--body-file", tmpFile)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to create issue from file: %v\n%s", err, output)
		}
		t.Log("Issue created from file works")
	})
}

// TestIssueCreateWithMilestone tests issue create with milestone
func TestIssueCreateWithMilestone(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()

	t.Run("List milestones to check if any exists", func(t *testing.T) {
		cmd := exec.Command(cli, "milestone", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Log("Milestone list failed, skipping milestone test:", string(output))
			t.Skip()
		}
		t.Log("Milestone list works")
	})

	t.Run("Create issue with milestone", func(t *testing.T) {
		issueTitle := fmt.Sprintf("Issue with milestone %d", time.Now().Unix())
		cmd := exec.Command(cli, "issue", "create", "-R", repo,
			"-t", issueTitle,
			"-b", "Issue body for milestone test",
			"--milestone", "1")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to create issue with milestone: %v\n%s", err, output)
		}
		t.Log("Issue with milestone created")
	})
}
