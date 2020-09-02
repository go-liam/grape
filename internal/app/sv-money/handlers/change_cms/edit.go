package change_cms

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/data/money/change"
	"io/ioutil"
	"net/http"
)

type SrvEdit struct {
	req    *ReqModel
	srv    change.Service
	result int64
}

func EditGin(c *gin.Context) {
	srv := new(SrvEdit)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &srv.req)
	srv.srv = change.Server
	c.JSON(http.StatusOK, srv.Edit())
}

func (e *SrvEdit) Edit() *response.APIResponse {
	if e.req.UserID <= 0 {
		return &response.APIResponse{Code: errorcode.RequestParameter, Message: errorcode.MsRequest, Data: response.DataItemNil}
	}
	item := GetModel(e.req)
	e.result, _ = e.srv.Update(item)
	if e.result > 0 {
		return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: e.req}
	}
	return &response.APIResponse{Code: errorcode.Database, Message: errorcode.MsDatabase, Data: e.req}
}