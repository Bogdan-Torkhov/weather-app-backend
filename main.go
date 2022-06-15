package main

import (
	"log"
	"os"

	"go-weather-app/serve"
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
	serve.Start()
}
