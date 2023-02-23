package main

import (
	dbconfig "github.com/PCBismarck/tiktok_server/cmd/video/db/dao"
)

// config your DB in ./tookit/init_db
// and run this file to create the tables used in this program
func main() {
	dbconfig.InitDB()
	dbconfig.DB.AutoMigrate(&dbconfig.Video{})
}