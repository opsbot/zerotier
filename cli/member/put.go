package member

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/opsbot/cli-go/utils"
	"github.com/opsbot/zerotier/api"

	"github.com/spf13/cobra"
)

// PutCommand returns a cobra command
func PutCommand() *cobra.Command {
	var network string
	var fileName string
	var fileData []byte

	cmd := &cobra.Command{
		Use:   "put",
		Short: "put network",
		PreRun: func(cmd *cobra.Command, args []string) {
			// get Stdin.Stat so we can check for piped inpput
			info, err := os.Stdin.Stat()
			if err != nil {
				log.Fatal(err)
			}

			if info.Size() > 0 { // fileData read from Stdin
				data, _ := ioutil.ReadAll(os.Stdin)
				fileData = []byte(string(data))
				return
			}

			if fileName != "" { // fileData read from file
				fileData = utils.FileRead(fileName)
				return
			}

			if cmd.Flag("username").Value.String() == "" {
				log.Fatal("you must provide a username")
			}
			if cmd.Flag("password").Value.String() == "" {
				log.Fatal("you must provide a password")
			}
			if cmd.Flag("venue").Value.String() == "" {
				log.Fatal("you must provide a venue id")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			var member api.Member
			network := cmd.Flag("network").Value.String()

			if fileData != nil {
				json.Unmarshal([]byte(fileData), &member)
			} else {
				member = api.Member{}
			}

			data := api.MemberPut(network, member)
			fmt.Println(data.Body)
		},
	}
	cmd.Flags().StringVarP(&network, "network", "n", "", "the network id")
	cmd.MarkFlagRequired("network")

	cmd.Flags().StringVarP(&fileName, "file", "f", "", "file path")
	return cmd
}
