package endpoints

import (
	"bot-backend/mongodb"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(ctx *gin.Context) {
	_ = mongodb.GetConnection()
	ctx.JSON(http.StatusOK, gin.H{"Hi": "there!"})
}
