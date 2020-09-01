package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/util/jwt2"
	"net/http"
	"strings"
)

type Refresh struct {
	token string
}

type RefreshResp struct {
	Token   string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

func RefreshGin(c *gin.Context) {
	srv := new(Refresh)
	srv.token = c.Request.Header.Get("Authorization")
	srv.token = strings.ReplaceAll(srv.token, "Bearer ", "")
	c.JSON(http.StatusOK, srv.data())
}

func (e *Refresh) data() *response.APIResponse {
	o := new(RefreshResp)
	o.Token, o.Refresh = jwt2.Refresh(e.token)
	if o.Refresh == "" || o.Token == "" {
		return &response.APIResponse{Code: errorcode.RefreshTokenExpired, Message: errorcode.RefreshTokenExpiredMsg, Data: o}
	}
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: o}
}
