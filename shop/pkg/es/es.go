package es

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"os"
)
type Order struct {
	Id       int64  `json:"userId"`
	GoodsId  int    `json:"goodsId"`
	UserId   int    `json:"userId"`
	UniId    string `json:"uniId"`
	Status   int    `json:"status"`
	CreateAt string `json:"createAt"`
}

type SearchRes struct {
	Data   []map[string]interface{}
	Status int
	Hits   int
	Took   int
}
var es *elasticsearch.Client

func InitEs() {
	esAddr := os.Getenv("es_addr")
	if esAddr == "" {
		esAddr = "192.168.93.10"
	}

	es, _ = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://" + esAddr +":9200",
		},
	})
}

func Search(index string, query map[string]interface{}) SearchRes {
	var (
		r  map[string]interface{}
		//wg sync.WaitGroup
	)

	// 3. Search for the indexed documents
	//
	// Build the request body.
	var buf bytes.Buffer
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"userId": 27063,
	//		},
	//	},
	//}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)

	var sourceData []map[string]interface{}
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
		sourceData = append(sourceData, hit.(map[string]interface{}))
	}

	searchRes := SearchRes{
		Hits: int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		Took: int(r["took"].(float64)),
		Data: sourceData,
	}

	return searchRes
}