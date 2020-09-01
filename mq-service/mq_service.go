package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"mq_service/jobs"
	"mq_service/model"
	mqservicegrpc "mq_service/proto"
	"net"
	_ "net/http/pprof"
	"os"
	//"mq_service/jobs"
	"mq_service/producer"
	"net/http"
	"net/url"
	"strconv"
)
type server struct {}
const (
	DEFAULT_TYPE = iota
)

func (s *server) CommitMsg(ctx context.Context,in *mqservicegrpc.SyncRequest) (*mqservicegrpc.SyncResponse, error) {
	var response = mqservicegrpc.SyncResponse{Data: make([]*mqservicegrpc.SyncData, 0, 10)}
	for i := 0; i < len(in.Data); i++{
		request := in.Data[i]
		switch request.Type {
		case DEFAULT_TYPE:
			// get something from local cache and set to response
			log.Println(request.Data)
			msgId, _ := strconv.Atoi(request.Data["msgId"])
			mqMsg, err := model.GetMqMsg(msgId)
			if err != nil {
				log.Fatal(err)
			}

			// 如果这里直接发送队列并等待队列ack的话会比较慢，所以也可以采用定时任务轮询可发送的机制
			// 先发送队列然后修改状态，这样要是反过来如果刚好宕机就不好追踪重试了
			// 既然如果是可以宕机重试，所以会有重发消息的风险，所以需要每个消息沿用一个uniId,在consumer端就依靠这个Id排重
			producer.Producer(mqMsg.Message, mqMsg.Key, mqMsg.Topic, mqMsg.UniId)
			err = model.UpdateMqMsgStatus(msgId, model.MSGSENDED)
			checkErr(err)

			resData := mqservicegrpc.SyncData{}
			resData.Data = make(map[string]string)
			resData.Data["msgId"] = strconv.Itoa(msgId)
			response.Data  = append(response.Data, &resData)
		}
	}
	return &response, nil
}

func (s *server)PreSendMsg(ctx context.Context, in *mqservicegrpc.SyncRequest) (*mqservicegrpc.SyncResponse, error){
	var response = mqservicegrpc.SyncResponse{Data: make([]*mqservicegrpc.SyncData, 0, 10)}
	for i := 0; i < len(in.Data); i++{
		request := in.Data[i]
		switch request.Type {
		case DEFAULT_TYPE:
			// get something from local cache and set to response
			log.Println(request.Data)
			msgId, _ := model.PreSendMqMsg(request.Data["fromService"], request.Data["message"], request.Data["topic"], request.Data["key"], request.Data["uniId"])
			resData := mqservicegrpc.SyncData{}
			resData.Data = make(map[string]string)
			resData.Data["msgId"] = strconv.Itoa(msgId)
			response.Data  = append(response.Data, &resData)
		}
	}
	return &response, nil
}

var db *sql.DB
var err error
var ctx context.Context
var cancel context.CancelFunc
var take_order_sevice_name string

func init() {
	producer.InitProducer()
	model.SetUp()
}

func main() {
	take_order_sevice_name = os.Getenv("take_order_service_name")
	if take_order_sevice_name == "" {
		take_order_sevice_name = "http://localhost:31226"
	} else {
		take_order_sevice_name += ":80"
	}

	// 定时任务处理
	ctx, cancel = context.WithCancel(context.Background())
	go jobs.Jobs(checkPre, 5, ctx) // 这个任务实际上应该是独立于mq_service的单独服务
	defer cancel()

	// rpc相关
	lis, err := net.Listen("tcp", ":31227")
	if err != nil {
		log.Println("start sync server error: %v", err)
	}
	s := grpc.NewServer()
	mqservicegrpc.RegisterSyncServer(s, &server{})

	//由于s.Serve方法是会一直阻塞住，所以我们需要起一个go routine来执行，在其停止后输出错误信息
	s.Serve(lis)

}

func test(w http.ResponseWriter, req *http.Request) {

}

func checkPre() {
	mqMsgList, _ := model.GetMqMsgList(map[string]interface{}{"status": model.MSGPRESEND})
	for _, mqMsg := range mqMsgList {
		var status string
		switch mqMsg.FromService {
		case "pay":
			status = postTo("http://localhost:31230/callBackCheck", mqMsg.Message, mqMsg.UniId, mqMsg.Topic)
			break
		case "take-order":
			status = postTo(take_order_sevice_name + "/callBackCheck", mqMsg.Message, mqMsg.UniId, mqMsg.Topic)
			break
		}

		if status == strconv.Itoa(model.MSGDELETE) {
			// 处理失败的直接存为对应状态,标明此条msg不再处理
			model.UpdateMqMsgStatus(mqMsg.ID, model.MSGDELETE)
		} else {
			// 这里可以选择重新发送到队列，也可以先去下游订单服务凭借uniId查询是否已经生成订单，如果已经生成，那么就直接改状态就好了
			// todo 这里需要调用order的查询接口查询是否有生成该订单

			// 如果下游订单也没有该记录，说明此时订单要么没创建成功，要么就还在队列里或者创建中，那么开始发送到队列，下游也需要做好幂等
			producer.Producer(mqMsg.Message, mqMsg.Key, mqMsg.Topic, mqMsg.UniId)
			model.UpdateMqMsgStatus(mqMsg.ID, model.MSGSENDED)
		}
	}
}

func postTo(postUrl string, data interface{}, uniId string, topic string) string {
	jsonData, _ := json.Marshal(data)

	params := url.Values{}
	params.Set("uniId", uniId)
	params.Set("message", string(jsonData))
	params.Set("topic",topic)
	resp, _ := http.PostForm(postUrl, params)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return string(body)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}