package network

import (
	"fmt"

	"github.com/opsbot/zerotier/api"
	"github.com/spf13/cobra"
)

// ListCommand returns a cobra command
func ListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list networks",
		Run: func(cmd *cobra.Command, args []string) {
			data := api.NetworkList()
			fmt.Println(data.Body)
		},
	}

	return cmd
}
