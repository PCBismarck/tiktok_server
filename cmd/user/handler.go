package main

import (
	"context"

	user "github.com/PCBismarck/tiktok_server/cmd/user/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

var usersLoginInfo = map[string]user.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   nil,
		FollowerCount: nil,
		IsFollow:      true,
	},
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	// TODO: Your code here...
	return
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// VerifyUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) VerifyUser(ctx context.Context, req *user.VerifyUserRequest) (resp *user.VerifyUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.VerifyUserResponse)
	resp.UserId = 12345678
	// resp.Token = req.Username + "@" + req.Password
	return
}
