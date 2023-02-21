package entities

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId      int64
	VideoId     int64 `gorm:"index"`
	CommentText string
}

type Video struct {
	gorm.Model
	VideoId       int64 `gorm:"index"`
	UserId        int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	Title         string
}

type User struct {
	gorm.Model
	Username      string
	FollowCount   int64
	FollowerCount int64
}
