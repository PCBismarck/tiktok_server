package main

import (
	"context"
	"github.com/PCBismarck/tiktok_server/cmd/relation/DAO"
	"github.com/PCBismarck/tiktok_server/cmd/relation/kitex_gen/user"
	"log"

	"github.com/PCBismarck/tiktok_server/cmd/relation/errno"
	"github.com/PCBismarck/tiktok_server/cmd/relation/kitex_gen/relation"
	"github.com/PCBismarck/tiktok_server/cmd/relation/service"
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
			return
		}
	}
	// 2-取消关注
	if req.ActionType == 2 {
		err = service.DisRelation(ctx, req.UserId, req.ToUserId)
		if err != nil {
			resp = service.CreatRelationActionResponse(errno.DBDisRelation)
			return
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
		return
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
		return
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
		return
	}
	resp = service.CreatFriendListResponse(FriendUser, errno.Success)
	return resp, err
}

// UserInfo implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) UserInfo(ctx context.Context, req *relation.UserInfoRequest) (resp *user.User, err error) {
	u, err := service.UserInfo(ctx, req.UserId)
	return u, err
}

// CreateUser implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) CreateUser(ctx context.Context, req *relation.CreateUserRequest) (resp bool, err error) {
	user0 := DAO.DgraphUser{Uid: req.UserId, Name: req.Name}
	resp = true
	err = DAO.UpsertUser(ctx, user0)
	if err != nil {
		log.Println("Mutate error")
		resp = false
	}
	return
}
