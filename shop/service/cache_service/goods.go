package cache_service

import (
	"strconv"
	"strings"

	"shop/pkg/e"
)

type Goods struct {
	ID            int
	Status 		  int
	CreateAt      string

	PageNum  int
	PageSize int
}

func (a *Goods) GetGoodsKey() string {
	return e.CACHE_GOODS + "_" + strconv.Itoa(a.ID)
}

func (a *Goods) GetGoodsListKey() string {
	keys := []string{
		e.CACHE_GOODS,
		"LIST",
	}

	if a.ID > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
	}
	if a.Status > 0 {
		keys = append(keys, strconv.Itoa(a.Status))
	}
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, "_")
}
