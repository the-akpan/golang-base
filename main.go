package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"golang-base/controllers"
	"golang-base/docs"
	"golang-base/models"
	"golang-base/utils"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	app := gin.Default()
	models.ConnectDatabase()

	app.GET(utils.BasepathV1+"/ping", GetHealth)
	controllers.UserRoutes(app)

	docs.SwaggerInfo.BasePath = utils.BasepathV1
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	app.Run(":" + port)
}

// GetHealth godoc
// @BasePath /
// PingExample godoc
// @Summary ping example
// @Schemes http https
// @Description do ping
// @Tags ping
// @Accept json
// @Produce json
// @Success 200 {object} utils.Message
// @Router /ping [get]
func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Message{Message: "pong"})
}
