package cmd

import (
	"fmt"
	"os"

	"github.com/weibaohui/atomgit-cli/internal/httpclient"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Work with AtomGit pull requests",
	Long: `Work with AtomGit pull requests.

ARGUMENTS
  A pull request can be supplied as argument in any of the following formats:
  - by number, e.g. "123"
  - by URL, e.g. "https://atomgit.com/OWNER/REPO/pulls/123"
  - by branch name, e.g. "feature-branch"

EXAMPLES
  $ atomgit pr list -R owner/repo
  $ atomgit pr view 123 -R owner/repo

LEARN MORE
  Use 'atomgit pr <subcommand> --help' for more information about a command.`,
}

var prListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pull requests in a repository",
	Long:  `List pull requests in a repository.`,
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
		s.Suffix = " Loading pull requests..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/pulls?state=%s&limit=%s", repo, state, limit))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list pull requests: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prViewCmd = &cobra.Command{
	Use:   "view <number>",
	Short: "View pull request details",
	Long:  `View details of a pull request.`,
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
		s.Suffix = " Loading pull request..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/pulls/%s", repo, number))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to view pull request: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	prListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prListCmd.Flags().StringP("state", "s", "open", "Filter by state: open, closed, all")
	prListCmd.Flags().StringP("limit", "L", "30", "Maximum number of PRs to list")

	prViewCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")

	prCmd.AddCommand(prListCmd)
	prCmd.AddCommand(prViewCmd)
}
