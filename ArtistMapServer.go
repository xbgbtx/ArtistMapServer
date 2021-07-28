package main

import (
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my web server!</h1>"))
}

func locationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Locations!</h1>"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/locations", locationsHandler )
	log.Fatal(http.ListenAndServe(":8080", mux))
}


