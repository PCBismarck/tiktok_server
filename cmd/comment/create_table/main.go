package main

import (
	dbconfig "github.com/PCBismarck/tiktok_server/cmd/comment/db_config"
	entity "github.com/PCBismarck/tiktok_server/cmd/comment/entities"
)

func main() {
	dbconfig.InitDB()
	dbconfig.DB.AutoMigrate(&entity.Video{})
}
