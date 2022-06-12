package main
import (
	"net/http"
	// "fmt"
	ya "github.com/Bogdan-Torkhov/go-yaweather-lib/weather"
	"log"
	"sync"
	"io"
	"encoding/json"
)

type objJson struct {
	ValueName string `json:"valueName"`
	Value int `json:"value"`
}

type iod []objJson

func newObjJson(we *ya.Weather) string {
	sl := []objJson{}
	sl = append(sl, objJson{"temp", we.Fact.Temp})
	sl = append(sl, objJson{"feelslike", we.Fact.FeelsLike})
	s, err := json.Marshal(sl)
	if err != nil { panic(err) }
	log.Println(string(s))
	return string(s)
}

type server struct {
	apiServe func(w http.ResponseWriter, _ *http.Request)
}

// HELP ME

func newServer() (c *server) {
	fakeData := `
	[
		{
			"valueName": "temp",
			"value": 37,
		},
		{
			"valueName": "feels",
			"value": 36,
		}
		
	]`
	c = new(server)
	c.apiServe = func(w http.ResponseWriter, _ *http.Request) {
		// weat, err := weather.NewWeather()
		io.WriteString(w, fakeData)
	}
	return
}

func (c *server) startServer() {
	var wg sync.WaitGroup
	http.HandleFunc("/api", c.apiServe)
	go func() {
	wg.Add(1)
	http.ListenAndServe(":80", nil)
	wg.Done()
	}()
	log.Println("Server :: OK")
	wg.Wait()
}

func start() {
	c := newServer()
	c.startServer()
}
