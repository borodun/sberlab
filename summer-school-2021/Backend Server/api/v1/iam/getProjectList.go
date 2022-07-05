package iam

import (
	errorcheck "backend/api/v1/error-check"
	"backend/api/v1/requester"
	"encoding/json"
	"fmt"
	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("info")

type Project struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type Projects struct {
	Error    string    `json:"error"`
	Count    int       `json:"count"`
	Projects []Project `json:"projects"`
}

func GetProjects() Projects {
	var reqUrl = fmt.Sprintf("https://iam.ru-moscow-1.hc.sbercloud.ru/v3/auth/projects")

	projectsList := requester.RequestGet(reqUrl)
	logger.Infof("Response from sber: " + projectsList)

	str := errorcheck.CheckError(projectsList)

	var projs Projects
	json.Unmarshal([]byte(projectsList), &projs)
	projs.Error = str
	projs.Count = len(projs.Projects)
	respString, _ := json.Marshal(projs)

	logger.Infof("Response to front: " + string(respString))
	return projs
}
