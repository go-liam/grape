package cms

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/response"
	"grape/pkg/service/syconfig/configclient"
	"net/http"
)

func Info(c *gin.Context) {
	name := c.Param("name")
	//println("name:", name)
	info, _ := configclient.Info(name, 0)
	c.JSON(http.StatusOK, response.APIResponseData(0, "OK", info))
}

func List(c *gin.Context) {
	name := c.Param("name")
	list, _ := configclient.List(name)
	c.JSON(http.StatusOK, response.APIResponseData(0, "OK", list))
}
