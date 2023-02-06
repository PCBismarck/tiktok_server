include "shared.thrift"

namespace go basic

// feed 
struct FeedRequest {
    1: optional i64 latest_time; // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: optional string token; // 可选参数，登录用户设置
}

struct FeedResponse {
    1: required i32 status_code; // 状态码，0-成功，其他值-失败
    2: optional string status_msg; // 返回状态描述
    3: repeated shared.Video videoLis3; // 视频列表
    4: optional i64 next_time; // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

// register
struct UserRegisterRequest {
    1: required string username // 注册用户名，最长32个字符
    2: required string password // 密码，最长32个字符
}

struct UserRegisterResponse {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: required i64 user_id // 用户id
    4: required string token // 用户鉴权token
}

// login
struct UserLoginRequest {
    1: required string username // 登录用户名
    2: required string password // 登录密码
}

struct UserLoginResponse {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: required i64 user_id // 用户id
    4: required string token // 用户鉴权token
}

//user info
struct UserRequest {
    1: required i64 user_id // 用户id
    2: required string token // 用户鉴权token
}

struct UserResponse {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: required shared.User user // 用户信息
}

// publish
struct PublishActionRequest {
    1: required string // 用户鉴权token
    2: required bytes // 视频数据
    3: required string // 视频标题
}

struct PublishActionResponse {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    1: optional string status_msg // 返回状态描述
}

// publish list
struct PublishListRequest {
  1: required i64 user_id // 用户id
  2: required string token // 用户鉴权token
}

struct PublishListResponse {
  1: required i32 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: repeated shared.Video videoList // 用户发布的视频列表
}