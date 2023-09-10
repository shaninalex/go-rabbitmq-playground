package applog

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func (applog *AppLog) defineQueues() error {
	q_new_user, err := applog.Channel.QueueDeclare(
		"q.new-user", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return err
	}
	applog.qNewUser = q_new_user

	q_user_login, err := applog.Channel.QueueDeclare(
		"q.user-login", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return err
	}
	applog.qUserLogin = q_user_login

	q_log, err := applog.Channel.QueueDeclare(
		"q.account.error.log", // name
		false,                 // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		return err
	}
	applog.qErrorLog = q_log

	return nil
}

func (applog *AppLog) defineConnections(rmq string) error {
	connectRabbitMQ, err := amqp.Dial(rmq)
	if err != nil {
		return err
	}
	applog.Connection = connectRabbitMQ

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		return err
	}
	applog.Channel = channelRabbitMQ
	return nil
}
