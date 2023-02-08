package rpc

import (
	"context"
	"fmt"
	"log"

	"github.com/PCBismarck/tiktok_server/cmd/api/biz/model/basic"
	"github.com/PCBismarck/tiktok_server/cmd/user/kitex_gen/user"
	"github.com/PCBismarck/tiktok_server/cmd/user/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)

var userClient userservice.Client

func initUser() {
	c, err := userservice.NewClient("tiktok", client.WithHostPorts(":9080"))
	if err != nil {
		log.Fatal(err)
	}
	userClient = c
}

// call Login in user service
func Login(ctx context.Context, c basic.UserLoginRequest) (resp *basic.UserLoginResponse) {
	req := new(user.LoginRequest)
	req.Username = c.Username
	req.Password = c.Password
	r, err := userClient.Login(ctx, req)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	resp = new(basic.UserLoginResponse)
	resp.StatusCode = int32(r.Base.StatusCode)
	resp.StatusMsg = r.Base.StatusMsg
	// resp.UserId = r.UserId
	resp.Token = r.Token
	return
}
