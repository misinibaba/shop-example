package shard

import (
	"database/sql"
	"log"
	"strconv"
)

var DbOrder1 *sql.DB
var DbOrder2 *sql.DB
var err error

func InitDb() {
	DbOrder1, err = sql.Open("mysql", "root:root@tcp(192.168.93.10:3306)/sell_order_1")
	if err != nil {
		log.Fatal(err)
	}

	DbOrder2, err = sql.Open("mysql", "root:root@tcp(192.168.93.10:3306)/sell_order_2")
	if err != nil {
		log.Fatal(err)
	}
}

func GetDbShardKey(pk int) int {
	// 决定根据userId进行划分，因为一般应用的话，会较多根据userId来进行查询，这样首先能保证一个userId在一个库里面，其次也少在映射表里面去找一次
	return pk % 2
}

func GetDbAndTable(shardKey int, pk int) (*sql.DB, string) {
	prefix := "order_"
	switch shardKey {
	case 0:
		tableName := prefix + strconv.Itoa(pk / 10000)
		return DbOrder1, tableName
		break
	case 1:
		tableName := prefix + strconv.Itoa(pk / 10000)
		return DbOrder2, tableName
		break
	}
	return nil, ""
}