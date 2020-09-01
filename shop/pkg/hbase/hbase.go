package hbase

import (
	"context"
	"fmt"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
	"os"
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
	getRequest, _ := hrpc.NewGetStr(context.Background(), "1", "1")
	fmt.Println(hBaseClient.Get(getRequest))
}

func GetData(table, rowKey string, prefix string)*hrpc.Result {
	pFilter := filter.NewPrefixFilter([]byte(prefix))

	getRequest, _ := hrpc.NewGetStr(context.Background(), table, rowKey, hrpc.Filters(pFilter))
	getRsp, _ := hBaseClient.Get(getRequest)
	return getRsp
}

func ScanRange(table, startRow, endRow string) hrpc.Scanner {
	scanRequest, _ := hrpc.NewScanRangeStr(context.Background(), table, startRow, endRow)
	scanRsp := hBaseClient.Scan(scanRequest)
	return scanRsp
}

func ReverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}