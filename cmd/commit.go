package cmd

import (
	"github.com/briandowns/spinner"
	"github.com/weibaohui/atomgit-cli/internal/httpclient"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Work with commits",
	Long:  `List and view commits.`,
}

var commitListCmd = &cobra.Command{
	Use:   "list",
	Short: "List commits",
	Long:  `List commits in a repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		sha, _ := cmd.Flags().GetString("sha")
		path, _ := cmd.Flags().GetString("path")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		apiPath := fmt.Sprintf("/repos/%s/commits", repo)
		if sha != "" {
			apiPath = fmt.Sprintf("/repos/%s/commits/%s", repo, sha)
		} else if path != "" {
			apiPath = fmt.Sprintf("/repos/%s/commits?sha=%s&path=%s", repo, sha, path)
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading commits..."
		s.Start()

		data, err := httpclient.Get(apiPath)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list commits: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var commitViewCmd = &cobra.Command{
	Use:   "view <sha>",
	Short: "View a commit",
	Long:  `View details of a specific commit.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		sha := args[0]

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = fmt.Sprintf(" Loading commit %s...", sha[:7])
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/commits/%s", repo, sha))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to view commit: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	commitListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	commitListCmd.Flags().StringP("sha", "", "", "SHA or branch name")
	commitListCmd.Flags().StringP("path", "", "", "Path to filter commits")

	commitViewCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")

	commitCmd.AddCommand(commitListCmd)
	commitCmd.AddCommand(commitViewCmd)
}
