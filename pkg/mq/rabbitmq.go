package mq

import "github.com/streadway/amqp"

var channel *amqp.Channel

func Connect(dsn string) {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		panic(err)
	}

	channel, err = conn.Channel()
	if err != nil {
		panic(err)
	}
}

func Channel() *amqp.Channel {
	return channel
}
