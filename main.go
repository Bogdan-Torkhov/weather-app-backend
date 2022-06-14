package main

import (
	"log"
	"os"
	SERVE "https://github.com/Bogdan-Torkhov/weather-app-backend/serve"
)

var Logio *log.Logger

func init() {
	f, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE, 0o755)
	if err != nil {
		panic(err)
	}

	Logio = log.New(f, "Debug :: ", log.Ldate)

	Logio.Print("Started!")
}

func main() {
	SERVE.start()
}
