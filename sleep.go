/*
Copyright 2019, 2021, Staffbase GmbH and contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.

You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	"time"
)

var sleeper = promauto.NewHistogram(prometheus.HistogramOpts{
	Name: "things_sleep_requested_seconds",
	Help: "The sleeping time the users requested.",
})

func sleep(w http.ResponseWriter, r *http.Request) {
	sleepParam := r.URL.Query().Get("sleep")
	// use build in time string parsing
	duration, err := time.ParseDuration(sleepParam)
	if err == nil {
		fmt.Printf("Sleeping for %s.\n", duration.String())
		time.Sleep(time.Duration(duration))
		sleeper.Observe(duration.Seconds())
		_, _ = fmt.Fprintf(w, "Sleeping for %s.\n", duration.String())
	} else {
		_, _ = fmt.Fprintf(w, "Could not parse input '%s'\n", sleepParam)
		_, _ = fmt.Fprint(w, "Try using localhost:8080/sleep?sleep=10s\n")
	}
}
