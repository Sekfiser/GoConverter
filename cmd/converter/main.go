package main

import (
	converter "Gonverter/app/service"
	"Gonverter/app/service/RabbitMQ"
	"log"
)

func main() {
	ch, _, err := RabbitMQ.DeclareQueue("fileToConvert", false, false, false, false, nil)
	RabbitMQ.FailOnError(err, "Failed to declare a queue")

	q, err := ch.QueueDeclare(
		"fileToConvert", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	RabbitMQ.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	RabbitMQ.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			file, _ := converter.OleConvert(string(d.Body))
			log.Printf(file)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
