package app

import (
	"context"
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// SIMPLE FANOUT CONSUMER
type RConsumer struct {
	ch *amqp.Channel
	q  amqp.Queue
}

type EventHandler func([]byte)

func (r *RConsumer) Listen(handler EventHandler) error {
	msgs, err := r.ch.Consume(
		r.q.Name, // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	if err != nil {
		return err
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			handler(d.Body)
		}
	}()

	log.Printf("System is operational. Ready to receive messages.")
	<-forever

	return nil
}

func CreateRConsumer(conn *amqp.Connection, exhangeName, queueName string) (*RConsumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		exhangeName, // name
		"fanout",    // type
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		true,      // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,      // queue name
		"",          // routing key
		exhangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &RConsumer{
		ch: ch,
		q:  q,
	}, nil
}

// SIMPLE FANOUT SENDER
func Push(conn *amqp.Connection, exchangeName string, message string) error {

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to open a channel")
		return err
	}
	err = ch.ExchangeDeclare(
		exchangeName, // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Println("Failed to declare an exchange")
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := message
	err = ch.PublishWithContext(ctx,
		exchangeName, // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		log.Println("Failed to publish a message")
		return err
	}

	log.Printf("Sent message to \"%s\" exchange. PAYLOAD: %s", exchangeName, body)
	return nil
}

// UTILS
func ConnectToRabbitMQ(connString string) (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial(connString) // "amqp://guest:guest@localhost:5672/"
		if err != nil {
			log.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			log.Println("Connected to RabbitMQ")
			connection = c
			break
		}

		if counts > 5 {
			log.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Waiting...")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}
