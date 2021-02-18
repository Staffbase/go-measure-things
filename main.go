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
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"net/http"
	"os"
)

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

func main() {
	fmt.Println("Starting a local server on http://localhost:8080 ...")
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/sleep", sleep)
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
