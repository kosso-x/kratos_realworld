package biz

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Token    string
	UserName string `gorm:"column:username"`
	Image    string
	Password string
	Bio      string
}
