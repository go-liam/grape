package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/middleware/router"
	"net/http"
)

type Information struct {
	jwtUser *router.User
	jwtFlag int
}

type InformationResp struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"name"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func InfoGin(c *gin.Context) {
	srv := new(Information)
	srv.jwtUser, srv.jwtFlag = router.GetJWTInfoByHeader(c)
	c.JSON(http.StatusOK, srv.data())
}

func (e *Information) data() *response.APIResponse {
	if e.jwtUser.UserID <= 0 {
		return &response.APIResponse{Code: e.jwtFlag, Message: errorcode.MsUsGetToken, Data: response.DataItemNil}
	}
	o := new(InformationResp)
	o.Nickname = "管理员"
	o.Username = "root"
	o.ID = conv.Int64ToString(e.jwtUser.UserID)
	o.Phone = "13845678901"
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: o}
}

/*
func InfoGinMock(c *gin.Context) {
	st := `{
 "id": "12345678901234567",
 "_id": "12345678901234567",
"phone": "13845678901",
      "name": "aaa"
}
`
	data := conv.StringToInterface(st)
	json := response.APIResponse{Code: 0, Message: "OK", Data: data}
	c.JSON(http.StatusOK, json)
}
*/
