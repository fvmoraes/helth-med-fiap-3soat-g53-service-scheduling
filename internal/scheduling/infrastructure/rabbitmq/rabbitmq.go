package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

func Init() {
	var err error
	Conn, err = amqp.Dial("amqp://fcplheqi:${RABBIT_PASSWORD}@albatross-01.rmq.cloudamqp.com:5672/fcplheqi")
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}

	Channel, err = Conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel: %v", err)
	}
}

func Close() {
	Channel.Close()
	Conn.Close()
}
