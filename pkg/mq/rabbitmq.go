package mq

import "github.com/streadway/amqp"

var conn *amqp.Connection

func Connect(dsn string) {
	var err error
	conn, err = amqp.Dial(dsn)
	if err != nil {
		panic(err)
	}
}

func Connection() *amqp.Connection {
	return conn
}
