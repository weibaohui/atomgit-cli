package cmd

import (
	"fmt"
	"os"

	"github.com/briandowns/spinner"
	"github.com/weibaohui/atomgit-cli/internal/httpclient"
	"github.com/spf13/cobra"
)

var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Work with webhooks",
	Long:  `List webhooks in a repository.`,
}

var hookListCmd = &cobra.Command{
	Use:   "list",
	Short: "List webhooks",
	Long:  `List all webhooks in a repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading webhooks..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/hooks", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list webhooks: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	hookListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	hookCmd.AddCommand(hookListCmd)
}
