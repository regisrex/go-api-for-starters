package main

import (
	"github.com/gin-gonic/gin"
	"gitub.com/regisrex/golang-apis/helpers"
	"gitub.com/regisrex/golang-apis/middlewares"
	auth_service "gitub.com/regisrex/golang-apis/services/auth"
	headlines_services "gitub.com/regisrex/golang-apis/services/headlines"
	users_service "gitub.com/regisrex/golang-apis/services/users"
)

func main() {
	helpers.LoadEnv()
	helpers.ConnectDB()
	// helpers.InitializeValidator()

	app := gin.Default()

	app.Use(gin.Recovery())

	app.GET("/ping", auth_service.Ping)
	app.POST("/auth/signup", auth_service.SignUp)
	app.POST("/auth/login", auth_service.Login)

	//  users controllers
	app.Use(middlewares.ValidateJwt())
	app.GET("/users", users_service.GetAllUsers)
	app.GET("/users/:id", users_service.GetSingleUser)
	app.PUT("/users/:id", users_service.UpdateUser)
	app.DELETE("/users/:id", users_service.DeleteUser)

	// headlines controllers
	app.POST("/headlines", headlines_services.CreateHeadline)
	app.GET("/headlines", headlines_services.GetAllHeadlines)
	app.GET("/headlines/:id", headlines_services.GetSingleHeadline)
	app.PUT("/headlines/:id", headlines_services.UpdateHeadline)
	app.DELETE("/headlines/:id", headlines_services.DeleteHeadline)
	app.Run()
}
