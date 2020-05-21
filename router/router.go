package router

import (
	"bot-backend/endpoints"

	"github.com/gin-gonic/gin"
)

const baseURL = "/api/v1/users"

// SetupRoutes declares routes
func SetupRoutes(router *gin.Engine) {
	router.GET("/", endpoints.RootHandler)
	router.GET(baseURL+"/searchByMobile", endpoints.SearchUserByMobile)
	router.GET(baseURL+"/searchUsersByLastName", endpoints.SearchUsersByLastName)
	router.GET(baseURL+"/getAccountStatus", endpoints.GetAccountStatus)
}
