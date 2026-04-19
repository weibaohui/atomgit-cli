package cmd

import (
	"github.com/briandowns/spinner"
	"github.com/weibaohui/atomgit-cli/internal/httpclient"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var orgCmd = &cobra.Command{
	Use:   "org",
	Short: "Work with organizations",
	Long:  `Get organization information and members.`,
}

var orgInfoCmd = &cobra.Command{
	Use:   "info [org]",
	Short: "Get organization info",
	Long:  `Get information about an organization.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		org := args[0]

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = fmt.Sprintf(" Loading org %s...", org)
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/orgs/%s", org))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get org info: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

var orgMembersCmd = &cobra.Command{
	Use:   "members <org>",
	Short: "List organization members",
	Long:  `List members of an organization.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		org := args[0]

		s := spinner.New(spinner.CharSets[14], 100*1000*1000)
		s.Suffix = fmt.Sprintf(" Loading members of %s...", org)
		s.Start()

		data, err := httpclient.Get(fmt.Sprintf("/orgs/%s/members", org))
		s.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list org members: %v\n", err)
			os.Exit(1)
			return nil
		}

		printJSON(data)
		return nil
	},
}

func init() {
	orgCmd.AddCommand(orgInfoCmd)
	orgCmd.AddCommand(orgMembersCmd)
}
