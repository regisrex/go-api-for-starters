package main

import (
	"github.com/gin-gonic/gin"
	"gitub.com/regisrex/golang-apis/helpers"
	"gitub.com/regisrex/golang-apis/middlewares"
	auth_service "gitub.com/regisrex/golang-apis/services/auth"
	users_service "gitub.com/regisrex/golang-apis/services/users"
)

func main() {
	helpers.LoadEnv()
	helpers.ConnectDB()
	helpers.InitializeValidator()

	app := gin.Default()

	app.Use(gin.Recovery())

	app.GET("/ping", auth_service.Ping)
	app.POST("/auth/signup", auth_service.SignUp)
	app.POST("/auth/login", auth_service.Login)

	// app.Use(middlewares.ValidateJwt())
	app.GET("/users", middlewares.ValidateJwt(), users_service.GetAllUsers)
	app.Run()
}
