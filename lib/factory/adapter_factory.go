package factory

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xylong/bingo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const (
	ElasticsearchAdapter = iota // es
	GormAdapter                 // gorm
	MongoAdapter                // mongo
	RedisAdapter                // redis
)

// AdapterType 适配类型
type AdapterType int

// CreateAdapter 创建适配器
type CreateAdapter func(*bingo.Config) interface{}

// CreateBoot 创建驱动
func CreateBoot(adapterType AdapterType) CreateAdapter {
	switch adapterType {
	case GormAdapter:
		return NewGorm()
	case MongoAdapter:
		return NewMongo()
	default:
		return nil
	}
}

// NewGorm 创建gorm
func NewGorm() CreateAdapter {
	return func(config *bingo.Config) interface{} {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			config.Get("mysql.user"), config.GetString("mysql.password"), config.Get("mysql.host"),
			config.Get("mysql.port"), config.Get("mysql.db"), config.Get("mysql.charset"))

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
			panic(err)
		}

		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(20)
		sqlDB.SetConnMaxLifetime(time.Second * 10)

		return db
	}
}

// NewMongo 创建mongo
func NewMongo() CreateAdapter {
	return func(config *bingo.Config) interface{} {
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

		if err = client.Ping(ctx, nil); err != nil {
			logrus.Errorf("ping mongo error: %s", err.Error())
		}

		return client
	}
}
