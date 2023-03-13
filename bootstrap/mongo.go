package bootstrap

import (
	"bingo-example/pkg/config"
	"bingo-example/pkg/database"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SetupMongo 设置mongo
func SetupMongo() {
	//mongodb://<dbuser>:<dbpassword>@ds041154.mongolab.com:41154/location
	database.ConnectMongo(options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%v",
			config.Get("mongo.user"),
			config.Get("mongo.password"),
			config.Get("mongo.host"),
			config.Get("mongo.port"),
		)))
}
