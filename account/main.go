package main

import (
	"account/app"
	"fmt"
	"os"
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
	app := app.App{}
	err := app.Initialize(POSTGRES_URL, RABBITMQ_URL)
	defer app.Log.Connection.Close()
	defer app.Log.Channel.Close()
	if err != nil {
		panic(err)
	}
	go app.ListenChannels()
	app.Run()
}
