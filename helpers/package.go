package helpers

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variables")
	}
}

var Database *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error")
	} else {
		Database = db
		log.Print("Database connected")
	}
}

var Validate *validator.Validate

func InitializeValidator() {
	vdator := validator.New()
	Validate = vdator

}
