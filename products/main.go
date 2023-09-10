package main

import (
	"os"
	"products/app"
)

func main() {
	app := app.App{}
	app.Initialize(os.Getenv("RABBITMQ_URL"))
	app.Run()
}
