package app

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	amqp "github.com/rabbitmq/amqp091-go"
)

// TODO: add rabbitmq exchange instead. For example:
// - "loggin" - for sending logging messages
// - "manage" - background tasks like update avatar link from filestorage service, or schedule payments etc...
type App struct {
	Collection         *mongo.Collection
	RabbitMQConnection *amqp.Connection
	router             *gin.Engine
}

func (app *App) Initialize(mongo_connection, rabbitmq_connection string) error {

	app.router = gin.Default()
	client := ConnectDB()
	app.Collection = GetCollection(client, "application")

	// Connect with RabbitMQ
	mq_connection, err := connectToRabbitMQ(rabbitmq_connection)
	if err != nil {
		return err
	}
	defer mq_connection.Close()
	app.RabbitMQConnection = mq_connection
	app.initializeRoutes()

	return nil
}

func (app *App) initializeRoutes() {
	app.router.GET("/ping", Ping)
	app.router.GET("/account/:sub", app.GetUser)
	app.router.POST("/account", app.CreateUser)
	app.router.PATCH("/account", app.UpdateUser)
}

func (app *App) Run() {
	app.router.Run("localhost:8000")
}
