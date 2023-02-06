namespace go interaction

// favorite
struct FavoriteActionRequest {
    1: string token // 用户鉴权token
    2: i64 videoId // 视频id
    3: i32 actionType // 1-点赞，2-取消点赞
}
struct FavoriteActionResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
}
// favorite list
struct FavoriteListRequest {
    1: i64 userId // 用户id
    2: string token // 用户鉴权token
}
struct FavoriteListResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: list<video> videoList // 用户点赞视频列表
}
// comment
struct CommentActionRequest {
    1: string token // 用户鉴权token
    2: i64 videoId // 视频id
    3: i32 actionType // 1-发布评论，2-删除评论
    4: optional string commentText // 用户填写的评论内容，在action_type=1的时候使用
    5: optional i64 commentId // 要删除的评论id，在action_type=2的时候使用
}
struct CommentActionResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: optional comment comment // 评论成功返回评论内容，不需要重新拉取整个列表
}
// comment list
struct CommentListRequest {
    1: string token // 用户鉴权token
    2: i64 videoId // 视频id
}
struct CommentListResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: list<comment> commentList // 评论列表
}