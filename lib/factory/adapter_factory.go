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
	GormAdapter  = iota // gorm
	MongoAdapter        // mongo
	RedisAdapter        // redis
	Elastic             // elasticSearch
)

// AdapterType 适配类型
type AdapterType int

// CreateAdapter 创建适配器
type CreateAdapter func(configure lib.Configure) interface{}

// CreateBoot 创建驱动
func CreateBoot(adapterType AdapterType) CreateAdapter {
	switch adapterType {
	case GormAdapter:
		return NewGorm()
	case MongoAdapter:
		return NewMongo()
	case Elastic:
		return NewEs()
	default:
		return nil
	}
}

// NewGorm 创建gorm
func NewGorm() CreateAdapter {
	return func(configure lib.Configure) interface{} {
		config := configure.Mysql()

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

// NewMongo 创建mongo
func NewMongo() CreateAdapter {
	return func(configure lib.Configure) interface{} {
		config := configure.Mongo()

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

func NewEs() CreateAdapter {
	return func(configure lib.Configure) interface{} {
		config := configure.Elastic()
		url := fmt.Sprintf("http://%s:%d", config.Host, config.Port)

		client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(config.Sniff))
		if err != nil {
			zap.L().Fatal("init elastic error", zap.Error(err))
		}

		return client
	}
}
