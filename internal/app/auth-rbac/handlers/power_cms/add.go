package power_cms

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"github.com/go-liam/util/uuid"
	"grape/configs/commom"
	"grape/configs/errorcode"
	"grape/internal/pkg/data/rbac/power"
	"io/ioutil"
	"net/http"
)

type SrvAdd struct {
	req    *ReqModel
	srv    power.Service
	result int64
}

func AddGin(c *gin.Context) {
	srv := new(SrvAdd)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &srv.req)
	srv.srv = power.Server
	srv.req.ID = uuid.UUID()
	c.JSON(http.StatusOK, srv.Add())
}

func (e *SrvAdd) Add() *response.APIResponse {
	if e.req.Name == "" {
		return &response.APIResponse{Code: errorcode.RequestParameter, Message: errorcode.MsRequest, Data: response.DataItemNil}
	}
	item := GetModel(e.req)
	item.ID = conv.StringToInt64(e.req.ID, 0) //uuid.AutoInt64ID()
	item.Status = commom.StatusDefault
	e.result, _ = e.srv.Create(item)
	if e.result > 0 {
		return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: e.req}
	}
	return &response.APIResponse{Code: errorcode.Database, Message: errorcode.MsDatabase, Data: e.req}
}
