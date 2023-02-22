package dbconfig

import (
	entity "github.com/PCBismarck/tiktok_server/cmd/comment/entities"
	"github.com/PCBismarck/tiktok_server/cmd/comment/kitex_gen/comment"
)

func CreateComment(videoId int64, userId int64, commentText string) (commentId int64, err error) {
	commentData := entity.Comment{
		UserId:      userId,
		VideoId:     videoId,
		CommentText: commentText,
	}
	result := DB.Create(&commentData)
	if result.Error != nil {
		return 0, result.Error
	}
	var video entity.Video
	DB.First(&video, videoId)
	video.CommentCount++
	DB.Save(&video)
	return int64(commentData.ID), nil
}

func DeleteComment(videoId int64, commentId int64) bool {
	result := DB.Where("id=?", commentId).Delete(&entity.Comment{})
	var video entity.Video
	//DB.First(&video, videoId)
	DB.Where("video_id=?", videoId).Find(&video)
	video.CommentCount--
	DB.Save(&video)
	return result.Error == nil
}

func GetUserInfoByUid(userId int64) *comment.UserInfo {
	a := GetUserByUid(userId)
	user := comment.UserInfo{
		UserId:        userId,
		Username:      a.Username,
		FollowCount:   &a.FollowCount,
		FollowerCount: &a.FollowerCount,
		IsFollow:      true,
	}
	return &user
}

func GetUserByUid(userId int64) (user *entity.User) {
	var a entity.User
	DB.First(&a, userId)
	return &a
}

func GetCommentsByVid(videoId int64) (res []entity.Comment, err error) {
	var commentList []entity.Comment
	result := DB.Where("video_id = ?", videoId).Find(&commentList)
	if result.Error != nil {
		return nil, result.Error
	}
	return commentList, nil
}
