package dao

import "gorm.io/gorm"

type query struct {
}

// lineNumber 构建分组行号
func (q *query) lineNumber(db *gorm.DB) *gorm.DB {
	return db.Select("@num:=0,@g:=''")
}
