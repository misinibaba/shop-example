package v1

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"shop/pkg/es"

	//"shop/pkg/es"
	"shop/pkg/hbase"
	"shop/pkg/shard"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"

	"shop/pkg/app"
	"shop/pkg/e"
	"shop/service/order_service"
)

// @Summary Get a single article
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles/{id} [get]
func GetOrder(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	orderService := order_service.Order{ID: id}
	exists, err := orderService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	order, err := orderService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}
	appG.C.HTML(http.StatusOK, "goods/list.tmpl", gin.H{
		"id": order.ID,
		"goodsId": order.GoodsId,
		"userId": order.UserId,
	})
}

// @Summary Get multiple articles
// @Produce  json
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles [get]
func GetOrderList(c *gin.Context) {
	appG := app.Gin{C: c}
	//id := com.StrTo(c.Param("id")).MustInt()

	type Order struct {
		ID            int64
		GoodsId       string
		UserId        string
		Status 		  int
		CreateAt      string
	}
	order := new(Order)
	orderList := []Order{}
	// 查询订单
	userId := 11067
	shardKey := shard.GetDbShardKey(userId)
	orderDB, shardTable := shard.GetDbAndTable(shardKey, userId)
	fmt.Println(shardTable)
	rows, _ := orderDB.Query("select id, goodsId, userId, status, createAt from " +  shardTable + " where userId = ? and status = ?",  userId, 1)
	for rows.Next() {
		if err := rows.Scan(&order.ID, &order.GoodsId, &order.UserId, &order.Status, &order.CreateAt); err != nil {
			log.Fatal(err)
		}
		orderList = append(orderList, *order)
	}

	appG.C.HTML(http.StatusOK, "order/user_list.tmpl", gin.H{
		"orderList": orderList,
	})
}

func GetOrderAllList(c *gin.Context) {
	userId := c.PostForm("user_id")
	//orderId := c.PostForm("order_id")
	describe := c.PostForm("describe")
	status := c.PostForm("status")
	timeType := c.PostForm("time_type")

	// --------------------组装hbase的开始结束key--------------
	nowTime := time.Now().Unix()
	var startRowKey string
	var endRowKey string

	// 根据条件组装startRow 和 endRow
	// 全部是开始为0结束为3，如果有status则是开始是status结束是status
	endStatusRowKey := status
	if status == "0" {
		endStatusRowKey = "3"
	}

	// 默认是开始为7个0结束为7个9，如果有userId则是开始是userId结束是userId
	userIdRowKey := strings.Repeat("0", 7)
	endUserIdRowKey := strings.Repeat("9", 7)
	if userId != "" {
		userIdRowKey = hbase.ReverseStr(fmt.Sprintf("%07s", userId))
		endUserIdRowKey = userIdRowKey
	}

	// 默认是开始为当前时间取反，结束时间为当前时间取反，如果有userId则是开始是userId结束是userId
	endTime := int64(0)
	if timeType == "1" {
		endTime = nowTime - (60 * 60 * 24 * 30)
	} else if timeType == "2" {
		endTime = nowTime - (60 * 60 * 24 * 30 * 365)
	}
	// 常规时间是一个月前比较小，然后翻转后就是比较大了，所以作为结束key
	endTimeRowKey := strconv.FormatInt(math.MaxInt64 - endTime, 10)
	startTimeRowKey := strconv.FormatInt(math.MaxInt64 - nowTime, 10)

	startRowKey = userIdRowKey + status + startTimeRowKey
	endRowKey = endUserIdRowKey + endStatusRowKey + endTimeRowKey
	fmt.Println("-----------------------------")
	fmt.Println("startrowkey= " + startRowKey)
	fmt.Println("endrowkey= " + endRowKey)
	fmt.Println("-----------------------------")
	// ------------------结束组装key---------------------


	var dataList []map[string]string
	//-----------------------没有describe直接搜索hbase------------------------
	if describe == "" {
		scanRsp := hbase.ScanRange("order", startRowKey, endRowKey)
		scanNext, _ := scanRsp.Next()
		for scanNext != nil {
			cellData := make(map[string]string)
			for _, v := range scanNext.Cells {
				cellData[string(v.Qualifier)] = string(v.Value)
			}
			dataList = append(dataList, cellData)

			scanNext, _ = scanRsp.Next()
		}
	} else {
		// -------------------------------有describe则用es作为index索引查找--------------------------
		match := map[string]interface{}{}
		match["describe"] = describe

		query := map[string]interface{}{
			"query": map[string]interface{}{
				"match": match,
			},
			"_source": []string{"hBaseKey"},
		}

		searchRes := es.Search("order", query)
		for _, v := range searchRes.Data {
			source := v["_source"]
			switch value := source.(type) {
				case map[string]interface{}:
					for _, hBaseKey := range value {
						// 获取filter
						pFilter := ""
						if userId != "" {
							pFilter += hbase.ReverseStr(fmt.Sprintf("%07s", userId))
						}

						// 根据hBaseKey查询
						getData := hbase.GetData("order", fmt.Sprintf("%v", hBaseKey), pFilter)
						cellData := make(map[string]string)
						for _, v := range getData.Cells {
							cellData[string(v.Qualifier)] = string(v.Value)
						}
						if len(cellData) != 0 {
							dataList = append(dataList, cellData)
						}
					}
			}
		}
	}

	appG := app.Gin{C: c}
	appG.C.HTML(http.StatusOK, "order/list.tmpl", gin.H{
		"orderList": dataList,
	})
}