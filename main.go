package main

import (
	"os"
	"web_test01/app"
)

func main() {
	// initialize the app
	app := app.InitializeApp()

	app.Listen(os.Getenv("ADDR"))
}
