package users_service

import (
	"log"

	"github.com/gin-gonic/gin"
	"gitub.com/regisrex/golang-apis/helpers"
	"gitub.com/regisrex/golang-apis/models"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	helpers.Database.Find(&users)

	user_id, _ := c.Get("user_id")
	user_email, _ := c.Get("user_email")
	log.Print(user_id, user_email)

	var user models.User
	helpers.Database.First(&user)

	c.JSON(200, gin.H{
		"users": users,
		"you":   user,
	})
}

func GetSingleUser(c *gin.Context) {
	user_id := c.Param("id")
	if user_id == "" {
		c.JSON(404, gin.H{
			"message": "User not found",
			"status":  404,
			"success": false,
		})
		c.Abort()
		return
	}

	var user models.User
	helpers.Database.Where("id = ?", user_id).First(&user)
	c.JSON(200, gin.H{
		"message": "User found",
		"status":  200,
		"success": true,
		"data":    user,
	})
	c.Abort()
	return

}

func UpdateUser(c *gin.Context) {

	user_id := c.Param("id")
	var updatePayload struct {
		Email    string
		Username string
	}
	c.Bind(&updatePayload)
	var user models.User
	helpers.Database.Where("id = ? ", user_id).First(&user)
	if updatePayload.Email != "" {
		user.Email = updatePayload.Email
	}
	if updatePayload.Username != "" {
		user.Username = updatePayload.Username
	}
	helpers.Database.Save(&user)

	c.JSON(200, gin.H{
		"message": "User updated successful",
		"success": true,
		"status":  200,
		"data":    user,
	})

}

func DeleteUser(c *gin.Context) {
	user_id := c.Param("id")
	if user_id == "" {
		c.JSON(404, gin.H{
			"message": "User not found",
			"status":  404,
			"success": false,
		})
		c.Abort()
		return
	}
	helpers.Database.Where(&models.User{ID: user_id}).Delete(&models.User{})
	c.JSON(200, gin.H{
		"message": "User deleted",
		"status":  200,
		"success": true,
	})
	c.Abort()
	return

}
