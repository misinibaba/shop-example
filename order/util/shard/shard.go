package shard

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"strconv"
)

var DbOrder1 *gorm.DB
var DbOrder2 *gorm.DB
var err error

type Order struct {
	ID      int64  `gorm:"primary_key" json:"id"`
	GoodsId int    `gorm:"column:goodsId" json:"goodsId"`
	UserId  int    `gorm:"column:userId" json:"userId"`
	UniId   string `gorm:"column:uniId" json:"uniId"`
	Status  int    `json:"status"`
}

func InitDb() {
	mysqlAddr := os.Getenv("mysql_addr")
	if mysqlAddr == "" {
		mysqlAddr = "192.168.93.10"
	}

	DbOrder1, err = gorm.Open("mysql", "root:root@tcp(" + mysqlAddr + ":3306)/sell_order_1")
	if err != nil {
		log.Fatal(err)
	}

	DbOrder2, err = gorm.Open("mysql", "root:root@tcp(" + mysqlAddr + ":3306)/sell_order_2")
	if err != nil {
		log.Fatal(err)
	}
}

func GetDbShardKey(pk int) int {
	// 决定根据userId进行划分，因为一般应用的话，会较多根据userId来进行查询，这样首先能保证一个userId在一个库里面，其次也少在映射表里面去找一次
	return pk % 2
}

func GetDbAndTable(shardKey int, pk int) (*gorm.DB, string) {
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

func SaveToDb(userId ,goodsId, status int, uniId string, snowId int64) (int64, error){
	shardKey := GetDbShardKey(userId)
	orderDB, shardTable := GetDbAndTable(shardKey, userId)
	fmt.Println("shard:" + strconv.Itoa(shardKey) + " user: " + strconv.Itoa(userId) + " order_" + shardTable)

	order := &Order{ID: snowId, GoodsId: goodsId, UserId: userId, UniId: uniId, Status: status}
	err := orderDB.Table(shardTable).Save(&order).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	return snowId, nil
}