package app

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	MQConnection *amqp.Connection
	MQChannel    *amqp.Channel
	MQQueue      *amqp.Queue
	router       *gin.Engine
	Collection   *mongo.Collection
}

func (app *App) Initialize(rabbitmqConnectionString string) error {
	return nil
}

func (app *App) Run() {
}

func (app *App) initializeRoutes() {
	store := persistence.NewInMemoryStore(time.Minute)
	app.router.GET("/products", cache.CachePage(store, time.Minute, app.ProductsList))
	app.router.GET("/product/:id", app.ProductDetail)
}
