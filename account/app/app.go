package app

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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
	client := ConnectDB(mongo_connection)
	app.Collection = GetCollection(client, "accounts")
	indexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "sub", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}

	indexName, err := app.Collection.Indexes().CreateMany(
		context.Background(),
		indexes,
	)
	if err != nil {
		return err
	}
	log.Printf("Indexes created: %v", indexName)

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
	app.router.POST("/account", app.CreateUser)
	app.router.GET("/account/:id", app.GetUser)
	app.router.PATCH("/account/:id", app.UpdateUser)
}

func (app *App) Run() {
	app.router.Run("localhost:8000")
}
