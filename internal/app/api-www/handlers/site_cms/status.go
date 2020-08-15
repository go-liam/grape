package site_cms

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/data/home_site/site"
	"io/ioutil"
	"net/http"
)

type SrvUpdateStatus struct {
	srv *site.SrvSite
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
	srv.srv = site.Server
	c.JSON(http.StatusOK, srv.UpdateStatus())
}

func (e *SrvUpdateStatus) UpdateStatus() *response.APIResponse {
	if e.req == nil || len(e.req.IDs) <= 0 {
		return &response.APIResponse{Code: errorcode.RequestParameter, Message: errorcode.MsRequest, Data: response.DataItemNil}
	}
	ids := conv.ArrayStringToInt64(e.req.IDs)
	e.srv.UpdateStatusByIDs(e.req.Status, ids)
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: e.req}
}
