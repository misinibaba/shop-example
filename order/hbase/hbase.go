package hbase

import (
	"context"
	"fmt"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
	"math"
	"os"
	"strconv"
	"sync"
	"time"
)
type Order struct {
	Id       string  `json:"id"`
	GoodsId  int     `json:"goodsId"`
	UserId   int     `json:"userId"`
	UniId    string  `json:"uniId"`
	Describe string  `json:"describe"`
	HBaseKey string  `json:"hBaseKey"`
	Status   uint8   `json:"status"`
	CreateAt string  `json:"createAt"`
}
var hBaseClient gohbase.Client

func InitHBase() {
	hbaseAddr := os.Getenv("hbase_addr")
	if hbaseAddr == "" {
		hbaseAddr = "192.168.93.10"
	}

	hBaseClient = gohbase.NewClient(hbaseAddr + ":2181")
}

func InputData(table, rowKey string, hValues map[string]map[string][]byte){
	var (
		//r  map[string]interface{}
		wg sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		putRequest, _ := hrpc.NewPutStr(context.Background(), table, rowKey, hValues)
		rsp, _ := hBaseClient.Put(putRequest)
		fmt.Println(rsp)
	}()
	wg.Wait()
}

func ReverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func GenerateRowKey(userId int) string {
	// rowkey:  reverse(userId)+status+(long.max - timestamp)+workerId // 这样刚好匹配查询时按照userId和status搜索的话可以在一个region里面找到,越高频的越向左;翻转userId是为了避免热区，且又能分块一块，时间戳那里是因为可以倒序
	// column family: d // 只设置一个无意义列族,越短越好
	// column qualify: goodsId, userId, uniId, describe, status, createAt // 基本跟es一样
	timeStr := strconv.FormatInt(math.MaxInt64 - time.Now().Unix(), 10)
	userIdRowKey := ReverseStr(fmt.Sprintf("%07d", userId))
	rowKey := userIdRowKey + strconv.Itoa(1) + timeStr
	return rowKey
}