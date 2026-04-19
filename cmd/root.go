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
  $ atomgit search repos golang
  $ atomgit user info
  $ atomgit branch list -R owner/repo
  $ atomgit commit list -R owner/repo
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
	rootCmd.AddCommand(userCmd)
	rootCmd.AddCommand(branchCmd)
	rootCmd.AddCommand(commitCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(labelCmd)
	rootCmd.AddCommand(eventCmd)
	rootCmd.AddCommand(milestoneCmd)
	rootCmd.AddCommand(orgCmd)
}
