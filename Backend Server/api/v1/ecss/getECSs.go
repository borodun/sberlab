package ecss

import (
	"backend/api/requester"
	"backend/api/v1/auth"
	"backend/api/v1/messages"
	"encoding/json"
	"fmt"
	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("getECSs")

type Server struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
}

type Servers struct {
	ErrorMessage string               `json:"error_msg"`
	Error        messages.ErrorWithID `json:"error"`
	Count        int                  `json:"count"`
	ServerList   []Server             `json:"servers"`
}

func GetECSs(limit string, offset string) Servers {
	var reqUrl = fmt.Sprintf("https://ecss.ru-moscow-1.hc.sbercloud.ru/v1/%s/ecss/detail?offset=%s&limit=%s", auth.Info.ProjectID, offset, limit)

	ecssList := requester.MakeRequest(reqUrl, auth.Info.AuthKeys.AKey, auth.Info.AuthKeys.SKey)
	logger.Infof("Response from sber: " + ecssList)
	var ecss Servers
	json.Unmarshal([]byte(ecssList), &ecss)
	respString, _ := json.Marshal(ecss)

	logger.Infof("Response to front: " + string(respString))
	return ecss
}
