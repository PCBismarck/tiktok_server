package rpc

import (
	"log"

	"github.com/PCBismarck/tiktok_server/cmd/user/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/internal/client"
)

var userClient userservice.Client

func initUser() {
	c, err := userservice.NewClient("tiktok", client.WithHostPorts(":9080"))
	if err != nil {
		log.Fatal(err)
	}
	userClient = c
}
