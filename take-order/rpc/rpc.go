package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	mqservicegrpc "take-order/proto"
)

type syncClient struct {
	client mqservicegrpc.SyncClient
	conn *grpc.ClientConn
	address string
}
var mq_service_service_name string

func Setup() {
	mq_service_service_name = os.Getenv("mq_service_service_name")
	if mq_service_service_name == "" {
		mq_service_service_name = "localhost:31227"
	} else {
		mq_service_service_name += ":80"
	}
}

func OpenSyncClient(address string) (*syncClient, error) {
	s := syncClient{}
	//grpc.WithInsecure用于关闭安全验证，因为我们是在docker内部环境里使用，不暴露在外网，就没有加安全认证了
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("----open client error %v, conn: %v\n", err, conn)
		return &s, err
	}
	s.conn = conn
	s.address = address
	s.client = mqservicegrpc.NewSyncClient(conn)
	return &s, nil
}

func CloseSyncClient(s *syncClient) {
	if s.conn != nil {
		s.conn.Close()
		s.conn = nil
		s.client = nil
	}
}

// key用于kafka那边判定发到哪个partition,uniId用于回调确认
func PreSend(data interface{}, key string, uniId string, topic string) string {
	jsonData, _ := json.Marshal(data)
	fmt.Println(mq_service_service_name)
	s, _ := OpenSyncClient(mq_service_service_name)
	defer CloseSyncClient(s)
	syncData := mqservicegrpc.SyncData{}
	syncData.Type = 0
	syncData.Data = make(map[string]string)
	syncData.Data["fromService"] = "take-order"
	syncData.Data["topic"] = topic
	syncData.Data["key"] = key
	syncData.Data["uniId"] = uniId
	syncData.Data["message"] = string(jsonData)
	dataList := []*mqservicegrpc.SyncData{}
	dataList = append(dataList, &syncData)

	resp, err := s.client.PreSendMsg(context.Background(), &mqservicegrpc.SyncRequest{
		Data: dataList,
	})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("client.Search err: deadline")
			}
		}
		log.Printf("client.Search err: %v \n", err)
	}
	resData := resp.GetData()
	log.Println(resData[0].Data["msgId"])
	return resData[0].Data["msgId"]
}

func CommitMsg(msgId string) {
	s, _ := OpenSyncClient(mq_service_service_name)
	defer CloseSyncClient(s)
	syncData := mqservicegrpc.SyncData{}
	syncData.Type = 0
	syncData.Data = make(map[string]string)
	syncData.Data["msgId"] = msgId
	dataList := []*mqservicegrpc.SyncData{}
	dataList = append(dataList, &syncData)
	resp, err := s.client.CommitMsg(context.Background(), &mqservicegrpc.SyncRequest{
		Data: dataList,
	})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("client.Search err: deadline")
			}
		}
		log.Printf("client.Search err: %v \n", err)
	}
	data := resp.GetData()
	fmt.Printf("%v", data)
}
