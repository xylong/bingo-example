package dao

import (
	"bingo-example/domain/entity/fruit"
	"gorm.io/gorm"
)

type FruitRepo struct {
	db *gorm.DB
}

func NewFruitRepo(db *gorm.DB) *FruitRepo {
	return &FruitRepo{db: db}
}

func (r *FruitRepo) GroupSearch() []*fruit.Fruit {
	var fruits []*fruit.Fruit

	sql := "SELECT type,name,view FROM\n(SELECT type,name,view,IF(@g=type,@num:=@num+1,@num:=1) as num,@g:=type FROM\n(SELECT type,name,view FROM fruits GROUP BY type,name ORDER BY type, view DESC) a,\n(SELECT @num:=0,@g:='') b) c WHERE num<=2"
	r.db.Raw(sql).Scan(&fruits)

	return fruits
}
