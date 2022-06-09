package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	owm "github.com/briandowns/openweathermap"
)

const proxy = "http://38.55.182.255:80"

const apiKey = "c782b17fcc6669fa73afe6ba68e7f4f9"

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
