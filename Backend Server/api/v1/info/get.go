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

var getConfigArray []entites.EntityGetInfo

const (
	queryParamOffset = "offset"
	queryParamLimit  = "limit"
	queryParamID     = "id"
	queryParamType   = "type"
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

	ws.Path("/get").Doc("Sb API version 1")

	ws.Route(ws.GET(fmt.Sprintf("/entities")).To(c.GetProjectEntities).
		Param(ws.QueryParameter(queryParamOffset, "Specifies a page number").DataType("integer")).
		Param(ws.QueryParameter(queryParamLimit, "Specifies the maximum number of ECSs on one page.").DataType("integer")).
		Doc("Returns ECSs list").
		Operation("GetProjectEntities"))

	ws.Route(ws.GET(fmt.Sprintf("/detail")).To(c.GetDetail).
		Param(ws.QueryParameter(queryParamID, "Specifies id of an entity").DataType("string")).
		Param(ws.QueryParameter(queryParamType, "Specifies the type of entity").DataType("string")).
		Doc("Returns detailed info about entity").
		Operation("GetDetail"))

	ws.Route(ws.GET(fmt.Sprintf("/projects")).To(c.GetProjects).
		Param(ws.BodyParameter("Keys", "Keys for auth").DataType("Keys")).
		Doc("Saves access and secret auth").
		Operation("GetProjects"))

	ws.Route(ws.GET(fmt.Sprintf("/token")).To(c.CreateToken).
		Doc("Creates token if it hasn't been created or expired").
		Operation("CreateToken"))

	container.Add(ws)

	getConfigArray = GetEntities()

	return c
}

func (c *Resource) GetDetail(request *restful.Request, response *restful.Response) {
	ID := request.QueryParameter(queryParamID)
	Type := request.QueryParameter(queryParamType)

	var details entites.Details
	var detailStr string
	var check bool
	for i := range getConfigArray {
		if getConfigArray[i].Type == Type {
			detailStr, check = requester.GetEntityDetail(&getConfigArray[i], ID)
			if check {
				details.Error = detailStr
				break
			}
			details.Details = detailStr
			break
		}
	}

	response.WriteEntity(details)
}

func GetEntities() []entites.EntityGetInfo {
	var util entites.EntityGetArray
	config, _ := ioutil.ReadFile("./api/v1/info/getConfig.json")
	err := json.Unmarshal(config, &util)
	if err != nil {
		return nil
	}
	return util.EntityGetInfos
}

func (c *Resource) GetProjectEntities(request *restful.Request, response *restful.Response) {
	requester.QueryParams.Limit = request.QueryParameter(queryParamLimit)
	requester.QueryParams.Offset = request.QueryParameter(queryParamOffset)

	var ents []entites.Entity
	var ent []entites.Entity
	var answerEnts entites.AnswerEntities
	var err string
	for i := 0; i < len(getConfigArray); i++ {
		if ent, err = requester.MakeUniRequest(&getConfigArray[i]); len(err) != 0 {
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
