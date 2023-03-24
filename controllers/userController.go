package controllers

import (
	"fmt"
	"net/http"

	"golang-base/models"
	"golang-base/utils"

	"github.com/gin-gonic/gin"
)

// ReqNewUser is the request body for creating a new user
type ReqNewUser struct {
	Email    string `json:"email" binding:"required,email"`
	Mobile   string `json:"mobile" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ReqUpdateUser is the request body for updating a user
type ReqUpdateUser struct {
	Email    string `json:"email" binding:"required,email"`
	Mobile   string `json:"mobile" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserRoutes is the route handler for user routes
func UserRoutes(incomingRoutes *gin.Engine) {
	rootPath := fmt.Sprintf("%s/%s", utils.BasepathV1, "users")
	rootPathID := fmt.Sprintf("%s/:id", rootPath)

	incomingRoutes.GET(rootPath, GetUsers)
	incomingRoutes.GET(rootPathID, GetUser)
	incomingRoutes.POST(rootPath, PostUser)
	incomingRoutes.PUT(rootPathID, PutUser)
	incomingRoutes.DELETE(rootPathID, DeleteUser)
}

// GetUsers godoc
// @BasePath /api/v1
// @Summary get all users
// @Schemes http https
// @Description get all users
// @Tags user getusers
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User

	if err := models.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser godoc
// @BasePath /api/v1
// @Summary get a single user
// @Schemes http https
// @Description get a single user
// @Tags user getusers
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /user [get]
// @Param id path string true "ID"
func GetUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// PostUser godoc
// @BasePath /api/v1
// @Summary create user
// @Schemes http https
// @Description create a single user
// @Tags user create post
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /user [post]
// @Param ReqNewUser body controllers.ReqNewUser true "New User"
func PostUser(c *gin.Context) {
	var user ReqNewUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := models.User{Email: user.Email, Mobile: user.Mobile, Username: user.Username, Password: user.Password}

	if err := models.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": newUser.ID})
}

// PutUser godoc
// @BasePath /api/v1
// @Summary create user
// @Schemes http https
// @Description create a single user
// @Tags user create post
// @Accept json
// @Produce json
// @Success 200 {object} utils.Message "User updated successfully"
// @Router /user [put]
// @Param ReqUpdateUser body controllers.ReqUpdateUser true "Update User"
func PutUser(c *gin.Context) {
	var user ReqUpdateUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var oldUser models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&oldUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	oldUser.Email = user.Email
	oldUser.Mobile = user.Mobile
	oldUser.Username = user.Username
	oldUser.Password = user.Password

	if err := models.DB.Save(&oldUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.Message{Message: "User updated successfully"})
}

// DeleteUser godoc
// @BasePath /api/v1
// @Summary delete a single user
// @Schemes http https
// @Description delete a single user
// @Tags user delete
// @Accept json
// @Produce json
// @Success 200 {object} utils.Message "User deleted successfully"
// @Router /user [delete]
// @Param id path string true "ID"
func DeleteUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := models.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": utils.Message{Message: "User deleted successfully"}})
}
