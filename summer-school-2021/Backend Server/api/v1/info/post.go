package info

import (
	"backend/api/v1/auth"
	"fmt"
	"github.com/emicklei/go-restful"
	"net/http"
)

type ProjID struct {
	ProjectID string
}

func (c *Resource) RegisterPost(container *restful.Container) *Resource {
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)

	ws.Path("/post").Doc("Sb API version 1")

	ws.Route(ws.POST(fmt.Sprintf("/saveid")).To(c.PostProjID).
		Param(ws.BodyParameter("ProjID", "Project id in cloud").DataType("ProjID")).
		Doc("Saves project id").
		Operation("PostProjID"))

	container.Add(ws)

	return c
}

func (c *Resource) PostProjID(req *restful.Request, resp *restful.Response) {
	projID := new(ProjID)
	err := req.ReadEntity(projID)
	if err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	auth.InfoAuth.ProjectID = projID.ProjectID
	logger.Infof("Saved project id: %s", projID.ProjectID)
	resp.WriteHeaderAndEntity(http.StatusOK, "Project id successfully saved")
}
