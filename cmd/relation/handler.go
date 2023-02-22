package main

import (
	"context"
	"tiktok_server-new/cmd/relation/errno"
	relation "tiktok_server-new/cmd/relation/kitex_gen/relation"
	service "tiktok_server-new/cmd/relation/service"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	// TODO: Your code here...
	// 鉴权操作未写

	if req.ActionType < 1 || req.ActionType > 2 {
		resp = service.CreatRelationActionResponse(errno.ErrorActionType)
		return resp, nil
	}

	// 1-关注
	if req.ActionType == 1 {
		err = service.NewRelation(ctx, req.UserId, req.ToUserId)
		if err != nil {
			resp = service.CreatRelationActionResponse(errno.DBNewRelation)
		}
	}
	// 2-取消关注
	if req.ActionType == 2 {
		err = service.DisRelation(ctx, req.UserId, req.ToUserId)
		if err != nil {
			resp = service.CreatRelationActionResponse(errno.DBDisRelation)
		}
	}

	resp = service.CreatRelationActionResponse(errno.Success)
	return resp, err
}

// FollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowList(ctx context.Context, req *relation.FollowListRequest) (resp *relation.FollowListResponse, err error) {
	// TODO: Your code here...

	// 鉴权操作未写
	FollowUser, err := service.FollowList(ctx, req.UserId)
	if err != nil {
		resp = service.CreatFollowListResponse(nil, errno.DBQueryFollowList)
	}
	resp = service.CreatFollowListResponse(FollowUser, errno.Success)
	return resp, err
}

// FollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowerList(ctx context.Context, req *relation.FollowerListRequest) (resp *relation.FollowerListResponse, err error) {
	// TODO: Your code here...

	// 鉴权操作未写

	FollowerUser, err := service.FollowerList(ctx, req.UserId)
	if err != nil {
		resp = service.CreatFollowerListResponse(nil, errno.DBQueryFollowerList)
	}
	resp = service.CreatFollowerListResponse(FollowerUser, errno.Success)
	return resp, err
}

// FriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FriendList(ctx context.Context, req *relation.FriendListRequest) (resp *relation.FriendListResponse, err error) {
	// TODO: Your code here...
	// 鉴权操作未写

	FriendUser, err := service.FriendList(ctx, req.UserId)
	if err != nil {
		resp = service.CreatFriendListResponse(nil, errno.DBQueryFriendList)
	}
	resp = service.CreatFriendListResponse(FriendUser, errno.Success)
	return resp, err
}
