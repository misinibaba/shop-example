package model

import (
	"github.com/jinzhu/gorm"
)

type Goods struct {
	ID            string `gorm:"primary_key" json:"id"`
	Title         string `json:"title"`
	Describe      string `json:"describe"`
	Count         int	 `json:"count"`
}

func (goods Goods) TableName() string {
	return "goods"
}

func GetGoods(Id string) (*Goods, error) {
	var goods Goods
	err := db.Where("id = ? ", Id).First(&goods).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &goods, nil
}

func GetGoodsCount(Id int) (int, error) {
	var count int
	if err := db.Model(&Goods{}).Where("id = ?", Id).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func BuyGoods(goodsId, userId, uniId string) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	tx.Model(Goods{ID:goodsId}).UpdateColumn("count", gorm.Expr("count - ?", 1))
	if err := tx.Table("goodsLog").Create(&GoodsLog{GoodsId:goodsId, UserId:userId, Type:1, Count:1, UniId:uniId, Status:1}).Error; err != nil {
		tx.Rollback()
	}

	tx.Commit()
}