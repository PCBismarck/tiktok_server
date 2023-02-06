include "shared.thrift"

namespace go basic

// feed 
struct FeedRequest {
    1: optional i64 latestTime // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: optional string token // 可选参数，登录用户设置
}
struct FeedResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: list<shared.Video> videoList // 视频列表
    4: optional i64 nextTime // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

// register
struct UserRegisterRequest {
    1: string username // 注册用户名，最长32个字符
    2: string password // 密码，最长32个字符
}
struct UserRegisterResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: i64 userId // 用户id
    4: string token // 用户鉴权token
}

// login
struct UserLoginRequest {
    1: string username // 登录用户名
    2: string password // 登录密码
}
struct UserLoginResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: i64 userId // 用户id
    4: string token // 用户鉴权token
}

//user info
struct UserRequest {
    1: i64 userId // 用户id
    2: string token // 用户鉴权token
}
struct UserResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: shared.User user // 用户信息
}

// publish
struct PublishActionRequest {
    1: string token // 用户鉴权token
    2: binary data // 视频数据
    3: string title // 视频标题
}
struct PublishActionResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
}

// publish list
struct PublishListRequest {
    1: i64 userId // 用户id
    2: string token // 用户鉴权token
}
struct PublishListResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: list<shared.Video> videoList // 用户发布的视频列表
}