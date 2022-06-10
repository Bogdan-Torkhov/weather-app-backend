package main

import (
	// "fmt"
	"log"
	"go-weather-app/weather"

	// owm "github.com/briandowns/openweathermap"
)

/*
const apiKey = "c782b17fcc6669fa73afe6ba68e7f4f9"

func main() {
	w, err := owm.NewCurrent("C", "ru", apiKey)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(apiKey)
	err = w.CurrentByName("Murmansk")
	if err != nil {
		panic(err)
	}
	fmt.Println(w.Weather)
}
*/

func main() {
	w, err := weather.NewWeather()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(w)
}
