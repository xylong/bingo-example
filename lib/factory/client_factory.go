package factory

import (
	"bingo-example/lib"
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
type createClient func(configure lib.Configure) interface{}

// ClientFactory 客户端工厂
type ClientFactory struct{}

// Create 创建连接客户端
func (f *ClientFactory) Create(createType CreateType) interface{} {
	lib.Init()
	switch createType {
	case GormClient:
		return newGorm()(lib.Mysql)
	case MongoClient:
		return newMongo()(lib.Mongo)
	case RedisClient:
		return newRedis()(lib.Redis)
	case ElasticClient:
		return newElastic()(lib.Elastic)
	default:
		return nil
	}
}

// newGorm 创建gorm
func newGorm() createClient {
	return func(configure lib.Configure) interface{} {
		config, _ := configure.(*lib.MysqlConfig)

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			config.User, config.Password, config.Host, config.Port, config.DB, config.Charset)

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

		return db
	}
}

func newMongo() createClient {
	return func(configure lib.Configure) interface{} {
		config, _ := configure.(*lib.MongoConfig)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		//mongodb://<dbuser>:<dbpassword>@ds041154.mongolab.com:41154/location
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(
			fmt.Sprintf("mongodb://%s:%s@%s:%d",
				config.User,
				config.Password,
				config.Host,
				config.Port)))
		if err != nil {
			zap.L().Fatal("init mongo error", zap.Error(err))
		}

		if err = client.Ping(ctx, nil); err != nil {
			zap.L().Error("ping mongo error", zap.Error(err))
		}

		return client
	}
}

func newRedis() createClient {
	return func(configure lib.Configure) interface{} {
		return nil
	}
}

func newElastic() createClient {
	return func(configure lib.Configure) interface{} {
		config, _ := configure.(*lib.ElasticConfig)
		url := fmt.Sprintf("http://%s:%d", config.Host, config.Port)

		client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(config.Sniff))
		if err != nil {
			zap.L().Fatal("init elastic error", zap.Error(err))
		}

		return client
	}
}
