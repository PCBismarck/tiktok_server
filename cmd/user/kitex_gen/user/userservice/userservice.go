// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	user "github.com/PCBismarck/tiktok_server/cmd/user/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateUser":      kitex.NewMethodInfo(createUserHandler, newUserServiceCreateUserArgs, newUserServiceCreateUserResult, false),
		"VerifyUser":      kitex.NewMethodInfo(verifyUserHandler, newUserServiceVerifyUserArgs, newUserServiceVerifyUserResult, false),
		"UserInfo":        kitex.NewMethodInfo(userInfoHandler, newUserServiceUserInfoArgs, newUserServiceUserInfoResult, false),
		"GetIDByUsername": kitex.NewMethodInfo(getIDByUsernameHandler, newUserServiceGetIDByUsernameArgs, newUserServiceGetIDByUsernameResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceCreateUserArgs)
	realResult := result.(*user.UserServiceCreateUserResult)
	success, err := handler.(user.UserService).CreateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCreateUserArgs() interface{} {
	return user.NewUserServiceCreateUserArgs()
}

func newUserServiceCreateUserResult() interface{} {
	return user.NewUserServiceCreateUserResult()
}

func verifyUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceVerifyUserArgs)
	realResult := result.(*user.UserServiceVerifyUserResult)
	success, err := handler.(user.UserService).VerifyUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceVerifyUserArgs() interface{} {
	return user.NewUserServiceVerifyUserArgs()
}

func newUserServiceVerifyUserResult() interface{} {
	return user.NewUserServiceVerifyUserResult()
}

func userInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserInfoArgs)
	realResult := result.(*user.UserServiceUserInfoResult)
	success, err := handler.(user.UserService).UserInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserInfoArgs() interface{} {
	return user.NewUserServiceUserInfoArgs()
}

func newUserServiceUserInfoResult() interface{} {
	return user.NewUserServiceUserInfoResult()
}

func getIDByUsernameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetIDByUsernameArgs)
	realResult := result.(*user.UserServiceGetIDByUsernameResult)
	success, err := handler.(user.UserService).GetIDByUsername(ctx, realArg.Username)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newUserServiceGetIDByUsernameArgs() interface{} {
	return user.NewUserServiceGetIDByUsernameArgs()
}

func newUserServiceGetIDByUsernameResult() interface{} {
	return user.NewUserServiceGetIDByUsernameResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateUser(ctx context.Context, req *user.CreateUserRequest) (r *user.CreateUserResponse, err error) {
	var _args user.UserServiceCreateUserArgs
	_args.Req = req
	var _result user.UserServiceCreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VerifyUser(ctx context.Context, req *user.VerifyUserRequest) (r *user.VerifyUserResponse, err error) {
	var _args user.UserServiceVerifyUserArgs
	_args.Req = req
	var _result user.UserServiceVerifyUserResult
	if err = p.c.Call(ctx, "VerifyUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserInfo(ctx context.Context, req *user.UserInfoRequest) (r *user.UserInfoResponse, err error) {
	var _args user.UserServiceUserInfoArgs
	_args.Req = req
	var _result user.UserServiceUserInfoResult
	if err = p.c.Call(ctx, "UserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetIDByUsername(ctx context.Context, username string) (r int64, err error) {
	var _args user.UserServiceGetIDByUsernameArgs
	_args.Username = username
	var _result user.UserServiceGetIDByUsernameResult
	if err = p.c.Call(ctx, "GetIDByUsername", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
