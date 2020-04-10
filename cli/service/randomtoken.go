package service

import (
	"fmt"

	"github.com/opsbot/zerotier/api"
	"github.com/spf13/cobra"
)

// RandomTokenCommand returns a cobra command
func RandomTokenCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "randomtoken",
		Short: "get a random token",
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		Run: func(cmd *cobra.Command, args []string) {
			data := api.RandomToken()
			fmt.Println(data.Body)
		},
	}

	return cmd
}
