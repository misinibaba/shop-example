package goods_service

import (
	"encoding/json"
	"shop/models"
	"shop/pkg/gredis"
	"shop/pkg/logging"
	"shop/service/cache_service"
)

type Goods struct {
	ID            int
	Title         string
	Describe      string
	Count         int
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

func (a *Goods) Get() (*models.Goods, error) {
	var cacheGoods *models.Goods

	cache := cache_service.Goods{ID: a.ID}
	key := cache.GetGoodsKey()

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheGoods)
			return cacheGoods, nil
		}
	}

	goods, err := models.GetGoods(a.ID)
	if err != nil {
		return nil, err
	}

	gredis.Set(key, goods, 3600)
	return goods, nil
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

func (a *Goods) Delete() error {
	return models.DeleteGoods(a.ID)
}

func (a *Goods) ExistByID() (bool, error) {
	return models.ExistGoodsByID(a.ID)
}

//func (a *Goods) GetCount() (int, error) {
//	return models.GetGoodsTotal(a.getMaps())
//}

//func (a *Goods) getMaps() map[string]interface{} {
//	maps := make(map[string]interface{})
//	maps["deleted_on"] = 0
//	if a.State != -1 {
//		maps["state"] = a.State
//	}
//	if a.TagID != -1 {
//		maps["tag_id"] = a.TagID
//	}
//
//	return maps
//}
