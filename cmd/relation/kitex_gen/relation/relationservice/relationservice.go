// Code generated by Kitex v0.4.4. DO NOT EDIT.

package relationservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	relation "tiktok_server-new/cmd/relation/kitex_gen/relation"
)

func serviceInfo() *kitex.ServiceInfo {
	return relationServiceServiceInfo
}

var relationServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "RelationService"
	handlerType := (*relation.RelationService)(nil)
	methods := map[string]kitex.MethodInfo{
		"RelationAction": kitex.NewMethodInfo(relationActionHandler, newRelationServiceRelationActionArgs, newRelationServiceRelationActionResult, false),
		"FollowList":     kitex.NewMethodInfo(followListHandler, newRelationServiceFollowListArgs, newRelationServiceFollowListResult, false),
		"FollowerList":   kitex.NewMethodInfo(followerListHandler, newRelationServiceFollowerListArgs, newRelationServiceFollowerListResult, false),
		"FriendList":     kitex.NewMethodInfo(friendListHandler, newRelationServiceFriendListArgs, newRelationServiceFriendListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "relation",
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

func relationActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceRelationActionArgs)
	realResult := result.(*relation.RelationServiceRelationActionResult)
	success, err := handler.(relation.RelationService).RelationAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceRelationActionArgs() interface{} {
	return relation.NewRelationServiceRelationActionArgs()
}

func newRelationServiceRelationActionResult() interface{} {
	return relation.NewRelationServiceRelationActionResult()
}

func followListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceFollowListArgs)
	realResult := result.(*relation.RelationServiceFollowListResult)
	success, err := handler.(relation.RelationService).FollowList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceFollowListArgs() interface{} {
	return relation.NewRelationServiceFollowListArgs()
}

func newRelationServiceFollowListResult() interface{} {
	return relation.NewRelationServiceFollowListResult()
}

func followerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceFollowerListArgs)
	realResult := result.(*relation.RelationServiceFollowerListResult)
	success, err := handler.(relation.RelationService).FollowerList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceFollowerListArgs() interface{} {
	return relation.NewRelationServiceFollowerListArgs()
}

func newRelationServiceFollowerListResult() interface{} {
	return relation.NewRelationServiceFollowerListResult()
}

func friendListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceFriendListArgs)
	realResult := result.(*relation.RelationServiceFriendListResult)
	success, err := handler.(relation.RelationService).FriendList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceFriendListArgs() interface{} {
	return relation.NewRelationServiceFriendListArgs()
}

func newRelationServiceFriendListResult() interface{} {
	return relation.NewRelationServiceFriendListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (r *relation.RelationActionResponse, err error) {
	var _args relation.RelationServiceRelationActionArgs
	_args.Req = req
	var _result relation.RelationServiceRelationActionResult
	if err = p.c.Call(ctx, "RelationAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowList(ctx context.Context, req *relation.FollowListRequest) (r *relation.FollowListResponse, err error) {
	var _args relation.RelationServiceFollowListArgs
	_args.Req = req
	var _result relation.RelationServiceFollowListResult
	if err = p.c.Call(ctx, "FollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowerList(ctx context.Context, req *relation.FollowerListRequest) (r *relation.FollowerListResponse, err error) {
	var _args relation.RelationServiceFollowerListArgs
	_args.Req = req
	var _result relation.RelationServiceFollowerListResult
	if err = p.c.Call(ctx, "FollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FriendList(ctx context.Context, req *relation.FriendListRequest) (r *relation.FriendListResponse, err error) {
	var _args relation.RelationServiceFriendListArgs
	_args.Req = req
	var _result relation.RelationServiceFriendListResult
	if err = p.c.Call(ctx, "FriendList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
