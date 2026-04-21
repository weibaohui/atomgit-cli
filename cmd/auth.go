package cmd

import (
	"fmt"
	"os"

	"github.com/weibaohui/atomgit-cli/internal/config"
	"github.com/weibaohui/atomgit-cli/internal/httpclient"

	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication",
	Long:  `Authenticate with AtomGit. Manage tokens and authentication status.`,
}

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to AtomGit",
	Long:  `Authenticate with AtomGit (stores token in user config).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		token, _ := cmd.Flags().GetString("token")
		if token == "" {
			fmt.Print("Paste your AtomGit token: ")
			fmt.Scanln(&token)
		}
		token = trimString(token)
		if token == "" {
			fmt.Fprintln(os.Stderr, "No token provided.")
			os.Exit(1)
			return nil
		}

		if err := config.SetToken(token); err != nil {
			return fmt.Errorf("failed to save token: %w", err)
		}
		fmt.Println("Token saved to your user config.")
		return nil
	},
}

var authLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from AtomGit",
	Long:  `Remove stored token.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.DeleteToken(); err != nil {
			return fmt.Errorf("failed to remove token: %w", err)
		}
		fmt.Println("Token removed.")
		return nil
	},
}

var authStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show auth status",
	Long:  `Show authentication status and token source.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		source := config.TokenSource()
		if source == "none" {
			fmt.Fprintf(os.Stderr, "Not authenticated. %s\n", config.UserConfigHint())
			os.Exit(1)
			return nil
		}

		user, err := fetchUsername()
		if err != nil {
			fmt.Printf("Authenticated (token from %s).\n", source)
			return nil
		}
		fmt.Printf("Logged in to api.atomgit.com as %s (%s)\n", user, source)
		return nil
	},
}

var authTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Print token hint",
	Long:  `Print a redacted token hint.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		token := config.GetToken()
		if token == "" {
			fmt.Fprintln(os.Stderr, "No token stored (try 'atg auth login').")
			os.Exit(1)
			return nil
		}
		fmt.Println(config.MaskToken(token))
		return nil
	},
}

func fetchUsername() (string, error) {
	user, err := httpclient.Get("/user")
	if err != nil {
		return "", err
	}

	if userMap, ok := user.(map[string]interface{}); ok {
		if login, ok := userMap["login"].(string); ok {
			return login, nil
		}
		if username, ok := userMap["username"].(string); ok {
			return username, nil
		}
		if name, ok := userMap["name"].(string); ok {
			return name, nil
		}
	}
	return "Unknown", nil
}

func trimString(s string) string {
	if len(s) == 0 {
		return s
	}
	start := 0
	end := len(s)
	for start < end && (s[start] == ' ' || s[start] == '\t' || s[start] == '\n' || s[start] == '\r') {
		start++
	}
	for end > start && (s[end-1] == ' ' || s[end-1] == '\t' || s[end-1] == '\n' || s[end-1] == '\r') {
		end--
	}
	return s[start:end]
}

func init() {
	authLoginCmd.Flags().StringP("token", "t", "", "API token")

	authCmd.AddCommand(authLoginCmd)
	authCmd.AddCommand(authLogoutCmd)
	authCmd.AddCommand(authStatusCmd)
	authCmd.AddCommand(authTokenCmd)
}
