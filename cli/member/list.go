package member

import (
	"fmt"

	"github.com/opsbot/zerotier/api"
	"github.com/spf13/cobra"
)

// ListCommand returns a cobra command
func ListCommand() *cobra.Command {
	var network string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "list members",
		Run: func(cmd *cobra.Command, args []string) {
			network := cmd.Flag("network").Value.String()
			data := api.MemberList(network)
			fmt.Println(data.Body)
		},
	}
	cmd.Flags().StringVarP(&network, "network", "n", "", "the network id")
	cmd.MarkFlagRequired("network")

	return cmd
}
