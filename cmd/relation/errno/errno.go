package errno

type ErrNo struct {
	ErrCode int
	ErrMsg  string
}

func NewErrNo(code int, msg string) ErrNo {
	return ErrNo{code, msg}
}

var (
	Success             = NewErrNo(200, "OK")
	ErrorActionType     = NewErrNo(201, "关注/取关操作类型错误")
	DBNewRelation       = NewErrNo(202, "关注里的数据库操作错误")
	DBDisRelation       = NewErrNo(203, "取关里的数据库操作错误")
	DBQueryFollowList   = NewErrNo(204, "查询关注列表里的数据库操作错误")
	DBQueryFollowerList = NewErrNo(204, "查询粉丝列表里的数据库操作错误")
	DBQueryFriendList   = NewErrNo(205, "查询好友列表里的数据库操作错误")
)
