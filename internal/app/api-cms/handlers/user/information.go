package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"net/http"
)

type Information struct {
}

/*
{
  "id": 1,
  "username": "root",
  "nickname": "范闲",
  "avatar": "https://pic3.zhimg.com/v2-fad5304306ca78c380d58fb628ed1386_400x224.jpeg",
  "email": "1312342604@qq.com",
  "groups": [
    {"id": 1, "name": "root", "info": "超级用户组"}
  ]
}
*/
type InformationResp struct {
	ID       string        `json:"id"`
	Username string        `json:"username"`
	Nickname string        `json:"nickname"`
	Avatar   string        `json:"avatar"`
	Email    string        `json:"email"`
	Groups   []*GroupsItem `json:"groups"`
}

type GroupsItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func InformationGin(c *gin.Context) {
	srv := new(Information)
	c.JSON(http.StatusOK, srv.data())
}

func (e *Information) data() *response.APIResponse {
	o := new(InformationResp)
	o.Nickname = "管理员"
	o.Username = "root"
	o.ID = "1234567890123456"
	o.Groups = make([]*GroupsItem, 0)
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: o}
}
