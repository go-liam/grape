package log_cms

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/data/sylog"
	"io/ioutil"
	"net/http"
)

type SrvUpdateStatus struct {
	srv sylog.Service
	req *reqStatus
}

type reqStatus struct {
	//IDs    []string `json:"IDs" `
	//Status int      `json:"status" `
	BeforeTime int64 `json:"beforeTime" `
}

func UpdateStatusGin(c *gin.Context) {
	srv := new(SrvUpdateStatus)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &srv.req)
	srv.srv = sylog.Server
	c.JSON(http.StatusOK, srv.UpdateStatus())
}

func (e *SrvUpdateStatus) UpdateStatus() *response.APIResponse {
	if e.req == nil {
		return &response.APIResponse{Code: errorcode.RequestParameter, Message: errorcode.MsRequest, Data: response.DataItemNil}
	}
	//ids := conv.ArrayStringToInt64(e.req.IDs)
	result, _ := e.srv.Delete(e.req.BeforeTime)
	if result > 0 {
		return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: response.DataItemNil}
	}
	return &response.APIResponse{Code: errorcode.Database, Message: errorcode.MsDatabase, Data: response.DataItemNil}
}
