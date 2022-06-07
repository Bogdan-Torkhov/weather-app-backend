package main

import (
	"log"
	"fmt"
	"os"
    owm "github.com/briandowns/openweathermap"
)

var apiKey = os.Getenv("OWM_API_KEY")
