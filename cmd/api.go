package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/weibaohui/atomgit-cli/internal/httpclient"

	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api <endpoint>",
	Short: "Make authenticated API requests",
	Long:  `Make authenticated API requests to AtomGit.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		endpoint := args[0]
		method, _ := cmd.Flags().GetString("method")
		data, _ := cmd.Flags().GetString("data")

		var body interface{}
		if data != "" {
			if err := json.Unmarshal([]byte(data), &body); err != nil {
				fmt.Fprintf(os.Stderr, "Invalid JSON data: %v\n", err)
				os.Exit(1)
				return nil
			}
		}

		headers := make(map[string]string)
		resp, err := httpclient.Request(endpoint, httpclient.HttpOptions{
			Method:  method,
			Body:    body,
			Headers: headers,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "API request failed: %v\n", err)
			os.Exit(1)
			return nil
		}
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
		fmt.Println(string(respBody))
		return nil
	},
}

func init() {
	apiCmd.Flags().StringP("method", "X", "GET", "HTTP method")
	apiCmd.Flags().StringP("data", "d", "", "JSON data for POST/PATCH/PUT")
}
