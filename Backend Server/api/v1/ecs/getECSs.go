package ecs

import (
	"backend/api/v1/auth"
	errorcheck "backend/api/v1/error-check"
	"backend/api/v1/requester"
	"encoding/json"
	"fmt"
	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("info")

type ECS struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
}

type ECSs struct {
	Error string `json:"error"`
	Count int    `json:"count"`
	ECSs  []ECS  `json:"servers"`
}

func GetECSs(limit string, offset string) ECSs {
	var reqUrl = fmt.Sprintf("https://ecs.ru-moscow-1.hc.sbercloud.ru/v1/%s/cloudservers/detail?offset=%s&limit=%s", auth.InfoAuth.ProjectID, offset, limit)

	ecssList := requester.RequestGet(reqUrl)
	logger.Infof("Response from sber: " + ecssList)

	var ecss ECSs

	var errCheck errorcheck.ErrorCheck
	json.Unmarshal([]byte(ecssList), &errCheck)
	str := errorcheck.CheckError(errCheck)

	json.Unmarshal([]byte(ecssList), &ecss)
	ecss.Error = str
	respString, _ := json.Marshal(ecss)

	logger.Infof("Response to front: " + string(respString))
	return ecss
}
