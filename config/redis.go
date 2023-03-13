package config

import "bingo-example/pkg/config"

func init() {
	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{
			"host":     config.Env("RDB_HOST", "127.0.0.1"),
			"port":     config.Env("RDB_PORT", 6379),
			"password": config.Env("RDB_PASSWORD", ""),
			"db":       config.Env("RDB_DB", 0),
		}
	})
}
