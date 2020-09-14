package health

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"net/http"
)

func Site(c *gin.Context) {
	st := `{"status":1,"list":[{"name":"mysql","status":1},{"name":"redis","status":1},{"name":"mysql","status":1}]}`
	data := conv.StringToInterface(st)
	json := response.APIResponse{Code: 0, Message: "OK", Data: data}
	c.JSON(http.StatusOK, json)
}
