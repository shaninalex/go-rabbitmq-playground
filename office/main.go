package main

import "office/app"

func main() {
	app := app.App{}
	app.Initialize("amqp://guest:guest@localhost:5672/")
	app.Run()
}
