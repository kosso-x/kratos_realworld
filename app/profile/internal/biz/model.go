package biz

import (
	"gorm.io/gorm"
)

type CelebUser struct {
	gorm.Model
	Email    string
	UserName string `gorm:"column:username"`
	Image    string
	Password string
	Bio      string
}

type Follows struct {
	UserId   int64 `gorm:"column:user_id"`
	FollowId int64 `gorm:"column:follow_id"`
}

type User struct {
	gorm.Model
	Email    string
	UserName string `gorm:"column:username"`
	Image    string
	Password string
	Bio      string
}
