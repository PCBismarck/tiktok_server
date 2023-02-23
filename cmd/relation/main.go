package main

import (
	"log"

	"github.com/PCBismarck/tiktok_server/cmd/relation/DAO"
	relation "github.com/PCBismarck/tiktok_server/cmd/relation/kitex_gen/relation/relationservice"
)

func main() {

	// InitDAO得在main里初始化
	err := DAO.InitDGO()
	if err != nil {
		log.Fatalln(err)
	}
	svr := relation.NewServer(new(RelationServiceImpl)) //server.WithServiceAddr(addr),

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
