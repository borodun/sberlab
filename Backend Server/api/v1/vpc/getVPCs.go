package vpc

import (
	"backend/api/v1/auth"
	errorcheck "backend/api/v1/error-check"
	"backend/api/v1/requester"
	"encoding/json"
	"fmt"
	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("info")

type VPC struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
}

type VPCs struct {
	Error string `json:"error"`
	Count int    `json:"count"`
	VPCs  []VPC  `json:"vpcs"`
}

func GetVPCs(limit string) VPCs {
	var reqUrl = fmt.Sprintf("https://vpc.ru-moscow-1.hc.sbercloud.ru/v1/%s/vpcs?limit=%s", auth.InfoAuth.ProjectID, limit)

	vpcsList := requester.MakeRequest(reqUrl)
	logger.Infof("Response from sber: " + vpcsList)

	var errCheck errorcheck.ErrorCheck
	json.Unmarshal([]byte(vpcsList), &errCheck)
	str := errorcheck.CheckError(errCheck)

	var vpcs VPCs
	json.Unmarshal([]byte(vpcsList), &vpcs)
	vpcs.Error = str
	respString, _ := json.Marshal(vpcs)

	logger.Infof("Response to front: " + string(respString))
	return vpcs
}
