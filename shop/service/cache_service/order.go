package cache_service

import (
	"strconv"
	"strings"

	"shop/pkg/e"
)

type Order struct {
	ID            int
	GoodsId       int
	UserId        int
	Status 		  int
	CreateAt      string

	PageNum  int
	PageSize int
}

func (a *Order) GetOrderKey() string {
	return e.CACHE_ORDER + "_" + strconv.Itoa(a.ID)
}

func (a *Order) GetOrderListKey() string {
	keys := []string{
		e.CACHE_ORDER,
		"LIST",
	}

	if a.ID > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
	}
	if a.UserId > 0 {
		keys = append(keys, strconv.Itoa(a.UserId))
	}
	if a.GoodsId > 0 {
		keys = append(keys, strconv.Itoa(a.GoodsId))
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
