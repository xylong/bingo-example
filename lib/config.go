package lib

import (
	"github.com/spf13/viper"
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

// Configuration 配置
type Configuration struct {
	*MysqlConfig
	*MongoConfig
}

func NewConfiguration() *Configuration {
	return &Configuration{new(MysqlConfig), new(MongoConfig)}
}

// Mysql mysql配置
func (c *Configuration) Mysql() *MysqlConfig {
	if err := viper.UnmarshalKey("mysql", c.MysqlConfig); err != nil {
		return nil
	}

	return c.MysqlConfig
}

// Mongo mongo配置
func (c *Configuration) Mongo() *MongoConfig {
	if err := viper.UnmarshalKey("mongo", c.MongoConfig); err != nil {
		return nil
	}

	return c.MongoConfig
}
