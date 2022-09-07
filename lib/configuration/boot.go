package configuration

import (
	"bingo-example/lib"
	"bingo-example/lib/factory"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// Boot 基础驱动
type Boot struct {
}

func NewBoot() *Boot {
	return new(Boot)
}

// Gorm 创建gorm实例
func (b *Boot) Gorm() *gorm.DB {
	return factory.CreateBoot(factory.GormAdapter)(lib.Config).(*gorm.DB)
}

// Mongo 创建mongo实例
func (b *Boot) Mongo() *mongo.Client {
	return factory.CreateBoot(factory.MongoAdapter)(lib.Config).(*mongo.Client)
}
