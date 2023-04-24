package main

import (
	"account/app"
	"log"
)

func main() {
	log.Println("Account server is running")

	app := app.App{}
	app.Initialize("mongodb://127.0.0.1:27017", "amqp://guest:guest@localhost:5672/")

	// need defer connections here, because in other case - thay close after Initialize end
	defer app.MQConnection.Close()
	defer app.MQChannel.Close()
	app.Run()
}
