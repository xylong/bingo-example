package dao

import (
	"bingo-example/domain/entity/fruit"
	"database/sql"
	"gorm.io/gorm"
)

type FruitRepo struct {
	db *gorm.DB
}

func NewFruitRepo(db *gorm.DB) *FruitRepo {
	return &FruitRepo{db: db}
}

// TopViewByType 按类型分组取浏览量最高的前n条
func (r *FruitRepo) TopViewByType(limit uint) []*fruit.Fruit {
	var fruits []*fruit.Fruit

	s := "SELECT type,name,view FROM\n" + // 4.每类取n条
		"(SELECT type,name,view,IF(@g=type,@num:=@num+1,@num:=1) as num,@g:=type FROM\n" + // 3.出现过的类型行号+1
		"(SELECT type,name,view FROM fruits ORDER BY type, view desc) a,\n" + // 1.按类型、浏览量分组排序
		"(SELECT @num:=0,@g:='') b) c WHERE num<=@n" // 2.构建行号

	r.db.Raw(s, sql.Named("n", limit)).Scan(&fruits)
	return fruits
}

// GroupSearch 分组取n条
func (r *FruitRepo) GroupSearch(group string, limit uint, order string) []*fruit.Fruit {
	var fruits []*fruit.Fruit

	sortQuery := r.db.Model(&fruit.Fruit{}).Select("type,name,view").Order(group).Order(order)
	lineQuery := r.db.Raw("SELECT @num:=0,@g:=''")
	mainQuery := r.db.Table("(?) as a,(?) as b", sortQuery, lineQuery).
		Select("type,name,view,IF(@g=type,@num:=@num+1,@num:=1) as num,@g:=type")
	r.db.Table("(?) as c", mainQuery).Select("type,name,view").Where("num<=?", limit).Find(&fruits)

	return fruits
}
