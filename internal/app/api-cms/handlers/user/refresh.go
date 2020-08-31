package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"net/http"
)

type Refresh struct {
	//req *RefreshReq
}

type RefreshResp struct {
	Token   string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

func RefreshGin(c *gin.Context) {
	srv := new(Refresh)
	c.JSON(http.StatusOK, srv.data())
}

func (e *Refresh) data() *response.APIResponse {
	o := new(RefreshResp)
	o.Token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZGVudGl0eSI6MSwic2NvcGUiOiJsaW4iLCJ0eXBlIjoiYWNjZXNzIiwiZXhw" +
		"IjoxNTk4ODAyNjgyfQ.imgpM44A_WZYTUZGm8kI-n0tgC7MREXcs717o3FNbbc"
	o.Refresh = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZGVudGl0eSI6MSwic2NvcGUiOiJsaW4iLCJ0eXBlIjoicmVmcmVzaCIsImV4cCI6" +
		"MTYwMTM5MTA4Mn0.69aH4T1B5w9cq6q6hfJt8fZZr5Gpys9HVwKCWmjMhZk"
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: o}
}
