namespace go comment

struct CommentInfo {
    1: required i64 CommentId;
    2: required UserInfo user;
    3: required string content;
    4: required string createDate;
}

struct UserInfo {
    1: required i64 UserId; // 用户id
    2: required string Username; // 用户名称
    3: optional i64 FollowCount; // 关注总数
    4: optional i64 FollowerCount; // 粉丝总数
    5: required bool IsFollow; // true-已关注，false-未关注
}

struct BaseResp {
    1: required i32 statusCode
    2: optional string statusMsg
}

struct CommentActionRequest {
    1: required i64 userId
    2: required i64 videoId
    3: required i32 actionType
    4: optional string commentText
    5: optional i64 commentId
}

struct CommentActionResponse {
    1: BaseResp base
    2: optional CommentInfo comment
}

struct CommentListRequest {
    1: i64 userId
    2: i64 videoId
}

struct CommentListResponse {
    1: BaseResp base
    2: list<CommentInfo> commentList
}

service CommentService {
    CommentActionResponse CommentAction (1: CommentActionRequest req)
    CommentListResponse CommentList (1: CommentListRequest req)
}