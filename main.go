package main

import (
	"log"
	"fmt"
	"os"
    owm "github.com/briandowns/openweathermap"
)

var apiKey = os.Getenv("OWM_API_KEY")


func main() {
	w, err := owm.NewCurrent("F", "ru", apiKey) // fahrenheit (imperial) with Russian output
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName("Murmansk")
	fmt.Println(w)
}
