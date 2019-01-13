package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Nice golang examples for prometheus:
//   https://prometheus.io/docs/guides/go-application/
var greetings = promauto.NewCounter(prometheus.CounterOpts{
	Name: "things_sent_greetings_total",
	Help: "How often the greetings where sent.",
})

var greetingsName = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "things_sent_greetings_personal_total",
	Help: "How often the greetings where sent to a specific person.",
}, []string{"name"})

func loadIndex(w http.ResponseWriter) {
	w.WriteHeader(200)
	index, err := ioutil.ReadFile("static/index.html")
	if err != nil {
		fmt.Println("ERROR: Unable to load static index file!")
		dir, _ := os.Getwd()
		fmt.Println("Current working dir is: ", dir)
	}
	_, _ = w.Write(index)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path
	name = strings.TrimPrefix(name, "/")
	if name == "" {
		loadIndex(w)
	} else if name == "favicon.ico" {
		w.WriteHeader(404)
	} else {
		message := "Hello " + name
		_, err := w.Write([]byte(message))
		if err == nil {
			greetings.Inc()
			greetingsName.WithLabelValues(name).Inc()
		}
	}
}

func main() {
	fmt.Println("Starting a local echo server on http://localhost:8080 ...")
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/sleep", sleep)
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
