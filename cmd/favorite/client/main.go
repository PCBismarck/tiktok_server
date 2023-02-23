package main

import (
	"context"
	"github.com/PCBismarck/tiktok_server/cmd/favorite/kitex_gen/favorite"
	"github.com/PCBismarck/tiktok_server/cmd/favorite/kitex_gen/favorite/favoriteservice"
	"github.com/cloudwego/kitex/client"
	"log"
)

func main() {
	//第一个是服务名字，第二个是指定服务端的地址
	c, err := favoriteservice.NewClient("favoriteservice", client.WithHostPorts("127.0.0.1:9010"))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.FavoriteAction(context.Background(), &favorite.FavoriteActionRequest{UserId: 10,
		VideoId: 1, ActionType: 2})
	//resp, err := c.FavoriteList(context.Background(), &favorite.FavoriteListRequest{UserId: 10})

	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(resp)
}
