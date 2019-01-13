package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	"strconv"
	"time"
)

var sleeper = promauto.NewSummary(prometheus.SummaryOpts{
	Name: "sleep_",
	Help: "The sleeping time the users requested.",
})

func sleep(w http.ResponseWriter, r *http.Request) {
	sleepParam := r.URL.Query().Get("sleep")
	duration, err := strconv.ParseFloat(sleepParam, 32)
	if err == nil {
		time.Sleep(time.Duration(duration))
		sleeper.Observe(duration)
		_, _ = fmt.Fprintf(w, "Sleept for %f milliseconds.\n", duration)
	} else {
		_, _ = fmt.Fprintf(w, "Could not parse input '%s'\n", sleepParam)
		_, _ = fmt.Fprint(w, "Try using localhost:8080/sleep?sleep=1000\n")
	}
}
