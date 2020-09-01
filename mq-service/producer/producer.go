package producer

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
)
var producer sarama.AsyncProducer

func InitProducer() {
	kafkaAddr := os.Getenv("kafka_addr")
	if kafkaAddr == "" {
		kafkaAddr = "192.168.93.6"
	}
	brokerList := []string{kafkaAddr + ":9092"}

	config := sarama.NewConfig()
	// 发送后需要等待队列发送ack，然后这边才会确认，否则会重试
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Version = sarama.V1_0_0_0
	var err error
	producer, err = sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	// We will just log to STDOUT if we're not able to produce messages.
	// Note: messages will only be returned here after all retry attempts are exhausted.
	go func() {
		for err := range producer.Errors() {
			log.Println("Failed to write access log entry:", err)
		}
	}()
}


// 在发送topic时可以指定partition数量，为1的话可以保证顺序性但性能就差点
func Producer(data string, key string, topic string, uniId string) {
	var enqueued, errors int
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("uniId"),
				Value: []byte(uniId),
			},
		},
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(data),
	}

	producer.Input() <- msg

	enqueued++
	log.Printf("Enqueued: %d; queue:%s; errors: %d\n", enqueued, data, errors)
}

