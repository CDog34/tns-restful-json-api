package main

import (
	"log"
	"net/http"
)

const (
	APP_VERSION = "0.0.1"
	API_VERSION = "0.0.1"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
