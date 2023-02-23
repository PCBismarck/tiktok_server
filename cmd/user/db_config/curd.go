package dbconfig

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func CreateAccount(username string, password string) (uid int64, err error) {
	salt := time.Now().Format("2006-01-02 15:04:05")
	user := Account{
		Username:      username,
		PasswordMD5:   fmt.Sprintf("%x", md5.Sum([]byte(password+salt))),
		Salt:          salt,
		FollowCount:   0,
		FollowerCount: 0,
	}
	result := DB.Create(&user)
	// TODO 这里需要RPC调用relation的CreateUser接口

	if result.Error != nil {
		return 0, result.Error
	}
	return int64(user.ID), nil
}

func QueryAccount(username string) (user *Account, existed bool) {
	var u Account
	// TODO 这里可以用RPC调用relation的UserInfo接口获得关注数和粉丝数

	result := DB.Where(
		"username = ?", username).First(&u)
	if result.Error != nil {
		return nil, false
	}
	return &u, true
}

func GetAccountInfoByUID(uid int64) (a *Account, err error) {
	var user Account
	res := DB.First(&user, uid)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}
	return &user, nil
}

func UpdateFollowCount(uid int64, count int64) bool {
	res := DB.Model(&Account{}).Where("id = ?", uid).Update("follow_count", count)
	return res.RowsAffected > 0
}

func UpdateFollowerCount(uid int64, count int64) bool {
	res := DB.Model(&Account{}).Where("id = ?", uid).Update("follower_count", count)
	return res.RowsAffected > 0
}

func FollowCountAdd(uid int64, toAdd int64) bool {
	res := DB.Model(&Account{}).Where("id = ?", uid).Update(
		"follow_count", gorm.Expr("follow_count + ?", toAdd))
	return res.RowsAffected > 0
}

func FollowerCountAdd(uid int64, toAdd int64) bool {
	res := DB.Model(&Account{}).Where("id = ?", uid).Update(
		"follower_count", gorm.Expr("follower_count + ?", toAdd))
	return res.RowsAffected > 0
}
