package cmd

import (
	"github.com/briandowns/spinner"
	"github.com/weibaohui/atomgit-cli/internal/httpclient"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var milestoneCmd = &cobra.Command{
	Use:   "milestone",
	Short: "Work with milestones",
	Long:  `List milestones in a repository.`,
}

var milestoneListCmd = &cobra.Command{
	Use:   "list",
	Short: "List milestones",
	Long:  `List all milestones in a repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading milestones..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/milestones", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list milestones: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	milestoneListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	milestoneCmd.AddCommand(milestoneListCmd)
}
