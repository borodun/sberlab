package vpc

import (
	"backend/api/v1/auth"
	errorcheck "backend/api/v1/error-check"
	"backend/api/v1/requester"
	"encoding/json"
	"fmt"
)

type NAT struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
}

type NATs struct {
	Error string `json:"error"`
	Count int    `json:"count"`
	NATs  []NAT  `json:"nat_gateways"`
}

func GetNATs(limit string) NATs {
	var reqUrl = fmt.Sprintf("https://nat.ru-moscow-1.hc.sbercloud.ru/v2/%s/nat_gateways?limit=%s", auth.InfoAuth.ProjectID, limit)

	natsList := requester.MakeRequest(reqUrl)
	logger.Infof("Response from sber: " + natsList)

	var errCheck errorcheck.ErrorCheck
	json.Unmarshal([]byte(natsList), &errCheck)
	str := errorcheck.CheckError(errCheck)

	var nats NATs
	json.Unmarshal([]byte(natsList), &nats)
	nats.Error = str

	respString, _ := json.Marshal(nats)

	logger.Infof("Response to front: " + string(respString))
	return nats
}
