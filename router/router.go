package router

import (
	"bot-backend/endpoints"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", endpoints.RootHandler)
}
