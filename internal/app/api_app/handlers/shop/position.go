package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"net/http"
)

// 根据经纬度获取位置详情
func PositionGin(c *gin.Context) {
	st := `{"address":"北京市昌平区北清路","city":"北京市","geohash":"40.10038,116.36867","latitude":"40.10038","longitude":"116.36867","name":"昌平区北七家修正大厦(北清路北)"}`
	data := conv.StringToInterface(st)
	json := response.APIResponse{Code: 0, Message: "OK", Data: data}
	c.JSON(http.StatusOK, json)
}
