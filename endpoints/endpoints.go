package endpoints

import (
	"bot-backend/mongodb"
	"bot-backend/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// RootHandler - dummy route
func RootHandler(ctx *gin.Context) {
	_ = mongodb.GetConnection()
	ctx.JSON(http.StatusOK, gin.H{"Hi": "there!"})
}

// SearchUserByMobile serach user by mobile
func SearchUserByMobile(ctx *gin.Context) {
	var userMobileJSON types.MobileJSONRequest
	if err := ctx.ShouldBindJSON(&userMobileJSON); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: "User not found",
			Status:  "Error",
		})
		return
	}
	user := mongodb.SearchUserByMobile(userMobileJSON)
	ctx.JSON(http.StatusOK, user)
}

// SearchUsersByLastName serach user by last name
func SearchUsersByLastName(ctx *gin.Context) {
	var userLastNameJSON types.UserLastNameJSONRequest
	if err := ctx.ShouldBindJSON(&userLastNameJSON); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: "Users not found.",
			Status:  "Error",
		})
		return
	}
	users := mongodb.SearchUserByLastName(userLastNameJSON)
	ctx.JSON(http.StatusOK, users)
}
