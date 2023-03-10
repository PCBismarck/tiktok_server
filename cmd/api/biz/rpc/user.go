package rpc

import (
	"context"
	"fmt"
	"log"

	"github.com/PCBismarck/tiktok_server/cmd/api/biz/model/basic"
	"github.com/PCBismarck/tiktok_server/cmd/api/biz/model/shared"
	"github.com/PCBismarck/tiktok_server/cmd/user/db_config/consts"
	"github.com/PCBismarck/tiktok_server/cmd/user/kitex_gen/user"
	"github.com/PCBismarck/tiktok_server/cmd/user/kitex_gen/user/userservice"

	"github.com/cloudwego/kitex/client"
)

var userClient userservice.Client

func initUser() {
	c, err := userservice.NewClient("tiktok", client.WithHostPorts(":9030"))
	if err != nil {
		log.Fatal(err)
	}
	userClient = c
}

// Register
// token will not be set here
// it will create user and run login auto
func Register(ctx context.Context, req basic.UserRegisterRequest) (resp *basic.UserRegisterResponse, err error) {
	resp = new(basic.UserRegisterResponse)
	cresp, err := userClient.CreateUser(ctx, (*user.CreateUserRequest)(&req))
	resp.StatusCode = int32(cresp.Base.StatusCode)
	resp.StatusMsg = cresp.Base.StatusMsg
	if err != nil {
		return
	}
	resp.UserId = cresp.UserId
	return
}

// call Login in user service
func VerifyUser(ctx context.Context, c basic.UserLoginRequest) (resp *basic.UserLoginResponse, err error) {
	req := new(user.VerifyUserRequest)
	req.Username = c.Username
	req.Password = c.Password
	r, err := userClient.VerifyUser(ctx, req)
	resp = new(basic.UserLoginResponse)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		resp.StatusCode = consts.FailureCode
		resp.StatusMsg = "verify user fail"
		return
	}
	resp.StatusCode = int32(r.Base.StatusCode)
	resp.StatusMsg = r.Base.StatusMsg
	resp.UserId = r.UserId
	return
}

func UserInfo(ctx context.Context, c basic.UserRequest) (resp *basic.UserResponse, err error) {
	req := new(user.UserInfoRequest)
	req.UserId = c.UserId
	uresp, err := userClient.UserInfo(ctx, req)
	resp = new(basic.UserResponse)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		resp.StatusCode = consts.FailureCode
		resp.StatusMsg = "fail"
		return
	}
	resp.StatusCode = consts.SuccessCode
	resp.StatusMsg = "success"
	resp.User = &shared.User{
		ID:            uresp.User.Id,
		Name:          uresp.User.Name,
		FollowCount:   *uresp.User.FollowCount,
		FollowerCount: *uresp.User.FollowerCount,
		IsFollow:      uresp.User.IsFollow,
	}
	return
}

func GetIDByUsername(ctx context.Context, username string) (uid int64, err error) {
	return userClient.GetIDByUsername(ctx, username)
}

func UpdateFollowCountWith(ctx context.Context, uid int64, num int64) (bool, error) {
	return userClient.SetFollowCount(ctx, uid, num)
}

func UpdateFollowerCountWith(ctx context.Context, uid int64, num int64) (bool, error) {
	return userClient.SetFollowerCount(ctx, uid, num)
}

func FollowCountAdd(ctx context.Context, uid int64, toAdd int64) (bool, error) {
	return userClient.FollowCountAdd(ctx, uid, toAdd)
}

func FollowerCountAdd(ctx context.Context, uid int64, toAdd int64) (bool, error) {
	return userClient.FollowCountAdd(ctx, uid, toAdd)
}
