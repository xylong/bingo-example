package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var mdb *mongo.Client

// Mongo 获取mongo
func Mongo() *mongo.Client {
	return mdb
}

// ConnectMongo 连接mongo
func ConnectMongo(clientOptions ...*options.ClientOptions) {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if mdb, err = mongo.Connect(ctx, clientOptions...); err != nil {
		panic(err)
	}

	if err = mdb.Ping(ctx, nil); err != nil {
		panic(err)
	}
}
