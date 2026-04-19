package cmd

import (
	"fmt"
	"os"
	"strings"

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
		labels, _ := cmd.Flags().GetString("label")
		assignee, _ := cmd.Flags().GetString("assignee")
		author, _ := cmd.Flags().GetString("author")
		draft, _ := cmd.Flags().GetBool("draft")
		search, _ := cmd.Flags().GetString("search")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading pull requests..."
		s.Start()

		url := fmt.Sprintf("/repos/%s/pulls?state=%s&limit=%s", repo, state, limit)
		if labels != "" {
			url += "&labels=" + labels
		}
		if assignee != "" {
			url += "&assignee=" + assignee
		}
		if author != "" {
			url += "&author=" + author
		}
		if draft {
			url += "&is_draft=true"
		}
		if search != "" {
			url += "&search=" + search
		}

		data, err := httpclient.Get(url)
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
		comments, _ := cmd.Flags().GetBool("comments")
		web, _ := cmd.Flags().GetBool("web")

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

		if web {
			if m, ok := data.(map[string]interface{}); ok {
				if htmlURL, ok := m["html_url"].(string); ok {
					fmt.Printf("Opening in browser: %s\n", htmlURL)
				}
			}
		}

		if comments {
			s2 := spinner.New(spinner.CharSets[14], 100*1000*1000)
			s2.Suffix = " Loading comments..."
			s2.Start()
			commentsData, err := httpclient.Get(fmt.Sprintf("/repos/%s/pulls/%s/comments", repo, number))
			s2.Stop()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to load comments: %v\n", err)
			} else {
				if m, ok := data.(map[string]interface{}); ok {
					m["comments"] = commentsData
				}
			}
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
	Long:  `List assignees on a pull request (uses requested_reviewers API).`,
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

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/pulls/%s/requested_reviewers", repo, number))
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
		assigneesStr, _ := cmd.Flags().GetString("assignee")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		if assigneesStr == "" {
			fmt.Fprintln(os.Stderr, "Assignee required. Use --assignee")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Adding assignees..."
		s.Start()

		reqBody := map[string]string{"assignees": assigneesStr}
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
		assigneesStr, _ := cmd.Flags().GetString("assignee")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		if assigneesStr == "" {
			fmt.Fprintln(os.Stderr, "Assignee required. Use --assignee")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Removing assignees..."
		s.Start()

		reqBody := map[string]string{"assignees": assigneesStr}
		err := httpclient.DeleteWithBody(fmt.Sprintf("/repos/%s/pulls/%s/assignees", repo, number), reqBody)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to remove assignees: %v\n", err)
			os.Exit(1)
			return nil
		}

		fmt.Println("Assignees removed successfully")
		return nil
	},
}

var prReviewersCmd = &cobra.Command{
	Use:   "reviewers <number>",
	Short: "List available reviewers",
	Long:  `List available reviewers for a pull request.`,
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
		s.Suffix = " Loading reviewers..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/pulls/%s/option_reviewers", repo, number))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list reviewers: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prAddReviewersCmd = &cobra.Command{
	Use:   "add-reviewers <number>",
	Short: "Add approval reviewers",
	Long:  `Add approval reviewers to a pull request.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		number := args[0]
		reviewers, _ := cmd.Flags().GetStringArray("reviewer")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Adding reviewers..."
		s.Start()

		reqBody := map[string][]string{"reviewers": reviewers}
		data, err := httpclient.Post(fmt.Sprintf("/repos/%s/pulls/%s/approval-reviewers", repo, number), reqBody)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to add reviewers: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prRemoveReviewersCmd = &cobra.Command{
	Use:   "remove-reviewers <number>",
	Short: "Remove approval reviewers",
	Long:  `Remove approval reviewers from a pull request.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		number := args[0]
		reviewers, _ := cmd.Flags().GetStringArray("reviewer")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Removing reviewers..."
		s.Start()

		reqBody := map[string][]string{"reviewers": reviewers}
		err := httpclient.DeleteWithBody(fmt.Sprintf("/repos/%s/pulls/%s/approval-reviewers", repo, number), reqBody)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to remove reviewers: %v\n", err)
			os.Exit(1)
			return nil
		}

		fmt.Println("Reviewers removed successfully")
		return nil
	},
}

var prOperateLogsCmd = &cobra.Command{
	Use:   "operate-logs <number>",
	Short: "Get PR operation logs",
	Long:  `Get operation logs for a pull request.`,
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
		s.Suffix = " Loading operate logs..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/pulls/%s/operate-logs", repo, number))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get operate logs: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prTestersCmd = &cobra.Command{
	Use:   "testers <number>",
	Short: "List PR testers",
	Long:  `List testers on a pull request.`,
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
		s.Suffix = " Loading testers..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/pulls/%s/option_testers", repo, number))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list testers: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prAddTestersCmd = &cobra.Command{
	Use:   "add-testers <number>",
	Short: "Add PR testers",
	Long:  `Add testers to a pull request.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		number := args[0]
		testers, _ := cmd.Flags().GetStringArray("tester")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Adding testers..."
		s.Start()

		reqBody := map[string][]string{"testers": testers}
		data, err := httpclient.Post(fmt.Sprintf("/repos/%s/pulls/%s/testers", repo, number), reqBody)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to add testers: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prRemoveTestersCmd = &cobra.Command{
	Use:   "remove-testers <number>",
	Short: "Remove PR testers",
	Long:  `Remove testers from a pull request.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		number := args[0]
		testers, _ := cmd.Flags().GetStringArray("tester")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Removing testers..."
		s.Start()

		reqBody := map[string][]string{"testers": testers}
		err := httpclient.DeleteWithBody(fmt.Sprintf("/repos/%s/pulls/%s/testers", repo, number), reqBody)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to remove testers: %v\n", err)
			os.Exit(1)
			return nil
		}

		fmt.Println("Testers removed successfully")
		return nil
	},
}

var prLinkedIssuesCmd = &cobra.Command{
	Use:   "linked-issues <number>",
	Short: "List linked issues",
	Long:  `List issues linked to a pull request.`,
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
		s.Suffix = " Loading linked issues..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/pulls/%s/issues", repo, number))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list linked issues: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prLinkIssueCmd = &cobra.Command{
	Use:   "link-issue <number>",
	Short: "Link issue to PR",
	Long:  `Link an issue to a pull request.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		number := args[0]
		issueNumber, _ := cmd.Flags().GetString("issue")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		if issueNumber == "" {
			fmt.Fprintln(os.Stderr, "Issue number required. Use --issue")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Linking issue..."
		s.Start()

		reqBody := map[string]string{"issue_number": issueNumber}
		data, err := httpclient.Post(fmt.Sprintf("/repos/%s/pulls/%s/linked-issues", repo, number), reqBody)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to link issue: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prUnlinkIssueCmd = &cobra.Command{
	Use:   "unlink-issue <number>",
	Short: "Unlink issue from PR",
	Long:  `Unlink an issue from a pull request.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		number := args[0]
		issueNumber, _ := cmd.Flags().GetString("issue")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		if issueNumber == "" {
			fmt.Fprintln(os.Stderr, "Issue number required. Use --issue")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Unlinking issue..."
		s.Start()

		err := httpclient.Delete(fmt.Sprintf("/repos/%s/pulls/%s/issues/%s", repo, number, issueNumber))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to unlink issue: %v\n", err)
			os.Exit(1)
			return nil
		}

		fmt.Println("Issue unlinked successfully")
		return nil
	},
}

var prCloseCmd = &cobra.Command{
	Use:   "close <number>",
	Short: "Close a pull request",
	Long:  `Close a pull request.`,
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
		s.Suffix = " Closing pull request..."
		s.Start()

		reqBody := map[string]string{"state": "closed"}
		data, err := httpclient.Patch(fmt.Sprintf("/repos/%s/pulls/%s", repo, number), reqBody)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to close pull request: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var prReopenCmd = &cobra.Command{
	Use:   "reopen <number>",
	Short: "Reopen a pull request",
	Long:  `Reopen a closed pull request.`,
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
		s.Suffix = " Reopening pull request..."
		s.Start()

		reqBody := map[string]string{"state": "open"}
		data, err := httpclient.Patch(fmt.Sprintf("/repos/%s/pulls/%s", repo, number), reqBody)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to reopen pull request: %v\n", err)
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
		bodyFile, _ := cmd.Flags().GetString("body-file")
		head, _ := cmd.Flags().GetString("head")
		base, _ := cmd.Flags().GetString("base")
		draft, _ := cmd.Flags().GetBool("draft")
		fill, _ := cmd.Flags().GetBool("fill")
		assigneesStr, _ := cmd.Flags().GetString("assignee")
		reviewersStr, _ := cmd.Flags().GetString("reviewer")
		labelsStr, _ := cmd.Flags().GetString("label")
		milestone, _ := cmd.Flags().GetString("milestone")
		web, _ := cmd.Flags().GetBool("web")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		if title == "" && !fill {
			fmt.Fprintln(os.Stderr, "Title required. Use --title")
			os.Exit(1)
			return nil
		}

		// Read body from file if specified
		if bodyFile != "" {
			content, err := os.ReadFile(bodyFile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to read body file: %v\n", err)
				os.Exit(1)
				return nil
			}
			body = string(content)
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Creating pull request..."
		s.Start()

		reqBody := map[string]interface{}{
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
		if draft {
			reqBody["draft"] = true
		}
		if fill {
			reqBody["fill"] = true
		}
		if assigneesStr != "" {
			reqBody["assignees"] = strings.Split(assigneesStr, ",")
		}
		if reviewersStr != "" {
			reqBody["reviewers"] = strings.Split(reviewersStr, ",")
		}
		if labelsStr != "" {
			reqBody["labels"] = strings.Split(labelsStr, ",")
		}
		if milestone != "" {
			reqBody["milestone"] = milestone
		}

		data, err := httpclient.Post(fmt.Sprintf("/repos/%s/pulls", repo), reqBody)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create pull request: %v\n", err)
			os.Exit(1)
			return nil
		}

		if web {
			if m, ok := data.(map[string]interface{}); ok {
				if htmlURL, ok := m["html_url"].(string); ok {
					fmt.Printf("Pull request created! Open in browser: %s\n", htmlURL)
				}
			}
		}

		printJSON(data)
		return nil
	},
}

func init() {
	prListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prListCmd.Flags().StringP("state", "s", "open", "Filter by state: open, closed, all")
	prListCmd.Flags().StringP("limit", "L", "30", "Maximum number of PRs to list")
	prListCmd.Flags().StringP("label", "l", "", "Filter by labels (comma-separated)")
	prListCmd.Flags().StringP("assignee", "a", "", "Filter by assignee")
	prListCmd.Flags().StringP("author", "A", "", "Filter by author")
	prListCmd.Flags().BoolP("draft", "d", false, "Filter by draft state")
	prListCmd.Flags().StringP("search", "S", "", "Search PRs by title/body")

	prViewCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prViewCmd.Flags().BoolP("comments", "c", false, "Show comments")
	prViewCmd.Flags().BoolP("web", "w", false, "Open in browser")

	prCommentsCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")

	prLabelsCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prAddLabelsCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prAddLabelsCmd.Flags().StringArrayP("label", "l", []string{}, "Label name (can specify multiple)")
	prRemoveLabelsCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prRemoveLabelsCmd.Flags().StringP("label", "l", "", "Label name to remove")

	prAssigneesCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prAddAssigneesCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prAddAssigneesCmd.Flags().StringP("assignee", "a", "", "Assignee username(s), comma-separated if multiple")
	prRemoveAssigneesCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prRemoveAssigneesCmd.Flags().StringP("assignee", "a", "", "Assignee username(s), comma-separated if multiple")

	prReviewersCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prAddReviewersCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prAddReviewersCmd.Flags().StringArrayP("reviewer", "r", []string{}, "Reviewer username (can specify multiple)")
	prRemoveReviewersCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prRemoveReviewersCmd.Flags().StringArrayP("reviewer", "r", []string{}, "Reviewer username (can specify multiple)")

	prOperateLogsCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")

	prTestersCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prAddTestersCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prAddTestersCmd.Flags().StringArrayP("tester", "t", []string{}, "Tester username (can specify multiple)")
	prRemoveTestersCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prRemoveTestersCmd.Flags().StringArrayP("tester", "t", []string{}, "Tester username (can specify multiple)")

	prLinkedIssuesCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prLinkIssueCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prLinkIssueCmd.Flags().StringP("issue", "i", "", "Issue number to link")
	prUnlinkIssueCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prUnlinkIssueCmd.Flags().StringP("issue", "i", "", "Issue number to unlink")

	prCloseCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	prReopenCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")

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
	prCreateCmd.Flags().String("body-file", "", "Read body from file")
	prCreateCmd.Flags().StringP("head", "", "", "Source branch")
	prCreateCmd.Flags().StringP("base", "b", "main", "Target branch")
	prCreateCmd.Flags().BoolP("draft", "d", false, "Create as draft")
	prCreateCmd.Flags().BoolP("fill", "f", false, "Auto-fill from commits (title from first commit, body from description)")
	prCreateCmd.Flags().StringP("assignee", "a", "", "Assignee username(s), comma-separated")
	prCreateCmd.Flags().StringP("reviewer", "r", "", "Reviewer username(s), comma-separated")
	prCreateCmd.Flags().StringP("label", "l", "", "Labels, comma-separated")
	prCreateCmd.Flags().String("milestone", "", "Milestone title or number")
	prCreateCmd.Flags().BoolP("web", "w", false, "Open in browser")

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
	prCmd.AddCommand(prReviewersCmd)
	prCmd.AddCommand(prAddReviewersCmd)
	prCmd.AddCommand(prRemoveReviewersCmd)
	prCmd.AddCommand(prOperateLogsCmd)
	prCmd.AddCommand(prTestersCmd)
	prCmd.AddCommand(prAddTestersCmd)
	prCmd.AddCommand(prRemoveTestersCmd)
	prCmd.AddCommand(prLinkedIssuesCmd)
	prCmd.AddCommand(prLinkIssueCmd)
	prCmd.AddCommand(prUnlinkIssueCmd)
	prCmd.AddCommand(prCloseCmd)
	prCmd.AddCommand(prReopenCmd)
	prCmd.AddCommand(prCommitsCmd)
	prCmd.AddCommand(prFilesCmd)
	prCmd.AddCommand(prMergeCmd)
	prCmd.AddCommand(prMergeStatusCmd)
	prCmd.AddCommand(prCreateCmd)
	prCmd.AddCommand(prUpdateCmd)
}
