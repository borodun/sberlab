package info

import (
	"backend/api/v1/auth"
	"backend/api/v1/ecss"
	"backend/api/v1/iam"
	"backend/api/v1/vpcs"
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/juju/loggo"
	"net/http"
)

var logger = loggo.GetLogger("info")

const (
	queryParamOffset = "offset"
	queryParamLimit  = "limit"
	queryParamStatus = "status"
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

	ws.Route(ws.GET(fmt.Sprintf("/ecs")).To(c.GetECSList).
		Param(ws.QueryParameter(queryParamOffset, "Specifies a page number").DataType("integer")).
		Param(ws.QueryParameter(queryParamLimit, "Specifies the maximum number of ECSs on one page.").DataType("integer")).
		Param(ws.QueryParameter(queryParamStatus, "Specifies the ECS status.").DataType("string")).
		Doc("Returns ECSs list").
		Operation("GetECSList"))

	ws.Route(ws.GET(fmt.Sprintf("/vpc")).To(c.GetVPCList).
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

func (c *Resource) GetECSList(request *restful.Request, response *restful.Response) {
	offset := request.QueryParameter(queryParamOffset)
	limit := request.QueryParameter(queryParamLimit)

	if auth.InfoAuth.Signer == nil {
		type Error struct {
			Error string `json:"error"`
		}
		var err Error
		err.Error = "Server has no saved keys"
		logger.Infof("Error: no saved keys")
		response.WriteEntity(err)
		return
	}

	ecssT := ecss.GetECSs(limit, offset)

	response.WriteEntity(ecssT)
}

func (c *Resource) GetVPCList(request *restful.Request, response *restful.Response) {
	limit := request.QueryParameter(queryParamLimit)

	if auth.InfoAuth.Signer == nil {
		type Error struct {
			Error string `json:"error"`
		}
		var err Error
		err.Error = "Server has no saved keys"
		logger.Infof("Error: no saved keys")
		response.WriteEntity(err)
		return
	}

	vpcsT := vpcs.GetVPCs(limit)
	response.WriteEntity(vpcsT)
}

func (c *Resource) PostKeys(req *restful.Request, resp *restful.Response) {
	authKeys := new(auth.Keys)
	err := req.ReadEntity(authKeys)
	if err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	auth.InfoAuth.Signer = authKeys

	errStr := iam.GetToken(authKeys.AKey, authKeys.SKey, authKeys.DomainName)
	fmt.Printf("Error with token: " + errStr)
	projs := iam.GetProjects()

	logger.Infof("Saved authKeys: access key: %s, secret key: %s", authKeys.AKey, authKeys.SKey)
	resp.WriteHeaderAndEntity(http.StatusOK, projs)
}

func (c *Resource) PostProjID(req *restful.Request, resp *restful.Response) {
	projID := new(auth.ProjID)
	err := req.ReadEntity(projID)
	if err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	auth.InfoAuth.ProjectID = projID.ProjectID
	logger.Infof("Saved project id: %s", projID.ProjectID)
	resp.WriteHeaderAndEntity(http.StatusOK, "Project id successfully saved")
}
