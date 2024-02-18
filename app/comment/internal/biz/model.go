package biz

import (
	"time"

	"gorm.io/gorm"
)

type AddComment struct {
	Body string
	Slug string
}

type Article struct {
	gorm.Model
	UserId      int64 `gorm:"column:user_id"`
	Slug        string
	Title       string
	Description string
	Body        string
}

type Comment struct {
	gorm.Model
	Body      string
	ArticleId int64 `gorm:"column:article_id"`
	UserId    int64 `gorm:"column:user_id"`
}

type CommentReply struct {
	Id        int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Body      string
	Author    Author
}

type Author struct {
	Username  string
	Bio       string
	Image     string
	Following bool
}
