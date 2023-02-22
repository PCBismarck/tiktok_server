include "shared.thrift"

namespace go interaction

// favorite
struct FavoriteActionRequest {
    1: string token (api.query="token")// 用户鉴权token
    2: i64 videoId (api.query="video_id")// 视频id
    3: i32 actionType (api.query="action_type")// 1-点赞，2-取消点赞
}
struct FavoriteActionResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: string statusMsg // 返回状态描述
}
// favorite list
struct FavoriteListRequest {
    1: i64 userId (api.query="user_id")// 用户id
    2: string token (api.query="token")// 用户鉴权token
}
struct FavoriteListResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: string statusMsg // 返回状态描述
    3: list<shared.Video> videoList // 用户点赞视频列表
}
// comment
struct CommentActionRequest {
    1: string token (api.query="token")// 用户鉴权token
    2: i64 videoId (api.query="video_id")// 视频id
    3: i32 actionType (api.query="action_type")// 1-发布评论，2-删除评论
    4: string commentText (api.query="comment_text")// 用户填写的评论内容，在action_type=1的时候使用
    5: i64 commentId (api.query="comment_id")// 要删除的评论id，在action_type=2的时候使用
}
struct CommentActionResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: string statusMsg // 返回状态描述
    3: shared.Comment comment // 评论成功返回评论内容，不需要重新拉取整个列表
}
// comment list
struct CommentListRequest {
    1: string token (api.query="token")// 用户鉴权token
    2: i64 videoId (api.query="video_id")// 视频id
}
struct CommentListResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: string statusMsg // 返回状态描述
    3: list<shared.Comment> commentList // 评论列表
}