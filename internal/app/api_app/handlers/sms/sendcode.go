package sms

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"net/http"
)

func SendGin(c *gin.Context) {
	st := `
{
 "phone": "138000000",
 "type": 1
}
`
	data := conv.StringToInterface(st)
	json := response.APIResponse{Code: 0, Message: "OK", Data: data}
	c.JSON(http.StatusOK, json)
}
