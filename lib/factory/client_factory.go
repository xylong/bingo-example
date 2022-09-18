package factory

import (
	"bingo-example/lib/config"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const (
	GormClient    CreateType = iota // gorm
	MongoClient                     // mongo
	RedisClient                     // redis
	ElasticClient                   // elasticSearch
)

// createClient 创建第三方客户端连接
type createClient func(chan<- interface{})

// ClientFactory 客户端工厂
type ClientFactory struct{}

// Create 创建连接客户端
func (f *ClientFactory) Create(createType CreateType) interface{} {
	c := make(chan interface{})

	go func() {
		switch createType {
		case GormClient:
			newGorm()(c)
		case MongoClient:
			newMongo()(c)
		case RedisClient:
			newRedis()(c)
		case ElasticClient:
			newElastic()(c)
		default:
			c <- struct{}{}
		}
	}()

	return <-c
}

// newGorm 创建gorm
func newGorm() createClient {
	return func(c chan<- interface{}) {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			config.Mysql.User, config.Mysql.Password, config.Mysql.Host, config.Mysql.Port, config.Mysql.DB, config.Mysql.Charset)

		db, err := gorm.Open(mysql.New(mysql.Config{
			DriverName:                "mysql",
			ServerVersion:             "8.0.13",
			DSN:                       dsn,
			Conn:                      nil,
			SkipInitializeWithVersion: true,
			DefaultStringSize:         191,
		}), &gorm.Config{
			Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold:             time.Second,
				Colorful:                  true,
				IgnoreRecordNotFoundError: false,
				LogLevel:                  logger.Info,
			}),
			SkipDefaultTransaction:                   false,
			DisableForeignKeyConstraintWhenMigrating: true,
		})

		if err != nil {
			zap.L().Fatal("init gorm error", zap.Error(err))
		}

		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(20)
		sqlDB.SetConnMaxLifetime(time.Second * 10)

		c <- db
	}
}

func newMongo() createClient {
	return func(c chan<- interface{}) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		//mongodb://<dbuser>:<dbpassword>@ds041154.mongolab.com:41154/location
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(
			fmt.Sprintf("mongodb://%s:%s@%s:%d",
				config.Mongo.User, config.Mongo.Password, config.Mongo.Host, config.Mongo.Port)))
		if err != nil {
			zap.L().Fatal("init mongo error", zap.Error(err))
		}

		if err = client.Ping(ctx, nil); err != nil {
			zap.L().Error("ping mongo error", zap.Error(err))
		}

		c <- client
	}
}

func newRedis() createClient {
	return func(c chan<- interface{}) {
		c <- struct{}{}
	}
}

func newElastic() createClient {
	return func(c chan<- interface{}) {
		url := fmt.Sprintf("http://%s:%d", config.Elastic.Host, config.Elastic.Port)

		client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(config.Elastic.Sniff))
		if err != nil {
			zap.L().Fatal("init elastic error", zap.Error(err))
		}

		c <- client
	}
}
