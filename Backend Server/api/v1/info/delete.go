package info

import (
	"backend/api/v1/entites"
	"backend/api/v1/requester"
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful"
	"io/ioutil"
	"net/http"
)

type EntityForDelete struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

var deleteConfigArray []entites.EntityDeleteInfo

func (c *Resource) RegisterDelete(container *restful.Container) *Resource {
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)

	ws.Path("/delete").Doc("Sb API version 1")

	ws.Route(ws.DELETE(fmt.Sprintf("")).To(c.DeleteEntity).
		Param(ws.BodyParameter("EntityForDelete", "Entity to delete").DataType("EntityForDelete")).
		Doc("Deletes entity").
		Operation("DeleteEntity"))

	container.Add(ws)

	deleteConfigArray = DeleteEntities()

	return c
}

func DeleteEntities() []entites.EntityDeleteInfo {
	var util entites.EntityDeleteArray
	config, _ := ioutil.ReadFile("./api/v1/info/deleteConfig.json")
	err := json.Unmarshal(config, &util)
	if err != nil {
		return nil
	}
	return util.EntityDeleteInfos
}

func (c *Resource) DeleteEntity(request *restful.Request, response *restful.Response) {
	entityForDelete := new(EntityForDelete)
	err := request.ReadEntity(entityForDelete)
	if err != nil {
		response.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	ID := entityForDelete.ID
	Type := entityForDelete.Type

	var answer entites.DeleteAnswer
	var detailStr string
	for i := range deleteConfigArray {
		if deleteConfigArray[i].Type == Type {
			detailStr = requester.DeleteEntity(&deleteConfigArray[i], ID)
			answer.Details = detailStr
			break
		}
	}

	response.WriteEntity(answer)
}
