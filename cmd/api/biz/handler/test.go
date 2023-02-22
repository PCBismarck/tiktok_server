package handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/PCBismarck/tiktok_server/cmd/api/biz/rpc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func FollowCountAddTest(ctx context.Context, c *app.RequestContext) {
	// uid := mw.JwtMiddleware.IdentityHandler(ctx, c).(shared.User).ID
	to_uid, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	action_type := c.Query("action_type")
	fmt.Println("follow test")
	// fmt.Printf("uid: %v\n", uid)
	fmt.Printf("to_uid: %v\n", to_uid)
	fmt.Printf("action_type: %v\n", action_type)
	var v int64
	if action_type == "1" {
		v = 1
	} else if action_type == "2" {
		v = -1
	} else {
		return
	}
	// rpc.FollowCountAdd(ctx, uid, v)
	ok, _ := rpc.FollowerCountAdd(ctx, to_uid, v)
	c.JSON(consts.StatusOK, map[string]interface{}{
		"status_code": !ok,
		"status_msg":  fmt.Sprintf("%d follow %d", -1, to_uid),
	})
}
