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
	UserId        int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	Title         string
}

type Account struct {
	gorm.Model
	Username      string `gorm:"unique"`
	PasswordMD5   string
	Salt          string
	FollowCount   int64 `gorm:"check: follow_count >= 0"`
	FollowerCount int64 `gorm:"check: follower_count >= 0"`
}
