package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/weibaohui/atomgit-cli/internal/httpclient"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Work with AtomGit repositories",
	Long:  `Manage repositories on AtomGit. Create, delete, list and view repos.`,
}

var repoCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new repository",
	Long:  `Create a new repository on AtomGit.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		description, _ := cmd.Flags().GetString("description")
		isPublic, _ := cmd.Flags().GetBool("public")
		isPrivate, _ := cmd.Flags().GetBool("private")

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Creating repository..."
		s.Start()

		body := map[string]interface{}{
			"name":        name,
			"description": description,
			"private":     isPrivate || !isPublic,
			"auto_init":   true,
		}

		result, err := httpclient.Post("/user/repos", body)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create repository: %v\n", err)
			os.Exit(1)
			return nil
		}

		resultMap := result.(map[string]interface{})
		fullName := getStringField(resultMap, "full_name")
		htmlURL := getStringField(resultMap, "html_url")

		fmt.Printf("Created repository %s\n\n  %s\n", fullName, htmlURL)
		return nil
	},
}

var repoDeleteCmd = &cobra.Command{
	Use:   "delete <repository>",
	Short: "Delete a repository",
	Long:  `Delete a repository from AtomGit.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := args[0]
		yes, _ := cmd.Flags().GetBool("yes")

		if !yes {
			fmt.Printf("Are you sure you want to delete %s? This cannot be undone. [y/N] ", repo)
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Aborted.")
				return nil
			}
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = fmt.Sprintf(" Deleting %s...", repo)
		s.Start()

		err := httpclient.Delete(fmt.Sprintf("/repos/%s", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to delete repository: %v\n", err)
			os.Exit(1)
			return nil
		}

		fmt.Printf("Deleted repository %s\n", repo)
		return nil
	},
}

var repoListCmd = &cobra.Command{
	Use:   "list",
	Short: "List your repositories",
	Long:  `List your repositories on AtomGit.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading repositories..."
		s.Start()

		data, err := httpclient.Get("/user/repos")
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list repos: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var repoViewCmd = &cobra.Command{
	Use:   "view [owner/repo]",
	Short: "View repository details",
	Long:  `View details of a repository.`,
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "Repository path required (e.g., owner/repo)")
			os.Exit(1)
			return nil
		}

		repoPath := args[0]
		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading repository..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s", repoPath))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to view repo: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func printJSON(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to format JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}

func getStringField(m map[string]interface{}, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}

func init() {
	repoCreateCmd.Flags().StringP("description", "d", "", "Description of the repository")
	repoCreateCmd.Flags().Bool("public", false, "Make the new repository public")
	repoCreateCmd.Flags().Bool("private", false, "Make the new repository private")
	repoDeleteCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt")

	repoCmd.AddCommand(repoCreateCmd)
	repoCmd.AddCommand(repoDeleteCmd)
	repoCmd.AddCommand(repoListCmd)
	repoCmd.AddCommand(repoViewCmd)
}
