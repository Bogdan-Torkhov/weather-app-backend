package weather
import (
	// owm "github.com/briandowns/openweathermap"
	// "errors"
	ya "github.com/Bogdan-Torkhov/go-yaweather-lib/weather"
	"log"
)

// const apiKey = "c782b17fcc6669fa73afe6ba68e7f4f9"

// For a time ::Start::

// const location = "Moscow"

// const c = "C" // celsius "C" or fahrenheit "F"

// const lang = "ru" // language to use

// ::End::

type Weather struct {
	Desc string
	Temp float64
	Feels float64
	WindSpeed float64
	AirHumidity uint
}

const yakey = "47ec4e13-b304-4e7b-bd7c-0ca727d87fbc"

func NewWeather() (w *Weather, err error) {
	w = new(Weather)
	err = nil
	we, err := ya.GetWeather(yakey, 45.0355, 38.975, "ru_RU")
	if err != nil { return }
	log.Println(we)
	return
}
