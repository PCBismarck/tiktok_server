package dao

import (
	"time"
)

func CreateVideo(videoId int64, userId int64, play_url string, cover_url string, favor_count int64, comment_count int64, title string) (err error) {
	video := Video{
		VideoId: 		videoId,
		UserId: 		userId,
		PlayUrl: 		play_url,
		CoverUrl: 		cover_url,
		FavoriteCount: 	favor_count,
		CommentCount: 	comment_count,
		Title: 			title,
	}
	result := DB.Create(&video)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteVideoByVid(vid int64) (err error) {
	result := DB.Where("id = ?", vid).Delete(&Video{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindVideoByVid(vid int64) (video *Video ,err error) {
	var v Video
	result := DB.Where("id = ?", vid).First(&v)
	if result.Error != nil {
		return nil, result.Error
	}
	return &v, nil
}

func FindVideoByUid(vid int64) (video []Video ,err error) {
	var vlist []Video
	result := DB.Where("user_id = ?", vid).Find(&vlist)
	if result.Error != nil {
		return nil, result.Error
	}
	return vlist, nil
}

func GetVideosForFeed(uid int64, latest int64) (video []Video, nexttime int64 ,err error) {
	time := time.Unix(latest, 0)
	var videoList []Video
	result := DB.Where("created_at < ?", time).Order("created_at desc").Limit(30).Find(&videoList)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return videoList, videoList[len(videoList) - 1].CreatedAt.Unix() ,nil
}

