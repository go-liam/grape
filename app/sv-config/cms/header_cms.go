package cms

import (
	"github.com/gin-gonic/gin"
	"grape/pkg/services/syconfig/configclient"
	"grape/pkg/util"
	"net/http"
)

func Info(c *gin.Context) {
	name := c.Param("name")
	//println("name:", name)
	info, _ := configclient.Info(name, 0)
	c.JSON(http.StatusOK, util.APIResponseData(0, "OK", info))
}

func List(c *gin.Context) {
	name := c.Param("name")
	list, _ := configclient.List(name)
	c.JSON(http.StatusOK, util.APIResponseData(0, "OK", list))
}
