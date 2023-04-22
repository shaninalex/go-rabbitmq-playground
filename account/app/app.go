package app

import (
	"context"
	"log"
	"math"
	"time"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	MongoConnection    *mongo.Client
	RabbitMQConnection *amqp.Connection
	router             *gin.Engine
}

func (app *App) Initialize(mongo_connection, rabbitmq_connection string) error {

	// Connect with MongoDB
	mongo_client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_connection))
	if err != nil {
		return err
	}

	app.MongoConnection = mongo_client

	defer func() {
		if err := mongo_client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Connect with RabbitMQ
	app.RabbitMQConnection, err = connectToRabbitMQ("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer app.RabbitMQConnection.Close()

	return nil
}

func connectToRabbitMQ(connectionString string) (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial(connectionString)
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
		log.Println("backing off...")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}
