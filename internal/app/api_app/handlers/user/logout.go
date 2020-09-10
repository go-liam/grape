package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/response"
	"net/http"
)

func LogoutGin(c *gin.Context) {
	json := response.APIResponse{Code: 0, Message: "OK", Data: nil}
	c.JSON(http.StatusOK, json)
}
