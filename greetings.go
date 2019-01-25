package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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
