package testing

import (
	"golang-base/models"
	"os"

	"github.com/gin-gonic/gin"
)

func SetUpMockDB() {
	os.Setenv("DBTYPE", "sqlite")
	os.Setenv("SQLITE", "mock.db")
	models.ConnectDatabaseMock()
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
