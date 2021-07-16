package requester

import (
	"backend/api/v1/auth"
	"backend/api/v1/entites"
	"fmt"
)

func DeleteEntity(uniReq *entites.EntityDeleteInfo, ID string) string {
	idRequired := uniReq.IDRequired
	var reqUrl string
	if idRequired {
		reqUrl = fmt.Sprintf("https://%s/%s/%s/%s", uniReq.Endpoint, auth.InfoAuth.ProjectID, uniReq.Link, ID)
	} else {
		reqUrl = fmt.Sprintf("https://%s/%s/%s", uniReq.Endpoint, uniReq.Link, ID)
	}

	logger.Infof("Request to sber: " + reqUrl)

	respStr := RequestDelete(reqUrl, "DELETE")

	logger.Infof("Response from sber: " + respStr)

	return respStr
}
