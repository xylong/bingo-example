package bootstrap

import (
	"bingo-example/pkg/config"
	"bingo-example/pkg/mq"
	"fmt"
)

func SetupRabbitmq() {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		config.Get("rabbitmq.user"),
		config.Get("rabbitmq.password"),
		config.Get("rabbitmq.host"),
		config.GetInt("rabbitmq.port"))

	mq.Connect(dsn)
}
