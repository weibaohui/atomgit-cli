package cmd

import (
	"github.com/briandowns/spinner"
	"github.com/weibaohui/atomgit-cli/internal/httpclient"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "Work with events",
	Long:  `List events in a repository.`,
}

var eventListCmd = &cobra.Command{
	Use:   "list",
	Short: "List events",
	Long:  `List all events in a repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading events..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/events", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list events: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	eventListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	eventCmd.AddCommand(eventListCmd)
}
