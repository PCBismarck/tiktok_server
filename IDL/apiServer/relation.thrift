include "shared.thrift"

namespace go relation

// relation action
struct RelationActionRequest {
    1: string token (api.query="token") // 用户鉴权token
    2: i64 toUserId (api.query="to_user_id") // 对方用户id
    3: i32 actionType (api.query="action_type") // 1-关注，2-取消关注
}

struct RelationActionResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
}

// follow list
struct RelationFollowListRequest {
    1: i64 userId (api.query="user_id") // 用户id
    2: string token (api.query="token") // 用户鉴权token
}

struct RelationFollowListResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: list<shared.User> userList // 用户信息列表
}

// follower list
struct RelationFollowerListRequest {
    1: i64 userId (api.query="user_id") // 用户id
    2: string token (api.query="token") // 用户鉴权token
}

struct RelationFollowerListResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: list<shared.User> userList // 用户列表
}

// friend list
struct RelationFriendListRequest {
    1: i64 userId (api.query="user_id") // 用户id
    2: string token (api.query="token") // 用户鉴权token
}

struct RelationFriendListResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: list<shared.FriendUser> userList // 用户列表
}

// chat list
struct MessageChatRequest {
    1: string token (api.query="token") // 用户鉴权token
    2: i64 toUserId (api.query="to_user_id") // 对方用户id
}

struct MessageChatResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: list<shared.Message> messageList // 消息列表
}


// message cation
struct MessageActionRequest {
    1: string token (api.query="token") // 用户鉴权token
    2: i64 toUserId (api.query="to_user_id") // 对方用户id
    3: i32 actionType (api.query="action_type") // 1-发送消息
    4: string content (api.query="content")// 消息内容
}

struct MessageActionResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
}