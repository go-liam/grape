package site

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"grape/configs/errorcode"
	"grape/internal/pkg/data/home_site/site"
	"net/http"
)

type SrvSiteInfo struct {
	id int64
	//req *ReqModel
	srv *site.SrvSite
}

func GetInfoGin(c *gin.Context) {
	srv := new(SrvSiteInfo)
	//body, _ := ioutil.ReadAll(c.Request.Body)
	//json.Unmarshal(body, &srv.req)
	srv.id = conv.StringToInt64(c.Param("id"), 0)
	srv.srv = site.Server
	c.JSON(http.StatusOK, srv.GetInfoSite())
}

func (e *SrvSiteInfo) GetInfoSite() *response.APIResponse {
	if e.id <= 0 {
		return &response.APIResponse{Code: errorcode.RequestParameter, Message: errorcode.MsRequest, Data: response.DataItemNil}
	}
	got, err := e.srv.FindOne(e.id)
	back := GetRespModel(got)
	if err != nil {
		return &response.APIResponse{Code: errorcode.Database, Message: errorcode.MsDatabase, Data: response.DataItemNil}
	}
	return &response.APIResponse{Code: errorcode.Success, Message: errorcode.MsSuccess, Data: back}
}
