package cmd

import (
	"fmt"
	"os"

	"github.com/briandowns/spinner"
	"github.com/weibaohui/atomgit-cli/internal/httpclient"
	"github.com/spf13/cobra"
)

var languageCmd = &cobra.Command{
	Use:   "language",
	Short: "Get repository languages",
	Long:  `Get programming languages used in a repository.`,
}

var languageListCmd = &cobra.Command{
	Use:   "list",
	Short: "List languages",
	Long:  `List programming languages used in a repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading languages..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/languages", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list languages: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	languageListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	languageCmd.AddCommand(languageListCmd)
}
