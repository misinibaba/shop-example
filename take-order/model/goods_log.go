package model

import "github.com/jinzhu/gorm"

type GoodsLog struct {
	ID            string `gorm:"primary_key" json:"id"`
	GoodsId       string `gorm:"column:goodsId" json:"goodsId"`
	UserId        string `gorm:"column:userId" json:"userId"`
	Type          int    `json:"type"`
	Count         int	 `json:"count"`
	UniId         string `gorm:"column:uniId" json:"uniId"`
	Status     	  int    `json:"status"`
}

func (goodsLog GoodsLog) TableName() string {
	return "goodsLog"
}

func GetGoodsLog(uniId string) (*GoodsLog, error) {
	var goodsLog GoodsLog
	err := db.Where("uniId = ? ", uniId).First(&goodsLog).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &goodsLog, nil
}