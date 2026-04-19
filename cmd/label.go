package cmd

import (
	"github.com/briandowns/spinner"
	"github.com/weibaohui/atomgit-cli/internal/httpclient"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var labelCmd = &cobra.Command{
	Use:   "label",
	Short: "Work with labels",
	Long:  `List labels in a repository.`,
}

var labelListCmd = &cobra.Command{
	Use:   "list",
	Short: "List labels",
	Long:  `List all labels in a repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading labels..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/labels", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list labels: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	labelListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	labelCmd.AddCommand(labelListCmd)
}
