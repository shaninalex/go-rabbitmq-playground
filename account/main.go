package main

import (
	"account/app"
	"log"
)

func main() {
	log.Println("Account server is running")

	app := app.App{}
	app.Initialize("mongodb://localhost:27017/application", "amqp://guest:guest@localhost:5672/")
	app.Run()
}
