package admin

import (
	"encoding/json"
	"fmt"

	"github.com/opsbot/cli-go/utils"

	log "github.com/sirupsen/logrus"


)

func getEndpoint(path string) string {
	log.Trace("getEndPoint")
	sync.CheckConfig()
	var url string
	url = fmt.Sprintf("%v/%v", sync.Config.AdminHost, path)
	log.Tracef("getEndpoint %v\n", url)
	return url
}

func getEndpoint(path string) string {
	log.Trace("getEndpoint")
	utils.PrettyPrint(sync.Config)
	sync.CheckConfig()
	url := fmt.Sprintf("%v/%v/%v", sync.Config.AdminHost, sync.Config.DB, path)
	log.Tracef("getEndpoint %v\n", url)
	return url
}

func showFullResponseBody(resp sync.APIResponse) {
	var data map[string]interface{}
	json.Unmarshal([]byte(resp.Body), &data)
	utils.PrettyPrint(data)
}
