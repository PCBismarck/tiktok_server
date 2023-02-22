package rpc

import (
	"context"
	"log"

	"github.com/PCBismarck/tiktok_server/cmd/favorite/kitex_gen/favorite"
	"github.com/PCBismarck/tiktok_server/cmd/favorite/kitex_gen/favorite/favoriteservice"
	"github.com/cloudwego/kitex/client"
)

var favoriteClient favoriteservice.Client

func initFavorite() {
	c, err := favoriteservice.NewClient("tiktok", client.WithHostPorts(":9010"))
	if err != nil {
		log.Fatal(err)
	}
	favoriteClient = c
}

func FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	resp, err = favoriteClient.FavoriteList(ctx, req)
	return
}

func FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	resp, err = favoriteClient.FavoriteAction(ctx, req)
	return
}
