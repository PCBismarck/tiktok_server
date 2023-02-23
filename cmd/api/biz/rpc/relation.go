package rpc

import (
	"context"
	"log"

	"github.com/PCBismarck/tiktok_server/cmd/relation/kitex_gen/relation"
	"github.com/PCBismarck/tiktok_server/cmd/relation/kitex_gen/relation/relationservice"

	"github.com/cloudwego/kitex/client"
)

var relationClient relationservice.Client

func initRelation() {
	// 目前与user的端口冲突，需要进行调整

	c, err := relationservice.NewClient("tiktok", client.WithHostPorts(":9040"))
	if err != nil {
		log.Fatal(err)
	}
	relationClient = c
}

func RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	resp, err = relationClient.RelationAction(ctx, req)
	return
}

func FollowList(ctx context.Context, req *relation.FollowListRequest) (resp *relation.FollowListResponse, err error) {
	resp, err = relationClient.FollowList(ctx, req)
	return
}

func FollowerList(ctx context.Context, req *relation.FollowerListRequest) (resp *relation.FollowerListResponse, err error) {
	resp, err = relationClient.FollowerList(ctx, req)
	return
}

func FriendList(ctx context.Context, req *relation.FriendListRequest) (resp *relation.FriendListResponse, err error) {
	resp, err = relationClient.FriendList(ctx, req)
	return
}
