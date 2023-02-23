package dbconfig

import (
	entity "github.com/PCBismarck/tiktok_server/cmd/favorite/entities"
	"github.com/PCBismarck/tiktok_server/cmd/favorite/kitex_gen/favorite"
)

func GetFavoriteList(userId int64) (res []entity.Favorite, err error) {
	var favoriteData []entity.Favorite
	result := DB.Where("user_id = ?", userId).Find(&favoriteData)
	if result.Error != nil {
		return nil, result.Error
	}
	return favoriteData, nil
}

func GetVideoByVid(videoId int64) (res *entity.Video, err error) {
	var video entity.Video
	result := DB.Where("id=?", videoId).Find(&video)
	if result.Error != nil {
		return nil, result.Error
	}
	return &video, nil
}

func GetUserInfoByUid(userId int64) *favorite.UserInfo {
	a := GetUserByUid(userId)
	user := favorite.UserInfo{
		UserId:        int64(a.ID),
		UserName:      a.Username,
		FollowCount:   &a.FollowCount,
		FollowerCount: &a.FollowerCount,
		IsFollow:      true,
	}
	return &user
}

func GetUserByUid(userId int64) (user *entity.Account) {
	var a entity.Account
	DB.First(&a, userId)
	return &a
}

func CreateFavorite(uId int64, vId int64, pUrl string, cUrl string, title string) bool {
	favoriteData := entity.Favorite{
		UserId:   uId,
		VideoId:  vId,
		PlayUrl:  pUrl,
		CoverUrl: cUrl,
		Title:    title,
	}
	result := DB.Create(&favoriteData)
	return result.Error == nil
}

func DeleteFavorite(uId int64, vId int64) bool {
	result := DB.Where("user_id=? AND video_id=?", uId, vId).Delete(&entity.Favorite{})
	return result.Error == nil
}
