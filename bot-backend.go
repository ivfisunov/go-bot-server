package main

import (
	"fmt"
	"os"

	"bot-backend/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var PORT string
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file.\nDefault PORT is 3333")
		PORT = "3333"
	}
	PORT = os.Getenv("PORT")

	// server setup
	server := gin.Default()
	router.SetupRoutes(server)
	server.Run(":" + PORT)
}
