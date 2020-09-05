package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/data/account/user"
	"grape/internal/pkg/data/rbac/power"
	"grape/internal/pkg/data/rbac/role"
	rbacUser "grape/internal/pkg/data/rbac/user"
	"grape/internal/pkg/middleware/router"
	"net/http"
)

type Permissions struct {
	jwtUser *router.User
	jwtFlag int
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
	srv.jwtUser, srv.jwtFlag = router.GetJWTInfoByHeader(c)
	c.JSON(http.StatusOK, srv.data())
}

func (e *Permissions) data() *response.APIResponse {
	if e.jwtUser.UserID <= 0 {
		return &response.APIResponse{Code: e.jwtFlag, Message: errorcode.MsUsGetToken, Data: response.DataItemNil}
	}
	us1, _ := rbacUser.Server.FindOne(e.jwtUser.UserID)
	us2, _ := user.Server.FindOne(e.jwtUser.UserID)
	ids1 := conv.StringToInt64Array(us1.RoleIDs)
	idsPower := role.Server.GetPowerIDsByIDs(ids1)
	lsPower, _ := power.Server.FindMultiByIDs(idsPower)
	o := new(PermissionsResp)
	o.Permissions = power.Server.GetTagByList(lsPower)
	//[]string{"rbac-menu", "rbac-user-menu","rbac-user-list","rbac-user-password"} //make([]string, 0)
	o.Flag = us1.Flag
	o.Nickname = us2.NickName
	o.Username = us2.Name
	o.ID = conv.Int64ToString(us1.ID)
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: o}
}
