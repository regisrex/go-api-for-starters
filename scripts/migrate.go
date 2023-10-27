package main

import (
	"log"

	"gitub.com/regisrex/golang-apis/helpers"
	"gitub.com/regisrex/golang-apis/models"
)

func main() {
	helpers.LoadEnv()
	helpers.ConnectDB()

	helpers.Database.Exec("DROP TABLE news_headlines")
	helpers.Database.AutoMigrate(&models.User{})
	helpers.Database.AutoMigrate(&models.NewsHeadline{})

	log.Print("Migration successful")

}
