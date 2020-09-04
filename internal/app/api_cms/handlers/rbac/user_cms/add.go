package user_cms

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"github.com/go-liam/util/uuid"
	"grape/configs/commom"
	"grape/configs/errorcode"
	"grape/internal/pkg/data/rbac/user"
	"grape/internal/pkg/service/account"
	"io/ioutil"
	"log"
	"net/http"
)

type SrvAdd struct {
	req    *ReqCreateModel
	srv    user.Service
	result int64
}

func AddGin(c *gin.Context) {
	srv := new(SrvAdd)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &srv.req)
	srv.srv = user.Server
	srv.req.ID = uuid.UUID()
	c.JSON(http.StatusOK, srv.Add())
}

func (e *SrvAdd) Add() *response.APIResponse {
	// user
	i := &account.RegModel{
		UserID: conv.StringToInt64(e.req.ID,0),
		NickName: "匿名",
		Password: e.req.Password,
		RegionID: 1,
		UserName: e.req.Username,
		Phone:    "",
		Email:    e.req.Email,
		Remark:   "公司员工",
		Extended: "", //`{"x":1,"b":"xxx"}`
	}
	v1, v2 := account.RegByNamePassword(i)
	if v1 <=0 || v2 != nil{
		log.Printf("[ERROR] %+v\n",v2)
		return &response.APIResponse{Code: errorcode.Database, Message: errorcode.MsDatabase, Data: e.req}
	}
	// rbac user
	u1:= &ReqModel{ID:e.req.ID,RoleIDs:e.req.RoleIDs}
	item := GetModel(u1)
	item.ID = conv.StringToInt64(e.req.ID, 0) //uuid.AutoInt64ID()
	item.Status = commom.StatusDefault
	e.result, _ = e.srv.Create(item)
	if e.result > 0 {
		return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: e.req}
	}
	return &response.APIResponse{Code: errorcode.Database, Message: errorcode.MsDatabase, Data: e.req}
}
