package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/service/jwt2"
	"io/ioutil"
	"net/http"
)

type Login struct {
	req *LoginReq
}

type LoginReq struct {
	Name      string `json:"username"`
	PassWord  string `json:"password"`
	ClientID  int    `json:"client"`
	LoginFlag int    `json:"login_flag"`
	//OutTime  int    `json:"time"` //过期时间秒  10 s - 1*60*60 (1小时)
}

type LoginResp struct {
	Token   string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

func LoginGin(c *gin.Context) {
	srv := new(Login)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &srv.req)
	c.JSON(http.StatusOK, srv.data())
}

func (e *Login) data() *response.APIResponse {
	o := jwt2.LoginToken(12345, e.req.ClientID, e.req.LoginFlag)
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: o}
}
