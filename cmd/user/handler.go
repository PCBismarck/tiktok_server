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

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	resp = new(user.LoginResponse)
	resp.UserId = 12345678
	resp.Token = req.Username + "@" + req.Password
	return
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	return
}
