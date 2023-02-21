package redis_config

import (
	"github.com/PCBismarck/tiktok_server/cmd/favorite/consts"
	"strconv"
)

func CreateFavoriteAction(videoId int64) (res int64, err error) {
	vid := videoId & (consts.HashNum - 1)
	key := "favorite:count:" + strconv.FormatInt(vid, 10)
	field := strconv.FormatInt(videoId, 10)
	res, err = Redisdb.HIncrBy(ctx, key, field, 1).Result()
	return res, err
}

func DeleteFavoriteAction(videoId int64) (res int64, err error) {
	vid := videoId & (consts.HashNum - 1)
	key := "favorite:count:" + strconv.FormatInt(vid, 10)
	field := strconv.FormatInt(videoId, 10)
	res, err = Redisdb.HIncrBy(ctx, key, field, -1).Result()
	return res, err
}
