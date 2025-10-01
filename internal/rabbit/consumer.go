package rabbit

import (
	"fmt"
	"log"
)

func Consumer(queue string, handler func(msg []byte) error) {
	_, err := Channel.QueueDeclare(
		queue,
		true,  // durable — очередь выживает после перезапуска сервера
		false, // auto-delete — очередь не удаляется автоматически
		false, // exclusive — не эксклюзивная
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Ошибка при создании очереди: %v", err)
	}

	msgs, err := Channel.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Ошибка при получении сообщений: %v", err)
	}

	go func() {
		for d := range msgs {
			fmt.Printf("Получено сообщение: %s\n", d.Body)
			if err := handler(d.Body); err != nil {
				log.Printf("Ошибка обработки: %v", err)
			}
		}
	}()
}
