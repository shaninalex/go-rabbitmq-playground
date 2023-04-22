package main

import (
	"account/app"
	"log"
)

func main() {
	log.Println("Account server is running")

	app := app.App{}
	app.Initialize("", "")
}
