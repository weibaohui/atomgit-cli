package cmd

import (
	"fmt"
	"os"

	"github.com/briandowns/spinner"
	"github.com/weibaohui/atomgit-cli/internal/httpclient"
	"github.com/spf13/cobra"
)

var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Work with webhooks",
	Long:  `List, create, view, update, delete and test webhooks.`,
}

var hookListCmd = &cobra.Command{
	Use:   "list",
	Short: "List webhooks",
	Long:  `List all webhooks in a repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Loading webhooks..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/hooks", repo))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list webhooks: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var hookCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a webhook",
	Long:  `Create a new webhook in a repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		url, _ := cmd.Flags().GetString("url")
		secret, _ := cmd.Flags().GetString("secret")
		events, _ := cmd.Flags().GetString("events")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}
		if url == "" {
			fmt.Fprintln(os.Stderr, "Webhook URL required (--url)")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Creating webhook..."
		s.Start()

		body := map[string]interface{}{
			"url": url,
		}
		if secret != "" {
			body["secret"] = secret
		}
		if events != "" {
			body["events"] = events
		}

		data, err := httpclient.Post(fmt.Sprintf("/repos/%s/hooks", repo), body)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create webhook: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var hookViewCmd = &cobra.Command{
	Use:   "view <id>",
	Short: "View a webhook",
	Long:  `View details of a specific webhook.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		id := args[0]

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = fmt.Sprintf(" Loading webhook %s...", id)
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/repos/%s/hooks/%s", repo, id))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to view webhook: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var hookUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a webhook",
	Long:  `Update a webhook.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		id := args[0]
		url, _ := cmd.Flags().GetString("url")
		secret, _ := cmd.Flags().GetString("secret")
		events, _ := cmd.Flags().GetString("events")

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		body := map[string]interface{}{}
		if url != "" {
			body["url"] = url
		}
		if secret != "" {
			body["secret"] = secret
		}
		if events != "" {
			body["events"] = events
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = fmt.Sprintf(" Updating webhook %s...", id)
		s.Start()

		data, err := httpclient.Patch(fmt.Sprintf("/repos/%s/hooks/%s", repo, id), body)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to update webhook: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var hookDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a webhook",
	Long:  `Delete a webhook.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		id := args[0]

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = fmt.Sprintf(" Deleting webhook %s...", id)
		s.Start()

		err := httpclient.Delete(fmt.Sprintf("/repos/%s/hooks/%s", repo, id))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to delete webhook: %v\n", err)
			os.Exit(1)
			return nil
		}

		fmt.Printf("Webhook %s deleted successfully.\n", id)
		return nil
	},
}

var hookTestCmd = &cobra.Command{
	Use:   "test <id>",
	Short: "Test a webhook",
	Long:  `Send a test event to a webhook.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := cmd.Flags().GetString("repo")
		id := args[0]

		if repo == "" {
			fmt.Fprintln(os.Stderr, "Repository required. Use -R owner/repo")
			os.Exit(1)
			return nil
		}

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = fmt.Sprintf(" Testing webhook %s...", id)
		s.Start()

		data, err := httpclient.Post(fmt.Sprintf("/repos/%s/hooks/%s/tests", repo, id), nil)
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to test webhook: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	hookListCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	hookCreateCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	hookCreateCmd.Flags().StringP("url", "", "", "Webhook URL (required)")
	hookCreateCmd.Flags().StringP("secret", "", "", "Webhook secret")
	hookCreateCmd.Flags().StringP("events", "", "*", "Events to trigger webhook")
	hookViewCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	hookUpdateCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	hookUpdateCmd.Flags().StringP("url", "", "", "Webhook URL")
	hookUpdateCmd.Flags().StringP("secret", "", "", "Webhook secret")
	hookUpdateCmd.Flags().StringP("events", "", "", "Events to trigger webhook")
	hookDeleteCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")
	hookTestCmd.Flags().StringP("repo", "R", "", "Select repository (OWNER/REPO)")

	hookCmd.AddCommand(hookListCmd)
	hookCmd.AddCommand(hookCreateCmd)
	hookCmd.AddCommand(hookViewCmd)
	hookCmd.AddCommand(hookUpdateCmd)
	hookCmd.AddCommand(hookDeleteCmd)
	hookCmd.AddCommand(hookTestCmd)
}
