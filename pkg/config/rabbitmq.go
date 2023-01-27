package config

import (
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMq() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panic(err)
	}
	return conn
}
