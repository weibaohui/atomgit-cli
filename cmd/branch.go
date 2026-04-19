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
	Long:  `List, create, view, protect and delete branches.`,
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

var branchCreateCmd = &cobra.Command{
	Use:   "create <branch>",
	Short: "Create a branch",
	Long:  `Create a new branch.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		branch := args[0]
		sha, _ := cmd.Flags().GetString("sha")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		body := map[string]interface{}{
			"branch_name": branch,
		}
		if sha != "" {
			body["refs"] = sha
		} else {
			body["refs"] = "main"
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = fmt.Sprintf(" Creating branch %s...", branch)
		s.Start()

		data, err := httpclient.Post(fmt.Sprintf("/repos/%s/branches", repo), body)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create branch: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var branchDeleteCmd = &cobra.Command{
	Use:   "delete <branch>",
	Short: "Delete a branch",
	Long:  `Delete a branch.`,
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
		s.Suffix = fmt.Sprintf(" Deleting branch %s...", branch)
		s.Start()

		err := httpclient.Delete(fmt.Sprintf("/repos/%s/branches/%s", repo, branch))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to delete branch: %v\n", err)
			os.Exit(1)
			return nil
		}

		fmt.Printf("Branch %s deleted successfully.\n", branch)
		return nil
	},
}

var branchProtectCmd = &cobra.Command{
	Use:   "protect <branch>",
	Short: "Protect a branch",
	Long:  `Set branch protection settings.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		branch := args[0]

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		body := map[string]interface{}{
			"wildcard": branch,
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = fmt.Sprintf(" Protecting branch %s...", branch)
		s.Start()

		data, err := httpclient.Put(fmt.Sprintf("/repos/%s/branches/setting/new", repo), body)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to protect branch: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var branchUnprotectCmd = &cobra.Command{
	Use:   "unprotect <branch>",
	Short: "Unprotect a branch",
	Long:  `Remove branch protection settings.`,
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
		s.Suffix = fmt.Sprintf(" Unprotecting branch %s...", branch)
		s.Start()

		err := httpclient.Delete(fmt.Sprintf("/repos/%s/branches/%s/setting", repo, branch))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to unprotect branch: %v\n", err)
			os.Exit(1)
			return nil
		}

		fmt.Printf("Branch %s unprotected successfully.\n", branch)
		return nil
	},
}

var branchProtectedListCmd = &cobra.Command{
	Use:   "protected-list",
	Short: "List protected branches",
	Long:  `List all protected branches in a repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading protected branches..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/protect_branches", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list protected branches: %v\n", err)
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
	branchCreateCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	branchCreateCmd.Flags().StringP("sha", "", "", "SHA to create branch from")
	branchDeleteCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	branchProtectCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	branchUnprotectCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	branchProtectedListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")

	branchCmd.AddCommand(branchListCmd)
	branchCmd.AddCommand(branchViewCmd)
	branchCmd.AddCommand(branchCreateCmd)
	branchCmd.AddCommand(branchDeleteCmd)
	branchCmd.AddCommand(branchProtectCmd)
	branchCmd.AddCommand(branchUnprotectCmd)
	branchCmd.AddCommand(branchProtectedListCmd)
}
