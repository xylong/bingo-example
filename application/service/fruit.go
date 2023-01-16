package service

import (
	"bingo-example/domain/entity/fruit"
	"bingo-example/infrastructure/dao"
	"gorm.io/gorm"
)

type FruitService struct {
	DB *gorm.DB `inject:"-"`
}

// Top 取前n条
func (s *FruitService) Top() map[string][]*fruit.Fruit {
	result := make(map[string][]*fruit.Fruit)

	fruits := dao.NewFruitRepo(s.DB).TopViewByType(2)
	for _, f := range fruits {
		result[f.Type] = append(result[f.Type], f)
		f.Type = ""
	}

	return result
}
