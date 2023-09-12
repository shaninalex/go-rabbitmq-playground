package app

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Logger struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	qLog       amqp.Queue
	context    context.Context
}

func InitLogger(rabbitmq_url string) *Logger {
	logger := &Logger{}
	connectRabbitMQ, err := amqp.Dial(rabbitmq_url)
	if err != nil {
		log.Println(err)
		return nil
	}
	logger.Connection = connectRabbitMQ
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		return nil
	}
	logger.Channel = channelRabbitMQ
	qLog, err := logger.Channel.QueueDeclare(
		"q.products.log", // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		return nil
	}
	logger.qLog = qLog
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logger.context = ctx
	return logger
}

type message struct {
	MessageType string `json:"type"`
	Payload     any    `json:"payload"`
}

func (l *Logger) Log(action_type string, payload any) {
	message := &message{
		MessageType: action_type,
		Payload:     payload,
	}

	message_str, err := json.Marshal(message)
	if err != nil {
		log.Println("can't marshal message", err)
	}
	l.Channel.PublishWithContext(l.context,
		"",               // exchange
		"q.products.log", // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message_str,
		},
	)

}
