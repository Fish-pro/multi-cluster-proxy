package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return gin.Recovery()
}

func SetUpMiddlewares(router *gin.Engine) {
	if router == nil {
		return
	}

	router.Use(Recovery())
}
