// Code generated by hertz generator.

package api

import (
	"context"

	basic "github.com/PCBismarck/tiktok_server/cmd/api/biz/model/basic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Feed .
// @router /douyin/feed/ [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req basic.FeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(basic.FeedResponse)
	resp.StatusCode = consts.StatusOK
	msg := req.Token
	resp.StatusMsg = msg

	c.JSON(consts.StatusOK, resp)
}

// UserInfo .
// @router /douyin/user/ [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req basic.UserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(basic.UserResponse)

	c.JSON(consts.StatusOK, resp)
}

// Register .
// @router /douyin/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req basic.UserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(basic.UserRegisterResponse)

	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /douyin/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req basic.UserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(basic.UserLoginResponse)
	rpc.userClient.Login()
	c.JSON(consts.StatusOK, resp)
}

// Publish .
// @router /douyin/publish/action/ [POST]
func Publish(ctx context.Context, c *app.RequestContext) {
	var err error
	var req basic.PublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(basic.PublishActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// PublishList .
// @router /douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req basic.PublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(basic.PublishListResponse)

	c.JSON(consts.StatusOK, resp)
}
