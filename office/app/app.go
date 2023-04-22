package app

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type App struct {
	MQConnection *amqp.Connection
	MQChannel    *amqp.Channel
	MQQueue      *amqp.Queue
}

func (app *App) Initialize(rabbitmqConnectionString string) error {
	mq_connection, err := connectToRabbitMQ(rabbitmqConnectionString)
	if err != nil {
		return err
	}

	ch, err := mq_connection.Channel()
	failOnError(err, "Failed to open a channel")
	q, err := ch.QueueDeclare(
		"register", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	app.MQConnection = mq_connection
	app.MQChannel = ch
	app.MQQueue = &q

	return nil
}

func (app *App) Run() {
	msgs, err := app.MQChannel.Consume(
		app.MQQueue.Name, // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
	// defer connections here!
}
