package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"golang-base/controllers"
	"golang-base/models"
)

func main() {
	app := gin.Default()
	models.ConnectDatabase()
	controllers.UserRoutes(app)

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	app.Run(":" + port)
}
