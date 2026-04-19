package e2e

import (
	"os"
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
