package router

import (
	"bot-backend/endpoints"

	"github.com/gin-gonic/gin"
)

const baseUrl = "/api/v1/users"

func SetupRoutes(router *gin.Engine) {
	router.GET("/", endpoints.RootHandler)
	router.GET(baseUrl+"/searchByMobile", endpoints.SearchUserByMobile)
	router.GET(baseUrl+"/searchUsersByLastName", endpoints.SearchUsersByLastName)
}
