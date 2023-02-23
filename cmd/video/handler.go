package main

import (
	"context"
	"time"

	userdao "github.com/PCBismarck/tiktok_server/cmd/user/db_config"

	videodao "github.com/PCBismarck/tiktok_server/cmd/video/db/dao"
	"github.com/PCBismarck/tiktok_server/cmd/user/kitex_gen/user"
	video "github.com/PCBismarck/tiktok_server/cmd/video/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	// TODO: Your code here...
	latest := req.LatestTime
	if latest == 0 {
		latest = time.Now().Unix()
	}
	//userid合法性校验
	uid := req.UserId
	resp = new(video.FeedResponse)
	if req.UserId < 0 {
		resp.BaseResp= &video.BaseResp{
			StatusCode: int64(video.ErrCode_ParamErrCode),
			StatusMessage: "Parameter is wrong with userId",
		}
		return 
	}
	// 获取视频流
	videos, nexttime, err := videodao.GetVideosForFeed(req.UserId, req.LatestTime)
	if err != nil {
		resp.BaseResp= &video.BaseResp{
			StatusCode: int64(video.ErrCode_ServiceErrCode),
			StatusMessage: "Feed Service Failed with unknown error",
		}
		return
	}
	// 封装为rpc接口格式
	videos_list := make([]*video.Video, 0)
	for _, v := range videos {
		var u user.User
		ac, err1:= userdao.GetAccountInfoByUID(uid)
		if err1 != nil {
			resp.BaseResp= &video.BaseResp{
				StatusCode: int64(video.ErrCode_ServiceErrCode),
				StatusMessage: "feed service failed because user not find",
			}
			return
		}
		u.SetId(int64(ac.ID))
		u.SetName(ac.Username)
		u.SetFollowCount(&ac.FollowCount)
		u.SetFollowerCount(&ac.FollowerCount)
		u.SetIsFollow(true)
		videos_list = append(videos_list, &video.Video{
			VideoId:       int64(v.ID),
			Author:        &u,
			Title:         v.Title,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    false,
		})
	}
	resp.NextTime = nexttime
	resp.VideoList = videos_list
	resp.BaseResp= &video.BaseResp{
		StatusCode: int64(video.ErrCode_SuccessCode),
		StatusMessage: "success with feed service!",
	}
	return 
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	// TODO: Your code here...
	//userid合法性校验
	uid := req.UserId
	resp = new(video.PublishActionResponse)
	if uid < 0 {
		resp.BaseResp= &video.BaseResp{
			StatusCode: int64(video.ErrCode_ParamErrCode),
			StatusMessage: "Parameter is wrong with userId",
		}
		return 
	}
	
	return
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...
	//userid合法性校验
	uid := req.UserId
	resp = new(video.PublishListResponse)
	if uid < 0 {
		resp.BaseResp= &video.BaseResp{
			StatusCode: int64(video.ErrCode_ParamErrCode),
			StatusMessage: "Parameter is wrong with userId",
		}
		return 
	}
	//获取用户个人视频列表
	videos , err:= videodao.FindVideoByUid(uid)
	if err != nil {
		resp.BaseResp = &video.BaseResp{
			StatusCode: int64(video.ErrCode_ServiceErrCode),
			StatusMessage: "FindVideoByUid service is failed with unknown error",
		}
		return 
	}
	//转换为rpc返回格式
	videos_list := make([]*video.Video, 0)
	for _, v := range videos {
		var u user.User
		ac, err1:= userdao.GetAccountInfoByUID(uid)
		if err1 != nil {
			resp.BaseResp= &video.BaseResp{
				StatusCode: int64(video.ErrCode_ServiceErrCode),
				StatusMessage: "FindVideoByUid service failed because user not find",
			}
			return
		}
		u.SetId(int64(ac.ID))
		u.SetName(ac.Username)
		u.SetFollowCount(&ac.FollowCount)
		u.SetFollowerCount(&ac.FollowerCount)
		u.SetIsFollow(true)
		videos_list = append(videos_list, &video.Video{
			VideoId:       int64(v.ID),
			Author:        &u,
			Title:         v.Title,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    false,
		})
	}
	resp.BaseResp = &video.BaseResp{
		StatusCode: int64(video.ErrCode_SuccessCode),
		StatusMessage: "success with publishlist service!",
	}
	resp.VideoList = videos_list
	return
}
