package cmd

import (
	"fmt"
	"os"

	"github.com/briandowns/spinner"
	"github.com/weibaohui/atomgit-cli/internal/httpclient"
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Work with releases",
	Long:  `List and view releases.`,
}

var releaseListCmd = &cobra.Command{
	Use:   "list",
	Short: "List releases",
	Long:  `List all releases in a repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading releases..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/releases", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list releases: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var releaseViewCmd = &cobra.Command{
	Use:   "view <tag>",
	Short: "View a release",
	Long:  `View details of a specific release by tag name.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		tag := args[0]

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = fmt.Sprintf(" Loading release %s...", tag)
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/releases/%s", repo, tag))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to view release: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	releaseListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	releaseViewCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")

	releaseCmd.AddCommand(releaseListCmd)
	releaseCmd.AddCommand(releaseViewCmd)
}
