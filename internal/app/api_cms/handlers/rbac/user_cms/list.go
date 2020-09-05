package user_cms

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/data/rbac/role"
	"grape/internal/pkg/data/rbac/user"
	"net/http"
)

type SrvList struct {
	srv      user.Service
	list     []*user.Model
	pageSize int // 分页
	current  int // 分页
	roleID   int
}

type RespList struct {
	List       []*RespModel         `json:"list"`
	Pagination *response.Pagination `json:"pagination"`
}

func GetListGin(c *gin.Context) {
	srv := new(SrvList)
	srv.srv = user.Server
	srv.pageSize = conv.StringToInt(c.Query("size"), 20)
	srv.current = conv.StringToInt(c.Query("current"), 1)
	srv.roleID = conv.StringToInt(c.Query("role"), 0)
	c.JSON(http.StatusOK, srv.GetList())
}

func (e *SrvList) GetList() *response.APIResponse {
	if e.current <= 0 {
		e.current = 1
	}
	if e.pageSize <= 0 {
		e.pageSize = 1
	}
	back := new(RespList)
	back.Pagination = &response.Pagination{PageSize: e.pageSize, Current: e.current}
	w := " "
	if e.roleID > 0 {
		w = fmt.Sprintf(" and role_ids <> '' and JSON_CONTAINS( role_ids, '%v') ", e.roleID)
	}
	s := &response.ListParameter{WhereSt: w, OrderSt: " order by id "}
	e.list, _ = e.srv.FindMulti(back.Pagination, s)
	lsRole, _ := role.Server.FindMulti(&response.Pagination{PageSize: 100000, Current: 1}, &response.ListParameter{})
	ls := make([]*RespModel, 0)
	for _, v := range e.list {
		i := GetRespModel(v)
		e.getRole(i, lsRole, v.RoleIDs)
		ls = append(ls, i)
	}
	back.List = ls
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: back}
}

func (e *SrvList) getRole(i *RespModel, ls []*role.Model, ids string) {
	ids64 := conv.StringToInt64Array(ids)
	more := role.Server.FindMoreByList(ls, ids64)
	lsG := make([]string, 0)
	for _, v := range more {
		lsG = append(lsG, v.Name)
	}
	i.RoleNames = lsG
}
