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

type LoginByPwd struct {
	req *LoginByPwdReq
}

type LoginByPwdReq struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	ClientID  int    `json:"client"`
	LoginFlag int    `json:"login_flag"`
}

func LoginByPwdGin(c *gin.Context) {
	srv := new(LoginByPwd)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &srv.req)
	c.JSON(http.StatusOK, srv.data())
}

func (e *LoginByPwd) data() *response.APIResponse {
	g1, _ := account.LoginByPassword(e.req.UserName, e.req.Password)
	if g1 <= 0 {
		return &response.APIResponse{Code: errorcode.LoginError, Message: errorcode.LoginErrorMsg, Data: response.DataItemNil}
	}
	o := jwt2.LoginToken(g1, e.req.ClientID, e.req.LoginFlag)
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: o}
}

/*
func LoginPwdGinMock(c *gin.Context) {
	st := `
{
"access_token":"access_token",
"refresh_token":"refresh_token",
 "id": "12345678901234567",
 "_id": "12345678901234567",
 "name": "aaa"
}
`
	data := conv.StringToInterface(st)
	json := response.APIResponse{Code: 0, Message: "OK", Data: data}
	c.JSON(http.StatusOK, json)
}
*/
