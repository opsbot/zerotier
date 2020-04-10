package member

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/opsbot/cli-go/utils"
	"github.com/opsbot/zerotier/api"
	"github.com/spf13/cobra"
)

// GetCommand returns a cobra command
func GetCommand() *cobra.Command {
	var network string
	var member string
	var outputDocument string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "get member",
		PreRun: func(cmd *cobra.Command, args []string) {
			if cmd.Flag("network").Value.String() == "" {
				log.Fatal("you must provide a network id")
			}
			if cmd.Flag("member").Value.String() == "" {
				log.Fatal("you must provide a member id")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			network := cmd.Flag("network").Value.String()
			member := cmd.Flag("member").Value.String()
			log.Tracef("get called for network id %v member id %v\n", network, member)

			data := api.MemberGet(network, member)

			if outputDocument != "" {
				utils.FileWrite(outputDocument, []byte(data.Body))
			} else {
				fmt.Println(data.Body)
			}
		},
	}
	cmd.Flags().StringVarP(&network, "network", "n", "", "the network id")
	cmd.MarkFlagRequired("network")

	cmd.Flags().StringVarP(&member, "member", "m", "", "the member id")
	cmd.MarkFlagRequired("member")

	cmd.Flags().StringVarP(&outputDocument, "output-document", "O", "", "Write output to a file")
	return cmd
}
