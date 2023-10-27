package models

import (
	"gitub.com/regisrex/golang-apis/helpers"
	"gorm.io/gorm"
)

type Role string

const (
	Reader  Role = "READER"
	Creator      = "CREATOR"
)

type User struct {
	ID       string `gorm:"type:varchar(255);primary_key"`
	Email    string
	Password string
	Username string
	Role     Role
}

type NewsHeadline struct {
	ID          string `gorm:"type:varchar(255);primary_key"`
	Title       string
	Quote       string
	Description string
	Body        string
	UserRefer   string
	User        User `gorm:"foreignKey:UserRefer"`
}

// func (u *User) AfterFind(tx *gorm.DB) (err error) {
// 	u.Password = ""
// 	return
// }

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	helpers.Database.Where("UserRefer  =  ?", u.ID).Delete(&NewsHeadline{})
	return
}

func (n *NewsHeadline) AfterFind(tx *gorm.DB) (err error) {
	var headlineOwner User
	helpers.Database.Where("id = ?", n.UserRefer).First(&headlineOwner)
	headlineOwner.Password = ""
	n.User = headlineOwner
	return

}
