package serve

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	ya "github.com/Bogdan-Torkhov/go-yaweather-lib/weather"
)

type objJson struct {
	ValueName string `json:"valueName"`
	Value     int    `json:"value"`
}

type iod []objJson

func newObjJson(we *ya.Weather) string {
	sl := []objJson{}
	sl = append(sl, objJson{"temp", we.Fact.Temp})
	sl = append(sl, objJson{"feelslike", we.Fact.FeelsLike})
	s, err := json.Marshal(sl)
	if err != nil {
		panic(err)
	}
	log.Println(string(s))
	return string(s)
}

type server struct {
	apiServe func(w http.ResponseWriter, _ *http.Request)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

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
		enableCors(&w)
		io.WriteString(w, fakeData)
	}
	return
}

func (c *server) startServer() {
	http.HandleFunc("/api", c.apiServe)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
	log.Println("Server :: OK")
}

func Start() {
	c := newServer()
	c.startServer()
}
