package requester

import (
	"backend/api/v1/auth"
	"backend/api/v1/entites"
	errorcheck "backend/api/v1/error-check"
	"encoding/json"
	"fmt"
)

func MakeUniRequest(uniReq *entites.EntityInfo) entites.Entities {
	var reqUrl = fmt.Sprintf("https://%s/%s/%s%s?limit=%s&offset=%s", uniReq.Endpoint, auth.InfoAuth.ProjectID, uniReq.Link, uniReq.Details, QueryParams.Limit, QueryParams.Offset)
	respStr := MakeRequest(reqUrl)

	var empty entites.Entities
	var errCheck errorcheck.ErrorCheck
	if err := json.Unmarshal([]byte(respStr), &errCheck); err != nil {
		empty.Error = err.Error()
		return empty
	}
	errStr := errorcheck.CheckError(errCheck)
	if len(errStr) != 0 {
		empty.Error = errStr
		return empty
	}

	var ent entites.Entities
	ent.Error = errStr

	switch uniReq.TypeInJSON {
	case "servers":
		ecs := new(entites.ECSs)
		if err := json.Unmarshal([]byte(respStr), ecs); err != nil {
			empty.Error = err.Error()
			return empty
		}
		ent.Entities = ecs.Servers
	case "vpcs":
		vpc := new(entites.VPCs)
		if err := json.Unmarshal([]byte(respStr), vpc); err != nil {
			empty.Error = err.Error()
			return empty
		}
		ent.Entities = vpc.VPCs
	}

	for i := 0; i < len(ent.Entities); i++ {
		ent.Entities[i].Type = uniReq.Type
	}

	return ent
}
