package main

import (
	"fmt"
	"testing"
	"time"

	dao "github.com/PCBismarck/tiktok_server/cmd/video/db/dao"
)

func TestFeed(t *testing.T) {
	dao.InitDB()
	fmt.Print(time.Unix(1799999999,0))
	result, nexttime, err := dao.GetVideosForFeed(123321, 1799999999)
	if result != nil && err == nil {
		fmt.Print("success!!!!!!!!!!!!!!!!")
		for i := 0; i < len(result); i++{
			fmt.Println(result[i].UserId)
			fmt.Println(nexttime)
		}
	}
}
func TestCreateVideo(t *testing.T) {
	dao.InitDB()
	dao.CreateVideo(102,123321,"192.168.10.100:/video/lzq3.avi","192.168.10.100:/cover/lzq3.png",1,1,"test")
}

func TestFindVideosByUid(t *testing.T) {
	dao.InitDB()
	result, err := dao.FindVideoByUid(123321)
	if result != nil && err == nil {
		fmt.Print("success!!!!!!!!!!!!!!!!!!!!!!!!!")
		for i := 0; i < len(result); i++{
			fmt.Println(result[i].UserId)
		}
	}
}
