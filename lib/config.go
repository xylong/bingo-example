package lib

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	Config = NewConfiguration()
}

// Config 配置
var Config *Configuration

// Configure 配置
type Configure interface {
	Mysql() *MysqlConfig
	Mongo() *MongoConfig
	Elastic() *ElasticConfig
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

// MongoConfig mongo配置
type MongoConfig struct {
	Host     string
	Port     int
	DB       string
	User     string
	Password string
}

// ElasticConfig es配置
type ElasticConfig struct {
	Host  string
	Port  int
	Sniff bool
}

// Configuration 配置
type Configuration struct {
	*MysqlConfig
	*MongoConfig
	*ElasticConfig
}

func NewConfiguration() *Configuration {
	return &Configuration{new(MysqlConfig), new(MongoConfig), new(ElasticConfig)}
}

// Mysql 获取mysql配置
func (c *Configuration) Mysql() *MysqlConfig {
	if err := viper.UnmarshalKey("mysql", c.MysqlConfig); err != nil {
		zap.L().Error("load mysql config error", zap.Error(err))
		return nil
	}

	return c.MysqlConfig
}

// Mongo 获取mongo配置
func (c *Configuration) Mongo() *MongoConfig {
	if err := viper.UnmarshalKey("mongo", c.MongoConfig); err != nil {
		zap.L().Error("load mongo config error", zap.Error(err))
		return nil
	}

	return c.MongoConfig
}

// Elastic 获取elastic配置
func (c *Configuration) Elastic() *ElasticConfig {
	if err := viper.UnmarshalKey("elastic", c.ElasticConfig); err != nil {
		zap.L().Error("load elastic config error", zap.Error(err))
		return nil
	}

	return c.ElasticConfig
}
