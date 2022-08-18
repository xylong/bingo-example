package lib

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xylong/bingo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// DB 数据库
type DB struct{}

func NewDB() *DB {
	return &DB{}
}

// Mongo 初始化mongodb
func (d DB) Mongo() *mongo.Client {
	config := ConfigPool.Get().(*bingo.Config)
	defer ConfigPool.Put(config)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//mongodb://<dbuser>:<dbpassword>@ds041154.mongolab.com:41154/location
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%d",
			config.Get("mongo.user"),
			config.GetString("mongo.password"),
			config.Get("mongo.host"),
			config.Get("mongo.port"))))
	if err != nil {
		logrus.Error(err)
	}

	return client
}
