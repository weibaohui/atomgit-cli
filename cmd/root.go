package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "atomgit",
	Short: "AtomGit CLI (gh-like)",
	Long: `AtomGit CLI for interacting with the AtomGit platform.

EXAMPLES
  $ atomgit repo create my-project --public
  $ atomgit repo list
  $ atomgit issue list -R owner/repo
  $ atomgit pr list -R owner/repo
  $ atomgit api /api/v1/user

LEARN MORE
  Use 'atomgit <command> --help' for more information about a command.`,
	Version: "0.1.0",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(repoCmd)
	rootCmd.AddCommand(issueCmd)
	rootCmd.AddCommand(prCmd)
	rootCmd.AddCommand(apiCmd)
}
