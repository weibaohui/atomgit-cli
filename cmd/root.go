package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "amc",
	Short: "AtomGit CLI (gh-like)",
	Long: `AtomGit CLI for interacting with the AtomGit platform.

EXAMPLES
  $ amc repo create my-project --public
  $ amc repo list
  $ amc issue list -R owner/repo
  $ amc pr list -R owner/repo
  $ amc search repos golang
  $ amc user info
  $ amc branch list -R owner/repo
  $ amc commit list -R owner/repo
  $ amc release list -R owner/repo
  $ amc api /api/v1/user

LEARN MORE
  Use 'amc <command> --help' for more information about a command.`,
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
	rootCmd.AddCommand(releaseCmd)
	rootCmd.AddCommand(tagCmd)
	rootCmd.AddCommand(forkCmd)
	rootCmd.AddCommand(starCmd)
	rootCmd.AddCommand(subscriberCmd)
	rootCmd.AddCommand(languageCmd)
	rootCmd.AddCommand(contributorCmd)
	rootCmd.AddCommand(hookCmd)
}
