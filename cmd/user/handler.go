package main

import (
	"context"
	"crypto/md5"
	"fmt"

	dbconfig "github.com/PCBismarck/tiktok_server/cmd/user/db_config"
	"github.com/PCBismarck/tiktok_server/cmd/user/db_config/consts"
	user "github.com/PCBismarck/tiktok_server/cmd/user/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.CreateUserResponse)
	uid, err := dbconfig.CreateAccount(req.Username, req.Password)
	if err != nil {
		resp.Base = &user.BaseResp{
			StatusCode: consts.FailureCode,
			StatusMsg:  "Create User Fail",
		}
		return
	}
	resp.UserId = uid
	resp.Base = &user.BaseResp{
		StatusCode: consts.SuccessCode,
		StatusMsg:  "Create User Fail",
	}
	return
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserInfoResponse)
	a := dbconfig.GetAccountInfoByUID(req.UserId)
	resp.User = &user.User{
		Id:            int64(a.ID),
		Name:          a.Username,
		FollowCount:   &a.FollowCount,
		FollowerCount: &a.FollowerCount,
		IsFollow:      false,
	}
	return
}

// VerifyUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) VerifyUser(ctx context.Context, req *user.VerifyUserRequest) (resp *user.VerifyUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.VerifyUserResponse)
	a, existed := dbconfig.QueryAccount(req.Username)
	if !existed {
		resp.Base = &user.BaseResp{
			StatusCode: consts.FailureCode,
			StatusMsg:  "User does not exist",
		}
		return
	}
	if !ComparePwd(req.Password, a.Salt, a.PasswordMD5) {
		resp.Base = &user.BaseResp{
			StatusCode: consts.FailureCode,
			StatusMsg:  "Wrong Password",
		}
		return
	}
	resp.UserId = int64(a.ID)
	resp.Base = &user.BaseResp{
		StatusCode: consts.SuccessCode,
		StatusMsg:  "Success",
	}
	return
}

// GetIDByUsername implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetIDByUsername(ctx context.Context, username string) (resp int64, err error) {
	// TODO: Your code here...
	return GetID(username), nil
}

func ComparePwd(pwd string, salt string, MD5 string) bool {
	return fmt.Sprintf("%x", md5.Sum([]byte(pwd+salt))) == MD5
}

func GetID(username string) (uid int64) {
	a, existed := dbconfig.QueryAccount(username)
	if !existed {
		return 0
	}
	return int64(a.ID)
}
