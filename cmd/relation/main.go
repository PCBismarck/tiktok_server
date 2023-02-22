package main

import (
	"log"
	"tiktok_server-new/cmd/relation/DAO"
	relation "tiktok_server-new/cmd/relation/kitex_gen/relation/relationservice"
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
