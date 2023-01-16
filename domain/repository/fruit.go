package repository

import "bingo-example/domain/entity/fruit"

type IFruitRepo interface {
	// TopViewByType 按类型分组取浏览量最高的前n条
	TopViewByType(uint) []*fruit.Fruit

	// GroupSearch 分组取n条
	GroupSearch(string, uint, bool) []*fruit.Fruit
}
