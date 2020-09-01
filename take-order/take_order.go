package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hashicorp/go-uuid"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"take-order/cache"
	"take-order/model"
	"take-order/rpc"
)

func init() {
	rpc.Setup()
	cache.SetUp()
	model.SetUp()
}

func main() {
	http.HandleFunc("/checkCount", checkCount)
	http.HandleFunc("/callBackCheck", callBackCheck)
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":31226", nil))
}

func test(w http.ResponseWriter, req *http.Request) {
}

func callBackCheck(w http.ResponseWriter, req *http.Request) {
	uniId := req.FormValue("uniId")
	goodsLog, _ := model.GetGoodsLog(uniId)
	if goodsLog.ID == "0" {
		fmt.Fprint(w, -1)
	} else {
		fmt.Fprint(w, 1)
	}
}

func checkCount(w http.ResponseWriter, req *http.Request) {
	//cpuF, _ := os.Create("cpu.out")
	//pprof.StartCPUProfile(w)
	//defer pprof.StopCPUProfile()
	//traceF, _ := os.Create("trace.out")
	//trace.Start(traceF)
	//defer trace.Stop()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	buyCount := 1
	goodsId := 1
	// 先获取分布式锁
	lockValue := cache.RedisLock()
	log.Println(lockValue)

	// 如果没锁的话，这里可能先查询出来后就被人改了
	count, err := model.GetGoodsCount(1)
	if err != nil {
		log.Fatal(err)
	}

	if count - buyCount < 0 {
		fmt.Fprint(w, "库存不足")
		return
	}
	// 货物减库存
	// 第一个方案：如果是订单需要即时生成的话，则可以先发送pre队列用snowId作为唯一键，使用mysql的xa事务保证两个库同时完成减库存和创建订单，然后再发送commit队列
	// 这里使用第二个方案：如果订单不需要即时生成，则先发送pre队列，然后扣减库存并添加扣减库存的log（或者查询binlog），最后发送commit队列，采用这个方案认为会快一点，并且更解耦

	// 如果是根据userId和time来作为对比的话，可能会在一秒内点击两次造成重复（可以从前端层面解决或者唯一索引）
	uniId, _ := uuid.GenerateUUID()
	userId := rand.Intn(99) + 1

	// 1.预发送消息
	productData := map[string]string{
		"uniId" : uniId,
		"userId": strconv.Itoa(userId),
		"goodsId": strconv.Itoa(goodsId),
		"count" : strconv.Itoa(buyCount),
	}
	msgId := rpc.PreSend(productData, strconv.Itoa(userId), uniId, "createOrder")

	// 2.执行业务逻辑
	model.BuyGoods(strconv.Itoa(goodsId), strconv.Itoa(userId), uniId)

	// 3.发到确认消息
	rpc.CommitMsg(msgId)

	//归还分布式锁
	cache.RedisUnlock("order:" + "1", lockValue)
}