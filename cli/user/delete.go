package user

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/opsbot/zerotier/api"
	"github.com/spf13/cobra"
)

// DeleteCommand returns a cobra command
func DeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete user",
		PreRun: func(cmd *cobra.Command, args []string) {
			if cmd.Flag("username").Value.String() == "" {
				log.Fatal("you must provide a username")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			data := api.UserDelete(cmd.Flag("username").Value.String())
			fmt.Println(data.Body)
		},
	}
	return cmd
}
