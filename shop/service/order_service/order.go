package order_service

import (
	"encoding/json"
	"shop/models"
	"shop/pkg/gredis"
	"shop/pkg/logging"
	"shop/service/cache_service"
)

type Order struct {
	ID            int
	GoodsId       string
	UserId        string
	Status 		  int
	CreateAt      string
}

//func (a *Goods) Add() error {
//	article := map[string]interface{}{
//		"tag_id":          a.TagID,
//		"title":           a.Title,
//		"desc":            a.Desc,
//		"content":         a.Content,
//		"created_by":      a.CreatedBy,
//		"cover_image_url": a.CoverImageUrl,
//		"state":           a.State,
//	}
//
//	if err := models.AddArticle(article); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (a *Goods) Edit() error {
//	return models.EditArticle(a.ID, map[string]interface{}{
//		"tag_id":          a.TagID,
//		"title":           a.Title,
//		"desc":            a.Desc,
//		"content":         a.Content,
//		"cover_image_url": a.CoverImageUrl,
//		"state":           a.State,
//		"modified_by":     a.ModifiedBy,
//	})
//}

func (a *Order) Get() (*models.Order, error) {
	var cacheOrder *models.Order

	cache := cache_service.Order{ID: a.ID}
	key := cache.GetOrderKey()

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheOrder)
			return cacheOrder, nil
		}
	}

	order, err := models.GetOrder(a.ID)
	if err != nil {
		return nil, err
	}

	gredis.Set(key, order, 3600)
	return order, nil
}

//func (a *Goods) GetAll() ([]*models.Article, error) {
//	var (
//		articles, cacheArticles []*models.Article
//	)
//
//	cache := cache_service.Article{
//		TagID: a.TagID,
//		State: a.State,
//
//		PageNum:  a.PageNum,
//		PageSize: a.PageSize,
//	}
//	key := cache.GetArticlesKey()
//	if gredis.Exists(key) {
//		data, err := gredis.Get(key)
//		if err != nil {
//			logging.Info(err)
//		} else {
//			json.Unmarshal(data, &cacheArticles)
//			return cacheArticles, nil
//		}
//	}
//
//	articles, err := models.GetArticles(a.PageNum, a.PageSize, a.getMaps())
//	if err != nil {
//		return nil, err
//	}
//
//	gredis.Set(key, articles, 3600)
//	return articles, nil
//}

func (a *Order) Delete() error {
	return models.DeleteGoods(a.ID)
}

func (a *Order) ExistByID() (bool, error) {
	return models.ExistOrderByID(a.ID)
}

func (a *Order) GetCount() (int, error) {
	return models.GetOrderTotal(a.getMaps())
}

func (a *Order) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if a.Status != 0 {
		maps["status"] = a.Status
	}

	return maps
}
