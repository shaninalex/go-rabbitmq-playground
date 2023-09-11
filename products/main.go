package main

import (
	"os"
	"products/app"
)

func main() {
	application, err := app.Initialize(
		os.Getenv("RABBITMQ_URL"),
		os.Getenv("POSTGRESQL_URL"),
	)
	if err != nil {
		panic(err)
	}

	application.Run()
}
