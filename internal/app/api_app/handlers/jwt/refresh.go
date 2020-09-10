package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/service/jwt2"
	"net/http"
	"strings"
)

type Refresh struct {
	token string
}

func RefreshGin(c *gin.Context) {
	e := new(Refresh)
	e.token = c.Request.Header.Get("Authorization")
	e.token = strings.ReplaceAll(e.token, "Bearer ", "")
	c.JSON(http.StatusOK, e.data())
}

func (e *Refresh) data() *response.APIResponse {
	o := jwt2.RefreshToken(e.token)
	if o.Refresh == "" || o.Token == "" {
		return &response.APIResponse{Code: errorcode.RefreshTokenExpired, Message: errorcode.RefreshTokenExpiredMsg, Data: o}
	}
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: o}
}
