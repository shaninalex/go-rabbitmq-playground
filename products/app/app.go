package app

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func (app *App) Initialize(rabbitmqConnectionString string) error {
	return nil
}

func (app *App) Run() {
}

func (app *App) initializeRoutes() {
	app.router.GET("/products", app.ProductsList)
	app.router.POST("/products", app.ProductsList)
	app.router.GET("/products/:id", app.ProductDetail)
	app.router.PATCH("/products/:id", app.ProductDetail)
	app.router.DELETE("/products/:id", app.ProductDetail)
}
