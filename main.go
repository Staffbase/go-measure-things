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
