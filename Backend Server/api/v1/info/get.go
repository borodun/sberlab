package info

import (
	"backend/api/v1/auth"
	"backend/api/v1/entites"
	"backend/api/v1/iam"
	"backend/api/v1/requester"
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/juju/loggo"
	"io/ioutil"
	"net/http"
	"os"
	"time"
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

func (c *Resource) RegisterGet(container *restful.Container) *Resource {
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)

	ws.Path("/info").Doc("Sb API version 1")

	ws.Route(ws.GET(fmt.Sprintf("/entities")).To(c.GetProjectEntities).
		Param(ws.QueryParameter(queryParamOffset, "Specifies a page number").DataType("integer")).
		Param(ws.QueryParameter(queryParamLimit, "Specifies the maximum number of ECSs on one page.").DataType("integer")).
		Param(ws.QueryParameter(queryParamStatus, "Specifies the ECS status.").DataType("string")).
		Doc("Returns ECSs list").
		Operation("GetProjectEntities"))

	ws.Route(ws.GET(fmt.Sprintf("/projects")).To(c.GetProjects).
		Param(ws.BodyParameter("Keys", "Keys for auth").DataType("Keys")).
		Doc("Saves access and secret auth").
		Operation("GetProjects"))

	ws.Route(ws.GET(fmt.Sprintf("/token")).To(c.CreateToken).
		Doc("Creates token if it hasn't been created or expired").
		Operation("CreateToken"))

	ws.Route(ws.POST(fmt.Sprintf("/projid")).To(c.PostProjID).
		Param(ws.BodyParameter("ProjID", "Project id in cloud").DataType("ProjID")).
		Doc("Saves project id").
		Operation("PostProjID"))

	container.Add(ws)

	return c
}

func GetEntities() []entites.EntityInfo {
	var utilArray entites.EntityArray
	config, _ := ioutil.ReadFile("./api/v1/info/config.json")
	err := json.Unmarshal(config, &utilArray)
	if err != nil {
		return nil
	}
	return utilArray.EntityInfos
}

func (c *Resource) GetProjectEntities(request *restful.Request, response *restful.Response) {
	requester.QueryParams.Limit = request.QueryParameter(queryParamLimit)
	requester.QueryParams.Offset = request.QueryParameter(queryParamOffset)

	var utilArray = GetEntities()
	var ents []entites.Entity
	var ent []entites.Entity
	var answerEnts entites.AnswerEntities
	var err string
	for i := 0; i < len(utilArray); i++ {
		if ent, err = requester.MakeUniRequest(&utilArray[i]); len(err) != 0 {
			continue
		}
		ents = append(ents, ent...)
	}
	answerEnts.Error = err
	answerEnts.AnswerEntities = ents

	response.WriteEntity(answerEnts)
}

func (c *Resource) CreateToken(req *restful.Request, resp *restful.Response) {
	if requester.Tok.CreationTime.Before(time.Now().Add(-1 * time.Hour)) {
		authLog := new(auth.Login)
		authLog.Login = os.Getenv("LOGIN")
		authLog.Password = os.Getenv("PASSWORD")
		authLog.DomainName = "fitnsu"
		auth.InfoAuth.Auth = authLog

		errStr := iam.GetToken(authLog.Login, authLog.Password, authLog.DomainName)
		logger.Infof("Token created: " + requester.Tok.Token)

		resp.WriteHeaderAndEntity(http.StatusOK, errStr)
		return
	}
	resp.WriteHeaderAndEntity(http.StatusOK, "")
}

func (c *Resource) GetProjects(req *restful.Request, resp *restful.Response) {
	projs := iam.GetProjects()
	resp.WriteHeaderAndEntity(http.StatusOK, projs)
}
