package main

import (
	"log"
	"net"

	"github.com/PCBismarck/tiktok_server/cmd/relation/DAO"
	relation "github.com/PCBismarck/tiktok_server/cmd/relation/kitex_gen/relation/relationservice"
	"github.com/cloudwego/kitex/server"
)

func main() {

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9040")
	if err != nil {
		panic(err)
	}
	// InitDAO得在main里初始化
	err = DAO.InitDGO()
	if err != nil {
		log.Fatalln(err)
	}
	svr := relation.NewServer(new(RelationServiceImpl), server.WithServiceAddr(addr)) //server.WithServiceAddr(addr),

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
