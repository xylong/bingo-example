package repository

import "bingo-example/domain/entity/fruit"

type IFruitRepo interface {
	GroupSearch() []*fruit.Fruit
}
