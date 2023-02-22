package main

import (
	"context"
	"github.com/PCBismarck/tiktok_server/cmd/favorite/consts"
	dbconfig "github.com/PCBismarck/tiktok_server/cmd/favorite/db_config"
	"github.com/PCBismarck/tiktok_server/cmd/favorite/kitex_gen/favorite"
	redisconfig "github.com/PCBismarck/tiktok_server/cmd/favorite/redis_config"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	userId := req.UserId
	videoId := req.VideoId
	actionType := req.ActionType

	resp = new(favorite.FavoriteActionResponse)

	switch actionType {
	case 1:
		if _, succeed := redisconfig.CreateFavoriteAction(videoId); succeed == nil {
			v, err := dbconfig.GetVideoByVid(videoId)
			if err != nil {
				resp.Base = &favorite.BaseResp{StatusCode: consts.SuccessCode, StatusMsg: "favorite action failed with GetVideoByVid"}
				_, _ = redisconfig.DeleteFavoriteAction(videoId)
				return resp, nil
			}
			ok := dbconfig.CreateFavorite(userId, videoId, v.PlayUrl, v.CoverUrl, v.Title)
			if !ok {
				resp.Base = &favorite.BaseResp{StatusCode: consts.SuccessCode, StatusMsg: "favorite action failed with CreateFavorite"}
				_, _ = redisconfig.DeleteFavoriteAction(videoId)
				return resp, nil
			}
			resp.Base = &favorite.BaseResp{
				StatusCode: consts.SuccessCode,
				StatusMsg:  "favorite action success",
			}
			return resp, nil
		}
	case 2:
		if _, succeed := redisconfig.DeleteFavoriteAction(videoId); succeed == nil {
			ok := dbconfig.DeleteFavorite(userId, videoId)
			if !ok {
				resp.Base = &favorite.BaseResp{StatusCode: consts.SuccessCode, StatusMsg: "favorite action failed with DeleteFavorite"}
				_, _ = redisconfig.CreateFavoriteAction(videoId)
				return resp, nil
			}
			resp.Base = &favorite.BaseResp{
				StatusCode: consts.SuccessCode,
				StatusMsg:  "favorite action success",
			}
			return resp, nil
		}
	}
	resp.Base = &favorite.BaseResp{
		StatusCode: consts.FailureCode,
		StatusMsg:  "favorite action failed with all",
	}
	return resp, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	// TODO: Your code here...
	userId := req.UserId

	resp = new(favorite.FavoriteListResponse)

	favoriteList, err := dbconfig.GetFavoriteList(userId)

	if err != nil {
		resp.Base = &favorite.BaseResp{
			StatusCode: consts.FailureCode,
			StatusMsg:  "Get favorite list failed",
		}
		return resp, nil
	}

	var videos []*favorite.VideoInfo
	for _, v := range favoriteList {
		vi, _ := dbconfig.GetVideoByVid(v.VideoId)
		newVideo := &favorite.VideoInfo{
			VideoId:       v.VideoId,
			User:          dbconfig.GetUserInfoByUid(vi.UserId),
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			Title:         v.Title,
			FavoriteCount: vi.FavoriteCount,
			CommentCount:  vi.CommentCount,
			IsFavorite:    true,
		}
		videos = append(videos, newVideo)
	}

	resp.Base = &favorite.BaseResp{
		StatusCode: consts.SuccessCode,
		StatusMsg:  "Get favorite list success",
	}
	resp.VideoList = videos
	return resp, nil
}
