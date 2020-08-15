package site_cms

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"github.com/go-liam/util/uuid"
	"grape/configs/commom"
	"grape/configs/errorcode"
	"grape/internal/pkg/data/home_site/site"
	"io/ioutil"
	"net/http"
)

type SrvAdd struct {
	req *ReqModel
	srv *site.SrvSite
}

func AddGin(c *gin.Context) {
	srv := new(SrvAdd)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &srv.req)
	srv.srv = site.Server
	c.JSON(http.StatusOK, srv.Add())
}

func (e *SrvAdd) Add() *response.APIResponse {
	if e.req.Title == "" {
		return &response.APIResponse{Code: errorcode.RequestParameter, Message: errorcode.MsRequest, Data: response.DataItemNil}
	}
	item := GetModel(e.req)
	item.ID = uuid.AutoInt64ID()
	item.Status = commom.StatusDefault
	e.srv.Create(item)
	e.req.ID = conv.Int64ToString(item.ID)
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: e.req}
}
