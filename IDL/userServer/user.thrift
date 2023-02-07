namespace go user

struct User {
    1: required i64 id; // 用户id
    2: required string name; // 用户名称
    3: optional i64 follow_count; // 关注总数
    4: optional i64 follower_count; // 粉丝总数
    5: required bool is_follow; // true-已关注，false-未关注
}

struct BaseResp {
    1: i64 statusCode
    2: string statusMsg
}

struct CreateUserRequest {
    1: string username (vt.min_size = "1", vt.pattern = "[0-9A-Za-z]+")
    2: string password (vt.min_size = "5")
}

struct CreateUserResponse {
    1: BaseResp base
    2: string userId
    3: string token
}

struct LoginRequest {
    1: string username (vt.min_size = "1", vt.pattern = "[0-9A-Za-z]+")
    2: string password (vt.min_size = "5")
}

struct LoginResponse {
    1: BaseResp base
    2: string userId
    3: string token
}

struct UserInfoRequest {
    1: string userId
    2: string token
}

struct UserInfoResponse {
    1: BaseResp base
    2: User user
}

service UserService {
    CreateUserResponse CreateUser (1: CreateUserRequest req)
    LoginResponse Login (1: LoginRequest req)
    UserInfoResponse UserInfo (1: UserInfoRequest req)
}