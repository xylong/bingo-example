package core

import (
	"bingo-example/lib/config"
	. "bingo-example/lib/factory"
	"github.com/olivere/elastic/v7"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// Client 客户端连接
type Client struct{}

func NewClient() *Client {
	config.Init()
	return new(Client)
}

// Gorm 创建gorm
func (c *Client) Gorm() *gorm.DB {
	return new(ClientFactory).Create(GormClient).(*gorm.DB)
}

// Mongo 创建mongo
func (c *Client) Mongo() *mongo.Client {
	return new(ClientFactory).Create(MongoClient).(*mongo.Client)
}

// Redis 创建redis客户端
func (c *Client) Redis() *redis.Client {
	return new(ClientFactory).Create(RedisClient).(*redis.Client)
}

// Elastic 创建elasticSearch
func (c *Client) Elastic() *elastic.Client {
	return new(ClientFactory).Create(ElasticClient).(*elastic.Client)
}
