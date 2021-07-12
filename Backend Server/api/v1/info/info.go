package info

import (
	"backend/api/v1/auth"
	"backend/api/v1/ecss"
	"backend/api/v1/messages"
	"backend/api/v1/vpcs"
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/gorilla/schema"
	"github.com/juju/loggo"
	"net/http"
)

var logger = loggo.GetLogger("ecss")
var decoder *schema.Decoder

const (
	pathParamProjectID = "project-id"
	queryParamOffset   = "offset"
	queryParamLimit    = "limit"
	queryParamStatus   = "status"
)

type Resource struct {
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

	ws.Route(ws.GET(fmt.Sprintf("/getinfo")).To(c.GetECSList).
		Param(ws.QueryParameter(queryParamOffset, "Specifies a page number").DataType("integer")).
		Param(ws.QueryParameter(queryParamLimit, "Specifies the maximum number of ECSs on one page.").DataType("integer")).
		Param(ws.QueryParameter(queryParamStatus, "Specifies the ECS status.").DataType("string")).
		Doc("Returns ECSs list").
		Operation("GetECSList"))

	ws.Route(ws.POST(fmt.Sprintf("/keys")).To(c.PostKeys).
		Param(ws.BodyParameter("Keys", "Keys for auth").DataType("Keys")).
		Doc("Saves access and secret authKeys").
		Operation("PostKeys"))

	ws.Route(ws.POST(fmt.Sprintf("/projid")).To(c.PostProjID).
		Param(ws.BodyParameter("ProjID", "Project id in cloud").DataType("ProjID")).
		Doc("Saves project id").
		Operation("PostProjID"))

	container.Add(ws)

	return c
}

func (c *Resource) GetECS(request *restful.Request, response *restful.Response) {
	logger.Infof("GetECS")

	endpoint := "use 'ecss/{project-id}/ecss'"

	response.WriteHeaderAndEntity(http.StatusOK, endpoint)
}

func (c *Resource) GetECSList(request *restful.Request, response *restful.Response) {
	logger.Infof("GetECSList")
	offset := request.QueryParameter(queryParamOffset)
	limit := request.QueryParameter(queryParamLimit)

	if auth.Info.AuthKeys == nil {
		error := new(messages.ErrorMsg)
		error.Error.Message = "Server has no saved keys"
		response.WriteEntity(error)
		return
	}

	ecss := ecss.GetECSs(limit, offset)
	vpcs := vpcs.GetVPCs(limit)
	var info Info
	info.Servers = ecss
	info.VPCs = vpcs

	response.WriteEntity(info)
}

func (c *Resource) PostKeys(req *restful.Request, resp *restful.Response) {
	logger.Infof("Saving authKeys")
	authKeys := new(auth.Keys)
	err := req.ReadEntity(authKeys)
	if err != nil { // bad request
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	auth.Info.AuthKeys = authKeys
	logger.Infof("Saved authKeys: access key: %s, secret key: %s", authKeys.AKey, authKeys.SKey)
	resp.WriteHeaderAndEntity(http.StatusOK, "Keys successfully saved")
}

func (c *Resource) PostProjID(req *restful.Request, resp *restful.Response) {
	logger.Infof("Saving project id")
	projID := new(auth.ProjID)
	err := req.ReadEntity(projID)
	if err != nil { // bad request
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	auth.Info.ProjectID = projID.ProjectID
	logger.Infof("Saved project id: %s", projID.ProjectID)
	resp.WriteHeaderAndEntity(http.StatusOK, "Project id successfully saved")
}
