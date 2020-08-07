package core

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter :
func SetupRouterTest() *gin.Engine {
	router := gin.Default()
	//router.Use(accessControlAllow())
	// other
	SetupRouter(router)
	return router
}
