package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"sync"
)

var once sync.Once

var (
	Mysql   *MysqlConfig
	Mongo   *MongoConfig
	Redis   *RedisConfig
	Elastic *ElasticConfig
)

func Init() {
	once.Do(func() {
		Mysql = new(MysqlConfig)
		if err := Mysql.load(); err != nil {
			zap.L().Fatal("load mysql config error", zap.Error(err))
		}

		Mongo = new(MongoConfig)
		if err := Mongo.load(); err != nil {
			zap.L().Fatal("load mongo config error", zap.Error(err))
		}

		Redis = new(RedisConfig)
		if err := Redis.load(); err != nil {
			zap.L().Fatal("load redis config error", zap.Error(err))
		}

		Elastic = new(ElasticConfig)
		if err := Elastic.load(); err != nil {
			zap.L().Fatal("load elastic config error", zap.Error(err))
		}
	})
}

// Configure 配置
type Configure interface {
	load() error
}

// MysqlConfig mysql配置
type MysqlConfig struct {
	Host     string
	Port     int
	DB       string
	User     string
	Password string
	Charset  string
}

func (c *MysqlConfig) load() error {
	return viper.UnmarshalKey("mysql", c)
}

// MongoConfig mongo配置
type MongoConfig struct {
	Host     string
	Port     int
	DB       string
	User     string
	Password string
}

func (c *MongoConfig) load() error {
	return viper.UnmarshalKey("mongo", c)
}

// RedisConfig redis配置
type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

func (c *RedisConfig) load() error {
	return viper.UnmarshalKey("redis", c)
}

// ElasticConfig es配置
type ElasticConfig struct {
	Host  string
	Port  int
	Sniff bool
}

func (c *ElasticConfig) load() error {
	return viper.UnmarshalKey("elastic", c)
}
