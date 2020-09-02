package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"io/ioutil"
	"net/http"
)

// 普通用户 -> 加入 管理员

type Register struct {
	req *RegisterReq
}

type RegisterReq struct {
	//Avatar   string `json:"avatar"`
	//Email    string `json:"email"`
	ID string `json:"id"`
	//Nickname string `json:"nickname"`
	//UserName string `json:"username"`
	RoleIDs string `json:"group_ids"` // 角色IDs
	Remark  string `json:"remark"`
	Flag    int8   `json:"admin"`
}

func RegisterGin(c *gin.Context) {
	srv := new(Register)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &srv.req)
	c.JSON(http.StatusOK, srv.data())
}

func (e *Register) data() *response.APIResponse {
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: e.req}
}
