package models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	Model
	GoodsId       string `json:"title"`
	UserId        string `json:"describe"`
	Status 		  int	 `json:"status"`
}

// ExistArticleByID checks if an article exists based on ID
func ExistOrderByID(id int) (bool, error) {
	var order Order
	err := db.Select("id").Where("id = ? ", id).First(&order).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if order.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetArticleTotal gets the total number of articles based on the constraints
func GetOrderTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Order{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetArticles gets a list of articles based on paging constraints
func GetOrderList(pageNum int, pageSize int, maps interface{}) ([]*Order, error) {
	var order []*Order
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&order).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return order, nil
}

// GetArticle Get a single article based on ID
func GetOrder(id int) (*Order, error) {
	var order Order
	err := db.Where("id = ? ", id).First(&order).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &order, nil
}

// EditArticle modify a single article
func EditGoods(id int, data interface{}) error {
	if err := db.Model(&Goods{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddArticle add a single article
//func AddGoods(data map[string]interface{}) error {
//	article := Goods{
//		Title:         data["title"].(string),
//		Describe:      data["describe"].(string),
//		Count:         data["count"].(int),
//		Status:        data["status"].(int),
//		CreateAt: 	   data["createAt"].(string),
//	}
//	if err := db.Create(&article).Error; err != nil {
//		return err
//	}
//
//	return nil
//}

// DeleteArticle delete a single article
//func DeleteGoods(id int) error {
//	if err := db.Where("id = ?", id).Delete(Article{}).Error; err != nil {
//		return err
//	}
//
//	return nil
//}

// CleanAllArticle clear all article
//func CleanAllGoods() error {
//	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{}).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
