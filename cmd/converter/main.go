package main

import (
	converter "Gonverter/app/service"
	"Gonverter/app/service/RabbitMQ"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	ch, _, err := RabbitMQ.DeclareQueue("fileToConvert", false, false, false, false, nil)
	RabbitMQ.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		"fileToConvert", // queue
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
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
