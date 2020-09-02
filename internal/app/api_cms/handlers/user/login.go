package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/service/account"
	"grape/internal/pkg/service/jwt2"
	"io/ioutil"
	"net/http"
)

type Login struct {
	req *LoginReq
}

type LoginReq struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	ClientID  int    `json:"client"`
	LoginFlag int    `json:"login_flag"`
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
	g1 ,_ := account.LoginByPassword(e.req.UserName,e.req.Password)
	if g1 <=0{
		return &response.APIResponse{Code: errorcode.LoginError, Message: errorcode.LoginErrorMsg, Data: response.DataItemNil}
	}
	o := jwt2.LoginToken(g1, e.req.ClientID, e.req.LoginFlag)
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: o}
}
