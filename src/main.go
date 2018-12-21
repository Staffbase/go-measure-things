package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strings"
)

var (
	greetings = promauto.NewCounter(prometheus.CounterOpts{
		Name: "things_send_greetings_total",
		Help: "How often the greetings where send.",
	})
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
	greetings.Inc()
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
