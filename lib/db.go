package lib

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

// DB 数据库
type DB struct{}

func NewDB() *DB {
	return &DB{}
}

func (d DB) Gorm() *gorm.DB {
	config := ConfigPool.Get().(*bingo.Config)
	defer ConfigPool.Put(config)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.Get("mysql.user"), config.Get("mysql.password"), config.Get("mysql.host"),
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

	if err = client.Ping(ctx, nil); err != nil {
		logrus.Errorf("ping mongo error: %s", err.Error())
	}

	return client
}
