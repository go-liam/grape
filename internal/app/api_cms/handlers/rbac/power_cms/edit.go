package power_cms

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/data/rbac/power"
	"io/ioutil"
	"net/http"
)

type SrvEdit struct {
	req    *ReqModel
	srv    power.Service
	result int64
	id int64
}

func EditGin(c *gin.Context) {
	srv := new(SrvEdit)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &srv.req)
	srv.srv = power.Server
	srv.id = conv.StringToInt64(c.Param("id"),0)
	c.JSON(http.StatusOK, srv.Edit())
}

func (e *SrvEdit) Edit() *response.APIResponse {
	if e.req.Name == "" || e.id <= 0 {
		return &response.APIResponse{Code: errorcode.RequestParameter, Message: errorcode.MsRequest, Data: response.DataItemNil}
	}
	item := GetModel(e.req)
	item.ID = e.id
	e.result, _ = e.srv.Update(item)
	e.req.ID = conv.Int64ToString(e.id)
	if e.result > 0 {
		return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: e.req}
	}
	return &response.APIResponse{Code: errorcode.Database, Message: errorcode.MsDatabase, Data: e.req}
}
