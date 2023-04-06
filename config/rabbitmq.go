package config

import "bingo-example/pkg/config"

func init() {
	config.Add("rabbitmq", func() map[string]interface{} {
		return map[string]interface{}{
			"host":     config.Env("MQ_HOST", "127.0.0.1"),
			"port":     config.Env("ES_PORT", 15672),
			"user":     config.Env("MQ_USER", ""),
			"password": config.Env("MQ_PASSWORD", ""),
		}
	})
}
