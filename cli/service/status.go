package service

import (
	"fmt"

	"github.com/opsbot/zerotier/api"
	"github.com/spf13/cobra"
)

// StatusCommand returns a cobra command
func StatusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "server status",
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		Run: func(cmd *cobra.Command, args []string) {
			data := api.Status()
			fmt.Println(data.Body)
		},
	}

	return cmd
}
