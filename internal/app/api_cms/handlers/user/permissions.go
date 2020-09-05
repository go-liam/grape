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
  "permissions": [
    {"id": 2, "name": "查询自己拥有的权限", "module": "用户"},
    {"id": 6, "name": "查询日志记录的用户", "module": "日志"}
  ]
*/
type PermissionsResp struct {
	Flag        int8     `json:"admin"`
	Avatar      string   `json:"avatar"`
	Email       string   `json:"email"`
	ID          string   `json:"id"`
	Nickname    string   `json:"nickname"`
	Username    string   `json:"username"`
	Permissions []string `json:"permissions"` //接管权限
}

func PermissionsGin(c *gin.Context) {
	srv := new(Permissions)
	c.JSON(http.StatusOK, srv.data())
}

func (e *Permissions) data() *response.APIResponse {
	o := new(PermissionsResp)
	o.Permissions = []string{"rbac-menu", "rbac-user-menu","rbac-user-list","rbac-user-password"} //make([]string, 0)
	o.Flag = 0
	o.Nickname = "管理员"
	o.Username = "root"
	o.ID = "1234567890123456"
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: o}
}
