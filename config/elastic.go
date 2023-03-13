package config

import "bingo-example/pkg/config"

func init() {
	config.Add("elastic", func() map[string]interface{} {
		return map[string]interface{}{
			"host":  config.Env("ES_HOST.HOST", "127.0.0.1"),
			"port":  config.Env("ES_PORT", 9200),
			"sniff": config.Env("ES_SNIFF", ""),
		}
	})
}
