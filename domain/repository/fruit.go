package repository

import "bingo-example/domain/entity/fruit"

type IFruitRepo interface {
	// GroupSearch 分组取n条
	GroupSearch(n uint) []*fruit.Fruit
}
