package main

import (
	"log"
	"net"

	dbconfig "github.com/PCBismarck/tiktok_server/cmd/user/db_config"
	user "github.com/PCBismarck/tiktok_server/cmd/user/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/server"
)

func main() {
	dbconfig.InitDB()
	addr, err := net.ResolveTCPAddr("tcp", ":9080")
	if err != nil {
		panic(err)
	}
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
