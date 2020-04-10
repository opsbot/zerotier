package member

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/opsbot/zerotier/api"
	"github.com/spf13/cobra"
)

// DeleteCommand returns a cobra command
func DeleteCommand() *cobra.Command {
	var network string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete member",
		PreRun: func(cmd *cobra.Command, args []string) {
			if cmd.Flag("network").Value.String() == "" {
				log.Fatal("you must provide a network id")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			data := api.NetworkDelete(cmd.Flag("network").Value.String())
			fmt.Println(data.Body)
		},
	}
	cmd.Flags().StringVarP(&network, "network", "n", "", "the network id")
	cmd.MarkFlagRequired("network")

	return cmd
}
