package applog

import (
	"account/app/models"
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AppLog struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel

	qNewUser   amqp.Queue
	qUserLogin amqp.Queue
	qErrorLog  amqp.Queue
	context    context.Context
}

func InitializeApplicationLog(rmq string) (*AppLog, error) {
	applog := &AppLog{}

	err := applog.defineConnections(rmq)
	if err != nil {
		return nil, err
	}

	err = applog.defineQueues()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	applog.context = ctx

	return applog, nil
}

func (applog *AppLog) UserCreated(user *models.NewUser) {
	user.Password = ""
	body, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	err = applog.send(body, applog.qNewUser.Name)
	if err != nil {
		log.Println(err)
	}
}

func (applog *AppLog) UserLoggined(userId int64) {
	body, err := json.Marshal(struct {
		Timestamp time.Time
		UserId    int64
	}{
		Timestamp: time.Now(),
		UserId:    userId,
	})
	if err != nil {
		log.Println(err)
	}
	err = applog.send(body, applog.qUserLogin.Name)
	if err != nil {
		log.Println(err)
	}
}

func (applog *AppLog) ErrorHappend(err error) {
	err = applog.send([]byte(err.Error()), applog.qErrorLog.Name)
	if err != nil {
		log.Println(err)
	}
}

func (applog *AppLog) send(body []byte, exchangeKey string) error {
	return applog.Channel.PublishWithContext(applog.context,
		"",          // exchange
		exchangeKey, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
