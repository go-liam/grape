package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/service/account"
	"io/ioutil"
	"net/http"
)

type CmsChangePassword struct {
	req *CmsChangePasswordReq
}

type CmsChangePasswordReq struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func CmsChangePasswordGin(c *gin.Context) {
	srv := new(CmsChangePassword)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &srv.req)
	c.JSON(http.StatusOK, srv.data())
}

func (e *CmsChangePassword) data() *response.APIResponse {
	uid := conv.StringToInt64(e.req.ID, 0)
	v1, _ := account.UpdatePassword(uid, e.req.Password)
	if v1 <= 0 {
		return &response.APIResponse{Code: errorcode.Database, Message: errorcode.MsDatabase, Data: nil}
	}
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: nil}
}
