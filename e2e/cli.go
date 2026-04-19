package e2e

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// getCLIPath returns the path to the atomgit CLI binary
func getCLIPath() string {
	// Try current directory first
	if _, err := os.Stat("./atomgit"); err == nil {
		return "./atomgit"
	}
	// Try parent directory (running from e2e subdir)
	if _, err := os.Stat("../atomgit"); err == nil {
		return "../atomgit"
	}
	// Fallback to current dir
	return "./atomgit"
}

// BuildCLI builds the CLI if needed and returns the path
func BuildCLI() (string, error) {
	cli := getCLIPath()
	if _, err := os.Stat(cli); err != nil {
		// Find project root (parent of e2e dir)
		execDir, _ := os.Getwd()
		projectRoot := filepath.Dir(execDir)
		cmd := exec.Command("go", "build", "-o", cli)
		cmd.Dir = projectRoot
		_, err = cmd.CombinedOutput()
		if err != nil {
			return "", err
		}
	}
	return cli, nil
}

// EnsureTestRepo creates the test repository if it doesn't exist
func EnsureTestRepo(owner, reponame string) error {
	cli, err := BuildCLI()
	if err != nil {
		return err
	}
	cmd := exec.Command(cli, "repo", "view", owner+"/"+reponame)
	out, err := cmd.CombinedOutput()
	if err != nil && !strings.Contains(string(out), "404") {
		// Create the repo
		cmd = exec.Command(cli, "repo", "create", owner+"/"+reponame, "--description", "E2E testing repository", "--private")
		cmd.Run()
	}
	return nil
}

// getTestRepo returns the test repository from env or default
func getTestRepo() string {
	if repo := os.Getenv("ATOMGIT_TEST_REPO"); repo != "" {
		return repo
	}
	return "weibaohui/atomgit-cli-e2e-test"
}

// Helper function to extract SHA from commit list output
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

// extractNumber extracts a number from JSON response
func extractNumber(s string) string {
	parts := strings.Split(s, `"number":`)
	if len(parts) < 2 {
		return ""
	}
	numStr := strings.TrimSpace(parts[1])
	for i, c := range numStr {
		if c < '0' || c > '9' {
			return numStr[:i]
		}
	}
	return numStr
}
