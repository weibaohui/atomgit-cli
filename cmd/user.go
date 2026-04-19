package cmd

import (
	"fmt"
	"os"

	"github.com/weibaohui/atomgit-cli/internal/httpclient"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Work with users",
	Long:  `Get user information, followers, and following.`,
}

var userInfoCmd = &cobra.Command{
	Use:   "info [username]",
	Short: "Get user information",
	Long:  `Get information about a user. If no username is provided, returns current user.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		username := ""
		if len(args) > 0 {
			username = args[0]
		}

		path := "/user"
		if username != "" {
			path = fmt.Sprintf("/users/%s", username)
		}

		data, err := httpclient.Get(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get user info: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var userFollowersCmd = &cobra.Command{
	Use:   "followers [username]",
	Short: "List user followers",
	Long:  `List followers of a user. If no username is provided, returns current user's followers.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		username := ""
		if len(args) > 0 {
			username = args[0]
		}

		path := "/user/followers"
		if username != "" {
			path = fmt.Sprintf("/users/%s/followers", username)
		}

		data, err := httpclient.Get(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get followers: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var userFollowingCmd = &cobra.Command{
	Use:   "following [username]",
	Short: "List users following",
	Long:  `List users a user is following. If no username is provided, returns current user's following.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		username := ""
		if len(args) > 0 {
			username = args[0]
		}

		path := "/user/following"
		if username != "" {
			path = fmt.Sprintf("/users/%s/following", username)
		}

		data, err := httpclient.Get(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get following: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	userCmd.AddCommand(userInfoCmd)
	userCmd.AddCommand(userFollowersCmd)
	userCmd.AddCommand(userFollowingCmd)
}
