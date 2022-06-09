package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	owm "github.com/briandowns/openweathermap"
)

const proxy = "http://38.55.182.255:80"

const apiKey = "6ab7fc8050c5668c33576f3a93d1deea"

func main() {
	i, err := url.Parse(proxy)
	if err != nil {
		panic(err)
	}
	cli := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(i)}}
	w, err := owm.NewCurrent("C", "ru", apiKey, owm.WithHttpClient(cli))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(apiKey)
	err = w.CurrentByName("Phoenix")
	if err != nil {
		panic(err)
	}
	fmt.Println(w.Weather)
}
