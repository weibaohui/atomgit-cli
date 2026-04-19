package e2e

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// TestUserInfoAndFollowersAndFollowing tests user info, followers, following
func TestUserInfoAndFollowersAndFollowing(t *testing.T) {
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
		t.Log("User info works")
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
		t.Log("Followers list works")
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
		t.Log("Following list works")
	})
}
