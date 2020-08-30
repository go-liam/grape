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
	PassWord string `json:"psw"`
	ClientID int    `json:"client"`
	OutTime  int    `json:"time"` //过期时间秒  10 s - 1*60*60 (1小时)
}

type LoginResp struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
}

func LoginGin(c *gin.Context) {
	srv := new(Login)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &srv.req)
	c.JSON(http.StatusOK, srv.data())
}

func (e *Login) data() *response.APIResponse {

	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: e.req}
}
