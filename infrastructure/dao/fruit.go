package dao

import (
	"bingo-example/domain/entity/fruit"
	"database/sql"
	"gorm.io/gorm"
)

type FruitRepo struct {
	db *gorm.DB
	*query
}

func NewFruitRepo(db *gorm.DB) *FruitRepo {
	return &FruitRepo{db: db}
}

func (r *FruitRepo) GroupSearch(n uint) []*fruit.Fruit {
	var fruits []*fruit.Fruit

	s := "SELECT type,name,view FROM\n" + // 4.每类取n条
		"(SELECT type,name,view,IF(@g=type,@num:=@num+1,@num:=1) as num,@g:=type FROM\n" + // 3.出现过的类型行号+1
		"(SELECT type,name,view FROM fruits ORDER BY type, view DESC) a,\n" + // 1.按类型、浏览量分组排序
		"(SELECT @num:=0,@g:='') b) c WHERE num<=@n" // 2.构建行号

	r.db.Raw(s, sql.Named("n", n)).Scan(&fruits)
	return fruits
}
