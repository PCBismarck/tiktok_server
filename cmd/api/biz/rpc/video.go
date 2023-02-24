package rpc

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/PCBismarck/tiktok_server/cmd/api/biz/model/basic"
	"github.com/PCBismarck/tiktok_server/cmd/api/biz/model/shared"
	"github.com/PCBismarck/tiktok_server/cmd/video/kitex_gen/video"

	//"github.com/PCBismarck/tiktok_server/cmd/user/kitex_gen/user"
	"github.com/PCBismarck/tiktok_server/cmd/video/kitex_gen/video/videoservice"
	"github.com/cloudwego/kitex/client"
)

const BaseUrl = "http://101.43.172.154:8888/static/"

var videoClient videoservice.Client

func initVideo() {
	c, err := videoservice.NewClient("tiktok", client.WithHostPorts(":8010"))
	if err != nil {
		log.Fatal(err)
	}
	videoClient = c
}

// Feed
// token will not be set here
// get 30 videos from mysql and order by desc,alse return the nexttime
func Feed(ctx context.Context, req basic.FeedRequest) (resp *basic.FeedResponse, err error) {

	// 构造creq请求
	var creq video.FeedRequest
	uid, err := strconv.ParseInt(req.Token, 10, 64)

	//转化
	creq.LatestTime = req.GetLatestTime()
	creq.UserId = uid

	cresp, err := videoClient.Feed(ctx, (*video.FeedRequest)(&creq))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	resp.NextTime = cresp.NextTime
	resp.VideoList = ChangeSharedFormat(cresp.GetVideoList())
	return
}

// PublishList
// token will not be set here
// according to the uid to find all videos which created by author
func PublishList(ctx context.Context, req basic.PublishListRequest) (resp *basic.PublishListResponse, err error) {

	//构造creq请求
	var creq video.PublishListRequest
	uid, err := strconv.ParseInt(req.Token, 10, 64)

	//转化
	creq.UserId = uid
	cresp, err := videoClient.PublishList(ctx, (*video.PublishListRequest)(&creq))

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	resp.VideoList = ChangeSharedFormat(cresp.VideoList)
	return
}

func PublishAction(ctx context.Context, req video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	resp, err = videoClient.PublishAction(ctx, &req)
	return
}

func ChangeSharedFormat(input []*video.Video) (res []*shared.Video) {
	var videoList []*shared.Video
	for _, v := range input {
		newAuthor := shared.NewUser()
		newAuthor.ID = v.Author.Id
		newAuthor.IsFollow = v.Author.IsFollow
		newAuthor.FollowCount = *v.Author.FollowCount
		newAuthor.FollowerCount = *v.Author.FollowerCount
		newAuthor.Name = v.Author.Name
		newVideo := &shared.Video{
			ID:            v.VideoId,
			Author:        newAuthor,
			PlayURL:       v.PlayUrl,
			CoverURL:      v.CoverUrl,
			Title:         v.Title,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
		}
		videoList = append(videoList, newVideo)
	}
	return videoList
}
