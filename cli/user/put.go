package user

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

// UpdateCommand returns a cobra command
func UpdateCommand() *cobra.Command {
	var fileName string
	var fileData []byte

	cmd := &cobra.Command{
		Use:   "put",
		Short: "update user",
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
			// channel := fmt.Sprintf("venue-%v", cmd.Flag("venue").Value.String())
			// disabled, _ := cmd.Flags().GetBool("disabled")

			var user api.User

			if fileData != nil {
				json.Unmarshal([]byte(fileData), &user)
			} else {
				user = api.User{}
			}

			data := api.UserPut(user)
			fmt.Println(data.Body)
		},
	}
	cmd.Flags().StringVarP(&fileName, "file", "f", "", "file path")
	return cmd
}
