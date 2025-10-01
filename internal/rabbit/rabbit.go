package rabbit

import (
	"log"

	"github.com/streadway/amqp"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

func InitRabbit(dsn string) {
	var err error
	Conn, err = amqp.Dial(dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к RabbitMQ: %v", err)
	}

	Channel, err = Conn.Channel()
	if err != nil {
		log.Fatalf("Ошибка создания канала: %v", err)
	}
}
