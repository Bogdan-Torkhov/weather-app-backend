package main

import (
	"log"
	// "go-weather-app/weather"
	"os"
	// "fmt"
)

var Logio *log.Logger

func init() {
	f, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	Logio = log.New(f, "Debug :: ", log.Ldate)
	

	Logio.Print("Started!")

}

func main() {
	// w, err := weather.NewWeather()
	// if err != nil {
		// log.Fatal(err)
	// }
	// newObjJson(w)
	// Logio.Println(w)
	start()
}
