package app

import (
	"log"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type App struct {
	Broker *amqp.Connection
	router *gin.Engine
}

func (app *App) Initialize() error {

	brokerConnection, err := ConnectToRabbitMQ("amqp://localhost:5672")
	if err != nil {
		log.Println(err)
	}

	app.Broker = brokerConnection

	return nil
}
