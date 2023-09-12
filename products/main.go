package main

import (
	"fmt"
	"os"
	"products/app"
)

var (
	POSTGRES_URL = fmt.Sprintf(
		"postgresql://%s:%s@%s:5432/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_DB"),
	)
	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
)

func main() {
	application, err := app.Initialize(
		RABBITMQ_URL,
		POSTGRES_URL,
	)
	if err != nil {
		panic(err)
	}

	application.Run()
}
