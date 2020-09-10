package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"net/http"
)

func InfoGin(c *gin.Context) {
	st := `{
 "id": "12345678901234567",
 "_id": "12345678901234567",
"phone": "13845678901",
      "name": "aaa"
}
`
	data := conv.StringToInterface(st)
	json := response.APIResponse{Code: 0, Message: "OK", Data: data}
	c.JSON(http.StatusOK, json)
}
