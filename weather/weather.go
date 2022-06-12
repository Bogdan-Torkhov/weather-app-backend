package weather
import (
	ya "github.com/Bogdan-Torkhov/go-yaweather-lib/weather"
)

const yakey = "47ec4e13-b304-4e7b-bd7c-0ca727d87fbc"

func NewWeather() (w *ya.Weather, err error) {
	w, err = ya.GetWeather(yakey, 45.0355, 38.975, "ru_RU")
	if err != nil { return }
	return
}
