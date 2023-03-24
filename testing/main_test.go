package testing

import (
	"golang-base/models"
	"os"

	"github.com/gin-gonic/gin"
)

func SetUpMockDB() {
	os.Setenv("DBTYPE", "sqlite")
	os.Setenv("Sqlite", "mock.db")
	models.ConnectDatabaseMock()
}

func RemoveMockDB() {
	os.Remove("mock.db")
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
