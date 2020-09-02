package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"io/ioutil"
	"net/http"
)

type ChangePassword struct {
	req *ChangePasswordReq
}

type ChangePasswordReq struct {
	//ID       string `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func ChangePasswordGin(c *gin.Context) {
	srv := new(ChangePassword)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &srv.req)
	c.JSON(http.StatusOK, srv.data())
}

func (e *ChangePassword) data() *response.APIResponse {
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: e.req}
}
