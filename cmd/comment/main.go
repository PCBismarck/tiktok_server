package main

import (
	dbconfig "github.com/PCBismarck/tiktok_server/cmd/comment/db_config"
	comment "github.com/PCBismarck/tiktok_server/cmd/comment/kitex_gen/comment/commentservice"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	dbconfig.InitDB()

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9020")
	if err != nil {
		panic(err)
	}
	svr := comment.NewServer(new(CommentServiceImpl), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
