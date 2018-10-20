package main

import (
	"os"
)

func main() {
	app := App{}
	app.Initialize(
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	app.Run(":8000")
}
