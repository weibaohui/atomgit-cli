package cmd

import (
	"github.com/briandowns/spinner"
	"github.com/weibaohui/atomgit-cli/internal/httpclient"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for resources",
	Long:  `Search for repositories, code, users, and issues.`,
}

var searchReposCmd = &cobra.Command{
	Use:   "repos <query>",
	Short: "Search repositories",
	Long:  `Search for repositories by keyword.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := args[0]
		limit, _ := cmd.Flags().GetString("limit")

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Searching..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/search/repositories?q=%s&limit=%s", query, limit))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to search repos: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var searchUsersCmd = &cobra.Command{
	Use:   "users <query>",
	Short: "Search users",
	Long:  `Search for users by keyword.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := args[0]
		limit, _ := cmd.Flags().GetString("limit")

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Searching..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/search/users?q=%s&limit=%s", query, limit))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to search users: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var searchCodeCmd = &cobra.Command{
	Use:   "code <query>",
	Short: "Search code",
	Long:  `Search for code by keyword.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := args[0]
		limit, _ := cmd.Flags().GetString("limit")

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = " Searching..."
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/search/code?q=%s&limit=%s", query, limit))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to search code: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	searchReposCmd.Flags().StringP("limit", "L", "30", "Maximum number of results")
	searchUsersCmd.Flags().StringP("limit", "L", "30", "Maximum number of results")
	searchCodeCmd.Flags().StringP("limit", "L", "30", "Maximum number of results")

	searchCmd.AddCommand(searchReposCmd)
	searchCmd.AddCommand(searchUsersCmd)
	searchCmd.AddCommand(searchCodeCmd)
}
