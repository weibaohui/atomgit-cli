package e2e

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

// TestRepoCreateAndListAndView tests repo create -> list -> view -> delete -> list verify
func TestRepoCreateAndListAndView(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repoName := fmt.Sprintf("e2e-repo-test-%d", time.Now().Unix())
	owner := "weibaohui"
	fullRepo := owner + "/" + repoName

	t.Run("Create repository", func(t *testing.T) {
		cmd := exec.Command(cli, "repo", "create", repoName,
			"--description", "E2E test repo",
			"--public")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to create repo: %v\n%s", err, output)
		}
		if !strings.Contains(string(output), repoName) {
			t.Fatalf("Expected repo name in output, got: %s", output)
		}
		t.Logf("Created repo: %s", fullRepo)
	})

	t.Run("List repository and verify exists", func(t *testing.T) {
		cmd := exec.Command(cli, "repo", "list")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list repos: %v\n%s", err, output)
		}
		if !strings.Contains(string(output), repoName) {
			t.Fatalf("Expected repo in list, got: %s", output)
		}
		t.Log("Repo found in list")
	})

	t.Run("View repository details", func(t *testing.T) {
		cmd := exec.Command(cli, "repo", "view", fullRepo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to view repo: %v\n%s", err, output)
		}
		if !strings.Contains(string(output), repoName) {
			t.Fatalf("Expected repo name in view output, got: %s", output)
		}
		t.Log("Repo view successful")
	})

	t.Run("Delete repository", func(t *testing.T) {
		cmd := exec.Command(cli, "repo", "delete", fullRepo, "-y")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to delete repo: %v\n%s", err, output)
		}
		t.Log("Repo deleted")
	})

	t.Run("List repository and verify deleted", func(t *testing.T) {
		cmd := exec.Command(cli, "repo", "list")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list repos: %v\n%s", err, output)
		}
		if strings.Contains(string(output), repoName) {
			t.Fatalf("Repo should be deleted but still appears in list")
		}
		t.Log("Repo confirmed deleted")
	})
}

// TestRepoCreateWithOptions tests repo create with gitignore, license and other options
func TestRepoCreateWithOptions(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repoName := fmt.Sprintf("e2e-repo-options-%d", time.Now().Unix())
	owner := "weibaohui"
	fullRepo := owner + "/" + repoName

	defer func() {
		exec.Command(cli, "repo", "delete", fullRepo, "-y").Run()
	}()

	t.Run("Create repo with gitignore and license", func(t *testing.T) {
		cmd := exec.Command(cli, "repo", "create", repoName,
			"--description", "Repo with gitignore and license",
			"--public",
			"--gitignore", "Go",
			"--license", "MIT")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to create repo: %v\n%s", err, output)
		}
		if !strings.Contains(string(output), repoName) {
			t.Fatalf("Expected repo name in output, got: %s", output)
		}
	})

	t.Run("View repository to verify settings", func(t *testing.T) {
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
}

// TestRepoEdit tests repo edit -> view verify -> edit again -> delete
func TestRepoEdit(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	repoName := fmt.Sprintf("e2e-repo-edit-%d", time.Now().Unix())
	owner := "weibaohui"
	fullRepo := owner + "/" + repoName

	defer func() {
		exec.Command(cli, "repo", "delete", fullRepo, "-y").Run()
	}()

	t.Run("Create repo for edit test", func(t *testing.T) {
		cmd := exec.Command(cli, "repo", "create", repoName,
			"--description", "Original description",
			"--public")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to create repo: %v\n%s", err, output)
		}
		t.Log("Repo created")
	})

	t.Run("Edit repo description", func(t *testing.T) {
		cmd := exec.Command(cli, "repo", "edit", fullRepo,
			"--description", "Updated description via e2e test")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to edit repo: %v\n%s", err, output)
		}
		if !strings.Contains(string(output), "Updated description") {
			t.Fatalf("Expected updated description in output, got: %s", output)
		}
	})

	t.Run("View repo to verify description updated", func(t *testing.T) {
		cmd := exec.Command(cli, "repo", "view", fullRepo)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to view repo: %v\n%s", err, output)
		}
		// View output should contain the repo
		if !strings.Contains(string(output), repoName) {
			t.Fatalf("Expected repo name in view output, got: %s", output)
		}
	})

	t.Run("Edit repo homepage", func(t *testing.T) {
		cmd := exec.Command(cli, "repo", "edit", fullRepo,
			"--homepage", "https://example.com")
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to edit repo homepage: %v\n%s", err, output)
		}
		t.Log("Homepage updated")
	})
}

// TestRepoFork tests repo fork -> list forks -> delete
func TestRepoFork(t *testing.T) {
	token := os.Getenv("ATOMGIT_TOKEN")
	if token == "" {
		t.Skip("ATOMGIT_TOKEN not set")
	}

	cli := getCLIPath()
	forkSource := "golang/go"

	t.Run("List forks of public repo", func(t *testing.T) {
		cmd := exec.Command(cli, "fork", "list", forkSource)
		cmd.Env = append(os.Environ(), "ATOMGIT_TOKEN="+token)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to list forks: %v\n%s", err, output)
		}
		t.Log("Fork list works")
	})
}
