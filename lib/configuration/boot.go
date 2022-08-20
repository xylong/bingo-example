package configuration

import (
	"bingo-example/lib"
	"bingo-example/lib/factory"
	"github.com/xylong/bingo"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// Boot 基础驱动
type Boot struct {
	config *bingo.Config
}

func NewBoot() *Boot {
	config := lib.ConfigPool.Get().(*bingo.Config)
	defer lib.ConfigPool.Put(config)

	return &Boot{config: config}
}

// Gorm 创建gorm实例
func (b *Boot) Gorm() *gorm.DB {
	return factory.CreateBoot(factory.GormAdapter)(b.config).(*gorm.DB)
}

// Mongo 创建mongo实例
func (b *Boot) Mongo() *mongo.Client {
	return factory.CreateBoot(factory.MongoAdapter)(b.config).(*mongo.Client)
}
