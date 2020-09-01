package model

import (
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
}

// 调用微服务填http://{微服务service的name}/
// 调用外部tcp服务(mysql)： "root:root@tcp(192.168.93.10:3306)/sell_system"
var db *gorm.DB
var err error

func SetUp() {
	//db.SetMaxOpenConns(1000)
	mysqlAddr := os.Getenv("mysql_addr")
	if mysqlAddr == "" {
		mysqlAddr = "192.168.93.10"
	}
	db, err = gorm.Open("mysql", "root:root@tcp(" + mysqlAddr + ":3306)/sell_goods")
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDb() {
	db.Close()
}