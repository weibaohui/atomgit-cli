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

// TestHookLifecycle tests webhook lifecycle: create -> list -> view -> update -> test -> delete
func TestHookLifecycle(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repo := getTestRepo()
	hookURL := fmt.Sprintf("https://example.com/webhook-%d", time.Now().Unix())
	var hookID string

	// Cleanup
	defer func() {
		if hookID != "" {
			exec.Command(cli, "hook", "delete", hookID, "-R", repo).Run()
		}
	}()

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
			hookID = "1"
		}
	})

	t.Run("List webhooks", func(t *testing.T) {
		cmd := exec.Command(cli, "hook", "list", "-R", repo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list webhooks: %v\n%s", err, output)
		}
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
		hookID = "" // Mark as deleted
	})
}
