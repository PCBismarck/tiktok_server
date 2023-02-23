include "../apiServer/shared.thrift"

namespace go video

enum ErrCode {
    SuccessCode                = 1000
    ServiceErrCode             = 1001
    ParamErrCode               = 1002
}

struct BaseResp {
    1: i64 status_code
    2: string status_message
}

struct Video {
    1: i64 video_id // 视频唯一标识
    2: shared.User author // 视频作者信息
    3: string play_url // 视频播放地址
    4: string cover_url; // 视频封面地址
    5: i64 favorite_count // 视频的点赞总数
    6: i64 comment_count // 视频的评论总数
    7: bool is_favorite // true-已点赞，false-未点赞
    8: string title // 视频标题
}

struct FeedRequest {
	// length of Message should be greater than or equal to 1
    1: i64 latest_time
    2: i64 user_id
}

struct FeedResponse {
    1: list<Video> video_list
    2: i64 next_time
    3: BaseResp base_resp
}

struct PublishActionRequest {
    1: i64 user_id
    2: binary data
    3: string title
}

struct PublishActionResponse {
    1: BaseResp base_resp
}

struct PublishListRequest {
    1: i64 user_id
}

struct PublishListResponse {
    1: list<Video> video_list
    2: BaseResp base_resp
}

service VideoService {
    FeedResponse Feed(1: FeedRequest req)
    PublishActionResponse PublishAction(1: PublishActionRequest req)
    PublishListResponse PublishList(1: PublishListRequest req)
}
