package bootstrap

import (
	"bingo-example/pkg/cache"
	"bingo-example/pkg/database"
	"bingo-example/pkg/es"
	"github.com/olivere/elastic/v7"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// Client 连接的客户端
type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

// Mysql mysql客户端
func (c *Client) Mysql() *gorm.DB {
	SetupDB()
	return database.DB()
}

// Redis redis客户端
func (c *Client) Redis() *redis.Client {
	SetupRedis()
	return cache.Redis()
}

// Elastic elastic客户端
func (c *Client) Elastic() *elastic.Client {
	SetupElastic()
	return es.ES()
}

// Mongo mongo客户端
func (c *Client) Mongo() *mongo.Client {
	SetupMongo()
	return database.Mongo()
}
