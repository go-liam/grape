package user_api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/data/rbac/user"
	"net/http"
)

type SrvInfo struct {
	id   int64
	srv  user.Service //*user.SrvSite
	info *user.Model
}

func GetInfoGin(c *gin.Context) {
	srv := new(SrvInfo)
	srv.id = conv.StringToInt64(c.Param("id"), 0)
	srv.srv = user.Server
	c.JSON(http.StatusOK, srv.GetInfo())
}

func (e *SrvInfo) GetInfo() *response.APIResponse {
	if e.id <= 0 {
		return &response.APIResponse{Code: errorcode.RequestParameter, Message: errorcode.MsRequest, Data: response.DataItemNil}
	}
	e.info, _ = e.srv.FindOne(e.id)
	back := GetRespModel(e.info)
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: back}
}
