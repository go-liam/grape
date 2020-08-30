package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"io/ioutil"
	"net/http"
)

type Login struct {
	req *LoginReq
}

type LoginReq struct {
	Name     string `json:"name"`
	PassWord string `json:"password"`
	ClientID int    `json:"client"`
	OutTime  int    `json:"time"` //过期时间秒  10 s - 1*60*60 (1小时)
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
	o := new(LoginResp)
	o.Token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZGVudGl0eSI6MSwic2NvcGUiOiJsaW4iLCJ0eXBlIjoiYWNjZXNzIiwiZXhw" +
		"IjoxNTk4ODAyNjgyfQ.imgpM44A_WZYTUZGm8kI-n0tgC7MREXcs717o3FNbbc"
	o.Refresh = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZGVudGl0eSI6MSwic2NvcGUiOiJsaW4iLCJ0eXBlIjoicmVmcmVzaCIsImV4cCI6" +
		"MTYwMTM5MTA4Mn0.69aH4T1B5w9cq6q6hfJt8fZZr5Gpys9HVwKCWmjMhZk"
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: o}
}
