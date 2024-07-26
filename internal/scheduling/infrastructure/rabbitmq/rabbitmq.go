package rabbitmq

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

func Init() {
	var err error

	// Lê a variável de ambiente
	password := os.Getenv("RABBIT_PASSWORD")
	if password == "" {
		log.Fatal("RABBIT_PASSWORD environment variable is not set")
	}

	amqpURL := "amqp://fcplheqi:" + password + "@albatross-01.rmq.cloudamqp.com:5672/fcplheqi"
	Conn, err = amqp.Dial(amqpURL)
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
