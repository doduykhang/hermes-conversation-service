package config

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMq(config *Config) *amqp.Connection {
	conn, err := amqp.Dial(getRabbitMQConnString(config.RabbitMQ))
	if err != nil {
		log.Panic(err)
	}
	return conn
}

func getRabbitMQConnString(rabbitMQ RabbitMQ) string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		rabbitMQ.Protocol,
		rabbitMQ.User,
		rabbitMQ.Password,
		rabbitMQ.Host,
		rabbitMQ.Port,
		rabbitMQ.VHost,
	)
}
