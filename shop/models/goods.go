package models

import (
	"github.com/jinzhu/gorm"
)

type Goods struct {
	Model
	Title         string `json:"title"`
	Describe      string `json:"describe"`
	Count         int	 `json:"count"`
}

// ExistArticleByID checks if an article exists based on ID
func ExistGoodsByID(id int) (bool, error) {
	var goods Goods
	err := db.Select("id").Where("id = ? ", id).First(&goods).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if goods.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetArticleTotal gets the total number of articles based on the constraints
func GetGoodsTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Goods{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetArticles gets a list of articles based on paging constraints
func GetGoodsList(pageNum int, pageSize int, maps interface{}) ([]*Goods, error) {
	var goods []*Goods
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&goods).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return goods, nil
}

// GetArticle Get a single article based on ID
func GetGoods(id int) (*Goods, error) {
	var goods Goods
	err := db.Where("id = ? ", id).First(&goods).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &goods, nil
}

// EditArticle modify a single article
//func EditGoods(id int, data interface{}) error {
//	if err := db.Model(&Goods{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// AddArticle add a single article
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
//
//// DeleteArticle delete a single article
func DeleteGoods(id int) error {
	if err := db.Where("id = ?", id).Delete(Goods{}).Error; err != nil {
		return err
	}

	return nil
}
//
//// CleanAllArticle clear all article
//func CleanAllGoods() error {
//	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{}).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
