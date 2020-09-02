package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/data/home_site/site"
	"grape/internal/pkg/middleware/router"
	"net/http"
)

type SrvList struct {
	srv      site.Service
	list     []*site.Model
	pageSize int // 分页
	current  int // 分页
	jwtUser  *router.User
	jwtFlag  int
}

type RespList struct {
	List       []*RespModel         `json:"list"`
	Pagination *response.Pagination `json:"pagination"`
}

type RespModel struct {
	ID          string      ` json:"id"`           // id
	CreatedAt   int64       ` json:"created_at"`   // 创建时间戳
	UpdatedAt   int64       ` json:"updated_at"`   // 更新时间戳
	Status      int8        `json:"status"  `      // 44 删除, 1 启用, 4 禁用
	LanguageID  int         `json:"languageID"  `  // 语言ID:1中文,2英语
	Title       string      `json:"title"  `       //标题
	Email       string      `json:"email"  `       //邮箱
	Description string      `json:"description"  ` //描述
	Extended    interface{} `json:"extended" `     // 扩展的
}

func GetRespModel(i *site.Model) *RespModel {
	o := new(RespModel)
	o.Title = i.Title
	o.Status = i.Status
	o.ID = conv.Int64ToString(i.ID)
	o.Email = i.Email
	o.Description = i.Description
	o.LanguageID = i.LanguageID
	o.Extended = conv.StringToInterface(i.Extended)
	return o
}

func GetListGin(c *gin.Context) {
	srv := new(SrvList)
	srv.srv = site.Server
	srv.pageSize = conv.StringToInt(c.Query("pageSize"), 20)
	srv.current = conv.StringToInt(c.Query("current"), 1)
	srv.jwtUser, srv.jwtFlag = router.GetJWTInfoByHeader(c)
	c.JSON(http.StatusOK, srv.GetList())
}

func (e *SrvList) GetList() *response.APIResponse {
	if e.jwtUser.UserID <= 0 {
		return &response.APIResponse{Code: e.jwtFlag, Message: errorcode.MsUsGetToken, Data: response.DataItemNil}
	}
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
