package health

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/svr_job/config"
	"net/http"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "1-Hello,It works.index "+config.ServerName)
}
