package cmd

import (
	"fmt"
	"os"

	"github.com/briandowns/spinner"
	"github.com/weibaohui/atomgit-cli/internal/httpclient"
	"github.com/spf13/cobra"
)

var starCmd = &cobra.Command{
	Use:   "star",
	Short: "Star a repository",
	Long:  `List stargazers of a repository.`,
}

var starListCmd = &cobra.Command{
	Use:   "list",
	Short: "List stargazers",
	Long:  `List users who starred a repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading stargazers..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/stargazers", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list stargazers: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	starListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	starCmd.AddCommand(starListCmd)
}
