include "shared.thrift"
include "basic.thrift"
include "interaction.thrift"
include "relation.thrift"

namespace go api

service Basic{
    // basic apis
    basic.FeedResponse Feed(1: basic.FeedRequest req) (api.get="/douyin/feed/")
    basic.UserResponse UserInfo(1: basic.UserRequest req) (api.get="/douyin/user/")
    basic.UserRegisterResponse Register(1: basic.UserRegisterRequest req) (api.post="/douyin/user/register/")
    basic.UserLoginResponse Login(1: basic.UserLoginRequest req) (api.post="/douyin/user/login/")
    basic.PublishActionResponse Publish(1: basic.PublishActionRequest req) (api.post="/douyin/publish/action/")
    basic.PublishListResponse PublishList(1: basic.PublishListRequest req) (api.get="/douyin/publish/list/")
}

service Interaction{
    // extra apis - I
    interaction.FavoriteActionResponse FavrotieAction(1: interaction.FavoriteActionRequest req) (api.post="/douyin/favorite/action/")
    interaction.FavoriteListResponse FavoriteList(1: interaction.FavoriteListRequest req) (api.get="/douyin/favorite/list/")
    interaction.CommentActionResponse CommentAction(1: interaction.CommentActionRequest req) (api.post="/douyin/comment/action/")
    interaction.CommentListResponse CommentList(1: interaction.CommentListRequest req) (api.get="/douyin/comment/list/")
}

service Relation{
    // extra apis - II
    relation.RelationActionResponse RelationAction(1: relation.RelationActionRequest req) (api.post="/douyin/relation/action/") 
    relation.RelationFollowerListResponse FollowList(1: relation.RelationFollowListRequest req) (api.get="/douyin/follow/list/")
    relation.RelationFollowerListResponse FollowerList(1: relation.RelationFollowerListRequest req) (api.get="/douyin/follower/list/")
    relation.RelationFriendListResponse FriendList(1: relation.RelationFriendListRequest req) (api.get="/douyin/friend/list/")
    relation.MessageChatResponse MessageChatList(1: relation.MessageChatRequest req) (api.get="/douyin/message/chat/")
    relation.MessageActionResponse MessageAction(1: relation.MessageActionRequest req) (api.post="/douyin/message/action/")    
}