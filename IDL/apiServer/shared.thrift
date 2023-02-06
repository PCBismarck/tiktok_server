namespace go shared

struct User {
    1: required i64 id; // 用户id
    2: required string name; // 用户名称
    3: optional i64 follow_count; // 关注总数
    4: optional i64 follower_count; // 粉丝总数
    5: required bool is_follow; // true-已关注，false-未关注
}

struct FriendUser {
    1: required i64 id; // 用户id
    2: required string name; // 用户名称
    3: optional i64 follow_count; // 关注总数
    4: optional i64 follower_count; // 粉丝总数
    5: required bool is_follow; // true-已关注，false-未关注
    6: optional string Message // 和该好友的最新聊天消息
    7: i64 msgType // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

struct Video {
    1: required i64 id; // 视频唯一标识
    2: required User author; // 视频作者信息
    3: required string play_url; // 视频播放地址
    4: required string cover_url; // 视频封面地址
    5: required i64 favorite_count; // 视频的点赞总数
    6: required i64 comment_count; // 视频的评论总数
    7: required bool is_favorite; // true-已点赞，false-未点赞
    8: required string title; // 视频标题
}

struct Comment {
    1: i64 id // 视频评论id
    2: User user // 评论用户信息
    3: string content // 评论内容
    4: string createDate // 评论发布日期，格式 mm-dd
}

struct Message {
    1: i64 id // 消息id
    2: i64 toUserId // 该消息接收者的id
    3: i64 fromUserId // 该消息发送者的id
    4: string content // 消息内容
    5: optional string createTime // 消息创建时间
}