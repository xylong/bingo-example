package service

import (
	"bingo-example/domain/entity/fruit"
	"bingo-example/infrastructure/dao"
	"gorm.io/gorm"
)

type FruitService struct {
	DB *gorm.DB `inject:"-"`
}

func (s *FruitService) Top() []*fruit.Fruit {
	return dao.NewFruitRepo(s.DB).GroupSearch()
}
