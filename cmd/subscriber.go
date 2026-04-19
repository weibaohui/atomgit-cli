package cmd

import (
	"fmt"
	"os"

	"github.com/briandowns/spinner"
	"github.com/weibaohui/atomgit-cli/internal/httpclient"
	"github.com/spf13/cobra"
)

var subscriberCmd = &cobra.Command{
	Use:   "subscriber",
	Short: "Work with subscribers",
	Long:  `List subscribers (watchers) of a repository.`,
}

var subscriberListCmd = &cobra.Command{
	Use:   "list",
	Short: "List subscribers",
	Long:  `List users watching a repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading subscribers..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/subscribers", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list subscribers: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	subscriberListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	subscriberCmd.AddCommand(subscriberListCmd)
}
