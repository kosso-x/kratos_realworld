package biz

import "gorm.io/gorm"

type User struct {
	Id       int64
	Email    string
	UserName string `gorm:"column:username"`
	Bio      string
}

type Article struct {
	gorm.Model
	UserId      int64 `gorm:"column:user_id"`
	Slug        string
	Title       string
	Description string
	Body        string
}

type ArticleReply struct {
	gorm.Model
	Author         *User `gorm:"-"`
	Slug           string
	Title          string
	Description    string
	Body           string
	Favorited      bool     `gorm:"-"`
	FavoritesCount int      `gorm:"-"`
	TagList        []string `json:"tagList" gorm:"-"`
}

type ArticleFavorite struct {
	ArticleId int64 `gorm:"column:article_id"`
	UserId    int64 `gorm:"column:user_id"`
}

type ArticleTag struct {
	ArticleId int64
	TagId     int64
}

type Tag struct {
	Id   int64
	Name string
}

type ListReq struct {
	Tag       string
	Author    string
	Favorited string
	Limit     int64
	Offset    int64
}

type FeedReq struct {
	Limit  int64
	Offset int64
}

type CreateReq struct {
	Title       string
	Description string
	Body        string
	TagList     []string
}
