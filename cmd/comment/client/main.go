package main

import (
	"context"
	"github.com/PCBismarck/tiktok_server/cmd/comment/kitex_gen/comment"
	"github.com/PCBismarck/tiktok_server/cmd/comment/kitex_gen/comment/commentservice"
	"github.com/cloudwego/kitex/client"
	"log"
)

func main() {
	//第一个是服务名字，第二个是指定服务端的地址
	c, err := commentservice.NewClient("commentservice", client.WithHostPorts("127.0.0.1:9020"))
	if err != nil {
		log.Fatal(err)
	}
	//var text string
	//var id int64 = 11

	//resp, err := c.CommentAction(context.Background(), &comment.CommentActionRequest{UserId: 10,
	//	VideoId: 1, ActionType: 2, CommentText: &text, CommentId: &id})
	resp, err := c.CommentList(context.Background(), &comment.CommentListRequest{VideoId: 1})

	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(resp)
}
