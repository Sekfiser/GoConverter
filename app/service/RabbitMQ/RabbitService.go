package RabbitMQ

import (
	"context"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"log"
	"time"
)

func DeclareQueue(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table) (*amqp.Channel, *amqp.Connection, error) {
	url := viper.GetString("RABBIT_URL")
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, err
	}
	ch, err := conn.Channel()
	if err != nil || ch == nil {
		return nil, nil, errors.New("error getting channel")
	}
	_, err = ch.QueueDeclare(
		name,       // name
		durable,    // durable
		autoDelete, // delete when unused
		exclusive,  // exclusive
		noWait,     // noWait
		args,       // arguments
	)
	return ch, conn, err
}

func MessageToFileQueue(fileString string) {
	url := viper.GetString("RABBIT_URL")
	conn, err := amqp.Dial(url)
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"fileToConvert", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	FailOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fileString),
		})
	FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", fileString)
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
