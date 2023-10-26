package main

import (
	"github.com/gin-gonic/gin"
	"gitub.com/regisrex/golang-apis/helpers"
	auth_service "gitub.com/regisrex/golang-apis/services/auth"
)

func main() {
	helpers.LoadEnv()
	helpers.ConnectDB()
	helpers.InitializeValidator()

	app := gin.Default()

	app.GET("/ping", auth_service.Ping)
	app.POST("/auth/signup", auth_service.SignUp)
	app.POST("/auth/login", auth_service.Login)

	app.Run()
}
