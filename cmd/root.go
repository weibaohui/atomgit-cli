package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "atg",
	Short: "AtomGit CLI (gh-like)",
	Long: `AtomGit CLI for interacting with the AtomGit platform.

EXAMPLES
  $ atg repo create my-project --public
  $ atg repo list
  $ atg issue list -R owner/repo
  $ atg pr list -R owner/repo
  $ atg search repos golang
  $ atg user info
  $ atg branch list -R owner/repo
  $ atg commit list -R owner/repo
  $ atg release list -R owner/repo
  $ atg api /api/v1/user

LEARN MORE
  Use 'atg <command> --help' for more information about a command.`,
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
