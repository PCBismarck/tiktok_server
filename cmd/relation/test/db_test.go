package test

import (
	"context"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/PCBismarck/tiktok_server/cmd/relation/DAO"
	"github.com/PCBismarck/tiktok_server/cmd/relation/kitex_gen/relation"
	"github.com/PCBismarck/tiktok_server/cmd/relation/kitex_gen/relation/relationservice"
	"github.com/PCBismarck/tiktok_server/cmd/relation/service"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// 测试数据库，先往里存点数据
func TestDgraph(t *testing.T) {
	ctx := context.Background()

	// 连接数据库:
	err := DAO.InitDGO()
	if err != nil {
		t.Error("init error")
	}

	// 创建模式
	//err = DAO.CreatDefaultSchema(ctx)
	//if err != nil {
	//	t.Error("CreatDefaultSchema error")
	//}
	user0 := DAO.DgraphUser{Uid: "2212", Name: "test1"}
	user1 := DAO.DgraphUser{Uid: "2323", Name: "test2"}
	user2 := DAO.DgraphUser{Uid: "2222", Name: "test3"}
	user3 := DAO.DgraphUser{Uid: "2223", Name: "test4"}
	err = DAO.UpsertUser(ctx, user0)
	if err != nil {
		t.Error("Mutate error")
	}
	err = DAO.UpsertUser(ctx, user1)
	if err != nil {
		t.Error("Mutate error")
	}
	err = DAO.UpsertUser(ctx, user2)
	if err != nil {
		t.Error("Mutate error")
	}
	err = DAO.UpsertUser(ctx, user3)
	if err != nil {
		t.Error("Mutate error")
	}
	uid1, _ := strconv.ParseInt(user1.Uid, 10, 64)
	uid2, _ := strconv.ParseInt(user2.Uid, 10, 64)
	err = service.NewRelation(ctx, uid1, uid2)
	if err != nil {
		t.Error("Follow error")
	}
}

// 测试三种功能，都是先构造请求后调用最下面的具体函数

func TestRelationAction(t *testing.T) {

	req := &relation.RelationActionRequest{UserId: 2323, ToUserId: 2212, ActionType: 2}
	_, err := RelationAction(context.Background(), req)
	if err != nil {
		t.Error(err)
		return
	}

}

func TestFollowList(t *testing.T) {
	req := &relation.FollowListRequest{UserId: 2212}
	_, err := FollowList(context.Background(), req)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestFollowerList(t *testing.T) {
	req := &relation.FollowerListRequest{UserId: 2323}
	_, err := FollowerList(context.Background(), req)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestFriendList(t *testing.T) {
	req := &relation.FriendListRequest{UserId: 2323}
	_, err := FriendList(context.Background(), req)
	if err != nil {
		t.Error(err)
		return
	}
}

// 测试用到的调用函数

func RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {

	c, err := relationservice.NewClient("tiktok_relation", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	resp, err = c.RelationAction(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	return
}

func FollowList(ctx context.Context, req *relation.FollowListRequest) (resp *relation.FollowListResponse, err error) {

	c, err := relationservice.NewClient("relationService", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	resp, err = c.FollowList(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	return
}

func FollowerList(ctx context.Context, req *relation.FollowerListRequest) (resp *relation.FollowerListResponse, err error) {

	c, err := relationservice.NewClient("relationService", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	resp, err = c.FollowerList(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	return
}

func FriendList(ctx context.Context, req *relation.FriendListRequest) (resp *relation.FriendListResponse, err error) {

	c, err := relationservice.NewClient("relationService", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	resp, err = c.FriendList(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	return
}
