namespace go favorite

struct VideoInfo {
  required i64 VideoId; // 视频唯一标识
  required UserInfo User; // 视频作者信息
  required string PlayUrl; // 视频播放地址
  required string CoverUrl; // 视频封面地址
  required i64 FavoriteCount; // 视频的点赞总数
  required i64 CommentCount; // 视频的评论总数
  required bool IsFavorite; // true-已点赞，false-未点赞
  required string Title; // 视频标题
}

struct UserInfo {
    1: required i64 UserId; // 用户id
    2: required string UserName; // 用户名称
    3: optional i64 FollowCount; // 关注总数
    4: optional i64 FollowerCount; // 粉丝总数
    5: required bool IsFollow; // true-已关注，false-未关注
}

struct BaseResp {
    1: i32 statusCode
    2: string statusMsg
}

struct FavoriteActionRequest {
    1: required i64 userId
    2: required i64 videoId
    3: required i32 actionType
}

struct FavoriteActionResponse {
    1: BaseResp base
}

struct FavoriteListRequest {
    1: required i64 userId
}

struct FavoriteListResponse {
    1: BaseResp base
    2: list<VideoInfo> videoList
}


service FavoriteService {
    FavoriteActionResponse FavoriteAction (1: FavoriteActionRequest req)
    FavoriteListResponse FavoriteList (1: FavoriteListRequest req)
}