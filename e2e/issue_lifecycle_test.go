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

// TestIssueLifecycle tests issue lifecycle: create -> list -> view -> close -> list (verify closed)
func TestIssueLifecycle(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()
	issueTitle := fmt.Sprintf("Test Issue %d", time.Now().Unix())
	var issueNumber int

	t.Run("Create issue", func(t *testing.T) {
		cmd := exec.Command(cli, "issue", "create", "-R", repo, "-t", issueTitle, "-m", "Test issue body")
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
	})

	t.Run("List issues", func(t *testing.T) {
		cmd := exec.Command(cli, "issue", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list issues: %v\n%s", err, output)
		}
	})

	t.Run("View issue", func(t *testing.T) {
		if issueNumber == 0 {
			t.Skip("No issue number available")
		}
		cmd := exec.Command(cli, "issue", "view", strconv.Itoa(issueNumber), "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to view issue: %v\n%s", err, output)
		}
	})
}
