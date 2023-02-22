package main

import (
	dbconfig "github.com/PCBismarck/tiktok_server/cmd/user/db_config"
)

// config your DB in ./tookit/init_db
// and run this file to create the tables used in this program
func main() {
	dbconfig.InitDB()
	dbconfig.DB.AutoMigrate(&dbconfig.Account{})
}
