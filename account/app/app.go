package app

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	SERVICE_ACCOUNT_PORT = os.Getenv("SERVICE_ACCOUNT_PORT")
)

type App struct {
	Router *gin.Engine
	DB     *sql.DB
}

func (app *App) Initialize(database_conn, rabbitmq_connection string) error {
	app.Router = gin.Default()

	db, err := sql.Open("postgres", database_conn)
	if err != nil {
		panic(err)
	}
	app.DB = db
	app.initializeRoutes()

	return nil
}

func (app *App) initializeRoutes() {
	app.Router.POST("/account", app.CreateUser)
	// app.router.GET("/account/:id", app.GetUser)
	// app.router.PATCH("/account/:id", app.UpdateUser)
	// TODO: login
}

func (app *App) Run() {
	app.Router.Run(fmt.Sprintf(":%s", SERVICE_ACCOUNT_PORT))
}
