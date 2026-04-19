package cmd

import (
	"fmt"
	"os"

	"github.com/weibaohui/atomgit-cli/internal/httpclient"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Work with AtomGit issues",
	Long: `Work with AtomGit issues.

ARGUMENTS
  An issue can be supplied as argument in any of the following formats:
  - by number, e.g. "123"
  - by URL, e.g. "https://atomgit.com/OWNER/REPO/issues/123"

EXAMPLES
  $ atomgit issue list -R owner/repo
  $ atomgit issue view 123 -R owner/repo

LEARN MORE
  Use 'atomgit issue <subcommand> --help' for more information about a command.`,
}

var issueListCmd = &cobra.Command{
	Use:   "list",
	Short: "List issues in a repository",
	Long:  `List issues in a repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		state, _ := cmd.Flags().GetString("state")
		limit, _ := cmd.Flags().GetString("limit")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading issues..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/issues?state=%s&limit=%s", repo, state, limit))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list issues: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var issueViewCmd = &cobra.Command{
	Use:   "view <number>",
	Short: "View issue details",
	Long:  `View details of an issue.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		number := args[0]

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading issue..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/issues/%s", repo, number))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to view issue: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	issueListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	issueListCmd.Flags().StringP("state", "s", "open", "Filter by state: open, closed, all")
	issueListCmd.Flags().StringP("limit", "L", "30", "Maximum number of issues to list")

	issueViewCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")

	issueCmd.AddCommand(issueListCmd)
	issueCmd.AddCommand(issueViewCmd)
}
