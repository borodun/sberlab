package cloudservers

import (
	"backend/api/requester"
	"backend/api/v1/vpcs"
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/gorilla/schema"
	"github.com/juju/loggo"
	"net/http"
)

var logger = loggo.GetLogger("cloudservers")
var decoder *schema.Decoder

const (
	pathParamProjectID = "project-id"
	queryParamOffset   = "offset"
	queryParamLimit    = "limit"
	queryParamStatus   = "status"
)

type Ecs struct {
	Name      string `json:"name"`
	CpuCount  int    `json:"cpuCount"`
	MemoryGib int    `json:"memoryGib"`
}

type Keys struct {
	AKey string `json:"aKey" description:"Access Key"`
	SKey string `json:"sKey" description:"Secret Key"`
}

type Server struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
}

type ErrorMsg struct {
	Error ErrorWithID `json:"error"`
}

type ErrorWithID struct {
	Message string `json:"message"`
}

type Servers struct {
	ErrorMessage string      `json:"error_msg"`
	Error        ErrorWithID `json:"error"`
	Count        int         `json:"count"`
	ServerList   []Server    `json:"servers"`
}

type Info struct {
	Servers Servers   `json:"ecs"`
	VPCs    vpcs.VPCs `json:"vpcs"`
}

type Resource struct {
	servers  Servers
	authKeys *Keys
}

func NewResource() *Resource {
	logger.SetLogLevel(loggo.INFO)
	return &Resource{}
}

func (c *Resource) Register(container *restful.Container) *Resource {
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)

	ws.Path("/info").Doc("Sb API version 1")

	ws.Route(ws.GET("/").To(c.GetECS).
		Doc("Returns v1 ECS endpoint").
		Operation("GetECS"))

	ws.Route(ws.GET(fmt.Sprintf("/{%s}/getinfo", pathParamProjectID)).To(c.GetECSList).
		Param(ws.PathParameter(pathParamProjectID, "project ID").DataType("string")).
		Param(ws.QueryParameter(queryParamOffset, "Specifies a page number").DataType("integer")).
		Param(ws.QueryParameter(queryParamLimit, "Specifies the maximum number of ECSs on one page.").DataType("integer")).
		Param(ws.QueryParameter(queryParamStatus, "Specifies the ECS status.").DataType("string")).
		Doc("Returns ECSs list").
		Operation("GetECSList"))

	ws.Route(ws.POST(fmt.Sprintf("/keys")).To(c.PostKeys).
		Param(ws.BodyParameter("Keys", "Keys for auth").DataType("Keys")).
		Doc("Saves access and secret authKeys").
		Operation("PostKeys"))

	container.Add(ws)

	return c
}

func (c *Resource) GetECS(request *restful.Request, response *restful.Response) {
	logger.Infof("GetECS")

	endpoint := "use 'ecs/{project-id}/cloudservers'"

	response.WriteHeaderAndEntity(http.StatusOK, endpoint)
}

func (c *Resource) GetECSList(request *restful.Request, response *restful.Response) {
	logger.Infof("GetECSList")

	projectID := request.PathParameter(pathParamProjectID)
	offset := request.QueryParameter(queryParamOffset)
	limit := request.QueryParameter(queryParamLimit)

	if c.authKeys == nil {
		error := new(ErrorMsg)
		error.Error.Message = "Server has no saved keys"
		response.WriteEntity(error)
		return
	}
	var reqUrl = fmt.Sprintf("https://ecs.ru-moscow-1.hc.sbercloud.ru/v1/%s/cloudservers/detail?offset=%s&limit=%s", projectID, offset, limit)
	logger.Infof("proj ID: %s, offset: %s, limit: %s, accessKey: %s, secretKey: %s", projectID, offset, limit, c.authKeys.AKey, c.authKeys.SKey)

	ecssList := requester.MakeRequest(reqUrl, c.authKeys.AKey, c.authKeys.SKey)
	logger.Infof("Response from sber: " + ecssList)
	var servers Servers
	json.Unmarshal([]byte(ecssList), &servers)
	respString, _ := json.Marshal(servers)
	logger.Infof("Response to front: " + string(respString))
	vpcs := vpcs.GetVPCs(limit, projectID, c.authKeys.AKey, c.authKeys.SKey)
	var info Info

	info.Servers = servers
	info.VPCs = vpcs

	response.WriteEntity(info)
}

func (c *Resource) PostKeys(req *restful.Request, resp *restful.Response) {
	logger.Infof("Saving authKeys")
	authKeys := new(Keys)
	err := req.ReadEntity(authKeys)
	if err != nil { // bad request
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	c.authKeys = authKeys
	logger.Infof("Saved authKeys: access key: %s, secret key: %s", authKeys.AKey, authKeys.SKey)
	resp.WriteHeaderAndEntity(http.StatusOK, "Keys successfully saved")
}
