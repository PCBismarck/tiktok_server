package service

import (
	"github.com/PCBismarck/tiktok_server/cmd/relation/errno"
	"github.com/PCBismarck/tiktok_server/cmd/relation/kitex_gen/relation"
	"github.com/PCBismarck/tiktok_server/cmd/relation/kitex_gen/user"
)

func CreatRelationActionResponse(err errno.ErrNo) (resp *relation.RelationActionResponse) {

	return &relation.RelationActionResponse{StatusCode: int64(err.ErrCode), StatusMsg: err.ErrMsg}
}

func CreatFollowListResponse(FollowUser []*user.User, err errno.ErrNo) (resp *relation.FollowListResponse) {

	return &relation.FollowListResponse{UserList: FollowUser, StatusCode: int64(err.ErrCode), StatusMsg: err.ErrMsg}
}

func CreatFollowerListResponse(FollowerUser []*user.User, err errno.ErrNo) (resp *relation.FollowerListResponse) {

	return &relation.FollowerListResponse{UserList: FollowerUser, StatusCode: int64(err.ErrCode), StatusMsg: err.ErrMsg}
}

func CreatFriendListResponse(FriendUser []*user.User, err errno.ErrNo) (resp *relation.FriendListResponse) {

	return &relation.FriendListResponse{UserList: FriendUser, StatusCode: int64(err.ErrCode), StatusMsg: err.ErrMsg}
}
