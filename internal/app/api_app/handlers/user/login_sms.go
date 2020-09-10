package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"net/http"
)

func LoginMsmGin(c *gin.Context) {
	st := `{
"access_token":"access_token",
"refresh_token":"refresh_token",
 "id": "12345678901234567",
 "_id": "12345678901234567",
      "phone": "13800000000"
}
`
	data := conv.StringToInterface(st)
	json := response.APIResponse{Code: 0, Message: "OK", Data: data}
	c.JSON(http.StatusOK, json)
}
