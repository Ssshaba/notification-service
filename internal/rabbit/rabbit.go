package rabbit

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

func InitRabbit(dsn string) {
	var err error
	for i := 0; i < 10; i++ {
		Conn, err = amqp.Dial(dsn)
		if err == nil {
			break
		}
		log.Printf("Не удалось подключиться к RabbitMQ, пробуем ещё... (%d/10)", i+1)
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		log.Fatalf("❌ Ошибка подключения к RabbitMQ после 10 попыток: %v", err)
	}

	Channel, err = Conn.Channel()
	if err != nil {
		log.Fatalf("❌ Ошибка создания канала: %v", err)
	}
}
