package user

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/opsbot/cli-go/utils"
	"github.com/opsbot/zerotier/api"
	"github.com/spf13/cobra"
)

// GetCommand returns a cobra command
func GetCommand() *cobra.Command {
	var outputDocument string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "get user",
		PreRun: func(cmd *cobra.Command, args []string) {
			if cmd.Flag("username").Value.String() == "" {
				log.Fatal("you must provide a username")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			username := cmd.Flag("username").Value.String()
			log.Tracef("get called for user id %v\n", username)

			data := api.UserGet(username)

			if outputDocument != "" {
				utils.FileWrite(outputDocument, []byte(data.Body))
			} else {
				fmt.Println(data.Body)
			}
		},
	}
	cmd.Flags().StringVarP(&outputDocument, "output-document", "O", "", "Write output to a file")
	return cmd
}
