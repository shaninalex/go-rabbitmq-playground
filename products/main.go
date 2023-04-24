package main

import (
	"products/app"
)

func main() {
	app := app.App{}
	app.Initialize("amqp://guest:guest@localhost:5672/")
	// defer app.MQConnection.Close()
	// defer app.MQChannel.Close()
	app.Run()
}
