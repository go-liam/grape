package role_api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/data/rbac/role"
	"net/http"
)

type SrvList struct {
	srv      role.Service
	list     []*role.Model
	pageSize int // 分页
	current  int // 分页
}

type RespList struct {
	List       []*RespModel         `json:"list"`
	Pagination *response.Pagination `json:"pagination"`
}

func GetListGin(c *gin.Context) {
	srv := new(SrvList)
	srv.srv = role.Server
	srv.pageSize = conv.StringToInt(c.Query("pageSize"), 20)
	srv.current = conv.StringToInt(c.Query("current"), 1)
	c.JSON(http.StatusOK, srv.GetList())
}

func (e *SrvList) GetList() *response.APIResponse {
	back := new(RespList)
	back.Pagination = &response.Pagination{PageSize: e.pageSize, Current: e.current}
	s := &response.ListParameter{WhereSt: " and 1=1 ", OrderSt: " order by id "}
	e.list, _ = e.srv.FindMulti(back.Pagination, s)
	ls := make([]*RespModel, 0)
	for _, v := range e.list {
		ls = append(ls, GetRespModel(v))
	}
	back.List = ls
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: back}
}
