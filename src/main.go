package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
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

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	_, err := w.Write([]byte(message))
	if err != nil {
		greetings.Inc()
		if message != "" {
			greetingsName.WithLabelValues(message).Inc()
		}
	}
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
