package site

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/data/home_site/site"
	"io/ioutil"
	"net/http"
)

type SrvSiteAdd struct {
	req *ReqModel
	srv *site.SrvSite
}

func CreateGin(c *gin.Context) {
	srv := new(SrvSiteAdd)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &srv.req)
	srv.srv = site.Server
	c.JSON(http.StatusOK, srv.CreateSite())
}

func (e *SrvSiteAdd) CreateSite() *response.APIResponse {
	if e.req.Title == "" {
		return &response.APIResponse{Code: errorcode.RequestParameter, Message: errorcode.MsRequest, Data: response.DataItemNil}
	}
	item := GetModel(e.req)
	id, err := e.srv.Create(item)
	if err != nil {
		return &response.APIResponse{Code: errorcode.Database, Message: errorcode.MsDatabase, Data: response.DataItemNil}
	}
	e.req.ID = id
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: e.req}
}
