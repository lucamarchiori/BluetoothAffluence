package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HttpHandler() {

}

func routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/store", store)
	err := http.ListenAndServe(":4000", r)
	log.Fatal(err)
	return r
}
