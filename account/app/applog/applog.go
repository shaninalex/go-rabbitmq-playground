package applog

import (
	"account/app/models"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AppLog struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func InitializeApplicationLog(rmq string) (*AppLog, error) {
	applog := &AppLog{}
	connectRabbitMQ, err := amqp.Dial(rmq)
	if err != nil {
		return nil, err
	}
	applog.Connection = connectRabbitMQ

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		return nil, err
	}

	applog.Channel = channelRabbitMQ
	return applog, nil
}

func (applog *AppLog) UserCreated(user *models.NewUser) {
	log.Println(user)
}

func (applog *AppLog) UserLoggined(userId int64) {
	log.Printf("User %d successfully loggined", userId)
}

func (applog *AppLog) ErrorHappend(err error) {
	log.Println("error happend: ", err.Error())
}
