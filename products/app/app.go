package app

import (
	"products/app/controllers"

	"github.com/gin-gonic/gin"
)

type App struct {
	router     *gin.Engine
	controller *controllers.ProductController
}

func Initialize(rabbitmq_url, postgres_url string) (*App, error) {
	app := &App{}
	app.router = gin.Default()
	pc, err := controllers.Init(postgres_url)
	if err != nil {
		return nil, err
	}
	app.controller = pc
	return app, nil
}

func (app *App) Run() {
}

func (app *App) initializeRoutes() {
	app.router.GET("/products", app.ProductsList)
	app.router.POST("/products", app.ProductCreate)
	app.router.GET("/products/:id", app.ProductDetail)
	app.router.PATCH("/products/:id", app.ProductPatch)
	app.router.DELETE("/products/:id", app.ProductDelete)
}
