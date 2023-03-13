package config

import (
	"bingo-example/pkg/helpers"
	"github.com/spf13/cast"
	viperLib "github.com/spf13/viper"
	"os"
)

var (
	viper *viperLib.Viper

	// ConfigFuncs 先加载到此数组，loadConfig 再动态生成配置信息
	ConfigFuncs map[string]ConfigFunc
)

// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]interface{}

func init() {
	viper = viperLib.New()
	viper.SetConfigType("env")
	viper.AddConfigPath(".")     // 环境变量配置文件查找的路径
	viper.SetEnvPrefix("appenv") // 设置环境变量前缀，用以区分 Go 的系统环境变量
	viper.AutomaticEnv()         // 读取环境变量（支持 flags）

	ConfigFuncs = make(map[string]ConfigFunc, 0)
}

// InitConfig 初始化配置信息，完成对环境变量以及 config 信息的加载
func InitConfig(env string) {
	// 1. 加载环境变量
	loadEnv(env)
	// 2. 注册配置信息
	loadConfig()
}

// 注册配置信息
func loadConfig() {
	for name, f := range ConfigFuncs {
		viper.Set(name, f())
	}
}

// 加载环境变量
func loadEnv(suffix string) {
	// 默认加载 .env 文件，如果有传参 --env=name 的话，加载 .env.name 文件
	envPath := ".env"
	if len(suffix) > 0 {
		filepath := ".env." + suffix
		if _, err := os.Stat(filepath); err != nil {
			panic(err)
		}

		// 如 .env.testing 或 .env.stage
		envPath = filepath
	}

	// 加载 env
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.WatchConfig()
}

// Env 读取环境变量，支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}
	return internalGet(envName)
}

// Add 新增配置项
func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

// Get 获取配置项
// 第一个参数 path 允许使用点式获取，如：app.name
// 第二个参数允许传参默认值
func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue...)
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	// config 或者环境变量不存在的情况
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}

		return nil
	}

	return viper.Get(path)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
