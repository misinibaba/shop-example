package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	_ "net/http/pprof"
	"order/cache"
	"order/es"
	"order/hbase"
	"order/util/shard"
	"order/util/snowflake"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type User struct {
	Id int `json:"id"`
	Name int `json:"name"`
	Score int `json:"score"`
	MqOffset int `json:"mqOffset"`
	Status int `json:"status"`
	CreateAt string `json:"createAt"`
}

type Consumer struct {
	ready chan bool
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// 不要将下面这些代码放进goroutine， 因为这个已经在goroutine中
	// 注意需要幂等，有时consumer消费性能低，老是不ack，或者ack网络波动没发送成功，mq那边到了重试时间进行重新发送就会重复发送
	// 同一个partition内的消息只能被同一个group中的一个consumer消费，要是partitions数量大于consumer,则某个consumer需要消费两个partition，则可能会出现offset重复

	// 如果是分库分表的话，唯一索性来实现幂等应该会有bug
	// 使用redis的话，就怕redis挂掉，所以就最大保险起见就是保证redis的高可用，以及数据库添加唯一索引

	// 就阻塞在下面这里来一条读一条
	for message := range claim.Messages() {
		log.Printf("Message claimed: key = %s, value = %v, topic = %s, partition = %v, offset = %v, header = %v", string(message.Key), string(message.Value), message.Topic, message.Partition, message.Offset, message.Headers)
		//mqOffset := message.Topic + "_" + strconv.Itoa(int(message.Partition)) + "_" + strconv.FormatInt(message.Offset, 10)

		// 获取订单流程的唯一id，如果重复说明是上游出现了重发，则需要抛出
		uniId := ""
		for _, v := range message.Headers {
			if string(v.Key) == "uniId" {
				uniId = string(v.Value)
			}
		}

		// 存在就返回
		if cache.IsExist(uniId) != 0 {
			session.MarkMessage(message, "")
		}

		// 查看消息
		msgValue := new(map[string]string)
		json.Unmarshal(message.Value, &msgValue)

		//status := rand.Intn(3) + 1
		status := 1
		// workerId应该是每个机器一个Id，从配置文件引入
		worker, _ := snowflake.NewWorker(1)
		snowId := worker.GetId()

		fmt.Println("-----------------mysql-------------")
		userId, _ := strconv.Atoi((*msgValue)["userId"])
		goodsId, _ := strconv.Atoi((*msgValue)["goodsId"])
		shard.SaveToDb(userId, goodsId, status, uniId, snowId)
		fmt.Println("--------------------------------")

		letterBytes := "abcdefghijklmnoqprstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		describe := make([]byte, 10)
		space := []byte(" ")
		for i := range describe {
			if i % 2 == 0 {
				describe[i] = space[0]
				continue
			}
			describe[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
		}

		// ----------------下面存到nosql需要改成监听binlog添加)--------------------
		//-------------添加到hbase
		fmt.Println("-----------------hbase-------------")
		hValues := map[string]map[string][]byte{
			"d": {
				"goodsId": []byte(strconv.Itoa(goodsId)),
				"userId": []byte(strconv.Itoa(userId)),
				"uniId": []byte(uniId),
				"describe": describe,
				"status": []byte(strconv.Itoa(status)),
			},
		}

		rowKey := hbase.GenerateRowKey(userId)
		hbase.InputData("order", rowKey, hValues)
		fmt.Println("-----------------------------------")


		//------------------添加到es,存hbase的key,当作全文索引用
		fmt.Println("-----------------es-------------")
		scoreLog := es.Order{
			Id: strconv.FormatInt(snowId, 10),
			GoodsId: goodsId,
			UserId: userId,
			UniId: uniId,
			Describe: string(describe),
			Status: uint8(status),
			HBaseKey: rowKey,
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		es.InputData("order", scoreLog, 0)
		fmt.Println("-----------------------------------")

		// 订单重复ID只存在60秒，主要是为了防止临时的接口宕机导致的重发，不存在太久因为需要释放内存
		// 不过为了防止还有那种隔了很久才重试的，此时缓存已经没有了，所以需要数据库查重
		// 目前失败重试的情况有：
		// 1.mq_service发送了mq，但没有设置状态，导致job那边每隔5秒扫描一次发现是presend状态，然后去前一个服务询问，如果前一个服务也全部宕机，那么就会一直卡住(因为需要确认是哪一步执行失败)，一直停到恢复为止，然后发过来就超时了
		// 2.(1)里面的去前一个服务询问后，然后返回已经处理，那么则可能是mq_service发送后没改状态，则会造成重发，这种情况一般就取决于job的轮询时间
		conn := cache.RedisConn.Get()
		conn.Do("set", uniId, 1, "EX", 60)
		defer conn.Close()

		// 对消息进行确认ack,然后mq那边就会确认offset，否则会重发
		session.MarkMessage(message, "")
	}
	return nil
}

func init() {
	shard.InitDb()
	es.InitEs()
	hbase.InitHBase()
	cache.SetUp()
}

func main() {
	kafkaAddr := os.Getenv("kafka_addr")
	if kafkaAddr == "" {
		kafkaAddr = "192.168.93.6"
	}
	broker := kafkaAddr + ":9092"
	// 每个group都收到会才会确认offset，但同属一个group的微服务只有一个会收到消息
	group := "score-group"
	topics := []string{"createOrder"}

	version, err := sarama.ParseKafkaVersion("2.3.0")
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	config := sarama.NewConfig()
	config.Version = version
	consumer := Consumer{
		ready: make(chan bool, 0),
	}

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup([]string{broker}, group, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	wg := &sync.WaitGroup{}
	// 消费逻辑放在goroutine中,方便main协程控制退出
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, topics, &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()
	<-consumer.ready
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	// 在这里阻塞住，不然就直接执行完了
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		// ctrl+c
		log.Println("terminating: via signal")
	}

	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}
