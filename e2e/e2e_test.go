package e2e

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Build CLI before running tests
	cli := getCLIPath()
	if _, err := os.Stat(cli); err != nil {
		// CLI not found, tests will skip or fail gracefully
	}
	os.Exit(m.Run())
}
