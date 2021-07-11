package vpcs

import (
	"backend/api/requester"
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("vpcs")
var decoder *schema.Decoder

type ErrorMsg struct {
	Error ErrorWithID `json:"error"`
}

type ErrorWithID struct {
	Message string `json:"message"`
}

type VPC struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
}

type VPCs struct {
	ErrorMessage string      `json:"error_msg"`
	Error        ErrorWithID `json:"error"`
	Count        int         `json:"count"`
	ServerList   []VPC       `json:"vpcs"`
}

func GetVPCs(limit string, prijectID string, aKey string, sKey string) VPCs {
	var reqUrl = fmt.Sprintf("https://vpc.ru-moscow-1.hc.sbercloud.ru/v1/%s/vpcs?limit=%s", prijectID, limit)

	vpcsList := requester.MakeRequest(reqUrl, aKey, sKey)
	logger.Infof("Response from sber: " + vpcsList)
	var vpcs VPCs
	json.Unmarshal([]byte(vpcsList), &vpcs)
	respString, _ := json.Marshal(vpcs)

	logger.Infof("Response to front: " + string(respString))
	return vpcs
}
