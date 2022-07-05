package requester

import (
	"backend/api/v1/auth"
	"backend/api/v1/entites"
	errorcheck "backend/api/v1/error-check"
	"fmt"
)

func GetEntityDetail(uniReq *entites.EntityGetInfo, ID string) (string, bool) {
	jsonType := uniReq.TypeInJSON
	var reqUrl string
	if jsonType == "loadbalancers" || jsonType == "listeners" || jsonType == "pools" {
		reqUrl = fmt.Sprintf("https://%s/%s/%s", uniReq.Endpoint, uniReq.Link, ID)
	} else {
		reqUrl = fmt.Sprintf("https://%s/%s/%s/%s", uniReq.Endpoint, auth.InfoAuth.ProjectID, uniReq.Link, ID)
	}

	logger.Infof("Request to sber: " + reqUrl)

	respStr := RequestGet(reqUrl)
	logger.Infof("Response from sber: " + respStr)

	errStr := errorcheck.CheckError(respStr)
	if len(errStr) != 0 {
		return jsonType + ": " + errStr, true
	}

	return respStr, false
}
