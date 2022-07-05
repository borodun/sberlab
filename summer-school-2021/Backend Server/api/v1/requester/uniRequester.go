package requester

import (
	"backend/api/v1/auth"
	"backend/api/v1/entites"
	errorcheck "backend/api/v1/error-check"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/juju/loggo"
	"strconv"
)

var logger = loggo.GetLogger("info")

var empty = make([]entites.Entity, 1)

func getBackendServrs(groupID string) ([]entites.Entity, string) {
	var reqUrl = fmt.Sprintf("https://elb.ru-moscow-1.hc.sbercloud.ru/v2.0/lbaas/pools/%s/members?limit=%s", groupID, QueryParams.Limit)
	logger.Infof("Request to sber: " + reqUrl)
	respStr := RequestGet(reqUrl)

	var err error
	errStr := errorcheck.CheckError(respStr)
	if len(errStr) != 0 {
		return empty, "Backend servers: " + errStr
	}

	ent := new(entites.BackendServers)
	if err = json.Unmarshal([]byte(respStr), ent); err != nil {
		return empty, "Backend servers: " + err.Error()
	}

	var ents entites.Entities
	for i := 0; i < len(ent.Entities); i++ {
		ents.Entities = append(ents.Entities, empty...)
		ents.Entities[i].Name = "Server " + strconv.Itoa(i+1)
		ents.Entities[i].ID = ent.Entities[i].ID
		ents.Entities[i].Status = ent.Entities[i].Status
		ents.Entities[i].Type = "backend"
	}

	return ents.Entities, ""
}

func MakeUniRequest(uniReq *entites.EntityGetInfo) ([]entites.Entity, string) {
	jsonType := uniReq.TypeInJSON
	var reqUrl string
	if jsonType == "loadbalancers" || jsonType == "listeners" || jsonType == "pools" {
		reqUrl = fmt.Sprintf("https://%s/%s%s?limit=%", uniReq.Endpoint, uniReq.Link, uniReq.Details, QueryParams.Limit)
	} else {
		reqUrl = fmt.Sprintf("https://%s/%s/%s%s?limit=%s", uniReq.Endpoint, auth.InfoAuth.ProjectID, uniReq.Link, uniReq.Details, QueryParams.Limit)
	}

	logger.Infof("Request to sber: " + reqUrl)

	respStr := RequestGet(reqUrl)
	//logger.Infof("Response from sber: " + respStr)

	var err error
	errStr := errorcheck.CheckError(respStr)
	if len(errStr) != 0 {
		return empty, jsonType + ": " + errStr
	}

	var ents entites.Entities
	switch jsonType {
	case "servers":
		ent := new(entites.ECSs)
		if err = json.Unmarshal([]byte(respStr), ent); err != nil {
			break
		}
		ents.Entities = ent.Entities
		break
	case "vpcs":
		ent := new(entites.VPCs)
		if err = json.Unmarshal([]byte(respStr), ent); err != nil {
			break
		}
		ents.Entities = ent.Entities
		break
	case "volumes":
		ent := new(entites.Volumes)
		if err = json.Unmarshal([]byte(respStr), ent); err != nil {
			break
		}
		ents.Entities = ent.Entities
		break
	case "subnets":
		ent := new(entites.Subnets)
		if err = json.Unmarshal([]byte(respStr), ent); err != nil {
			break
		}
		ents.Entities = ent.Entities
		break
	case "security_groups":
		ent := new(entites.SGs)
		if err = json.Unmarshal([]byte(respStr), ent); err != nil {
			break
		}
		ents.Entities = ent.Entities
		break
	case "nat_gateways":
		ent := new(entites.NATs)
		if err = json.Unmarshal([]byte(respStr), ent); err != nil {
			break
		}
		ents.Entities = ent.Entities
		break
	case "snat_rules":
		ent := new(entites.SNATs)
		if err = json.Unmarshal([]byte(respStr), ent); err != nil {
			break
		}
		for i := 0; i < len(ent.Entities); i++ {
			ents.Entities = append(ents.Entities, empty...)
			ents.Entities[i].Name = ent.Entities[i].Name
			ents.Entities[i].ID = ent.Entities[i].ID
			ents.Entities[i].Status = ent.Entities[i].Status
		}
		break
	case "dnat_rules":
		ent := new(entites.DNATs)
		if err = json.Unmarshal([]byte(respStr), ent); err != nil {
			break
		}
		for i := 0; i < len(ent.Entities); i++ {
			str := strconv.Itoa(ent.Entities[i].External) + "->" + strconv.Itoa(ent.Entities[i].Internal)
			ents.Entities = append(ents.Entities, empty...)
			ents.Entities[i].Name = str
			ents.Entities[i].ID = ent.Entities[i].ID
			ents.Entities[i].Status = ent.Entities[i].Status
		}
		break
	case "publicips":
		ent := new(entites.EIPs)
		if err = json.Unmarshal([]byte(respStr), ent); err != nil {
			break
		}
		ents.Entities = ent.Entities
		break
	case "loadbalancers":
		ent := new(entites.ELBs)
		if err = json.Unmarshal([]byte(respStr), ent); err != nil {
			break
		}
		for i := 0; i < len(ent.Entities); i++ {
			ents.Entities = append(ents.Entities, empty...)
			ents.Entities[i].Name = ent.Entities[i].Name
			ents.Entities[i].ID = ent.Entities[i].ID
			ents.Entities[i].Status = ent.Entities[i].Status
		}
		break
	case "listeners":
		ent := new(entites.Listeners)
		if err = json.Unmarshal([]byte(respStr), ent); err != nil {
			break
		}
		for i := 0; i < len(ent.Entities); i++ {
			ents.Entities = append(ents.Entities, empty...)
			ents.Entities[i].Name = ent.Entities[i].Name
			ents.Entities[i].ID = ent.Entities[i].ID
			ents.Entities[i].Status = "Listening on port: " + strconv.Itoa(ent.Entities[i].Port)
		}
		break
	case "pools":
		ent := new(entites.Pools)
		if err = json.Unmarshal([]byte(respStr), ent); err != nil {
			break
		}
		for i := 0; i < len(ent.Entities); i++ {
			ents.Entities = append(ents.Entities, empty...)
			ents.Entities[i].Name = ent.Entities[i].Name
			ents.Entities[i].ID = ent.Entities[i].ID
			ents.Entities[i].Status = "Protocol: " + ent.Entities[i].Protocol
			ents.Entities[i].Type = uniReq.Type
		}
		for i := 0; i < len(ent.Entities); i++ {
			servers, _ := getBackendServrs(ent.Entities[i].ID)
			ents.Entities = append(ents.Entities, servers...)
		}
		break

	default:
		err = errors.New("No such entity: " + uniReq.TypeInJSON)
		break
	}

	if err != nil {
		return empty, err.Error()
	}

	for i := 0; i < len(ents.Entities) && jsonType != "pools"; i++ {
		ents.Entities[i].Type = uniReq.Type
	}

	entsStr, _ := json.Marshal(ents)
	logger.Infof("Entities: " + string(entsStr))

	return ents.Entities, ""
}
