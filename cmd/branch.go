package cmd

import (
	"fmt"
	"os"

	"github.com/briandowns/spinner"
	"github.com/weibaohui/atomgit-cli/internal/httpclient"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Work with branches",
	Long:  `List, create, and delete branches.`,
}

var branchListCmd = &cobra.Command{
	Use:   "list",
	Short: "List branches",
	Long:  `List all branches in a repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading branches..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/branches", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list branches: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var branchViewCmd = &cobra.Command{
	Use:   "view <branch>",
	Short: "View branch details",
	Long:  `View details of a specific branch.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		branch := args[0]

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = fmt.Sprintf(" Loading branch %s...", branch)
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/branches/%s", repo, branch))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to view branch: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	branchListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	branchViewCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")

	branchCmd.AddCommand(branchListCmd)
	branchCmd.AddCommand(branchViewCmd)
}
