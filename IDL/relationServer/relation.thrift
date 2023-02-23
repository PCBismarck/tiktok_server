namespace go relation
include "../userServer/user.thrift"


struct RelationActionRequest {
  1: string token                           // 用户鉴权token
  2: i64 user_id                            // 自己用户id
  3: i64 to_user_id                         // 对方用户id
  4: i64 action_type                     // 1-关注，2-取消关注
}

struct RelationActionResponse {
  1: i64 status_code                     // 状态码，0-成功，其他值-失败
  2: string status_msg                  // 返回状态描述
}

struct FollowListRequest {
  1: i64 user_id                             // 用户id
  2: string token                            // 用户鉴权token
}

struct FollowListResponse {
  1: i64 status_code                     // 状态码，0-成功，其他值-失败
  2: string status_msg                 // 返回状态描述
  3: list <user.User> user_list                        // 用户信息列表
}


struct FollowerListRequest{
  1: i64 user_id                             // 用户id
  2: string token                            // 用户鉴权token
}
struct UserInfoRequest{
  1: i64 user_id                             // 用户id
  2: string token                            // 用户鉴权token
}
struct FollowerListResponse {
  1: i64 status_code                     // 状态码，0-成功，其他值-失败
  2: string status_msg                  // 返回状态描述
  3: list <user.User> user_list                         // 用户列表
}

struct FriendListRequest {
  1: i64 user_id                             // 用户id
  2: string token                            // 用户鉴权token
}

struct CreateUserRequest {
  1: string user_id                             // 用户id
  2: string name                            // 用户姓名
}

struct FriendListResponse {
  1: i64 status_code                     // 状态码，0-成功，其他值-失败
  2: string status_msg                  // 返回状态描述
  3: list <user.User> user_list                         // 用户列表
}

service RelationService {
    RelationActionResponse RelationAction(1:RelationActionRequest req)  // 用户关注/取关服务
    FollowListResponse FollowList(1:FollowListRequest req)  // 获取用户关注列表
    FollowerListResponse FollowerList(1:FollowerListRequest req)  // 获取用户粉丝列表
    FriendListResponse FriendList(1:FriendListRequest req)  // 获取用户好友列表
    user.User UserInfo(1:UserInfoRequest req)  // 获取用户好友列表
    bool CreateUser(1:CreateUserRequest req) //创建用户
}
