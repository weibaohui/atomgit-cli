package cmd

import (
	"fmt"
	"os"

	"github.com/briandowns/spinner"
	"github.com/weibaohui/atomgit-cli/internal/httpclient"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Work with tags",
	Long:  `List tags in a repository.`,
}

var tagListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tags",
	Long:  `List all tags in a repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading tags..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/tags", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list tags: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	tagListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	tagCmd.AddCommand(tagListCmd)
}
