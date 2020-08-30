package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"net/http"
)

type Permissions struct {
}

/*
admin: true
avatar: null
email: null
id: 1
nickname: "root"
permissions: []
*/
type PermissionsResp struct {
	Admin       bool     `json:"admin"`
	Avatar      string   `json:"avatar"`
	Email       string   `json:"email"`
	ID          string   `json:"id"`
	Nickname    string   `json:"nickname"`
	Permissions []string `json:"permissions"`
}

func PermissionsGin(c *gin.Context) {
	srv := new(Permissions)
	c.JSON(http.StatusOK, srv.data())
}

func (e *Permissions) data() *response.APIResponse {
	o := new(PermissionsResp)
	o.Permissions = make([]string, 0)
	o.Admin = true
	o.Nickname = "Admin"
	o.ID = "1234567890123456"
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: o}
}
