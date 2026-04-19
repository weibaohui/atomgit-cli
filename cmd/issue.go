package cmd

import (
	"fmt"
	"os"
	"strings"

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
  $ atomgit issue create -R owner/repo --title "Bug fix" --body "Description"

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
		labels, _ := cmd.Flags().GetString("label")
		assignee, _ := cmd.Flags().GetString("assignee")
		author, _ := cmd.Flags().GetString("author")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading issues..."
		s.Start()

		url := fmt.Sprintf("/repos/%s/issues?state=%s&limit=%s", repo, state, limit)
		if labels != "" {
			url += "&labels=" + labels
		}
		if assignee != "" {
			url += "&assignee=" + assignee
		}
		if author != "" {
			url += "&author=" + author
		}

		data, err := httpclient.Get(url)
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
		comments, _ := cmd.Flags().GetBool("comments")

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

		if comments {
			s2 := spinner.New(spinner.CharSets[14], 100*1000*1000)
			s2.Suffix = " Loading comments..."
			s2.Start()
			commentsData, err := httpclient.Get(fmt.Sprintf("/repos/%s/issues/%s/comments", repo, number))
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

var issueCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an issue",
	Long:  `Create a new issue on AtomGit.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		title, _ := cmd.Flags().GetString("title")
		body, _ := cmd.Flags().GetString("body")
		bodyFile, _ := cmd.Flags().GetString("body-file")
		assigneesStr, _ := cmd.Flags().GetString("assignee")
		labelsStr, _ := cmd.Flags().GetString("label")
		milestone, _ := cmd.Flags().GetString("milestone")
		web, _ := cmd.Flags().GetBool("web")

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
		s.Suffix = " Creating issue..."
		s.Start()

		reqBody := map[string]interface{}{
			"title": title,
		}
		if body != "" {
			reqBody["body"] = body
		}
		if assigneesStr != "" {
			reqBody["assignees"] = strings.Split(assigneesStr, ",")
		}
		if labelsStr != "" {
			reqBody["labels"] = strings.Split(labelsStr, ",")
		}
		if milestone != "" {
			reqBody["milestone"] = milestone
		}

		data, err := httpclient.Post(fmt.Sprintf("/repos/%s/issues", repo), reqBody)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create issue: %v\n", err)
			os.Exit(1)
			return nil
		}

		if web {
			if m, ok := data.(map[string]interface{}); ok {
				if htmlURL, ok := m["html_url"].(string); ok {
					fmt.Printf("Issue created! Open in browser: %s\n", htmlURL)
				}
			}
		}

		printJSON(data)
		return nil
	},
}

func init() {
	issueListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	issueListCmd.Flags().StringP("state", "s", "open", "Filter by state: open, closed, all")
	issueListCmd.Flags().StringP("limit", "L", "30", "Maximum number of issues to list")
	issueListCmd.Flags().StringP("label", "l", "", "Filter by labels (comma-separated)")
	issueListCmd.Flags().StringP("assignee", "a", "", "Filter by assignee")
	issueListCmd.Flags().StringP("author", "A", "", "Filter by author")

	issueViewCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	issueViewCmd.Flags().BoolP("comments", "c", false, "Show comments")

	issueCreateCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	issueCreateCmd.Flags().StringP("title", "t", "", "Issue title (required)")
	issueCreateCmd.Flags().StringP("body", "b", "", "Issue body/description")
	issueCreateCmd.Flags().String("body-file", "", "Read body from file")
	issueCreateCmd.Flags().StringP("assignee", "a", "", "Assignee username(s), comma-separated")
	issueCreateCmd.Flags().StringP("label", "l", "", "Labels, comma-separated")
	issueCreateCmd.Flags().StringP("milestone", "m", "", "Milestone title or number")
	issueCreateCmd.Flags().BoolP("web", "w", false, "Open in browser")

	issueCmd.AddCommand(issueListCmd)
	issueCmd.AddCommand(issueViewCmd)
	issueCmd.AddCommand(issueCreateCmd)
}