package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	Reader  Role = "READER"
	Creator      = "CREATOR"
)

type User struct {
	ID       uuid.UUID `gorm:"type:varchar(255);primary_key"`
	Email    string
	Password string
	Username string
	Role     Role
}

type NewsHeadline struct {
	gorm.Model
	Title       string
	Quote       string
	Description string
	Body        string
	UserRefer   uuid.UUID
	User        User `gorm:"references:UserRefer"`
}
