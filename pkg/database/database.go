package database

import (
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

// Connect 连接数据库
func Connect(dbConfig gorm.Dialector, p logger.Interface) {
	var err error

	// 使用 gorm.Open 连接数据库
	if db, err = gorm.Open(dbConfig, &gorm.Config{
		Logger:                                   p,
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		panic(err)
	}

	// 获取底层的 sqlDB
	if sqlDB, err = db.DB(); err != nil {
		panic(err)
	}
}

// DB 获取db
func DB() *gorm.DB {
	return db
}

// SQLDB 获取底层db
func SQLDB() *sql.DB {
	return sqlDB
}
