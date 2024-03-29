package app

import (
	"fmt"
	"os"
	"products/app/controllers"

	"github.com/gin-gonic/gin"
)

type App struct {
	router     *gin.Engine
	controller *controllers.ProductController
	logger     *Logger
	ch_log     chan string
}

func Initialize(rabbitmq_url, postgres_url string) (*App, error) {
	app := &App{}
	app.router = gin.Default()
	pc, err := controllers.Init(postgres_url)
	if err != nil {
		return nil, err
	}
	app.ch_log = make(chan string)
	logger := InitLogger(rabbitmq_url)
	app.logger = logger
	app.controller = pc
	app.initializeRoutes()
	return app, nil
}

func (app *App) initializeRoutes() {
	app.router.GET("/products", app.ProductsList)
	app.router.POST("/products", app.ProductCreate)
	app.router.GET("/products/:id", app.ProductDetail)
	app.router.PATCH("/products/:id", app.ProductPatch)
	app.router.DELETE("/products/:id", app.ProductDelete)
}

func (app *App) Run() {
	app.router.Run(fmt.Sprintf(":%s", os.Getenv("PRODUCTS_SERVICE_PORT")))
}
