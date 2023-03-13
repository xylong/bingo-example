package config

import "bingo-example/pkg/config"

func init() {
	config.Add("mongo", func() map[string]interface{} {
		return map[string]interface{}{
			"host":     config.Env("MDB_HOST", "127.0.0.1"),
			"port":     config.Env("MDB_PORT", 27017),
			"user":     config.Env("MDB_USER", "root"),
			"password": config.Env("MDB_PASSWORD", ""),
			"database": config.Env("MDB_DATABASE", ""),
		}
	})
}
