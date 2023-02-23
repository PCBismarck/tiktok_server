package main

import (
	dbconfig "github.com/PCBismarck/tiktok_server/cmd/favorite/db_config"
	entity "github.com/PCBismarck/tiktok_server/cmd/favorite/entities"
)

func main() {
	dbconfig.InitDB()
	//dbconfig.DB.AutoMigrate(&entity.Favorite{})
	dbconfig.DB.AutoMigrate(&entity.Video{})
}
