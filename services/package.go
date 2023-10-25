package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitub.com/regisrex/golang-apis/helpers"
	"gitub.com/regisrex/golang-apis/models"
	utils_passwords "gitub.com/regisrex/golang-apis/utils/passwords"
)

func SignUp(c *gin.Context) {
	var body struct {
		Username        string      `validate:"required"`
		Email           string      `validate:"required,email"`
		Password        string      `validate:"required"`
		ConfirmPassword string      `validate:"required"`
		Role            models.Role `validate:"required"`
	}

	c.Bind(&body)
	validationError := helpers.Validate.Struct(body)
	if validationError != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid information given",
			"success": false,
			"status":  406,
		})
		return
	}
	if body.Password != body.ConfirmPassword {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"message": "Passwords mismatch",
			"success": false,
			"status":  406,
		})
		return
	}
	hashedPassword, _ := utils_passwords.HashPassword(body.Password)
	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: hashedPassword,
		Role:     body.Role,
	}

	usersWithSameEmail := helpers.Database.Where("email = ?", user.Email).Find(&models.User{}).RowsAffected
	if usersWithSameEmail != 0 {
		c.AbortWithStatusJSON(406, gin.H{
			"message": "Email taken",
			"success": false,
			"status":  406,
		})
		return
	}
	user.ID = uuid.New()

	helpers.Database.Create(&user)
	c.AbortWithStatusJSON(200, gin.H{
		"message": "User created successfully",
		"success": true,
		"status":  200,
		"data":    map[string]any{"id": user.ID},
	})
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}
