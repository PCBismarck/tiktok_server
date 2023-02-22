package entities

import (
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserId   int64 `gorm:"index:idx_uid_vid"`
	VideoId  int64 `gorm:"index:idx_uid_vid"`
	PlayUrl  string
	CoverUrl string
	Title    string
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
	UserId        int64
	Username      string `gorm:"index:unique"`
	PasswordMD5   string
	Salt          string
	FollowCount   int64 `gorm:"check: follow_count >= 0"`
	FollowerCount int64 `gorm:"check: follower_count >= 0"`
}
