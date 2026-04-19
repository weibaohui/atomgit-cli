package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/weibaohui/atomgit-cli/internal/httpclient"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Work with AtomGit repositories",
	Long:  `Manage repositories on AtomGit. Create, delete, list, clone, fork and manage repos.`,
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
		gitignore, _ := cmd.Flags().GetString("gitignore")
		license, _ := cmd.Flags().GetString("license")
		includeAllBranches, _ := cmd.Flags().GetBool("include-all-branches")
		push, _ := cmd.Flags().GetBool("push")
		addReadme, _ := cmd.Flags().GetBool("add-readme")
		clone, _ := cmd.Flags().GetBool("clone")
		org, _ := cmd.Flags().GetString("org")

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Creating repository..."
		s.Start()

		body := map[string]interface{}{
			"name":      name,
			"private":   isPrivate || !isPublic,
			"auto_init": true,
		}
		if description != "" {
			body["description"] = description
		}
		if gitignore != "" {
			body["gitignore_template"] = gitignore
		}
		if license != "" {
			body["license_template"] = license
		}
		if includeAllBranches {
			body["include_all_branches"] = true
		}
		if addReadme {
			body["auto_init"] = false // disable auto_init if adding README manually
		}

		var result interface{}
		var err error

		if org != "" {
			// Create under organization
			result, err = httpclient.Post(fmt.Sprintf("/orgs/%s/repos", org), body)
		} else {
			result, err = httpclient.Post("/user/repos", body)
		}

		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create repository: %v\n", err)
			os.Exit(1)
			return nil
		}

		resultMap := result.(map[string]interface{})
		fullName := getStringField(resultMap, "full_name")
		htmlURL := getStringField(resultMap, "html_url")
		sshURL := getStringField(resultMap, "ssh_url")
		cloneURL := getStringField(resultMap, "clone_url")

		fmt.Printf("Created repository %s\n\n  %s\n", fullName, htmlURL)

		// Clone if requested
		if clone {
			s2 := spinner.New(spinner.CharSets[14], 100*1000*1000)
			s2.Suffix = " Cloning repository..."
			s2.Start()

			gitURL := sshURL
			if gitURL == "" {
				gitURL = cloneURL
			}

			gitArgs := []string{"clone"}
			if push {
				gitArgs = append(gitArgs, "-b", "main")
			}
			gitArgs = append(gitArgs, gitURL, name)

			gitCmd := exec.Command("git", gitArgs...)
			gitCmd.Stdout = os.Stdout
			gitCmd.Stderr = os.Stderr
			err := gitCmd.Run()
			s2.Stop()

			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to clone repository: %v\n", err)
				// Don't exit with error, repo was created successfully
			} else {
				fmt.Printf("Cloned to %s/\n", name)

				// Add remote and push if requested
				if push {
					fmt.Println("Pushing to remote...")
					gitPush := exec.Command("git", "-C", name, "push", "-u", "origin", "main")
					gitPush.Stdout = os.Stdout
					gitPush.Stderr = os.Stderr
					if err := gitPush.Run(); err != nil {
						fmt.Fprintf(os.Stderr, "Failed to push: %v\n", err)
					}
				}
			}
		}

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

var repoCloneCmd = &cobra.Command{
	Use:   "clone <repository>",
	Short: "Clone a repository",
	Long:  `Clone a repository from AtomGit.`,
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var repo string
		if len(args) > 0 {
			repo = args[0]
		}

		if repo == "" {
			// Try to determine repo from git remote
			fmt.Println("No repository specified, trying to clone from current directory...")
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Getting repository info..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get repository: %v\n", err)
			os.Exit(1)
			return nil
		}

		resultMap := data.(map[string]interface{})
		fullName := getStringField(resultMap, "full_name")
		sshURL := getStringField(resultMap, "ssh_url")
		cloneURL := getStringField(resultMap, "clone_url")
		defaultBranch := getStringField(resultMap, "default_branch")
		if defaultBranch == "" {
			defaultBranch = "main"
		}

		gitURL := sshURL
		if gitURL == "" {
			gitURL = cloneURL
		}

		s2 := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s2.Suffix = fmt.Sprintf(" Cloning %s...", fullName)
		s2.Start()

		// Use the repository name as the clone directory
		parts := strings.Split(fullName, "/")
		repoName := parts[len(parts)-1]

		gitArgs := []string{"clone"}
		if defaultBranch != "main" {
			gitArgs = append(gitArgs, "-b", defaultBranch)
		}
		gitArgs = append(gitArgs, gitURL, repoName)

		gitCmd := exec.Command("git", gitArgs...)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		err = gitCmd.Run()
		s2.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to clone repository: %v\n", err)
			os.Exit(1)
			return nil
		}

		fmt.Printf("Cloned %s to %s/\n", fullName, repoName)
		return nil
	},
}

var repoForkCmd = &cobra.Command{
	Use:   "fork <repository>",
	Short: "Fork a repository",
	Long:  `Fork a repository on AtomGit.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := args[0]
		org, _ := cmd.Flags().GetString("org")
		clone, _ := cmd.Flags().GetBool("clone")

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = fmt.Sprintf(" Forking %s...", repo)
		s.Start()

		body := map[string]interface{}{}
		if org != "" {
			body["organization"] = org
		}

		result, err := httpclient.Post(fmt.Sprintf("/repos/%s/forks", repo), body)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to fork repository: %v\n", err)
			os.Exit(1)
			return nil
		}

		resultMap := result.(map[string]interface{})
		fullName := getStringField(resultMap, "full_name")
		htmlURL := getStringField(resultMap, "html_url")
		sshURL := getStringField(resultMap, "ssh_url")
		cloneURL := getStringField(resultMap, "clone_url")

		fmt.Printf("Forked repository %s\n\n  %s\n", fullName, htmlURL)

		// Clone if requested
		if clone {
			s2 := spinner.New(spinner.CharSets[14], 100*1000*1000)
			s2.Suffix = " Cloning forked repository..."
			s2.Start()

			gitURL := sshURL
			if gitURL == "" {
				gitURL = cloneURL
			}

			parts := strings.Split(fullName, "/")
			repoName := parts[len(parts)-1]

			gitArgs := []string{"clone", gitURL, repoName}
			gitCmd := exec.Command("git", gitArgs...)
			gitCmd.Stdout = os.Stdout
			gitCmd.Stderr = os.Stderr
			err := gitCmd.Run()
			s2.Stop()

			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to clone forked repository: %v\n", err)
			} else {
				fmt.Printf("Cloned to %s/\n", repoName)
			}
		}

		return nil
	},
}

var repoEditCmd = &cobra.Command{
	Use:   "edit <repository>",
	Short: "Edit repository settings",
	Long:  `Edit repository settings on AtomGit.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := args[0]
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		homepage, _ := cmd.Flags().GetString("homepage")
		isPrivate, _ := cmd.Flags().GetBool("private")
		hasIssues, _ := cmd.Flags().GetBool("enable-issues")
		hasWiki, _ := cmd.Flags().GetBool("enable-wiki")
		hasProjects, _ := cmd.Flags().GetBool("enable-projects")

		body := map[string]interface{}{}
		if name != "" {
			body["name"] = name
		}
		if cmd.Flags().Changed("description") {
			body["description"] = description
		}
		if cmd.Flags().Changed("homepage") {
			body["homepage"] = homepage
		}
		if cmd.Flags().Changed("private") {
			body["private"] = isPrivate
		}
		if cmd.Flags().Changed("enable-issues") {
			body["has_issues"] = hasIssues
		}
		if cmd.Flags().Changed("enable-wiki") {
			body["has_wiki"] = hasWiki
		}
		if cmd.Flags().Changed("enable-projects") {
			body["has_projects"] = hasProjects
		}

		if len(body) == 0 {
			fmt.Fprintln(os.Stderr, "No changes specified")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = fmt.Sprintf(" Updating %s...", repo)
		s.Start()

		data, err := httpclient.Patch(fmt.Sprintf("/repos/%s", repo), body)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to update repository: %v\n", err)
			os.Exit(1)
			return nil
		}

		fmt.Println("Repository updated:")
		printJSON(data)
		return nil
	},
}

var repoSyncCmd = &cobra.Command{
	Use:   "sync <repository>",
	Short: "Sync a forked repository",
	Long:  `Sync a forked repository with its parent.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := args[0]
		branch, _ := cmd.Flags().GetString("branch")

		if branch == "" {
			branch = "main"
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Getting fork info..."
		s.Start()

		// First get the repo to find the parent
		data, err := httpclient.Get(fmt.Sprintf("/repos/%s", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get repository: %v\n", err)
			os.Exit(1)
			return nil
		}

		resultMap := data.(map[string]interface{})
		parent := resultMap["parent"]
		if parent == nil {
			fmt.Fprintf(os.Stderr, "%s is not a fork\n", repo)
			os.Exit(1)
			return nil
		}

		parentMap := parent.(map[string]interface{})
		parentFullName := getStringField(parentMap, "full_name")
		parentSSHURL := getStringField(parentMap, "ssh_url")
		if parentSSHURL == "" {
			parentSSHURL = getStringField(parentMap, "clone_url")
		}

		fmt.Printf("Syncing %s from %s\n", repo, parentFullName)

		// Check if we have a local clone
		parts := strings.Split(repo, "/")
		repoName := parts[len(parts)-1]

		if _, err := os.Stat(repoName); os.IsNotExist(err) {
			// No local clone, offer to clone
			fmt.Printf("No local clone found. Would you like to clone %s first? [y/N] ", repo)
			var confirm string
			fmt.Scanln(&confirm)
			if confirm == "y" || confirm == "Y" {
				s2 := spinner.New(spinner.CharSets[14], 100*1000*1000)
				s2.Suffix = " Cloning..."
				s2.Start()
				gitCmd := exec.Command("git", "clone", parentSSHURL, repoName)
				gitCmd.Stdout = os.Stdout
				gitCmd.Stderr = os.Stderr
				err := gitCmd.Run()
				s2.Stop()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to clone: %v\n", err)
					os.Exit(1)
					return nil
				}
			} else {
				fmt.Println("Cannot sync without a local clone")
				return nil
			}
		}

		// Add upstream remote if not exists
		gitRemoteCheck := exec.Command("git", "-C", repoName, "remote", "get-url", "upstream")
		_, err = gitRemoteCheck.Output()
		if err != nil {
			gitAddUpstream := exec.Command("git", "-C", repoName, "remote", "add", "upstream", parentSSHURL)
			gitAddUpstream.Stdout = os.Stdout
			gitAddUpstream.Stderr = os.Stderr
			if err := gitAddUpstream.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to add upstream: %v\n", err)
			}
		}

		// Fetch upstream
		s3 := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s3.Suffix = " Fetching upstream..."
		s3.Start()
		gitFetch := exec.Command("git", "-C", repoName, "fetch", "upstream")
		gitFetch.Stdout = os.Stdout
		gitFetch.Stderr = os.Stderr
		err = gitFetch.Run()
		s3.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to fetch upstream: %v\n", err)
			os.Exit(1)
			return nil
		}

		// Merge from upstream/branch
		s4 := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s4.Suffix = fmt.Sprintf(" Merging %s...", branch)
		s4.Start()
		gitMerge := exec.Command("git", "-C", repoName, "merge", fmt.Sprintf("upstream/%s", branch))
		gitMerge.Stdout = os.Stdout
		gitMerge.Stderr = os.Stderr
		err = gitMerge.Run()
		s4.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to merge: %v\n", err)
			fmt.Println("You may need to resolve conflicts manually")
			os.Exit(1)
			return nil
		}

		// Push to origin
		s5 := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s5.Suffix = " Pushing to origin..."
		s5.Start()
		gitPush := exec.Command("git", "-C", repoName, "push", "origin", branch)
		gitPush.Stdout = os.Stdout
		gitPush.Stderr = os.Stderr
		err = gitPush.Run()
		s5.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to push: %v\n", err)
			os.Exit(1)
			return nil
		}

		fmt.Printf("Successfully synced %s\n", repo)
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
	repoCreateCmd.Flags().StringP("gitignore", "g", "", "Specify .gitignore template")
	repoCreateCmd.Flags().StringP("license", "l", "", "Specify license template (e.g., MIT)")
	repoCreateCmd.Flags().Bool("include-all-branches", false, "Include all branches")
	repoCreateCmd.Flags().Bool("push", false, "Push local commits after clone")
	repoCreateCmd.Flags().Bool("add-readme", false, "Add README")
	repoCreateCmd.Flags().Bool("clone", false, "Clone after create")
	repoCreateCmd.Flags().String("org", "", "Create under organization")

	repoDeleteCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt")

	repoForkCmd.Flags().String("org", "", "Fork to organization")
	repoForkCmd.Flags().Bool("clone", false, "Clone after fork")

	repoEditCmd.Flags().StringP("name", "n", "", "Rename repository")
	repoEditCmd.Flags().StringP("description", "d", "", "Description")
	repoEditCmd.Flags().String("homepage", "", "Homepage URL")
	repoEditCmd.Flags().Bool("private", false, "Make private")
	repoEditCmd.Flags().Bool("enable-issues", false, "Enable issues")
	repoEditCmd.Flags().Bool("enable-wiki", false, "Enable wiki")
	repoEditCmd.Flags().Bool("enable-projects", false, "Enable projects")

	repoSyncCmd.Flags().StringP("branch", "b", "main", "Branch to sync")

	repoCmd.AddCommand(repoCreateCmd)
	repoCmd.AddCommand(repoDeleteCmd)
	repoCmd.AddCommand(repoListCmd)
	repoCmd.AddCommand(repoViewCmd)
	repoCmd.AddCommand(repoCloneCmd)
	repoCmd.AddCommand(repoForkCmd)
	repoCmd.AddCommand(repoEditCmd)
	repoCmd.AddCommand(repoSyncCmd)
}