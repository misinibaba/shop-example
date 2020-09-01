package main

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hashicorp/go-uuid"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"take-order-fast/cache"
	localSpike2 "take-order-fast/cache/localSpike"
	remoteSpike2 "take-order-fast/cache/remoteSpike"
	"take-order-fast/producer"
	"take-order-fast/util"
)

var (
	localSpike localSpike2.LocalSpike
	remoteSpike remoteSpike2.RemoteSpikeKeys
)

func init() {
	kafkaAddr := os.Getenv("kafka_addr")
	if kafkaAddr == "" {
		kafkaAddr = "192.168.93.6"
	}
	producer.InitProducer([]string{kafkaAddr + ":9092"})

	cache.SetUp()
	localSpike = localSpike2.LocalSpike{
		LocalInStock:     30,
		LocalSalesVolume: 0,
	}
	remoteSpike = remoteSpike2.RemoteSpikeKeys{
		SpikeOrderHashKey:  "goods_nums",
		TotalInventoryKey:  "goods_total_nums",
		QuantityOfOrderKey: "goods_sold_nums",
	}
}

func main() {
	http.HandleFunc("/checkCount", checkCount)
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":31226", nil))
}

func test(w http.ResponseWriter, req *http.Request) {
}

func checkCount(w http.ResponseWriter, req *http.Request) {
	// 使用redis减库存这个策略就是无法保证原子操作（不过应该可以使用lua脚本做更多的事，比如库存扣掉同时调用api）
	// 无法原子操作的也可以先确定哪一个操作比较重要就放前面，比如先扣钱然后记录扣钱日志，但是宕机导致没发货，这时然后后台有一个脚本在扫描是否有扣钱但没发货的，然后补上
	// 如果有复制的减库存逻辑的话，还是不适合在redis里面弄，可能会有其他的联动操作，所以先看这些联动操作能不能滞后完成，不能的话还是需要使用mysql来

	// 如果先扣远程redis,然后再发送订单，则宕机的话可能造成少卖
	// 如果先扣本地缓存，然后发送订单，然后扣远程redis可能造成超卖

	// 如果是写入redis队列，则首先会多占用一定内存，然后是如果redis宕机则丢失掉队列

	// 库存处理
	if !(localSpike.LocalDeductionStock() && remoteSpike.RemoteDeductionStock()) {
		localSpike.LocalSalesVolume = localSpike.LocalInStock + 1
		util.RespJson(w, -1, "已售罄", nil)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	// 发送创建订单消息
	uniId, _ := uuid.GenerateUUID()
	userId := rand.Intn(99) + 1
	buyCount := 1
	goodsId := 1
	productData := map[string]string{
		"uniId" : uniId,
		"userId": strconv.Itoa(userId),
		"goodsId": strconv.Itoa(goodsId),
		"count" : strconv.Itoa(buyCount),
	}
	jsonData, _ := json.Marshal(productData)
	producer.Producer(string(jsonData), strconv.Itoa(userId), "createOrder", uniId)
}