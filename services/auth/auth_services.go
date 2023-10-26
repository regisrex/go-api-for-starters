package services

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gitub.com/regisrex/golang-apis/helpers"
	"gitub.com/regisrex/golang-apis/models"
	utils_passwords "gitub.com/regisrex/golang-apis/utils/passwords"
)

type JwtPayload struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	jwt.RegisteredClaims
}

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

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtPayload{
		user.ID,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}).SignedString([]byte(os.Getenv("JWT_PRIVATE")))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	helpers.Database.Create(&user)
	c.JSON(200, gin.H{
		"message": "Sign up successfully",
		"success": true,
		"status":  200,
		"data":    map[string]interface{}{"token": token},
	})
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	c.Bind(&body)
	var user models.User
	helpers.Database.Where("email = ? ", body.Email).First(&user)

	passwordsMatch := utils_passwords.ComparePassword(body.Password, user.Password)
	if passwordsMatch != true {
		c.JSON(406, gin.H{
			"message": "Invalid crentials",
			"success": false,
			"status":  406,
		})
		return
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtPayload{
		user.ID,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}).SignedString([]byte(os.Getenv("JWT_PRIVATE")))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, gin.H{
		"message": "Logged in successfully",
		"success": true,
		"status":  200,
		"data":    map[string]interface{}{"token": token},
	})
}
