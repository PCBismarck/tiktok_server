package dbconfig

import (
	"crypto/md5"
	"fmt"
	"time"
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
	if result.Error != nil {
		return 0, result.Error
	}
	return int64(user.ID), nil
}

func QueryAccount(username string) (user *Account, existed bool) {
	var u Account
	result := DB.Where(
		"username = ?", username).First(&u)
	if result.Error != nil {
		return nil, false
	}
	return &u, true
}

func GetAccountInfoByUID(uid int64) (a *Account) {
	var user Account
	DB.First(&user, uid)
	return &user
}
