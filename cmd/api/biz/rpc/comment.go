package rpc

import (
	"context"
	"log"

	"github.com/PCBismarck/tiktok_server/cmd/comment/kitex_gen/comment"
	"github.com/PCBismarck/tiktok_server/cmd/comment/kitex_gen/comment/commentservice"
	"github.com/cloudwego/kitex/client"
)

var commentClient commentservice.Client

func initComment() {
	c, err := commentservice.NewClient("tiktok", client.WithHostPorts(":9020"))
	if err != nil {
		log.Fatal(err)
	}
	commentClient = c
}
func CommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	resp, err = commentClient.CommentAction(ctx, req)
	return
}

func CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	resp, err = commentClient.CommentList(ctx, req)
	return
}
