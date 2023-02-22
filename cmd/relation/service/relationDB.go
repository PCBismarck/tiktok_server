package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"tiktok_server-new/cmd/relation/DAO"
	"tiktok_server-new/cmd/relation/kitex_gen/user"
)

func Check() error {
	if DAO.DGO == nil {
		return errors.New("DGO not init")
	}
	return nil
}

func NewRelation(ctx context.Context, uid int64, toUid int64) error {

	if err := Check(); err != nil {
		log.Println(err.Error())
	}
	return DAO.Set(ctx, nil, []byte(fmt.Sprintf(`<%d> <follows> <%d> .`, uid, toUid)))
}

func DisRelation(ctx context.Context, uid int64, toUid int64) error {

	if err := Check(); err != nil {
		log.Println(err.Error())
	}
	return DAO.Delete(ctx, nil, []byte(fmt.Sprintf(`<%d> <follows> <%d> .`, uid, toUid)))
}

func FollowList(ctx context.Context, uid int64) ([]*user.User, error) {

	if err := Check(); err != nil {
		log.Println(err.Error())
	}
	u, err := DAO.GetFollowList(ctx, uid)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	us := make([]*user.User, len(u.Follows))
	for i, uu := range u.Follows {
		uid, _ := strconv.ParseInt(uu.Uid, 0, 64)
		us[i] = &user.User{Id: uid, Name: uu.Name, FollowCount: uu.FollowsCount, FollowerCount: uu.FollowersCount, IsFollow: true}
	}
	return us, nil
}

func FollowerList(ctx context.Context, uid int64) ([]*user.User, error) {

	if err := Check(); err != nil {
		log.Println(err.Error())
	}
	var chooseUser, _ = DAO.GetFollowerList(ctx, uid)
	followerUser := make([]*user.User, 0, len(chooseUser.Followers))

	m := make(map[string]struct{})
	for i := range chooseUser.Follows {
		m[chooseUser.Follows[i].Uid] = struct{}{}
	}
	for i, value := range chooseUser.Followers {
		uid, _ := strconv.ParseInt(value.Uid, 0, 64)
		followerUser = append(followerUser, &user.User{Id: uid, Name: value.Name, FollowCount: value.FollowsCount, FollowerCount: value.FollowersCount})
		if _, ok := m[chooseUser.Followers[i].Uid]; ok {
			followerUser[len(followerUser)-1].IsFollow = true
		}
	}
	return followerUser, nil

}
func FriendList(ctx context.Context, uid int64) ([]*user.User, error) {

	if err := Check(); err != nil {
		log.Println(err.Error())
	}
	var u, _ = DAO.GetFriendList(ctx, uid)

	us := make([]*user.User, len(u.Friends))
	for i, uu := range u.Friends {
		uid, _ := strconv.ParseInt(uu.Uid, 0, 64)
		us[i] = &user.User{Id: uid, Name: uu.Name, FollowCount: uu.FollowsCount, FollowerCount: uu.FollowersCount, IsFollow: true}
	}
	return us, nil

}
