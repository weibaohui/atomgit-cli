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

var prMergeCmd = &cobra.Command{
	Use:   "merge <number>",
	Short: "Merge a pull request",
	Long:  `Merge a pull request.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		number := args[0]
		mergeMethod, _ := cmd.Flags().GetString("method")
		commitTitle, _ := cmd.Flags().GetString("title")
		commitMessage, _ := cmd.Flags().GetString("message")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Merging pull request..."
		s.Start()

		reqBody := map[string]string{}
		if mergeMethod != "" {
			reqBody["merge_method"] = mergeMethod
		}
		if commitTitle != "" {
			reqBody["commit_title"] = commitTitle
		}
		if commitMessage != "" {
			reqBody["commit_message"] = commitMessage
		}

		data, err := httpclient.Put(fmt.Sprintf("/repos/%s/pulls/%s/merge", repo, number), reqBody)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to merge pull request: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prCommentsCmd = &cobra.Command{
	Use:   "comments <number>",
	Short: "List PR review comments",
	Long:  `List review comments on a pull request.`,
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
		s.Suffix = " Loading comments..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/pulls/%s/comments", repo, number))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list comments: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prLabelsCmd = &cobra.Command{
	Use:   "labels <number>",
	Short: "List PR labels",
	Long:  `List labels on a pull request.`,
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
		s.Suffix = " Loading labels..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/pulls/%s/labels", repo, number))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list labels: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prAddLabelsCmd = &cobra.Command{
	Use:   "add-labels <number>",
	Short: "Add PR labels",
	Long:  `Add labels to a pull request.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		number := args[0]
		labels, _ := cmd.Flags().GetStringArray("label")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Adding labels..."
		s.Start()

		reqBody := map[string][]string{"labels": labels}
		data, err := httpclient.Post(fmt.Sprintf("/repos/%s/pulls/%s/labels", repo, number), reqBody)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to add labels: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prRemoveLabelsCmd = &cobra.Command{
	Use:   "remove-labels <number>",
	Short: "Remove PR labels",
	Long:  `Remove labels from a pull request.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		number := args[0]
		label, _ := cmd.Flags().GetString("label")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		if label == "" {
			fmt.Fprintln(os.Stderr, "Label name required. Use --label")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Removing label..."
		s.Start()

		err := httpclient.Delete(fmt.Sprintf("/repos/%s/pulls/%s/labels/%s", repo, number, label))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to remove label: %v\n", err)
			os.Exit(1)
			return nil
		}

		fmt.Println("Label removed successfully")
		return nil
	},
}

var prAssigneesCmd = &cobra.Command{
	Use:   "assignees <number>",
	Short: "List PR assignees",
	Long:  `List assignees on a pull request.`,
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
		s.Suffix = " Loading assignees..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/pulls/%s/assignees", repo, number))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list assignees: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prAddAssigneesCmd = &cobra.Command{
	Use:   "add-assignees <number>",
	Short: "Add PR assignees",
	Long:  `Add assignees to a pull request.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		number := args[0]
		assignees, _ := cmd.Flags().GetStringArray("assignee")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Adding assignees..."
		s.Start()

		reqBody := map[string][]string{"assignees": assignees}
		data, err := httpclient.Post(fmt.Sprintf("/repos/%s/pulls/%s/assignees", repo, number), reqBody)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to add assignees: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prRemoveAssigneesCmd = &cobra.Command{
	Use:   "remove-assignees <number>",
	Short: "Remove PR assignees",
	Long:  `Remove assignees from a pull request.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		number := args[0]
		assignees, _ := cmd.Flags().GetStringArray("assignee")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Removing assignees..."
		s.Start()

		reqBody := map[string][]string{"assignees": assignees}
		data, err := httpclient.Patch(fmt.Sprintf("/repos/%s/pulls/%s/assignees", repo, number), reqBody)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to remove assignees: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prCommitsCmd = &cobra.Command{
	Use:   "commits <number>",
	Short: "List PR commits",
	Long:  `List commits in a pull request.`,
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
		s.Suffix = " Loading commits..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/pulls/%s/commits", repo, number))
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

var prFilesCmd = &cobra.Command{
	Use:   "files <number>",
	Short: "List changed files",
	Long:  `List changed files in a pull request.`,
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
		s.Suffix = " Loading files..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/pulls/%s/files", repo, number))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list files: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prMergeStatusCmd = &cobra.Command{
	Use:   "merge-status <number>",
	Short: "Check PR merge status",
	Long:  `Check if a pull request is merged.`,
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
		s.Suffix = " Checking merge status..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/pulls/%s/merge", repo, number))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to check merge status: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prUpdateCmd = &cobra.Command{
	Use:   "update <number>",
	Short: "Update a pull request",
	Long:  `Update an existing pull request.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		number := args[0]
		title, _ := cmd.Flags().GetString("title")
		body, _ := cmd.Flags().GetString("body")
		state, _ := cmd.Flags().GetString("state")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Updating pull request..."
		s.Start()

		reqBody := map[string]string{}
		if title != "" {
			reqBody["title"] = title
		}
		if body != "" {
			reqBody["body"] = body
		}
		if state != "" {
			reqBody["state"] = state
		}

		data, err := httpclient.Patch(fmt.Sprintf("/repos/%s/pulls/%s", repo, number), reqBody)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to update pull request: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a pull request",
	Long:  `Create a new pull request.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		title, _ := cmd.Flags().GetString("title")
		body, _ := cmd.Flags().GetString("body")
		head, _ := cmd.Flags().GetString("head")
		base, _ := cmd.Flags().GetString("base")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		if title == "" {
			fmt.Fprintln(os.Stderr, "Title required. Use --title")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Creating pull request..."
		s.Start()

		reqBody := map[string]string{
			"title": title,
		}
		if body != "" {
			reqBody["body"] = body
		}
		if head != "" {
			reqBody["head"] = head
		}
		if base != "" {
			reqBody["base"] = base
		}

		data, err := httpclient.Post(fmt.Sprintf("/repos/%s/pulls", repo), reqBody)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create pull request: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	prCommentsCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")

	prLabelsCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prAddLabelsCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prAddLabelsCmd.Flags().StringArrayP("label", "l", []string{}, "Label name (can specify multiple)")
	prRemoveLabelsCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prRemoveLabelsCmd.Flags().StringP("label", "l", "", "Label name to remove")

	prAssigneesCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prAddAssigneesCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prAddAssigneesCmd.Flags().StringArrayP("assignee", "a", []string{}, "Assignee username (can specify multiple)")
	prRemoveAssigneesCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prRemoveAssigneesCmd.Flags().StringArrayP("assignee", "a", []string{}, "Assignee username (can specify multiple)")

	prCommitsCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prFilesCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")

	prMergeCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prMergeCmd.Flags().StringP("method", "m", "", "Merge method (merge, squash, rebase)")
	prMergeCmd.Flags().StringP("title", "t", "", "Commit title")
	prMergeCmd.Flags().StringP("message", "", "", "Commit message")

	prMergeStatusCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")

	prCreateCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prCreateCmd.Flags().StringP("title", "t", "", "Pull request title (required)")
	prCreateCmd.Flags().StringP("body", "m", "", "Pull request body/description")
	prCreateCmd.Flags().StringP("head", "", "", "Source branch")
	prCreateCmd.Flags().StringP("base", "b", "main", "Target branch")

	prUpdateCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prUpdateCmd.Flags().StringP("title", "t", "", "Pull request title")
	prUpdateCmd.Flags().StringP("body", "m", "", "Pull request body/description")
	prUpdateCmd.Flags().StringP("state", "s", "", "Pull request state (open, closed)")

	prCmd.AddCommand(prListCmd)
	prCmd.AddCommand(prViewCmd)
	prCmd.AddCommand(prCommentsCmd)
	prCmd.AddCommand(prLabelsCmd)
	prCmd.AddCommand(prAddLabelsCmd)
	prCmd.AddCommand(prRemoveLabelsCmd)
	prCmd.AddCommand(prAssigneesCmd)
	prCmd.AddCommand(prAddAssigneesCmd)
	prCmd.AddCommand(prRemoveAssigneesCmd)
	prCmd.AddCommand(prCommitsCmd)
	prCmd.AddCommand(prFilesCmd)
	prCmd.AddCommand(prMergeCmd)
	prCmd.AddCommand(prMergeStatusCmd)
	prCmd.AddCommand(prCreateCmd)
	prCmd.AddCommand(prUpdateCmd)
}
