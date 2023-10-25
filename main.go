package main

import (
	"github.com/gin-gonic/gin"
	"gitub.com/regisrex/golang-apis/helpers"
	"gitub.com/regisrex/golang-apis/services"
)

func main() {
	helpers.LoadEnv()
	helpers.ConnectDB()
	helpers.InitializeValidator()

	app := gin.Default()

	app.GET("/ping", services.Ping)
	app.POST("/signup", services.SignUp)

	app.Run()
}
