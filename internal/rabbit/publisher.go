package rabbit

import (
	"log"

	"github.com/streadway/amqp"
)

func Publish(queue string, body []byte) error {
	_, err := Channel.QueueDeclare(
		queue,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	err = Channel.Publish(
		"",    // exchange
		queue, // routing key (имя очереди)
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Ошибка отправки сообщения: %v", err)
	}
	return err
}
