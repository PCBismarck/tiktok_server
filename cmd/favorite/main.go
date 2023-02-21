package main

import (
	dbconfig "github.com/PCBismarck/tiktok_server/cmd/favorite/db_config"
	favorite "github.com/PCBismarck/tiktok_server/cmd/favorite/kitex_gen/favorite/favoriteservice"
	redisconfig "github.com/PCBismarck/tiktok_server/cmd/favorite/redis_config"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	dbconfig.InitDB()
	redisconfig.InitRedis()

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9010")
	if err != nil {
		panic(err)
	}
	svr := favorite.NewServer(new(FavoriteServiceImpl), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
