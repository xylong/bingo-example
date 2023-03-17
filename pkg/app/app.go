package app

import "bingo-example/pkg/config"

// IsLocal 判断是否为本地开发环境
func IsLocal() bool {
	return config.Get("app.env") == "local"
}

// IsProduction 判断是否为生产环境
func IsProduction() bool {
	return config.Get("app.env") == "production"
}

// IsTesting 判断是否为测试环境
func IsTesting() bool {
	return config.Get("app.env") == "testing"
}
