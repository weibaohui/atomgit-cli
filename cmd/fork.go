package cmd

import (
	"fmt"
	"os"

	"github.com/briandowns/spinner"
	"github.com/weibaohui/atomgit-cli/internal/httpclient"
	"github.com/spf13/cobra"
)

var forkCmd = &cobra.Command{
	Use:   "fork",
	Short: "Fork a repository",
	Long:  `Fork a repository to your account.`,
}

var forkCreateCmd = &cobra.Command{
	Use:   "create <repo>",
	Short: "Fork a repository",
	Long:  `Fork a repository to your account.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := args[0]

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = fmt.Sprintf(" Forking %s...", repo)
		s.Start()

		data, err := httpclient.Post(fmt.Sprintf("/repos/%s/forks", repo), nil)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to fork repository: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var forkListCmd = &cobra.Command{
	Use:   "list <repo>",
	Short: "List forks",
	Long:  `List all forks of a repository.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := args[0]

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading forks..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/forks", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list forks: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	forkCmd.AddCommand(forkCreateCmd)
	forkCmd.AddCommand(forkListCmd)
}
