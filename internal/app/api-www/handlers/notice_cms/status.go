package notice_cms

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/data/home_site/notice"
	"io/ioutil"
	"net/http"
)

type SrvUpdateStatus struct {
	srv notice.Service
	req *reqStatus
}

type reqStatus struct {
	IDs    []string `json:"IDs" `
	Status int      `json:"status" `
}

func UpdateStatusGin(c *gin.Context) {
	srv := new(SrvUpdateStatus)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &srv.req)
	srv.srv = notice.Server
	c.JSON(http.StatusOK, srv.UpdateStatus())
}

func (e *SrvUpdateStatus) UpdateStatus() *response.APIResponse {
	if e.req == nil || len(e.req.IDs) <= 0 {
		return &response.APIResponse{Code: errorcode.RequestParameter, Message: errorcode.MsRequest, Data: response.DataItemNil}
	}
	ids := conv.ArrayStringToInt64(e.req.IDs)
	result, _ := e.srv.UpdateStatusByIDs(e.req.Status, ids)
	if result > 0 {
		return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: response.DataItemNil}
	}
	return &response.APIResponse{Code: errorcode.Database, Message: errorcode.MsDatabase, Data: response.DataItemNil}
}
