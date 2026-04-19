package cmd

import (
	"fmt"
	"os"

	"github.com/briandowns/spinner"
	"github.com/weibaohui/atomgit-cli/internal/httpclient"
	"github.com/spf13/cobra"
)

var contributorCmd = &cobra.Command{
	Use:   "contributor",
	Short: "Work with contributors",
	Long:  `List contributors to a repository.`,
}

var contributorListCmd = &cobra.Command{
	Use:   "list",
	Short: "List contributors",
	Long:  `List contributors to a repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading contributors..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/contributors", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list contributors: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var contributorStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Get contributor statistics",
	Long:  `Get contributor statistics for a repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading contributor stats..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/contributors-statistic", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get contributor stats: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	contributorListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	contributorStatsCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")

	contributorCmd.AddCommand(contributorListCmd)
	contributorCmd.AddCommand(contributorStatsCmd)
}
