package es

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"os"
	"strings"
	"sync"
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
var es *elasticsearch.Client
func InitEs() {
	esAddr := os.Getenv("es_addr")
	if esAddr == "" {
		esAddr = "192.168.93.10"
	}

	es, _ = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://"+ esAddr +":9200",
		},
	})
}

func InputData(index string, data Order, id int) {
	var (
		//r  map[string]interface{}
		wg sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		orderJson, _ := json.Marshal(data)
		//Set up the request object.
		req := esapi.IndexRequest{
			Index:      index,
			//DocumentID: strconv.Itoa(id),
			Body:       strings.NewReader(string(orderJson)),
			Refresh:    "true",
		}

		// Perform the request with the client.
		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			buf := new(bytes.Buffer)
			buf.ReadFrom(res.Body)
			log.Printf("[%s] Error: %s", res.Status(), buf.String())
		} else {
			// Deserialize the response into a map.
			var r map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
				log.Printf("Error parsing the response body: %s", err)
			} else {
				// Print the response status and indexed document version.
				log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
			}
		}
	}()

	wg.Wait()
}