package app

import (
	"account/app/applog"
	"account/app/models"
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
	Router   *gin.Engine
	DB       *sql.DB
	ch_user  chan *models.NewUser
	ch_login chan int64
	ch_error chan error
	Log      *applog.AppLog
}

func (app *App) Initialize(database_conn, rabbitmq_conn string) error {
	app.Router = gin.Default()
	db, err := sql.Open("postgres", database_conn)
	if err != nil {
		return err
	}

	app.DB = db
	app.initializeRoutes()
	applog, err := applog.InitializeApplicationLog(rabbitmq_conn)
	if err != nil {
		return err
	}

	app.Log = applog
	app.ch_user = make(chan *models.NewUser)
	app.ch_login = make(chan int64)
	app.ch_error = make(chan error)

	return nil
}

func (app *App) initializeRoutes() {
	app.Router.POST("/account", app.CreateUser)
	app.Router.POST("/account/login", app.Login)

	private := app.Router.Group("/user")
	private.Use(AuthMiddleware())
	{
		private.GET("/:id", app.GetUser)
	}
}

func (app *App) Run() {
	app.Router.Run(fmt.Sprintf(":%s", SERVICE_ACCOUNT_PORT))
}

func (app *App) ListenChannels() {
	for {
		select {
		case user := <-app.ch_user:
			go app.Log.UserCreated(user)
		case user_id := <-app.ch_login:
			go app.Log.UserLoggined(user_id)
		case err := <-app.ch_error:
			go app.Log.ErrorHappend(err)
		}
	}
}
