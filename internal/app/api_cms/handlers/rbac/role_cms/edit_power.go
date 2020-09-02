package role_cms

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/data/rbac/role"
	"io/ioutil"
	"net/http"
)

type SrvEditPower struct {
	req    *ReqModel
	srv    role.Service
	result int64
	id     int64
}

func EditPowerGin(c *gin.Context) {
	srv := new(SrvEditPower)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &srv.req)
	srv.srv = role.Server
	srv.id = conv.StringToInt64(c.Param("id"), 0)
	c.JSON(http.StatusOK, srv.Edit())
}

func (e *SrvEditPower) Edit() *response.APIResponse {
	println("id=", e.id)
	println("req=", e.req.PowerIDS)
	if e.req.PowerIDS == "" || e.id <= 0 {
		return &response.APIResponse{Code: errorcode.RequestParameter, Message: errorcode.MsRequest, Data: response.DataItemNil}
	}
	item := GetModel(e.req)
	i, _ := e.srv.FindOne(e.id)
	i.PowerIDS = item.PowerIDS
	e.result, _ = e.srv.Update(i)
	o := GetRespModel(i)
	if e.result > 0 {
		return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: o}
	}
	return &response.APIResponse{Code: errorcode.Database, Message: errorcode.MsDatabase, Data: o}
}
