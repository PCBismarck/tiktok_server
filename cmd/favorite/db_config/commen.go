package dbconfig

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlConn struct {
	user      string
	pwd       string
	protocol  string
	address   string
	port      string
	dbname    string
	charset   string
	parseTime string
}

func (c *mysqlConn) GetDSN() string {
	return fmt.Sprintf("%v:%v@%v(%v:%v)/%v?charset=%v&parseTime=%v",
		c.user, c.pwd, c.protocol, c.address, c.port, c.dbname, c.charset, c.parseTime)
}

var DB *gorm.DB

func InitDB() {
	dsn := (&mysqlConn{
		user:      "root",
		pwd:       "douyinDY_123",
		protocol:  "tcp",
		address:   "101.43.172.154",
		dbname:    "douyin",
		port:      "3306",
		charset:   "utf8mb4",
		parseTime: "True",
	}).GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}
