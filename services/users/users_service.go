package users_service

import (
	"log"

	"github.com/gin-gonic/gin"
	"gitub.com/regisrex/golang-apis/helpers"
	"gitub.com/regisrex/golang-apis/models"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	helpers.Database.Select("id", "email", "username", "role").Find(&users)

	user_id, _ := c.Get("user_id")
	user_email, _ := c.Get("user_email")
	log.Print(user_id, user_email)

	var user models.User
	helpers.Database.Select("id", "email", "username", "role").First(&user)

	c.JSON(200, gin.H{
		"users": users,
		"you":   user,
	})
}
