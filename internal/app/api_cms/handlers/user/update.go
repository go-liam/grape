package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"io/ioutil"
	"net/http"
)

type UpdateUser struct {
	req *UpdateUserReq
}

type UpdateUserReq struct {
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
}

func UpdateUserGin(c *gin.Context) {
	srv := new(UpdateUser)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &srv.req)
	c.JSON(http.StatusOK, srv.data())
}

func (e *UpdateUser) data() *response.APIResponse {
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: e.req}
}
